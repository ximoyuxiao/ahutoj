#ifndef __MYDB__H__
#define __MYDB__H__
#include<mysql/mysql.h>
#include<string>
using std::string;
class mysqlDB{
private:
    static mysqlDB* mydb;
    string host;
    string user;
    string pass;
    string db;
    int port;
    mysqlDB();
    mysqlDB(const char* host,const char* user,const char* pass,const char* db,int port = 3306);
public:
    static void initConn(const char* host,const char* user,const char* pass,const char* db,int port = 3306);
    static mysqlDB* getInstance();
    int getDatabase(MYSQL* mysql);
    void CloseDatabase(MYSQL* mysql,MYSQL_RES* res);
};

#endif