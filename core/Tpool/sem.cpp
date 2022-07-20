#include"sem.h"
sem::sem(/* args */)
{
    sem_init(&m_sem,0,0);
}
sem::sem(int val)
{
    sem_init(&m_sem,0,val);
}
sem::~sem()
{
    sem_destroy(&m_sem);
}
bool sem::post()
{
    return sem_post(&m_sem) == 0;
}
bool sem::wait()
{
    return sem_wait(&m_sem) == 0;
}