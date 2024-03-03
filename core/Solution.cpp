#include "Solution.h"
#include "result.h"
#include "Tpool/locker.h"
#include "mlog.h"
#include"judgeClient.h"
#include<iostream>
using namespace my;
Solution* Solution::solution = nullptr;
Solution::Solution()
{
    this->redis = nullptr;
    mq = nullptr;
}

Solution::~Solution()
{
    if(redis != nullptr)
        redis->close();
    delete redis;
    if(mq){
        delete mq;
    }
}

Solution* Solution::GetInstance(){
    if(solution == nullptr){
        return solution = new Solution();
    }
    return solution;
}

void Solution::Destory(){
    if(solution) delete solution;
}

bool Solution::init(readConfig* rcf)
{
    char host[128]="";
    if (redis == nullptr){
        redis = new MyRedis();
        rcf->getCOnfigString(host,"REDIS","host");
        string host1 = host;
        int port = rcf->getCOnfigInt("REDIS","port");
        char password[128]="";
        rcf->getCOnfigString(password,"REDIS","password");
        ILOG("host:%s,port:%d\n",host,port);
        redis->connect(host1,port,password);
    }
    if(mq == nullptr){
        mq = new RabbitMQ(RMQ_HOST,atoi(RMQ_PORT),RMQ_USER,RMQ_PASS);
    }
    return true;
}

static vector<string> splite(string key,char sp){
    vector<string> ret;
    string s = "";
    for(int i = 0;i< static_cast<int>(key.size());i++){
        if(key[i] != sp){
            s+= key[i];
        }else{
            ret.push_back(s);
            s = "";
        }
    }
    if(s != ""){
        ret.push_back(s);
    }
    return ret;
}

void Solution::GetProblemInfo(Solve* solve){
    auto value = redis->getString(solve->Pid());
    ILOG("value:%s",value.c_str());
    if (value == ""){
        char sql[256]="";
        sprintf(sql,"select LimitTime,LimitMemory,SpjJudge from Problem where PID=\'%s\'",solve->Pid().c_str());
        ILOG(sql);
        auto db = mysqlDB::getInstance();
        MYSQL mysql;
        db->getDatabase(&mysql);
        mysql_query(&mysql,sql);
        MYSQL_RES *res = mysql_store_result(&mysql);
        if(res == NULL)
        {
            db->CloseDatabase(&mysql,nullptr);
            return ;
        }
        int rows = mysql_num_rows(res);
        if(!rows){
            db->CloseDatabase(&mysql,res);
            return ;
        }    
        MYSQL_ROW row;
        if((row = mysql_fetch_row(res))){
            if(row[0]){
                solve->LimitTime(atoll(row[0]));
            }
            if(row[1]){
                solve->LimitMemory(atoll(row[1]));
            }
            if(row[2]){
                solve->setSpjJudge(atoi(row[2]));
            }else{
                solve->setSpjJudge(-1);
            }
            char temp[128];
            sprintf(temp,"%s,%s,%s",row[0],row[1],row[2]);
            redis->setString(solve->Pid(),temp);
        }
    }else{
        auto values = splite(value,',');
        solve->LimitTime(atoll(values[0].c_str()));
        solve->LimitMemory(atoll(values[1].c_str()));
        solve->setSpjJudge(-1);
        if(values.size() != 2)
            solve->setSpjJudge(atoi(values[2].c_str()));
    }
    ILOG("ltime:%lld,memory:%lld,spj:%d",solve->LimitTime(),solve->LimitMemory(),solve->getSpjJudge());
}
class soulution_run:public worker
{
private:
    judgeClient* jc;
    Solve* solve;
public:
    sem wait;
public:
    soulution_run(Solve* solve){
        this->jc =new judgeClient(solve);
        this->solve = solve;
        this->solve->setjudgeID(this->solve->Sid());
    };
    virtual void run(){
        jc->judge();
        DLOG("judge complete Result:%s",runningres[solve->Sres()]);
        if(solve->Sres() == OJ_JUDGE){
            solve->Sres(OJ_FAILED);
        }
        auto solution = Solution::GetInstance();
        solution->commitSolveToQueue(solve);
        solution->ReleaseSolve(solve);
        wait.post();
    }
    judgeClient* getJudgeClient(){
        return jc;
    }
    virtual ~soulution_run(){
        if(jc != nullptr)
            delete jc;
    };
};

using json = nlohmann::json;
void Solution::Process(amqp_envelope_t amqp){
    string data = (char*)(amqp.message.body.bytes);
    data[amqp.message.body.len] = '\0';
    std::cout<<data<<std::endl;
    json j = nlohmann::json::parse(data);
    Solve* solve = new Solve();
    solve->from_json(j);
    solution->GetProblemInfo(solve);
    solution->commitSolveToQueue(solve);
    auto excute = shared_ptr<soulution_run>(new soulution_run(solve));
    auto pool =  threadpool::getPool();
    pool->excute(excute);
    excute->wait.wait();
}
void Solution::LoopSolve(){
    while(true){
        Consumer consumer = mq->createConsumer(INNERJUDGE);
        consumer.consumeMessage(Process);
        if(ret){
            sleep(5);//延时重连，避免浪费资源
        }
    }
}

char sql[102400] ="";
// 等待优化
locker sqlLock;
bool Solution::commitSolveToDb(Solve* solve){
    // insert into Submit values (null,#{pid},#{uid},#{cid},#{judgeid},#{source},#{lang},'Judgeing',0,0,#{submitTime})
    sqlLock.lock();
    sprintf(sql,"update Submit set JudgeID=%d,ResultACM='%s',UseTime=%lld,UseMemory=%lld,PassSample=%lld,SampleNumber=%lld where SID=%d",
        solve->getjudgeID(),
        runningres[solve->Sres()],
        solve->getUsetime(),
        solve->getuseMemory(),
        solve->getPassSample(),
        solve->getSampleNumber(),
        solve->Sid()
    );
    ILOG("update mysql:%s",sql);
    auto db = mysqlDB::getInstance();
    MYSQL mysql;
    db->getDatabase(&mysql);
    int res = mysql_query(&mysql,sql) == 0;
    if(solve->Sres() == OJ_CE){
        sprintf(sql,"insert into CEINFO values(%d,'%s')",solve->Sid(),solve->ceInfo().c_str());
        ILOG("insert mysql:%s",sql);
        mysql_query(&mysql,sql);
    }
    sqlLock.unlock();
    db->CloseDatabase(&mysql,nullptr);
    DLOG("Close DB:%d",solve->Sid());
    return res;
}

void Solution::commitSolveToQueue(Solve* solve){
    json j;
    solve->to_json(j);
    Producer pro = mq->createProducer();
    auto data = j.dump();
    pro.sendMessage(JUDGERESULT,(void*)data.c_str(),data.size());
    j.clear();
    data.clear();
    if(solve->Sres() == OJ_CE){
        solve->to_ceJson(j);
        data = j.dump();
        pro.sendMessage(JUDGECE,(void*)data.c_str(),data.size());
    }
}

void Solution::ReleaseSolve(Solve* solve){
    DLOG("---begin solve");
    if(solve != nullptr)
        delete solve;
    DLOG("---end delete solve!\n-----");
}