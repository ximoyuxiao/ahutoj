#ifndef READCONFIG_H__
#define READCONFIG_H__
#include<map>
#include<cstring>
#include<vector>

#include"Conf.h"
using namespace std;

class readConfig
{
private:
    char* configdir;
    vector<Conf*> myconfs;
    const char* conferr; 
private:
    int trim(char *,const char*);
    bool ishead(const char *str);
    int readline(char* res,int MAXBUFF,FILE *cnf);
public:
    readConfig();
    readConfig(const char* configdr);

    bool config_init();
    bool config_update(const char* head,const char* key,const char* value);
    
    bool closeConfig();

    int getCOnfigInt(const char* head,const char* key);
    int getCOnfigString(char* res,const char* head,const char* key);

    bool getConfgdir(char* dir);
    void setConfigdir(const char*);
    const char* getConferr();
    ~readConfig();
};
#endif