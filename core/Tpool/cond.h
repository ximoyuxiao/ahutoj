#ifndef COND_H__
#define COND_H__
#include<pthread.h>
class cond
{
private:
    pthread_cond_t m_cond;
public:
    cond();
    ~cond();
    bool wait(pthread_mutex_t *mutex);
    bool timewait(pthread_mutex_t *mutex,struct timespec t);
    bool signal();
    bool broadcast();
};


#endif