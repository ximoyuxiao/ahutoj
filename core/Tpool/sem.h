#ifndef SEM_H__
#define SEM_H__
#include<semaphore.h>
class sem
{
private:
    sem_t m_sem;
public:
    sem();
    ~sem();
    sem(int value);
    bool post();
    bool wait();
};

#endif