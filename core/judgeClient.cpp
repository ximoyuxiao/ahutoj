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
#include <sys/ptrace.h>
#include <sys/user.h>
#include<mysql/mysql.h>
#include "judgeClient.h"
#include "mlog.h"
#include "okcalls.h"

#define MAXBUFF 1024
#define BUFFER_SIZE 4096
#define STD_MB 1024*1024 // 1M = 1024K = 1024*1024B

#ifdef __x86_64__      //64位x86寄存器
#define REG_SYSCALL orig_rax
#define REG_RET rax
#define REG_ARG0 rdi
#define REG_ARG1 rsi
#endif
using namespace std;
using namespace my;

static pid_t pid ;
static sighandler_t oldAlarmHandle;
static sighandler_t oldHupHandle;
static SubRes status = OJ_JUDGE;
static double cpu_compensation = 1.0;
const int call_array_size = CALL_ARRAY_SIZE;
unsigned int call_id = 0;
int call_counter[call_array_size] = {0};
static char LANG_NAME[BUFFER_SIZE];
static void init_syscalls_limits(lanuage lang){
    memset(call_counter, 0, sizeof(call_counter));
    switch (lang)
    {
    case CPP:
    case CPP11:
    case CPP17:
    case C:{
        for (int i = 0; i == 0 || LANG_CV[i]; i++)
			call_counter[LANG_CV[i]] = HOJ_MAX_LIMIT;
        break;
    }
    case PYTHON3:{
        for (int i = 0; i == 0 || LANG_YV[i]; i++)
			call_counter[LANG_YV[i]] = HOJ_MAX_LIMIT;
        break;
    }
    default:{
        for (int i = 0; i < call_array_size; i++)
			call_counter[i] = 0;
    }
        break;
    }
}

