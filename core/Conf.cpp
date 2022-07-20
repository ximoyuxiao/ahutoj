#include<iostream>
#include<cstring>
#include<map>
#include<cstdlib>
#include"Conf.h"
using namespace std;
bool cmp::operator()(const char* a,const char*b)const {
        return (strcmp(a,b)<0);
}
Conf::Conf(const char* head = NULL)
{
    this->head = new char[strlen(head) + 1];
    strcpy(this->head,head);
}

void Conf::setItem(const char* key,const char* value)
{
    int keylen = strlen(key);
    int valuelen = strlen(value);
    map<char*,char*,cmp>::iterator it = config_dic.begin();
    while (it != config_dic.end())
    {
        if(strcmp(key,it->first)==0)
        {
            delete []it->second;
            it->second = new char[valuelen + 1];
            strcpy(it->second,value);
            return ;
        }
        it++;
    }
    char* cpkey = new char[keylen + 1];
    char* cpval = new char[valuelen + 1];
    config_dic.insert(pair<char*,char*>(strcpy(cpkey,key),strcpy(cpval,value)));
}

int Conf::getValue(char* res,const char*key)
{
    map<char*,char*,cmp>::iterator it = config_dic.begin();
    while(it != config_dic.end())
    {

        if(strcmp(key,it->first) == 0)
        {
            strcpy(res,it->second);
            return strlen(res);
        }
        it++;
    }
    return NOT_FIND_KEY;
}
int Conf::getvaluetoInt(const char* key)
{
    char str[1024]={0};
    if(!getValue(str,key))
        return NOT_FIND_KEY;
    int res = 0;
    for(int i = 0;str[i]!='\0';i++)
    {
        if(str[i]>='0' &&str[i]<='9')
            res = res*10 + (str[i]-'0');
        else
            return NOT_NUMBER;
    } 
    return res;
}

bool Conf::isHead(const char* SourceHead){
    return (strcmp(head,SourceHead)==0);
}

Conf::~Conf()
{
    delete []head;
    head = NULL;
    map<char*,char*,cmp>::iterator it = config_dic.begin();
    while(it!=config_dic.end())
    {
        it++;
        map<char*,char*,cmp>::iterator temp = it;
        it =  config_dic.erase(it);    
        delete []temp->first;
        delete []temp->second;
    }
}