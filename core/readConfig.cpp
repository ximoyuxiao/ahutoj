#include<iostream>
#include<cstring>
#include"readConfig.h"
using namespace std;

readConfig::readConfig(){
    configdir = NULL;
    conferr = NULL;
}
readConfig::readConfig(const char* configdir)
{
    setConfigdir(configdir);
    conferr = NULL;
}

readConfig::~readConfig()
{
    closeConfig();
}

int readConfig::trim(char* res,const char* source)
{
    int j = 0;
    for(int i = 0;source[i]!='\0';i++)
    {
        if(source[i]!=' ')
            res[j++] = source[i];
    }
    
    res[j] =  '\0';
    return  j;
}

bool readConfig::ishead(const char *str)
{
    int len = strlen(str);
    return (str[0] == '[' && str[len-1]==']');
}

int readConfig::readline(char* res,int MAXBUFF,FILE *cnf){
    
    if(fgets(res,MAXBUFF,cnf))
    {
        int len = strlen(res);
        if(res[len-1] == '\n')
            res[len-1] = '\0';
        if (len >= 2 && res[len -2] == '\r'){
            res[len-2] = '\0';
        }
        return len - 1;
    }
    return -1;
}

bool readConfig::config_init()
{
    const int MAXBUFF = 1024;
    char buff[1024]={0};
    if(configdir ==NULL)
    {
        this->conferr = "配置文件未选择";
        return false;
    }
    FILE* cnf  = fopen(this->configdir,"r");
    if(cnf ==NULL)
    {
        this->conferr = "没有该配置文件";
        return false;
    }
    int k = -1;
    while(readline(buff,MAXBUFF,cnf)>0)
    {
        char res[1024] = {0};
        trim(res,buff);
        if(ishead(res))
        {
            res[strlen(res)-1] = '\0';
            myconfs.push_back(new Conf(res + 1)); 
            k++;
        }
        else
        {
            char key[1024]={0};
            char value[1024] ={0};
            int t1 = 0,t2 = 0,i = 0;
            for(;res[i]!='=' && res[i] != '\0';i++)
                key[t1++] = res[i];
            i++;
            for(;res[i] != '\0';i++)
                value[t2++] = res[i];
            if(strcmp(key,"") && strcmp(value,"") && k>=0)
            {
                this->config_update(myconfs[k]->head,key,value);
            }
        }
    }
    return true;
}
bool readConfig::config_update(const char* head,const char* key,const char* value)
{
    for(auto &conf:this->myconfs){
        if(conf->isHead(head))
        {
            conf->setItem(key,value);
            return true;
        }
    }
    conferr = "找不到配置头";
    return false;
}

bool readConfig::closeConfig()
{
    conferr = NULL;
    if(configdir)
    {
        delete []configdir;
        configdir =NULL;
    }
    myconfs.clear();
    return true;
}

int readConfig::getCOnfigInt(const char* head,const char* key)
{
    for(auto &conf:this->myconfs){
        if(conf->isHead(head))
        {
            int res = conf->getvaluetoInt(key);
            if(res == NOT_FIND_KEY)
                conferr = "找不到关键字";
            if(res == NOT_NUMBER)
                conferr = "该配置项非整数";
            return res;
        }
    }
    conferr = "找不到配置项";
    return NOT_FIND_HEAD;
}

int readConfig::getCOnfigString(char* res,const char* head,const char* key)
{
    for(auto conf:this->myconfs){
        if(conf->isHead(head))
        {
            int len = conf->getValue(res,key);
            if(len == NOT_NUMBER)
                conferr = "找不到关键字";
            return len;
        }
    }
    conferr ="找不到配置项";
    return NOT_FIND_HEAD;
}

bool readConfig::getConfgdir(char* dir){
    if(this->configdir == NULL)
        return false;
    strcpy(dir,this->configdir);
    return true;
}

void readConfig::setConfigdir(const char* configdir){
    if(this->configdir)
        delete []this->configdir;
    this->configdir = new char[strlen(configdir) + 1];
    strcpy(this->configdir,configdir);
}

const char* readConfig::getConferr(){
    return this->conferr;
}
