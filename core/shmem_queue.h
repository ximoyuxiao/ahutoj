#ifndef SHMEM_QUEUE_H__
#define SHMEM_QUEUE_H__
#include<sys/ipc.h>
#include<sys/shm.h>
#include<pthread.h>
template<class ValueType>
class ShmemQueue{
private:
    key_t key;
    int shmid;
    
    ValueType* addr;
    int begin;
    int end;
    size_t qsize;
    size_t Maxsize;
    pthread_spinlock_t spinlock;

    void nextVal(int &now);
public:
    ShmemQueue(const char* __pathname,size_t QueueSize = 20);
    ShmemQueue(key_t key,size_t QueueSize = 20);
    int create_queue();
    bool get_queue();
    ValueType*  top();
    bool        push(const ValueType &type);
    ValueType*  pop();
    size_t size();
    bool empty();
};
#endif