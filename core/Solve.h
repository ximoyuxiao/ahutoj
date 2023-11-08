#ifndef SOLVE_H__
#define SOLVE_H__
#include<vector>
#include<mysql/mysql.h>
#include<string>
#include "result.h"
#include <nlohmann/json.hpp>
using std::string;
class Solve
{
private:
    string        problemID;      //问题编号
    int           solutionID;     //提交编号
    string        UserID;         //用户ID
    int           CompleteID;     //竞赛ID 
    int           JudgeID;        //判题机ID
    string        source;         //代码
    long long     PassSample;     //样例通过数<->样例WA on test
    long long     SampleNumber;    //样例总数
    int           Sim;            //相似度检测结果
    long long     limitTime;      //极限运行时间
    long long     limitMemory;    //极限运行内存
    lanuage       lang;           //使用语言
    SubRes        res;            //运行结果
    int           judgeID;        //JudgeID
    long long     usetime;        //运行时间
    long long     usememory;      //运行内存
    string        ceinfo;         //错误信息
    int           SpjJudge;      //是否开启特判(1:on/-1:off)
    void intTostr(char* args,int num);
public:
    void to_json(nlohmann::json& j);
    void to_ceJson(nlohmann::json &j);
    void from_json(nlohmann::json& j);
    Solve(string problemID="",int solutionID=0,const char *source="",int limitTime=0,int limitMeory=0,lanuage lang=C, int spj = -1);
    Solve(Solve &solve);
    ~Solve();
    bool operator<(const Solve& b) const;
    Solve& operator=(const Solve &s);
    string  Pid();
    void Pid(string pid);
    int  Sid();
    void Sid(int sid);
    string  Uid();
    void Uid(string uid);
    int Cid();
    void Cid(int cid);
    long SubmitTime();
    void SubmitTime(long submitTime);
    string Source();
    void Source(string code);
    long long LimitTime();
    void LimitTime(long long limit);
    long long  LimitMemory();
    void LimitMemory(long long limit);
    lanuage Lang();
    void Lang(lanuage lang);
    SubRes Sres();
    void   Sres(SubRes subres);
    int getjudgeID();
    void setjudgeID(int judgeId);
    long long getUsetime();
    void setUsetime(long long usetime);
    long long getuseMemory();
    void setUseMemory(long long useMemory);
    string ceInfo();
    void ceInfo(string ceinfo);
    void getargs(char *args[8]);
    int getSpjJudge();
    void setSpjJudge(int spjJudge);
    long long getPassSample();
    void incPassSample();
    long long getSampleNumber();
    void setSampleNumber();
};

#endif