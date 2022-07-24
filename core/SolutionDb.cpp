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
        ILOG("host:%s,port:%d\n",host,port);
        redis->connect(host1,port);
    }
    return true;
}

vector<Solve*> SolutionDb::getSolve(){
    vector<Solve*> ret;
    char sql[256] = "";
    sprintf(sql,"select sid,pid,uid,cid,source,lang from Submit where result='pendding' or result = 'rejuged'");
    auto db = mysqlDB::getInstance();
    MYSQL mysql;
    db->getDatabase(&mysql);
    mysql_query(&mysql,sql);
    MYSQL_RES *res = mysql_store_result(&mysql);
    if(res == NULL)
    {
        return ret;
    }
    int rows = mysql_num_rows(res);
    if(!rows) return ret;
    MYSQL_ROW row;
    while((row = mysql_fetch_row(res)))
    {
       Solve* solve = new  Solve();
       solve->Cid(atoi(row[0]));
       solve->Pid(atoi(row[1]));
       solve->Uid(atoi(row[2]));
       solve->Cid(atoi(row[3]));
       solve->Source(row[4]);
       solve->Sres(OJ_JUDGE);
       solve->Lang((lanuage)atoi(row[5]));
       ret.push_back(solve);
       commitSolveToDb(solve);
    }
    mysql_free_result(res);
    mysql_close(&mysql);
    return ret;
}

bool SolutionDb::commitSolveToDb(Solve* solve){
    // insert into Submit values (null,#{pid},#{uid},#{cid},#{judgeid},#{source},#{lang},'Judgeing',0,0,#{submitTime})
    char sql[256] ="";
    sprintf(sql,"update Submit judgeid=%d,result=%s,usetime=%d,memory=%d where sid=%d",
        solve->getjudgeID(),
        runningres[solve->Sres()],
        solve->getUsetime(),
        solve->getuseMemory(),
        solve->Cid()
    );
    ILOG("update mysql:%s",sql);
    auto db = mysqlDB::getInstance();
    MYSQL mysql;
    db->getDatabase(&mysql);
    int res = mysql_query(&mysql,sql) == 0;
    if(solve->Sres() == OJ_CE){
        sprintf(sql,"insert into CEINFO values(%d,%s)",solve->Cid(),solve->ceInfo().c_str());
        mysql_query(&mysql,sql);
    }
    db->CloseDatabase(&mysql,nullptr);
    return res;
}
void SolutionDb::ReleaseSolve(Solve* solve){
    delete solve;
}