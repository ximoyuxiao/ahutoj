#include<iostream>

#include<cstdlib>
#include<dirent.h>
#include<regex.h>
#include<cstring>
#include<unistd.h>
#include<signal.h>
#include<fcntl.h>

#include<sys/time.h>
#include<sys/resource.h>
#include<sys/types.h>
#include<sys/stat.h>
#include<sys/wait.h>

#include<mysql/mysql.h>
#include "judgeClient.h"
#include "mlog.h"

#define MAXBUFF 1024
using namespace std;
using namespace my;
static pid_t pid ;
static sighandler_t oldAlarmHandle;
static sighandler_t oldHupHandle;
static SubRes status = JUDGE;
static void Alam_Handle(int s){
    if(!kill(pid,SIGKILL))
        status  = TLE;  
    struct itimerval time;
    time.it_interval.tv_usec =0;
    time.it_interval.tv_sec = 0;
    time.it_value.tv_sec = 0;
    time.it_value.tv_usec = 0;
    setitimer(ITIMER_REAL,&time,NULL);
    signal(SIGALRM,oldAlarmHandle);
}
static void kiil_handle(int s)
{
    if(s == SIGKILL)
    {
        exit(100);
    }
}
static void upsig_handle(int s)
{
    if(status !=TLE)
        wait(NULL);
    else
        wait(NULL);
    signal(SIGHUP,oldHupHandle);
    
}
bool judgeClient::checkSource(){
    return true;
}
bool judgeClient::compile()
{
    char comp[MAXBUFF];
    char sourceFile[128];
    switch (solve->Lang())
    {
        case C:
        {
            sprintf(sourceFile,"%s/%d.c",dir,solve->Pid());
            FILE* fp = fopen(sourceFile,"w");
            fprintf(fp,"%s",solve->Source().c_str());
            fclose(fp);
            sprintf(comp,"gcc %s -o %s/main 2>err.txt",sourceFile,dir);
            system(comp);
            break;
        }
        case CPP:
        {
            sprintf(sourceFile,"%s/%d.cpp",dir,solve->Pid());
            DLOG("SourceFile:%s",sourceFile);
            FILE* fp = fopen(sourceFile,"w");
            fprintf(fp,"%s",solve->Source().c_str());
            fclose(fp);
            sprintf(comp,"g++ %s -o %s/main 2>%s/err.txt",sourceFile,dir,dir);
            system(comp);
            break;
        }
        case CPP11:
        {
            sprintf(sourceFile,"%s/%d.cpp",dir,solve->Pid());
            DLOG("SourceFile:%s",sourceFile);
            FILE* fp = fopen(sourceFile,"w");
            fprintf(fp,"%s",solve->Source().c_str());
            DLOG("source:\n%s",solve->Source().c_str());
            fclose(fp);
            sprintf(comp,"g++ %s -o %s/main -std=c++11 2> %s/err.txt",sourceFile,dir,dir);
            DLOG("%s",comp);
            system(comp);
            break;
        }
        case CPP17:
        {
            sprintf(sourceFile,"%s/%d.cpp",dir,solve->Pid());
            FILE* fp = fopen(sourceFile,"w");
            fprintf(fp,"%s",solve->Source().c_str());
            fclose(fp);
            sprintf(comp,"g++ %s -o %s/main -std=c++17 2>%s/err.txt",sourceFile,dir,dir);
            system(comp);
            break;
        }
        case JAVA:
        {
            sprintf(sourceFile,"%s/Main.java",dir);
            FILE* fp = fopen(sourceFile,"w");
            fprintf(fp,"%s",solve->Source().c_str());
            fclose(fp);
            sprintf(comp,"javac %s %s/Main.class 2>%s/err.txt",sourceFile,dir,dir);
            system(comp);
            break;
        }
        case PYTHON3:
        {
            return false;
            break;
        }
        default:
            break;
    }
    DLOG("compile.txt");
    char ceFile[128]="";
    sprintf(ceFile,"%s/err.txt",dir);
    if(getFileSize(ceFile)!=0)
        return false;
    return true;
}

