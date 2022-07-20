#include"redis.h"
Redis::Redis():cont(nullptr),reply(nullptr){}
Redis::~Redis()
{            
}

bool Redis::connect(std::string host, int port)
{
    this->cont = redisConnect(host.c_str(), port);
    if(this->cont != nullptr && this->cont->err)
        return false;
    return true;
}

std::string Redis::getString(std::string key)
{
    this->reply = (redisReply*)redisCommand(this->cont, "GET %s", key.c_str());
    if(reply->str == nullptr)
        return "";
    std::string str = this->reply->str;
    freeReplyObject(this->reply);
    reply = nullptr;
    return str;
}

void Redis::setString(std::string key, std::string value)
{
    redisCommand(this->cont, "SET %s %s", key.c_str(), value.c_str());
}
void Redis::setExpire(std::string key,long long second)
{
    redisCommand(this->cont,"expire %s %lld",key.c_str(),second);
}
//列表
void Redis::lpush(std::string key,std::string value)
{
    redisCommand(this->cont,"lpush %s %s",key.c_str(),value.c_str());
}
std::string Redis::rpop(std::string key)
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
void Redis::setbit(string key,int offset,int value)
{
    redisCommand(this->cont,"setbit %s %d %d",key.c_str(),offset,value);
}
bool Redis::getbit(string key,int offset)
{
    this->reply = (redisReply*)redisCommand(this->cont,"getbit %s %d",key.c_str(),offset);
    long long res = reply->integer;
    freeReplyObject(this->reply);
    reply = nullptr;
    return res;
}
long long Redis::bitcount(std::string key)
{
    this->reply = (redisReply*)redisCommand(this->cont,"bitcount %s",key.c_str());
    long long res = reply->integer;
    freeReplyObject(this->reply);
    reply = nullptr;
    return res;
}
//事物相关
bool Redis::addmulti()
{
    redisCommand(this->cont,"multi");
    return true;
}
bool Redis::exec()
{
    redisCommand(this->cont,"exec");
    return true;
}
bool Redis::discard()
{
    redisCommand(this->cont,"discard");
    return true;
}
bool Redis::close()
{
    if(cont != nullptr)
        redisFree(cont);
    cont = nullptr;
    if(reply != nullptr)
        freeReplyObject(reply);
    this->reply = nullptr;          
}