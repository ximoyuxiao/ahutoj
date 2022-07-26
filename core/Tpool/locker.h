#ifndef  LOCKER_H__
#define LOCKER_H__
#include<pthread.h>
class locker
{
private:
public:
    locker();
    ~locker();
    bool lock();
    bool unlock();
    pthread_mutex_t* getMutex();
private:
    pthread_mutex_t m_mutex;
};


#endif