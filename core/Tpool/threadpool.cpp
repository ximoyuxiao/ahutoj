#include"threadpool.h"
#include<exception>
#include<iostream>
using std::cout;
using std::endl;
typedef struct threadInfo
{
    bool iscore;
    long currentTime;
    long liveTime;
    shared_ptr<worker> work;
    bool islive;
    long lastTime;
    pthread_t tid;
    int  pos;
    static void* start(void* args);
    threadInfo(int iscore,long currentTime,long liveTime,shared_ptr<worker> work):
    iscore(iscore),currentTime(currentTime),liveTime(liveTime),work(work),islive(true)
    {}
}threadInfo_st;
threadpool* threadpool::mythredpool = NULL;

threadpool::threadpool(int core,int max_core,long livetime,size_t listlen,void (*rejectFunc)(worker *worker)):
core_thread(core),max_thread(max_core),thread_num(0),liveTime(livetime),listlen(listlen),islive(true),rejectFunc(rejectFunc)
{
    if(core<= 0 || max_core<=0)
        throw std::exception();
    threads.resize(max_core,NULL);
}
threadpool::~threadpool()
{
    for(threadInfo_ptr temp:threads)
    {
        threadInfo* info = (threadInfo*)temp; 
        if(info == NULL)    continue;
        info->islive = false;
    }
    int num = thread_num;
    for(int i = 0;i<num;i++) thread_sem.post();
}
int threadpool::getFreePos()
{
    for(int i =  0;i<max_thread;i++)
    {
        if(threads[i] == NULL)
            return i;
    }
    return -1;
}
bool threadpool::createThread(bool isCore,shared_ptr<worker> work)
{
    threadInfo* tinfo = new threadInfo(isCore,time(NULL),this->liveTime,work);
    if(thread_num >= max_thread)    return false;
    this->thread_loker.lock();
    pthread_t tid;
    tinfo->pos = getFreePos();
    pthread_create(&tid,NULL,tinfo->start,static_cast<void*>(tinfo));
    pthread_detach(tid);
    tinfo->tid = tid;
    this->threads[tinfo->pos] = tinfo;
    this->thread_num++;
    this->thread_loker.unlock();
    return true;
}
bool threadpool::delThread(threadInfo_ptr thread)
{
    threadInfo* tinfo = (threadInfo_st*)thread;
    int pos = tinfo->pos;
    this->thread_loker.lock();
    threads[pos] = NULL;
    delete tinfo;
    this->thread_num--;
    this->thread_loker.unlock();
    return true;
}
threadpool* threadpool::getPool(int core,int max_core,long livetime,size_t listlen,void (*rejectFunc)(worker *worker))
{
    if(mythredpool == NULL )
        mythredpool = new threadpool(core,max_core,livetime,listlen,rejectFunc);
    return mythredpool;
}
threadpool* threadpool::getPool()
{
    return mythredpool;
}
bool threadpool::cancleThread(threadInfo_ptr thread)
{
    return delThread(thread);
}
bool threadpool::excute(shared_ptr<worker> work)
{
    // cout<<"workersQueue.size()"<<workersQueue.size()<<endl;
    if(!islive) return false;
    if(thread_num < core_thread)
    {
        createThread(true,work);
    }
    else if(workersQueue.size() < listlen)
    {
        queue_locker.lock();
        workersQueue.push(work);
        queue_locker.unlock();
        thread_sem_post();
    }
    else if(thread_num < max_thread)
    {
        createThread(false,work);
    }
    else
    {
        rejectFunc(work.get());
        return false;
    }
    return true;
}
void threadpool::defaultRejectFunc(worker *worker)
{
    std::cout<<"队列已满，任务执行失败。";
    return ;
}
void* threadInfo::start(void* args)
{
    threadInfo_st* tinfo = (threadInfo_st*)args;
    threadpool* mypool = threadpool::getPool();
    while(tinfo->islive)
    {
        if(tinfo->work.get() !=nullptr)
        {
            tinfo->currentTime = tinfo->lastTime = time(NULL); 
            tinfo->work->run();
            tinfo->work.reset();
        }
        mypool->thread_sem_wait();
        tinfo->currentTime = time(NULL);
        if(!mypool->live() || ( tinfo->currentTime > tinfo->lastTime + tinfo->liveTime && tinfo->iscore == false))
        {
            tinfo->islive = false;
            mypool->thread_sem_post();
            break;
        }   
        tinfo->work = mypool->get_worker();
    }
    mypool->cancleThread(tinfo);
    pthread_exit(NULL);
}
bool threadpool::live()
{
    return islive;
}

bool threadpool::thread_sem_wait()
{
    return thread_sem.wait();
}
bool threadpool::thread_sem_post()
{
    return thread_sem.post();
}
bool threadpool::thread_locker_lock()
{
    return thread_loker.lock();
}
bool threadpool::thread_locker_unlocker()
{
    return thread_loker.unlock();
}
bool threadpool::queue_locker_lock()
{
    return queue_locker.lock();
}
bool threadpool::queue_locker_unlocker()
{
    return queue_locker.unlock();
}
shared_ptr<worker> threadpool::get_worker()
{
    queue_locker.lock();
    shared_ptr<worker> work = workersQueue.front();
    workersQueue.pop();
    queue_locker.unlock();
    return work;
}
int  threadpool::getListlen()
{
    return workersQueue.size();
}
int  threadpool::getThreadNum()
{
    return thread_num;
}

#include<unistd.h>
void threadpool::destory()
{
    mythredpool->islive = false;
    mythredpool->queue_locker.lock();
    while(!mythredpool->workersQueue.empty())
        mythredpool->workersQueue.pop();
    mythredpool->queue_locker.unlock();
    for(threadInfo_ptr temp:mythredpool->threads)
    {
        threadInfo* info = (threadInfo*)temp; 
        if(info == NULL)    continue;
        info->islive = false;
    }
    mythredpool->thread_sem.post();
    while(mythredpool->thread_num);
    delete mythredpool;
    mythredpool = nullptr;
}