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
    if (redis == nullptr){
        char host[128]="";
        rcf->getCOnfigString(host,"REDIS","host");
        int port = rcf->getCOnfigInt("REDIS","port");
        redis->connect(host,port);
    }
    return true;
}

vector<Solve*> SolutionDb::getSolve(){
    char sql[256] = "";
    sprintf(sql,"select sid,pid,uid,cid,source,lang from Submit where result='Judgeing'");
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
    db->CloseDatabase(&mysql,nullptr);
    return res;
}
void SolutionDb::ReleaseSolve(Solve* solve){
    delete solve;
}