static void setfreeTimer(){
    struct itimerval time;
    time.it_interval.tv_usec =0;
    time.it_interval.tv_sec = 0;
    time.it_value.tv_sec = 0;
    time.it_value.tv_usec = 0;
    setitimer(ITIMER_REAL,&time,NULL);
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

int get_proc_status(int pid, const char *mark)
{
	FILE *pf;
	char fn[BUFFER_SIZE], buf[BUFFER_SIZE];
	int ret = 0;
	sprintf(fn, "/proc/%d/status", pid);
	pf = fopen(fn, "re");
	int m = strlen(mark);
	while (pf && fgets(buf, BUFFER_SIZE - 1, pf))
	{

		buf[strlen(buf) - 1] = 0;
		if (strncmp(buf, mark, m) == 0)
		{
			if(1!=sscanf(buf + m + 1, "%d", &ret)) printf("proc read fail\n");
		}
	}
	if (pf)
		fclose(pf);
	return ret;
}

int get_page_fault_mem(struct rusage &ruse, pid_t &pidApp)
{
	//java use pagefault
	int m_vmpeak, m_vmdata, m_minflt;
	m_minflt = ruse.ru_minflt * getpagesize();
	if (0 && DEBUG)
	{
		m_vmpeak = get_proc_status(pidApp, "VmPeak:");
		m_vmdata = get_proc_status(pidApp, "VmData:");
		printf("VmPeak:%d KB VmData:%d KB minflt:%d KB\n", m_vmpeak, m_vmdata,
			   m_minflt >> 10);
	}
	return m_minflt;
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
        int tempmemory = 0;
        int status, sig, exitcode;
        int usetime = 0;
        int useMemory = 0;
        struct user_regs_struct reg;
	    struct rusage ruse;
        int first = true;   
        while(1){
            wait4(pid,&status,__WALL,&ruse); //等待子进程切换内核态（调用系统API或者运行状态变化）
            // 这一段也不知道干嘛的
            if (first){
                ptrace(PTRACE_SETOPTIONS, pid, NULL, PTRACE_O_TRACESYSGOOD | PTRACE_O_TRACEEXIT
                        |PTRACE_O_EXITKILL|PTRACE_O_TRACECLONE|PTRACE_O_TRACEFORK|PTRACE_O_TRACEVFORK);
            }

            // 获取程序运行内存
            tempmemory = get_proc_status(pid,"VmPeak:") << 10;
            if (tempmemory > useMemory)
			    useMemory = tempmemory;
            if (useMemory > this->solve->LimitMemory() * STD_MB){
                DLOG("res:MLE userMemory:%d",useMemory);
                result = OJ_MLE;
                ptrace(PTRACE_KILL, pid, NULL, NULL); //杀死子进程
                break;
            }

            // 子进程已经退出 ，返回值不为0则判RE
            if (WIFEXITED(status)) { 
                exitcode = WEXITSTATUS(status);
                if(exitcode){
                    DLOG("res:RE exitcode:%d",exitcode);
                    result = OJ_RE;
                }
                break;
            }
            exitcode = WEXITSTATUS(status);

            if(exitcode == 0x05 || exitcode == 0 || exitcode == 133);  //休眠或者IO 啥也不做
            else{
                if(result == OJ_AC){
                    switch (exitcode)                  // 根据退出的原因给出判题结果
				{
				case SIGCHLD:
				case SIGALRM:
					setfreeTimer();
				case SIGKILL:
				case SIGXCPU:
                    DLOG("res:TLE signal:%d",exitcode);
					result = OJ_TLE;
					usetime = solve->LimitTime() * 1000;
					break;
				case SIGXFSZ:
                    DLOG("res:OLE signal:%d",exitcode);
					result = OJ_OLE;
					break;
				default:
                    DLOG("res:RE signal:%d",exitcode);
					result = OJ_RE;
				}
                }
                ptrace(PTRACE_KILL, pid, NULL, NULL);    // 杀死出问题的进程
                break;
            }
            //  如果程序退出 并且检测到退出信号
            if (WIFSIGNALED(status))
            {
                /*  WIFSIGNALED: if the process is terminated by signal
                    *  由外部信号触发的进程终止
                    *  psignal(int sig, char *s)，like perror(char *s)，print out s, with error msg from system of sig  
                    * sig = 5 means Trace/breakpoint trap
                    * sig = 11 means Segmentation fault
                    * sig = 25 means File size limit exceeded
                    */
                sig = WTERMSIG(status);
                if (result == OJ_AC)
                {
                    switch (sig)      //根据原因给出结论
                    {
                    case SIGCHLD:
                    case SIGALRM:
                        setfreeTimer();
                    case SIGKILL:
                    case SIGXCPU:
                        DLOG("res:TLE,sig:%d",sig);
                        result = OJ_TLE;
                        break;
                    case SIGXFSZ:
                        DLOG("res:OLE,sig:%d",sig);
                        result = OJ_OLE;
                        break;
                    default:
                        DLOG("res:RE,sig:%d",sig);
                        result = OJ_RE;
                    }
                }
                break;
            }

            //禁用 sysCall
            call_id=ptrace(PTRACE_GETREGS, pid, NULL, &reg);
			call_id = ((unsigned int)reg.REG_SYSCALL) % call_array_size;

            if (call_counter[call_id])
			{
				call_counter[call_id]--;
			}
			else
			{
                DLOG("call syscall forbiden! callid:%d",call_id);
				result = OJ_RE;
				ptrace(PTRACE_KILL, pid, NULL, NULL);
			}
			call_id=0;
            // 等待下一次陷入中断
            ptrace(PTRACE_SYSCALL, pid, NULL, NULL);
            first = false;
        }
        ptrace(PTRACE_KILL, pid, NULL, NULL);    // 杀死出问题的进程
        usetime += (ruse.ru_utime.tv_sec * 1000 + ruse.ru_utime.tv_usec / 1000) * cpu_compensation; // 统计用户态耗时，在更快速的CPU上加以cpu_compensation倍数放大
        usetime += (ruse.ru_stime.tv_sec * 1000 + ruse.ru_stime.tv_usec / 1000) * cpu_compensation; // 统计内核态耗时，在更快速的CPU上加以cpu_compensation倍数放大
        solve->setUsetime(usetime + solve->getUsetime());
        kill(pid,SIGKILL);
    }
    else
    {
        lanuage lang = this->solve->Lang();
        // 默认使用UTF-8编码
        char * const envp[]={(char * const )"PYTHONIOENCODING=utf-8",
                    (char * const )"LANG=zh_CN.UTF-8",
                    (char * const )"LANGUAGE=zh_CN.UTF-8",
                    (char * const )"LC_ALL=zh_CN.utf-8",NULL};
        // 输入输出重定向
        close(STDOUT_FILENO);
        close(STDIN_FILENO);
        int rfd = open(runFile,O_RDONLY);
        int wfd = open(resFile,O_RDWR|O_CREAT,0777);
        dup2(rfd,STDIN_FILENO);
        dup2(wfd,STDOUT_FILENO);
        // 当发生系统调用的情况下,父进程可以跟踪子进程
        ptrace(PTRACE_TRACEME, 0, NULL, NULL);
        // 限制 运行时间为
        itimerval time;
        time.it_interval.tv_usec =0;
        time.it_interval.tv_sec = 0;
        time.it_value.tv_sec = this->solve->LimitTime();  
        time.it_value.tv_usec = 1000;
        setitimer(ITIMER_REAL,&time,NULL);
        char path[128]={0};
        switch(lang){
        case CPP:
        case CPP11:
        case CPP17:
        case C:
            sprintf(path,"%s/main",dir);
            execle(path,"main",NULL,envp);
            break;
        case PYTHON3:    // python暂时还未完全  支持
			execle("/usr/bin/python3", "/usr/bin/python3", "main.py",NULL,envp);
        }
        exit(-1);
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
    if(result !=OJ_AC)
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
            result = OJ_PE;
        }
        else
        {
            result = OJ_WA;
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
    status = OJ_JUDGE;
    sprintf(dir,COMPDIR,solve->getjudgeID());
    mkdir(dir,0777);
    while(true)
    {
        switch (this->Jstat)
        {
            case J_CHECK:{
                ILOG("J_CHECK");
                if(checkSource()){
                    Jstat = J_GETFILE;
                }
                else
                {
                    solve->Sres(OJ_RE);
                    Jstat = J_FAILED;
                }
                break;
            }
            
            case J_GETFILE:{
                ILOG("J_GETFILE");
                getFiles();
                Jstat = J_COMPILE;
                break;
            }
            
            case J_COMPILE:{
                ILOG("J_COMPILE");
                if(compile()){
                    Jstat = J_RUNNING;
                }
                else{
                    Jstat = J_CE;
                    solve->Sres(OJ_CE);
                }
                break;
            }
            
            case J_RUNNING:{
                ILOG("J_RUNNING");
                SubRes res  = OJ_AC;
                char resoutfile[128];
                sprintf(resoutfile,"%s/ans",dir);
                init_syscalls_limits(this->solve->Lang());
                for(int i = 0;i<inputFiles.size();i++){
                    DLOG("runnning:%s",inputFiles[i].c_str());
                    running(res,inputFiles[i].c_str(),resoutfile);
                    DLOG("runned:%s",outputFiles[i].c_str());
                    if(res != OJ_AC)
                        break;
                    cmpFIle(res,resoutfile,outputFiles[i].c_str());
                    if(res != OJ_AC)
                        break;
                }
                if(res != OJ_AC){
                    Jstat = J_FAILED;
                    solve->Sres(res);
                }
                else{
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
            
            case J_SUCESS:{
                DLOG("%d:AC",solve->Sid());
                solve->Sres(OJ_AC);
                char del[1024] = "";
                sprintf(del,"rm -rf %s",dir);
                system(del);
                return true;
            }
            
            case J_FAILED:{
                DLOG("%d:%s",solve->Sid(),runningres[solve->Sres()]);
                char del[1024] = "";
                sprintf(del,"rm -rf %s",dir);
                system(del);
                return false;
            }
            
            default:{
                char del[1024] = "";
                sprintf(del,"rm -rf %s",dir);
                system(del);
                return false;
            }
        }
    }

}

judgeClient::judgeClient(Solve *solve){
    Jstat = J_CHECK;
    this->solve = solve;
}