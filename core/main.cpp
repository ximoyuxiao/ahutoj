#include<iostream>
#include<vector>
#include <string>
#include <chrono>
#include <thread>

#include<unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

#include <mysql/mysql.h>
#include "result.h"
#include "readConfig.h"
#include "mlog.h"
#include "Redis/redis.h"
#include "Tpool/threadpool.h"
#include "Solve.h"
#include "judgeClient.h"
#include "mydb.h"
#include "SolutionDb.h"
using std::vector;
using std::cout;
using std::cin;
using std::endl;
using namespace my;
SolutionDb solutionDB;
static int init_daemon();
class soulution_run:public worker
{
private:
    judgeClient* jc;
    Solve* solve;
public:
    soulution_run(Solve* solve){
        this->jc =new judgeClient(solve);
        this->solve = solve;
    };
    virtual void run(){
        jc->judge();
        ILOG("judge complete Result:%s",runningres[solve->Sres()]);
        solutionDB.commitSolveToDb(solve);
        solutionDB.ReleaseSolve(solve);
    }
    judgeClient* getJudgeClient(){
        return jc;
    }
    virtual ~soulution_run(){
        if(jc != nullptr)
            delete jc;
    };
};


static bool init_db(readConfig *rcf,Redis* redis);
static bool init_Solve_pool(readConfig *rcf,threadpool **tp);/*此处实际上是创建一个解决*/
int main(int argc, char **argv)
{
    /*挂起守护进程*/
   init_daemon();
   
    /*读配置项目*/
    readConfig *rcf = new readConfig(CONF);
    threadpool *pool = nullptr; 
    Redis redis;
    if(rcf->config_init())
    {
        mlog* log = mlog::init(LOGPATH);
        /*初始化数据库*/
        if(!init_db(rcf,&redis))
        {
            ELOG("db init error");
            mlog::destory();
            delete rcf;
            exit(-2);
        }

        if(!solutionDB.initDB(rcf))
        {
            ELOG("solution db init error");
            mlog::destory();
            delete rcf;
            exit(-2);
        }

        /*初始化线程池*/
        if(!init_Solve_pool(rcf,&pool))
        {
            ELOG("database pool init error");
            delete rcf;
            mlog::destory();
            exit(-3);
        }
        /*判题线程*/
        /*这一块后期优化就考虑采用redis*/
        while(true)
        {
            vector<Solve*> solution = solutionDB.getSolve();
            while(!solution.empty())
            {
                Solve* last = solution.back();
                pool->excute(
                    shared_ptr<soulution_run>(
                        new soulution_run(
                            last
                            )
                    )
                );
                solution.pop_back();
            }
            sleep(3);
        }   
    } 
    else
    {
        fprintf(stderr,"call config_init faild,err = %s",rcf->getConferr());
        delete rcf;
        exit(-1);
    }
    
    delete rcf;
    return 0;
}

int init_daemon(void)
{
    pid_t pid = fork();
    if (pid < 0)
        return -1;
    if(pid)
        exit(0);
    
    close(0);
    close(1);
    close(2);
    
    int fd = open("/dev/null",O_RDWR);
    dup2( fd, 0 );
	dup2( fd, 1 );
	dup2( fd, 2 );
	if ( fd > 2 )
		close( fd );
    setsid();
    return 0;
}

bool init_DB(readConfig *rcf,Redis *redis)
{
    char host[128],user[128],pwd[128],dbase[128];
    int port;
    rcf->getCOnfigString(host,"MYSQL","host");
    rcf->getCOnfigString(user,"MYSQL","user");
    rcf->getCOnfigString(pwd,"MYSQL","password");
    rcf->getCOnfigString(dbase,"MYSQL","db");
    port = rcf->getCOnfigInt("MYSQL","port");
    char msg[1024];
    ILOG(msg,"%s %s %s %s %d\n",host,user,pwd,dbase,port);
    mysqlDB::initConn(host,user,pwd,dbase,port);
    return false;
}

bool init_Solve_pool(readConfig *rcf,threadpool **tp)
{
    int core = rcf->getCOnfigInt("Thread","Core");
    int maxThread = rcf->getCOnfigInt("Thread","MaxThread");
    int maxqueue = rcf->getCOnfigInt("Thread","MaxQueue");
    int livetime = rcf->getCOnfigInt("Thread","LiveTime");
    *tp = threadpool::getPool(core,maxThread,livetime,maxqueue);
    if(*tp) return true;
    return false;
}
