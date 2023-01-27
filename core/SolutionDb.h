#ifndef SOLUTION_DB_H__
#define SOLUTION_DB_H__
#include "Redis/redis.h"
#include "mydb.h"
#include"Solve.h"
#include"readConfig.h"
#include <vector>
using std::vector;
class SolutionDb
{
private:
    MyRedis* redis;
    mysqlDB* mydb;
    void GetProblemInfo(Solve*);
public:
    SolutionDb();
    ~SolutionDb();
    bool initDB(readConfig* rcf);
    vector<Solve*> getSolve();
    bool commitSolveToDb(Solve* solve);
    void ReleaseSolve(Solve* solve);
    
};


#endif