long long judgeClient::getFileSize(const char * filepath)
{
    struct stat mystat;
    stat(filepath,&mystat);
    return (long long)mystat.st_size;
}

bool judgeClient::running(SubRes &result,const char * runFile,const char *resFile)
{

    pid = fork();
    if(pid < 0)
    {
        DLOG("fork:%s",strerror(errno));
        return false;
    }
    if(pid)
    {
        oldAlarmHandle =  signal(SIGALRM,Alam_Handle);
        // oldMemHandle = signal(SIGCHLD,Alam_Handle);
        itimerval time;
        time.it_interval.tv_usec =0;
        time.it_interval.tv_sec = 0;
        time.it_value.tv_sec = this->solve->LimitTime();
        time.it_value.tv_usec = 1000;
        setitimer(ITIMER_REAL,&time,NULL);
        // setrlimit()
        int ret = 0;
        wait(&ret);
        if(ret != 0)
            result =RE;
        if(status != JUDGE)
            result = status;
        // prlimit();
    }
    else
    {
        close(STDOUT_FILENO);
        close(STDIN_FILENO);
        int rfd = open(runFile,O_RDONLY);
        int wfd = open(resFile,O_RDWR|O_CREAT,0777);
        dup2(rfd,STDIN_FILENO);
        dup2(wfd,STDOUT_FILENO);
        signal(SIGKILL,kiil_handle);
        char path[128]={0};
        sprintf(path,"%s/main",dir);
        int ret = 0;
        ret = execlp(path,"main",NULL);
        exit(WEXITSTATUS(ret));
    }
    

}
bool judgeClient::getFiles()
{
    char path[MAXBUFF];
    sprintf(path,"%s/%d/",DATAPATH,solve->Pid());
    ILOG("%s",path);
    inputFiles.clear();
    DIR *dir = opendir(path);
    dirent *dirp;
    if(dir ==NULL)
    {
        ELOG("inputFiles:%s",strerror(errno));     
        return false;    
    }
    else
    {
        while( (dirp = readdir(dir) )!=NULL)
        {
            int len = strlen(dirp->d_name);
            if( len< 3)
                continue;
            if(strcmp(dirp->d_name + len -3,".in") == 0)
            {
                DLOG("file:%s",dirp->d_name);
                string infile = dirp->d_name;
                string outfile = infile.substr(0,len - 2) + "out";
                inputFiles.push_back(path + infile);
                outputFiles.push_back(path + outfile);
            }
        }
        closedir(dir);   
    }
    return true;
}
bool judgeClient::judgePE(FILE*source,FILE *res)
{
    fseek(source,SEEK_SET,0);
    fseek(res,SEEK_SET,0);
    char sourcech=0,resch=0;
    bool tail = true;
    while(tail)
    {
        while(sourcech = fgetc(source))
        {
            if(sourcech == '\n' || sourcech ==' ')  continue;
            if(sourcech == -1)
            {
                tail = false;
                break;
            }
            break;
        }
        while(resch= fgetc(res))
        {
            if(resch == '\n' || resch ==' ')  continue;
            if(resch == -1)
            {
                tail = false;
                break;
            }
            break;
        } 
        if(tail)
        {
            if(resch != sourcech)
                return false;
        }
    }
    return sourcech == resch;
}
bool judgeClient::cmpFIle(SubRes &result,char *myfile,const char* sourceFile)
{
    if(result !=AC)
        return false;
    char diffFile[128];
    sprintf(diffFile,"%s/diff",dir);
    char cmd[128];
    sprintf(cmd,"diff %s %s > %s",myfile,sourceFile,diffFile);
    DLOG(cmd);
    system(cmd);
    if(getFileSize(diffFile))
    {
        FILE* fp1 = fopen(myfile,"r");
        FILE* fp2 = fopen(sourceFile,"r");
        if(judgePE(fp1,fp2))
        {
            result = PE;
        }
        else
        {
            result = WA;
        }
        fclose(fp1);
        fclose(fp2);
        return false;
    }
    return true;
}
Solve* judgeClient::GetSolve(){
    return this->solve;
}
Solve* judgeClient::SetSolve(Solve* solve){
    this->solve = solve;
}
bool judgeClient::judge()
{
    status = JUDGE;
    sprintf(dir,COMPDIR,solve->getjudgeID());
    mkdir(dir,0777);
    while(true)
    {
        switch (this->Jstat)
        {
            case J_CHECK:
            {
                ILOG("J_CHECK");
                if(checkSource()){
                    Jstat = J_GETFILE;
                }
                else
                {
                    solve->Sres(CE);
                    Jstat = J_CE;
                }
                break;
            }
            case J_GETFILE:
            {
                ILOG("J_GETFILE");
                getFiles();
                Jstat = J_COMPILE;
                break;
            }
            case J_COMPILE:
            {
                ILOG("J_COMPILE");
                if(compile()){
                    Jstat = J_RUNNING;
                }
                else{
                    Jstat = J_CE;
                    solve->Sres(CE);
                }
                break;
            }
            case J_RUNNING:
            {
                ILOG("J_RUNNING");
                SubRes res  = AC;
                char resoutfile[128];
                sprintf(resoutfile,"%s/ans",dir);
                for(int i = 0;i<inputFiles.size();i++)
                {
                    DLOG("runnning:%s",inputFiles[i].c_str());
                    running(res,inputFiles[i].c_str(),resoutfile);
                    DLOG("runned:%s",outputFiles[i].c_str());
                    if(res != AC)
                        break;
                    cmpFIle(res,resoutfile,outputFiles[i].c_str());
                    if(res != AC)
                        break;
                    
                }
                if(res != AC)
                {
                    Jstat = J_FAILED;
                    solve->Sres(res);
                }
                else
                {
                    Jstat = J_SUCESS;
                }
                break;
            }
            case J_CE:{
                //处理CE事件。。。
                DLOG("%d:CE",solve->Sid());
                Jstat = J_FAILED;
                break;
            }
            case J_SUCESS:
            {
                DLOG("%d:AC",solve->Sid());
                solve->Sres(AC);
                char del[1024] = "";
                sprintf(del,"rm -rf %s",dir);
                system(del);
                return true;
            }
            case J_FAILED:
            {
                DLOG("%d:%s",solve->Sid(),runningres[solve->Sres()]);
                char del[1024] = "";
                sprintf(del,"rm -rf %s",dir);
                system(del);
                return false;
            }
            default:
                char del[1024] = "";
                sprintf(del,"rm -rf %s",dir);
                system(del);
                return false;
        }
    }

}
judgeClient::judgeClient(Solve *solve){
    Jstat = J_CHECK;
    this->solve = solve;
}
// static void judgeClient_test()
// {
//     Solve *solve;
//     solve = new Solve(1,1,"#include<iostream>\nusing namespace std;\nint main(){\n\tfor(int i=0;;i++){int a =0;}\n\texit(0);\n}",1,128,CPP11);
//     judgeClient juc(*solve);
//     status = JUDGE;
//     juc.judge();
//     return ;
// }
// int main(int argc, char const *argv[])
// {
//     mlog::init("./log");
//     judgeClient_test();
//     mlog::destory();
//     // ShmemQueue<Solve> squeue(IPC_PATH);
//     // squeue.get_queue();
//     // while(true)
//     // {
//     //     //这边应该用信号量那一套、
//     //     while(!solve)
//     //     {
//     //         solve = squeue.pop();
//     //     }
//     //     judgeClient juc(*solve);
//     //     juc.judge();
//     //     solve = nullptr;
//     // }
//     return 0;
// }
