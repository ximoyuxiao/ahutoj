#ifndef THREAD_POOL_H__
#define THREAD_POOL_H__
#include<pthread.h>
#include<vector>
#include<queue>
#include<memory>

#include"cond.h"
#include"locker.h"
#include"sem.h"
#include"worker.h"

using std::vector;
using std::queue;
using std::shared_ptr;
typedef void* threadInfo_ptr;
class threadpool
{
private:
    static threadpool* mythredpool;
    
    int core_thread;                            // 核心线程
    int max_thread;                             //最大线程数
    int thread_num;                             //线程总数
    int liveTime;                               //存活时间
    size_t listlen;
    bool islive;
    vector<threadInfo_ptr> threads;                  //线程数组
    locker thread_loker;
    sem    thread_sem;
    queue<std::shared_ptr<worker> > workersQueue;                //任务队列
    locker queue_locker;

    void (*rejectFunc)(worker *worker);         //拒绝策略 函数指针
private:
    threadpool(int core,int max_core,long livetime,size_t listlen,void (*rejectFunc)(worker *worker) = defaultRejectFunc);
    threadpool(threadpool&)=delete;
    static void defaultRejectFunc(worker *worker);
    bool createThread(bool isCore,shared_ptr<worker> work);
    bool delThread(threadInfo_ptr thread);
    int getFreePos();
    
public:
    bool excute(shared_ptr<worker>);
    static threadpool* getPool(int core,int max_core,long livetime,size_t listlen,void (*rejectFunc)(worker *worker) = defaultRejectFunc);
    static threadpool* getPool();
    bool cancleThread(threadInfo_ptr thread);

    bool thread_sem_wait();
    bool thread_sem_post();
    bool thread_locker_lock();
    bool thread_locker_unlocker();
    bool queue_locker_lock();
    bool queue_locker_unlocker();
    shared_ptr<worker> get_worker();
    int getListlen();
    int getThreadNum();
    static void destory();
    bool live();
    ~threadpool();
};

#endif