#ifndef JUDGECLIENT_H__
#define JUDGECLIENT_H__
#include"result.h"
#include"Solve.h"
#include<string>
#include<vector>
using namespace std;

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
    JSTAT Jstat;
private:
    bool checkSource();
    bool compile();
    long long getFileSize(const char * filepath);
    bool running(SubRes &result,const char * runFile,const char *resFile);
    bool getFiles();
    bool judgePE(FILE*source,FILE *res);
    bool cmpFIle(SubRes &result,char *myfile,const char* sourceFile);
public:
    bool judge();
    Solve* GetSolve();
    void SetSolve(Solve* solve);
    judgeClient(Solve *solve);
};


#endif