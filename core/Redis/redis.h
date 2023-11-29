#ifndef REDIS_H__
#define REDIS_H__
#include<hiredis/hiredis.h>
#include<string>
using std::string;
class MyRedis
{
public:
 
    MyRedis();
    ~MyRedis();
    bool connect(std::string host = "redis", int port=6379,const char* password ="");
    //字符串
    std::string getString(std::string key);
    void setString(std::string key, std::string value);
    //过期
    void setExpire(std::string key,long long second);
    //列表
    void lpush(std::string key,std::string value);
    std::string rpop(std::string key);
    //位图
    void setbit(string key,int offset,int value);
    bool getbit(string key,int offset);
    long long bitcount(std::string key);
    //事物相关
    bool addmulti();
    bool exec();
    bool discard();
    bool close();
private:
    redisContext* cont;
    redisReply* reply;
};
#endif