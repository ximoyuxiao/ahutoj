#ifndef JUDGECLIENT_H__
#define JUDGECLIENT_H__
#include"result.h"
#include"Solve.h"
#include"Language.h"
#include<string>
#include<vector>
using namespace std;
//const int call_array_size = 512;
enum JSTAT{
    J_CHECK,
    J_GETFILE,
    J_COMPILE,
    J_RUNNING,
    J_CE,
    J_SUCESS,
    J_FAILED
};
class judgeClient
{
private:
    Solve *solve;
    char dir[56];
    vector<string> inputFiles;
    vector<string> outputFiles;
    Language* language;
    JSTAT Jstat;
    unsigned int call_id;
    int call_counter[call_array_size];
private:
    bool checkSource();
    bool compile();
    long long getFileSize(const char * filepath);
    //运行runFile文件，
    bool running(SubRes &result,const char * runFile,const char *resFile,long long &useMemory,long long &useTime);
    bool getFiles();
    bool judgePE(FILE*source,FILE *res);
    bool cmpFIle(SubRes &result,const char *myfile,const char* sourceFile);
    void init_syscalls_limits(lanuage lang);
public:
    bool judge();
    Solve* GetSolve();
    void SetSolve(Solve* solve);
    judgeClient(Solve *solve);
    ~judgeClient();
};


#endif