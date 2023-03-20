#ifndef SOLUTION_H__
#define SOLUTION_H__
#include "Redis/redis.h"
#include "mydb.h"
#include"Solve.h"
#include"readConfig.h"
#include <vector>
#include"rabbitmq/rabbitmq.h"
using std::vector;
class Solution
{
private:
    static Solution* solution;
    MyRedis* redis;
    RabbitMQ* mq;
    void GetProblemInfo(Solve*);
    Solution();
    ~Solution();
public:
    static Solution* GetInstance();
    static void Destory();
    bool init(readConfig* rcf);
    void LoopSolve();
    bool commitSolveToDb(Solve* solve);
    void commitSolveToQueue(Solve* solve);
    void ReleaseSolve(Solve* solve);
    static void Process(amqp_envelope_t);
};


#endif