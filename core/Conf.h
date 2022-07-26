#ifndef CONF_H__
#define CONF_H__
#include<map>
#define NOT_NUMBER      -1000
#define NOT_FIND_KEY    -1001
#define NOT_FIND_HEAD   -1002
using namespace std;
struct cmp{
    bool operator()(const char* a,const char*b) const;    
};
struct Conf{
    char* head;
    map<char*,char*,cmp> config_dic;
    Conf(const char* head);
    void setItem(const char* key,const char* value);
    int getValue(char* res,const char* key);
    int getvaluetoInt(const char* key);
    bool isHead(const char* SourceHead);
    ~Conf();
};
#endif