#include"shmem_queue.h"
template<class ValueType>
ShmemQueue<ValueType>::ShmemQueue(const char* __pathname,size_t QueueSize = 20)
{
    key = ftok(__pathname,0x2560);
    this->Maxsize = QueueSize;
    addr = nullptr;
    begin = end = 0;
    qsize = 0;
    shmid = -1;
    pthread_spin_init(&spinlock,PTHREAD_PROCESS_SHARED);
}
template<class ValueType>
ShmemQueue<ValueType>::ShmemQueue(key_t key,size_t Queuesize = 20)
{
    this->key = key;
    this->Maxsize = Queuesize;
    addr = nullptr;
    begin = end = 0;
    qsize = 0;
    shmid = -1;
    pthread_spin_init(&spinlock,PTHREAD_PROCESS_SHARED);
}
template<class ValueType>
int ShmemQueue<ValueType>::create_queue(){
    shmid = shmget(key,sizeof(ValueType)*Maxsize,IPC_EXCL|IPC_CREAT);
    this->addr = shmat(shmid,NULL,0);
    return shmid;
}
template<class ValueType>
bool ShmemQueue<ValueType>::get_queue(){
    shmid = shmget(key,sizeof(ValueType)*Maxsize,IPC_CREAT);
    this->addr = shmat(shmid,NULL,0);
    return shmid;
}
template<class ValueType>
ValueType*  ShmemQueue<ValueType>::top(){
    return addr[begin];
}
template<class ValueType>
bool ShmemQueue<ValueType>::push(const ValueType &type){
    pthread_spin_lock(&spinlock);
    if(size() == Maxsize)
    {
        pthread_spin_unlock(&spinlock);
        return false;
    }
    addr[end] = type;
    nextVal(end);
    qsize++;
    pthread_spin_unlock(&spinlock);
    return true;
}
template<class ValueType>
ValueType*  ShmemQueue<ValueType>::pop(){
    pthread_spin_lock(&spinlock);
    int now = begin;
    nextVal(begin);
    qsize--;
    pthread_spin_unlock(&spinlock);
    return addr[now];
}
template<class ValueType>
size_t ShmemQueue<ValueType>::size(){
    return qsize;
}
template<class ValueType>
bool ShmemQueue<ValueType>::empty(){
    return qsize == 0;
}
template<class ValueType>
void ShmemQueue<ValueType>::nextVal(int &now){
    now = (now + 1)%Maxsize;
}