#include"mlog.h"
#include<cstdlib>
#include<cstring>
#include<string>
#include <unistd.h>
#include<sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include<iostream>
using namespace std;
using namespace my;
#define FFLUSH_TIME 30
mlog* mlog::log = nullptr;
void my::createlogBlock(int tags,string fmt,int line,const char* file,const char* func,...)
{
    char buff[2048]="";
    va_list vs;
    logBlock block;
    block.line = std::to_string(line);
    block.file = file;
    block.function = func;
    block.level = tags;
    va_start(vs,func);
    vsnprintf(buff,2048,fmt.c_str(),vs);
    va_end(vs);
    block.log_info = buff;
    mlog* log = log->getInstance();
    log->commit(block);
   
}
void* mlog::fflush_file(void* args)
{
    mlog* log = static_cast<mlog*>(args);
    while (log->live())
    {
        usleep(10000);
        if(log->file->buff != "")
            log->file->buff_fflush();
    }
    return NULL;
}
void* mlog::write_log_thread(void* args)
{
    mlog* log = static_cast<mlog*>(args);
    string level[] ={"INFO","DEBUG","WARMMING","ERROR"};
    time_t current_time = time(NULL);
    time_t last_time = time(NULL);
    while(log->live() || log->blockqueue.size())
    {
        logBlock str = log->get_block_log();
        string logstr ="[" + level[str.level] +  "]" + log->gettimeString() + "[flie:"+ str.file + "][line:"+ str.line +"][function:" +str.function +"]"+ str.log_info + "\n";
        log->file->write_line(logstr);
        current_time = time(NULL);
        #if LOG_DEBUG
            log->file->buff_fflush();
        #else
            if(current_time-last_time > FFLUSH_TIME || str.level >= WARMMING)
            {
                log->file->buff_fflush();
                last_time = time(NULL);
            }
        #endif
    }
    return NULL;
}

int logfile::write_line(string str)
{
    int len = str.size();
    if(offset + len > filesize)
    {
        this->buff_fflush();
        openfile();
    }
    offset += len;
    buff += str;
    return len;
}

void logfile::openfile()
{
    if(fd != -1)
        close(fd);
    string date = this->getDateString();
    if(date == string(todate))
    {
        this->filecnt++;
        filename =path +  date + + "_" + std::to_string(filecnt) + ".log";
    }
    else
        filename =path +  date + ".log";
    todate = date;
    this->fd = open(filename.c_str(),O_WRONLY|O_CREAT|O_APPEND,0666);
    buff = "";
    this->offset = 0;
}

void logfile::buff_fflush()
{
    // cout<<buff<<endl;
    write(fd,buff.c_str(),buff.size());
    buff = "";
}

logfile::~logfile()
{
    if(fd != -1)
    {
        buff_fflush();
        close(fd);
    }
}

logBlock  mlog::get_block_log()
{
    block_lock.lock();
    while(blockqueue.empty())
        block_cond.wait(block_lock.getMutex());
    logBlock ret = blockqueue.front();
    blockqueue.pop();
    block_lock.unlock();
    return ret;
}

string logfile::getDateString()
{
    char date[128];
    time_t now;
    time(&now);
    tm *phare = localtime(&now);
    sprintf(date,"%d-%02d-%02d",phare->tm_year + 1900,phare->tm_mon + 1,phare->tm_mday);
    return date;
}

string mlog::gettimeString()
{
    char ti[128];
    time_t now;
    time(&now);
    tm *phare = localtime(&now);
    sprintf(ti,"[%02d:%02d:%02d]",phare->tm_hour,phare->tm_min,phare->tm_sec);
    return string(ti);
}

//线程不安全 默认只会被调用一次
mlog* mlog::init(string path,int filesize)
{
    mkdir(path.c_str(),0777);
    if(log == nullptr)
        log = new mlog(path,filesize);
    return log;
}

mlog::mlog(string path,int filesize)
{
    if(path.back() != '/' && path.back() !='\\')
        path += '/';
    this->islive = true;
    this->filesize = filesize;
    this->file =new logfile();
    this->file->buff.resize(filesize);
    file->path = path;
    file->offset =  0;
    file->todate = "";
    file->filename = "";
    file->filesize = filesize;
    this->file->fd = -1;
    file->openfile();
    pthread_create(&tid,NULL,write_log_thread,(void*)this);
    pthread_create(&fftid,NULL,fflush_file,(void*)this);
}

mlog* mlog::getInstance()
{
    return log;
}

void mlog::destory()
{
    ILOG("destroy");
    log->islive = false;
    log->block_cond.broadcast();
    pthread_join(log->tid,NULL);
    pthread_join(log->fftid,NULL);
    delete log;
    log = nullptr;
}

bool mlog::live()
{
    return islive;
}

bool mlog::commit(logBlock loginfo)
{
    if(live())
    {
        block_lock.lock();
        blockqueue.push(loginfo);
        block_lock.unlock();
        block_cond.broadcast();
        return true;
    }
    return false;
}

mlog::~mlog()
{ 
    delete file;
}
