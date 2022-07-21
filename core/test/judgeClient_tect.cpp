#include "../mlog.h"
#include "../Solve.h"
#include "../judgeClient.h"
#include "../result.h"
#include "../readConfig.h"
using namespace my;
int allPass = 0;
int all = 0;
SubRes status;
static void judgeClient_AC(){
    all++;
    printf("Test AC:");
    Solve* solve = nullptr;
    solve = solve = new Solve(1,1,"#include<iostream>\nusing namespace std;\nint main(){\n\tcout<<\"hello world\";\n\texit(0);\n}",1,128,CPP11);
    judgeClient juc(solve);
    status = JUDGE;
    juc.judge();
    if(solve->Sres() != AC){
        printf("Failed res=%s\n",runningres[solve->Sres()]);
        return ;
    }
    puts("pass");
    allPass++;
}
static void judgeClient_PE(){
    all++;
    printf("Test PE:");
    Solve* solve = nullptr;
    solve = solve = new Solve(1,1,"#include<iostream>\nusing namespace std;\nint main(){\n\tcout<<\"hello world\"<<endl;\n\texit(0);\n}",1,128,CPP11);
    judgeClient juc(solve);
    status = JUDGE;
    juc.judge();
    if(solve->Sres() != PE){
        printf("Failed res=%s\n",runningres[solve->Sres()]);
        return ;
    }
    puts("pass");
    allPass++;
}
static void judgeClient_RE(){
    all++;
    printf("Test RE:");
    Solve* solve = nullptr;
    solve = solve = new Solve(1,1,"#include<iostream>\nusing namespace std;\nint main(){\n\tint a=1,b=0;\n\tcout<<a/b;\n\texit(0);\n}",1,128,CPP11);
    judgeClient juc(solve);
    status = JUDGE;
    juc.judge();
    if(solve->Sres() != RE){
        printf("Failed res=%s\n",runningres[solve->Sres()]);
        return ;
    }
    puts("pass");
    allPass++;
}
static void judgeClient_WA(){
     all++;
    printf("Test WA:");
    Solve* solve = nullptr;
    solve = solve = new Solve(1,1,"#include<iostream>\nusing namespace std;\nint main(){\n\tcout<<\"world\";\n\texit(0);\n}",1,128,CPP11);
    judgeClient juc(solve);
    status = JUDGE;
    juc.judge();
    if(solve->Sres() != WA){
        printf("Failed res=%s\n",runningres[solve->Sres()]);
        return ;
    }
    puts("pass");
    allPass++;
}
static void judgeClient_TLE(){
    all++;
    printf("Test TLE:");
    Solve *solve;
    solve = new Solve(1,1,"#include<iostream>\nusing namespace std;\nint main(){\n\tfor(int i=0;;i++){int a =0;}\n\texit(0);\n}",1,128,CPP11);
    judgeClient juc(solve);
    status = JUDGE;
    juc.judge();
    if(solve->Sres() != TLE){
        printf("Failed res=%s\n",runningres[solve->Sres()]);
        return ;
    }
    puts("pass");
    allPass++;
}
static void judgeClient_MLE(){
     all++;
    printf("Test MLE:");
    Solve* solve = nullptr;
    solve = solve = new Solve(1,1,"#include<iostream>\n#include<cstdlib>\nusing namespace std;\nint main(){\n\tmalloc(256*1024*1024);\n\texit(0);\n}",1,128,CPP11);
    judgeClient juc(solve);
    status = JUDGE;
    juc.judge();
    if(solve->Sres() != MLE){
        printf("Failed res=%s\n",runningres[solve->Sres()]);
        return ;
    }
    puts("pass");
    allPass++;
}
static void judgeClient_test()
{
    judgeClient_AC();
    judgeClient_TLE();
    judgeClient_MLE();
    judgeClient_PE();
    judgeClient_RE();
    judgeClient_WA();
    printf("you all test:%d,you pass:%d\nyou pass rate:%.2lf%%\n",all,allPass,1.0*allPass/all*100);
}
int main()
{
    mlog::init("./log");
    judgeClient_test();
    mlog::destory();
    return 0;
}