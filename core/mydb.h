#ifndef __MYDB__H__
#define __MYDB__H__
#include<mysql/mysql.h>
#include<string>
using std::string;
#define MAX_CONN 5
enum mysql_status_t{
    SQL_UNINIT,
    SQL_FREE,
    SQL_BUSY,
    SQL_DESTORY,
};
struct mysql_item{
    MYSQL mysql;
    mysql_status_t status;
};
class mysqlDB{
private:
    static mysqlDB* mydb;
    mysql_item mysqls[MAX_CONN];
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