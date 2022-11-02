#include"mydb.h"
mysqlDB* mysqlDB::mydb = nullptr;

void mysqlDB::initConn(const char* host,const char* user,const char* pass,const char* db,int port){
    mydb = new mysqlDB(host,user,pass,db,port);
    return ;
}

mysqlDB::mysqlDB(){
    host = nullptr;
    user = nullptr;
    pass = nullptr;
    db = nullptr;
    port = 0;
}

mysqlDB::mysqlDB(const char* host,const char* user,const char* pass,const char* db,int port)
:host(host),user(user),pass(pass),db(db),port(port)
{
    for(int i = 0;i<MAX_CONN;i++){
        mysqls[i].status = SQL_UNINIT;
    }
}

mysqlDB* mysqlDB::getInstance(){
    return mydb;
}

int mysqlDB::getDatabase(MYSQL* mysql)
{
    mysql_init(mysql);
    if(mysql_real_connect(mysql,host.c_str(),user.c_str(),pass.c_str(),db.c_str(),port,NULL,0))
    {
        mysql_query(mysql,"set names utf8mb4");
        return 0;
    }
    return mysql_errno(mysql);
}

void mysqlDB::CloseDatabase(MYSQL* mysql,MYSQL_RES* res)
{
    if(res != nullptr)
        mysql_free_result(res);
    if(mysql != nullptr)
        mysql_close(mysql);
}

