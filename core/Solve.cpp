#include<iostream>
#include<vector>


#include"Solve.h"
#include"readConfig.h"
Solve::Solve(string problemID,int solutionID,const char *source,int limitTime,int limitMeory,lanuage lang, int spj){
    this->problemID = problemID;
    this->solutionID = solutionID;
    this->source = source;
    this->limitTime = limitTime;
    this->limitMemory = limitMeory;
    this->lang = lang;
    this->res = OJ_JUDGE;
    this->SpjJudge = spj;
    PassSample = 0;     //样例通过数
    SampleNumber = 0;
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
    this->PassSample = solve.PassSample;
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
    this->PassSample = solve.PassSample;
    return *this;
}
void Solve::getargs(char* args[8])
{
//Solve(int problemID, int solutionID, const char *source, int limitTime, int limitMeory, int lang)
    strcpy(args[0],"judgeclient"); 
    strcpy(args[1],problemID.c_str());
    intTostr(args[2],solutionID);
    strcpy(args[3],source.c_str());
    intTostr(args[4],limitTime);
    intTostr(args[5],limitMemory);
    intTostr(args[6],lang);
    intTostr(args[7],getjudgeID());
}
void longlongTostr(char* args,long long num){
    sprintf(args,"%lld",num);
    return ;
}
void Solve::intTostr(char*args,int num)
{
    sprintf(args,"%d",num);
    return ;
}

string  Solve::Pid(){
    return problemID;
}
void Solve::Pid(string pid){
    this->problemID = pid;
}
int  Solve::Sid(){
    return solutionID;
}
void Solve::Sid(int sid){
    this->solutionID = sid;
}
string  Solve::Uid(){
    return this->UserID;
}
void Solve::Uid(string uid){
    this->UserID = uid;
}
int Solve::Cid(){
    return this->CompleteID;
}
void Solve::Cid(int cid){
    this->CompleteID = cid;
}
string Solve::Source(){
    return source;
}
void Solve::Source(string code){
    this->source = code;
}
long long Solve::LimitTime(){
    return limitTime;
}
void Solve::LimitTime(long long limit){
    this->limitTime = limit;
}
long long  Solve::LimitMemory(){
    return limitMemory;
}
void Solve::LimitMemory(long long limit){
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
long long Solve::getUsetime(){
    return this->usetime;
}
void Solve::setUsetime(long long usetime){
    this->usetime = usetime;
}
long long Solve::getuseMemory()
{
    return this->usememory;
}
void Solve::setUseMemory(long long useMemory){
    this->usememory = useMemory;
}
string Solve::ceInfo(){
    return ceinfo;
}
void Solve::ceInfo(string ceinfo)
{
    this->ceinfo = ceinfo;
}
int Solve::getSpjJudge(){
    return this->SpjJudge;
}
void Solve::setSpjJudge(int spjJudge){
    this->SpjJudge = spjJudge;
}
long long Solve::getPassSample(){
    return this->PassSample;
}
void Solve::incPassSample(){
    this->PassSample++;
}
void Solve::setSampleNumber(size_t sampleNumber) {
    this->PassSample=sampleNumber;
}
long long Solve::getSampleNumber() {
    return this->SampleNumber;
}
void Solve::to_json(nlohmann::json& j){
    j = nlohmann::json
    {
        {"SID", Sid()}, 
        {"PID", Pid()},
        {"UID",Uid()},
        {"CID",Cid()},
        {"JudgeID",getjudgeID()},
        {"Source",Source()},
        {"Lang",Lang()},
        {"ResultACM",runningres[Sres()]},
        {"PassSample",getPassSample()},
        {"Sim",Sim},
        {"UseTime",usetime},
        {"UseMemory",usememory},
    };
}

void Solve::to_ceJson(nlohmann::json &j){
    j = nlohmann::json
    {
        {"SID", Sid()}, 
        {"info",ceinfo},
    };
}

void Solve::from_json(nlohmann::json& j){
    j.at("SID").get_to(solutionID);
    j.at("PID").get_to(problemID);
    j.at("UID").get_to(UserID);
    j.at("CID").get_to(CompleteID);
    j.at("JudgeID").get_to(JudgeID);
    j.at("Source").get_to(source);
    j.at("Lang").get_to(lang);
    Sres(OJ_JUDGE);
    PassSample = 0;
    Sim = 0;
}