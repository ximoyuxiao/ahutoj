#ifndef SOLVE_H__
#define SOLVE_H__
#include<vector>
#include<mysql/mysql.h>
#include<string>
#include "result.h"
using std::string;
class Solve
{
private:
    int     problemID;      //问题编号
    int     solutionID;     //提交编号
    int     UserID;         //用户ID
    int     CompleteID;     //竞赛ID 
    string  source;         //代码
    int     limitTime;      //极限运行时间
    int     limitMemory;    //极限运行内存
    lanuage lang;           //使用语言
    SubRes  res;            //运行结果
    int     judgeID;        //JudgeID
    int     usetime;        //运行时间
    int     usememory;      //运行内存
    string  ceinfo;         //错误信息
    long    submitTime;     //提交时间
   
    void intTostr(char* args,int num);
public:
    Solve(int problemID=0,int solutionID=0,const char *source="",int limitTime=0,int limitMeory=0,lanuage lang=C);
    Solve(Solve &solve);
    ~Solve();
    bool operator<(const Solve& b) const;
    Solve& operator=(const Solve &s);
    int  Pid();
    void Pid(int pid);
    int  Sid();
    void Sid(int sid);
    int  Uid();
    void Uid(int uid);
    int Cid();
    void Cid(int cid);
    long SubmitTime();
    void SubmitTime(long submitTime);
    string Source();
    void Source(string code);
    int LimitTime();
    void LimitTime(int limit);
    int  LimitMemory();
    void LimitMemory(int limit);
    lanuage Lang();
    void Lang(lanuage lang);
    SubRes Sres();
    void   Sres(SubRes subres);
    int getjudgeID();
    void setjudgeID(int judgeId);
    int getUsetime();
    void setUsetime(int usetime);
    int getuseMemory();
    void setUseMemory(int useMemory);
    string ceInfo();
    void ceInfo(string ceinfo);
    void getargs(char *args[8]);
};
#endif