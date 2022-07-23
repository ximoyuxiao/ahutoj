#include<iostream>
#include<vector>


#include"Solve.h"
#include"readConfig.h"
Solve::Solve(int problemID,int solutionID,const char *source,int limitTime,int limitMeory,lanuage lang){
    this->problemID = problemID;
    this->solutionID = solutionID;
    this->source = source;
    this->limitTime = limitTime;
    this->limitMemory = limitMeory;
    this->lang = lang;
    this->res = OJ_JUDGE;
    judgeID = 0;        //JudgeID
    usetime = 0;        //运行时间
    usememory = 0;      //运行内存
    ceinfo ="";
}
Solve::Solve(Solve &solve)
{
    this->problemID  = solve.problemID;
    this->solutionID = solve.solutionID;
    this->limitTime = solve.limitTime;
    this->limitMemory = solve.limitMemory;
    this->source = solve.source;
    this->lang = solve.lang;
    this->judgeID = solve.judgeID;
    this->usememory = solve.usememory;
    this->usetime = solve.usetime;
    this->ceinfo = solve.ceinfo;
    this->res  =  res;
}
Solve::~Solve()
{
}
bool Solve::operator<(const Solve& b) const
{
    return this->solutionID < b.solutionID;    
}
Solve& Solve::operator=(const Solve &solve)
{
     this->problemID  = solve.problemID;
    this->solutionID = solve.solutionID;
    this->limitTime = solve.limitTime;
    this->limitMemory = solve.limitMemory;
    this->source = solve.source;
    this->lang = solve.lang;
    this->judgeID = solve.judgeID;
    this->usememory = solve.usememory;
    this->usetime = solve.usetime;
    this->ceinfo = solve.ceinfo;
    this->res  =  res;
}
void Solve::getargs(char* args[8])
{
//Solve(int problemID, int solutionID, const char *source, int limitTime, int limitMeory, int lang)
    strcpy(args[0],"judgeclient"); 
    intTostr(args[1],problemID);
    intTostr(args[2],solutionID);
    strcpy(args[3],source.c_str());
    intTostr(args[4],limitTime);
    intTostr(args[5],limitMemory);
    intTostr(args[6],lang);
    intTostr(args[7],getjudgeID());
}
void Solve::intTostr(char*args,int num)
{
    sprintf(args,"%d",num);
    return ;
}

int  Solve::Pid(){
    return problemID;
}
void Solve::Pid(int pid){
    this->problemID = pid;
}
int  Solve::Sid(){
    return solutionID;
}
void Solve::Sid(int sid){
    this->solutionID = sid;
}
int  Solve::Uid(){
    return this->UserID;
}
void Solve::Uid(int uid){
    this->UserID = uid;
}
int Solve::Cid(){
    return this->CompleteID;
}
void Solve::Cid(int cid){
    this->CompleteID = cid;
}
long Solve::SubmitTime(){
    return submitTime;
}
void Solve::SubmitTime(long submitTime){
    this->submitTime = submitTime;
}
string Solve::Source(){
    return source;
}
void Solve::Source(string code){
    this->source = code;
}
int Solve::LimitTime(){
    return limitTime;
}
void Solve::LimitTime(int limit){
    this->limitTime = limit;
}
int  Solve::LimitMemory(){
    return limitMemory;
}
void Solve::LimitMemory(int limit){
    this->limitMemory = limit;
}
lanuage Solve::Lang(){
    return lang;
}
void Solve::Lang(lanuage lang){
    this->lang = lang;
}
SubRes Solve::Sres(){
    return res;
}
void  Solve::Sres(SubRes subres){
    this->res = subres;
}
int Solve::getjudgeID(){
    return this->judgeID;
}
void Solve::setjudgeID(int judgeId)
{
    this->judgeID = judgeId;
}
int Solve::getUsetime(){
    return usetime;
}
void Solve::setUsetime(int usetime){
    this->usememory = usetime;
}
int Solve::getuseMemory()
{
    return this->usememory;
}
void Solve::setUseMemory(int useMemory){
    this->usememory = useMemory;
}
string Solve::ceInfo(){
    return ceinfo;
}
void Solve::ceInfo(string ceinfo)
{
    this->ceinfo = ceinfo;
}