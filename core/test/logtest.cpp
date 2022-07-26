#include<iostream>
#include<unistd.h>
#include<time.h>
#include"../mlog.h"
using namespace std;
using namespace my;
int main(int argc, char const *argv[])
{
    mlog* log = mlog::init("./log/");
    long long start = time(NULL);
    for(int i = 0;i<1e6;i++)
    {
        DLOG("test DEBUG %d",i);
        ILOG("test INFO %d",i);
        WLOG("test WARRING %d",i);
        ELOG("test ERROR");
    }
    long long end = time(NULL);
    cout<<end - start;
    mlog::destory();
    pthread_exit(NULL);
}
