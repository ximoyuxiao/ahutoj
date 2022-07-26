#include"threadpool.h"
#include"worker.h"
#include"unistd.h"
#include<iostream>
using namespace std;
class test:public worker{
public:
    int i;
    static locker all_lock;
    test(int i):i(i){}
    void run()
    {
        all_lock.lock();
        threadpool* tp =  threadpool::getPool();
        cout<<"this thread:"<<pthread_self()<<"\tworker id:"<<i<<"\n";
        cout<<"thread size:"<<tp->getThreadNum()<<"\tqueue len:"<<tp->getListlen()<<endl<<endl;
        all_lock.unlock();
    }
};
locker test::all_lock = locker();
int main()
{
    threadpool* tp =  threadpool::getPool(5,10,(size_t)10,20);
    for(int i = 0;i<30;i++)
    {
        tp->excute(shared_ptr<worker>(new test(i)));
    }
    sleep(20);
    for(int i = 0;i<5;i++)
    {
        tp->excute(shared_ptr<worker>(new test(i +  30)));
        usleep(100);
    }
    for(int i = 0;i<1e9;i++);
    tp->destory();
    pthread_exit(NULL);
}