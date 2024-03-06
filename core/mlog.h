#ifndef MLOG_H__
#define MLOG_H__
#include<string>
#include<queue>
#include<stdarg.h>
#include <type_traits>
#include<unistd.h>
#include"Tpool/threadpool.h"
#define FILESIZE 4 * 1024 * 1024
#define LOG_DEBUG 1
using std::string;
namespace my
{
    enum{INFO = 0,DEBUG,WARMMING,ERROR};
    struct logBlock{
        string line;
        string file;
        string function;
        string log_info;
        int level;
    };
    
    void createlogBlock(int tags,string fmt,int line,const char* file,const char* func,...);    
    #define ILOG(fmt,...) createlogBlock(INFO,fmt,__LINE__,__FILE__,__FUNCTION__,##__VA_ARGS__) 
    #define DLOG(fmt,...) createlogBlock(DEBUG,fmt,__LINE__,__FILE__,__FUNCTION__,##__VA_ARGS__) 
    #define WLOG(fmt,...) createlogBlock(WARMMING,fmt,__LINE__,__FILE__,__FUNCTION__,##__VA_ARGS__) 
    #define ELOG(fmt,...) createlogBlock(ERROR,fmt,__LINE__,__FILE__,__FUNCTION__,##__VA_ARGS__) 

    //日志文件
    struct logfile{
        string path;
        string filename;
        string todate;
        string buff;
        size_t offset;
        size_t filesize;
        size_t filecnt;
        int fd;

        int write_line(string str);
        void buff_fflush();
        string getDateString();
        void openfile();
        ~logfile();
    };

    class mlog
    {
    private:
        static mlog* log;
        static locker log_locker;
        queue<logBlock> blockqueue;
        locker block_lock;
        cond   block_cond;
        logfile* file;
        int filesize;
        bool islive;
        pthread_t tid,fftid;
    private:
        mlog(string path,int filesize);
        
        string gettimeString();
        logBlock get_block_log();
        static void* write_log_thread(void* args);
        static void* fflush_file(void* args);
    public:
        mlog() = delete;
        mlog(const mlog&) = delete;
        static mlog* init(string path,int filesize = FILESIZE);
        static mlog* getInstance();
        static void destory();
        bool live();
        bool commit(logBlock loginfo);
        ~mlog();
    };
    template<typename F, typename... Args>
    void retry(bool flag,std::string errMsg,F&& f, Args&&... args)
    {   
        using RetType = decltype(std::forward<F>(f)(std::forward<Args>(args)...));
        static_assert(std::is_same<RetType, bool>::value, "Return type of f must be bool");

        while (true) {
            bool ret = f(args...);
            if (ret!=flag) {
                ELOG("%s",errMsg);
                sleep(5);
                continue;
            }
            break;
        }
    }
    template<typename F, typename... Args>
    void retry(const std::string& errMsg, F&& f, Args&&... args)
    {
        retry(true, errMsg, std::forward<F>(f), std::forward<Args>(args)...);
    }
}
#endif