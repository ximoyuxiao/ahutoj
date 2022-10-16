#include "SolutionDb.h"
#include "result.h"
#include "mlog.h"
using namespace my;
SolutionDb::SolutionDb()
{
    this->redis = nullptr;
}

SolutionDb::~SolutionDb()
{
    if(redis != nullptr)
        redis->close();
    delete redis;
}

bool SolutionDb::initDB(readConfig* rcf)
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
    return true;
}
static vector<string> splite(string key,char sp){
    vector<string> ret;
    string s = "";
    for(int i = 0;i<key.size();i++){
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
void SolutionDb::GetSolveLimit(Solve* solve){
    auto value = redis->getString(to_string(solve->Pid()));
    if (value == ""){
        char sql[256]="";
        sprintf(sql,"select LimitTime,LimitMemory from Problem where PID=%d",solve->Pid());
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
            solve->LimitTime(atoll(row[0]));
            solve->LimitMemory(atoll(row[1]));
            char temp[128];
            sprintf(temp,"%s,%s",row[0],row[1]);
            redis->setString(to_string(solve->Pid()),temp);
        }
    }else{
        auto values = splite(value,',');
        solve->LimitTime(atoll(values[0].c_str()));
        solve->LimitMemory(atoll(values[1].c_str()));
    }
   
}
vector<Solve*> SolutionDb::getSolve(){
    vector<Solve*> ret;
    char sql[256] = "";
    sprintf(sql,"select SID,PID,UID,CID,Source,Lang From Submit where IsOriginJudge = 0 and (Result='PENDING' or Result = 'REJUDGING')");
    auto db = mysqlDB::getInstance();
    MYSQL mysql;
    db->getDatabase(&mysql);
    mysql_query(&mysql,sql);
    MYSQL_RES *res = mysql_store_result(&mysql);
    if(res == NULL)
    {
        db->CloseDatabase(&mysql,nullptr);
        return ret;
    }
    int rows = mysql_num_rows(res);
    if(!rows){
        db->CloseDatabase(&mysql,res);
        return ret;
    }    
    MYSQL_ROW row;
    while((row = mysql_fetch_row(res)))
    {
       Solve* solve = new  Solve();
       solve->Sid(atoi(row[0]));
       solve->Pid(atoi(row[1]));
       solve->Uid(atoi(row[2]));
       solve->Cid(atoi(row[3]));
       solve->Source(row[4]);
       solve->Sres(OJ_JUDGE);
       solve->Lang((lanuage)atoi(row[5]));
       ret.push_back(solve);
       GetSolveLimit(solve);
       commitSolveToDb(solve);
    }
    db->CloseDatabase(&mysql,res);
    return ret;
}

char sql[102400] ="";
bool SolutionDb::commitSolveToDb(Solve* solve){
    // insert into Submit values (null,#{pid},#{uid},#{cid},#{judgeid},#{source},#{lang},'Judgeing',0,0,#{submitTime})
    sprintf(sql,"update Submit set JudgeID=%d,Result='%s',UseTime=%lld,UseMemory=%lld where SID=%d",
        solve->getjudgeID(),
        runningres[solve->Sres()],
        solve->getUsetime(),
        solve->getuseMemory(),
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
    db->CloseDatabase(&mysql,nullptr);
    DLOG("Close DB:%d",solve->Sid());
    return res;
}
void SolutionDb::ReleaseSolve(Solve* solve){
    DLOG("---begin solve");
    if(solve != nullptr)
        delete solve;
    DLOG("---end delete solve!\n-----");
}