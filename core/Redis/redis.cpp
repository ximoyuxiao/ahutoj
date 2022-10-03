#include"redis.h"
MyRedis::MyRedis():cont(nullptr),reply(nullptr){}
MyRedis::~MyRedis()
{            
}
#include<iostream>
bool MyRedis::connect(std::string host, int port,const char* password)
{
    this->cont = redisConnect(host.c_str(),port);
    if(password != "")
    redisCommand(this->cont,"auth %s",password);
    if(this->cont != nullptr && this->cont->err)
        return false;
    return true;
}

std::string MyRedis::getString(std::string key)
{
    this->reply = (redisReply*)redisCommand(this->cont, "GET %s", key.c_str());
    if(reply->str == nullptr)
        return "";
    std::string str = this->reply->str;
    freeReplyObject(this->reply);
    reply = nullptr;
    return str;
}

void MyRedis::setString(std::string key, std::string value)
{
    redisCommand(this->cont, "SET %s %s", key.c_str(), value.c_str());
}
void MyRedis::setExpire(std::string key,long long second)
{
    redisCommand(this->cont,"expire %s %lld",key.c_str(),second);
}
//列表
void MyRedis::lpush(std::string key,std::string value)
{
    redisCommand(this->cont,"lpush %s %s",key.c_str(),value.c_str());
}
std::string MyRedis::rpop(std::string key)
{
    this->reply = (redisReply*)redisCommand(this->cont,"rpop %s",key.c_str());
    if(reply->str == nullptr)
        return "";
    std::string str = reply->str;
    freeReplyObject(this->reply);
    reply = nullptr;
    return str;
}
//位图
void MyRedis::setbit(string key,int offset,int value)
{
    redisCommand(this->cont,"setbit %s %d %d",key.c_str(),offset,value);
}
bool MyRedis::getbit(string key,int offset)
{
    this->reply = (redisReply*)redisCommand(this->cont,"getbit %s %d",key.c_str(),offset);
    long long res = reply->integer;
    freeReplyObject(this->reply);
    reply = nullptr;
    return res;
}
long long MyRedis::bitcount(std::string key)
{
    this->reply = (redisReply*)redisCommand(this->cont,"bitcount %s",key.c_str());
    long long res = reply->integer;
    freeReplyObject(this->reply);
    reply = nullptr;
    return res;
}
//事物相关
bool MyRedis::addmulti()
{
    redisCommand(this->cont,"multi");
    return true;
}
bool MyRedis::exec()
{
    redisCommand(this->cont,"exec");
    return true;
}
bool MyRedis::discard()
{
    redisCommand(this->cont,"discard");
    return true;
}
bool MyRedis::close()
{
    if(cont != nullptr)
        redisFree(cont);
    cont = nullptr;
    if(reply != nullptr)
        freeReplyObject(reply);
    this->reply = nullptr;          
    return true;
}