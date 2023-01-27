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
    solve =  new Solve("1",1,"#include<iostream>\nusing namespace std;\nint main(){\n\tint a;\n\tint b;\n\tcin >> a >> b;\n\tcout<<(a / (float)b + 0.001);\n\texit(0);\n}",1000,128,CPP11, 1);
    judgeClient juc(solve);
    status = OJ_JUDGE;
    juc.judge();
    if(solve->Sres() != OJ_AC){
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
    solve  = new Solve("1",1,"#include<iostream>\nusing namespace std;\nint main(){\n\tcout<<\"hello world\"<<endl<<endl;\n\texit(0);\n}",1000,128,CPP11);
    judgeClient juc(solve);
    status = OJ_JUDGE;
    juc.judge();
    if(solve->Sres() != OJ_PE){
        printf("Failed res=%s\n",runningres[solve->Sres()]);
        return ;
    }
    puts("pass");
    allPass++;
    delete solve;
}

static void judgeClient_RE(){
    all++;
    printf("Test RE:");
    Solve* solve = nullptr;
    solve =  new Solve("1",1,"#include<iostream>\nusing namespace std;\nint main(){\n\tint a=1,b=0;\n\tcout<<a/b;\n\texit(0);\n}",1000,128,CPP11);
    judgeClient juc(solve);
    status = OJ_JUDGE;
    juc.judge();
    if(solve->Sres() != OJ_RE){
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
    solve =  new Solve("1",1,"#include<iostream>\nusing namespace std;\nint main(){\n\tcout<<\"world\";\n\texit(0);\n}",1000,128,CPP11);
    judgeClient juc(solve);
    status = OJ_JUDGE;
    juc.judge();
    if(solve->Sres() != OJ_WA){
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
    solve = new Solve("1",1,"#include<iostream>\nusing namespace std;\nint main(){\n\tfor(int i=0;;i++){int a =0;}\n\texit(0);\n}",1000,128,CPP11);
    judgeClient juc(solve);
    status = OJ_JUDGE;
    juc.judge();
    if(solve->Sres() != OJ_TLE){
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
    solve =  new Solve("1",1,"#include<iostream>\n#include<cstdlib>\nusing namespace std;\nint main(){\n\tauto arr= malloc(256*1024*1024);\n\texit(0);\n}",1000,128,CPP11);
    judgeClient juc(solve);
    status = OJ_JUDGE;
    juc.judge();
    if(solve->Sres() != OJ_MLE){
        printf("Failed res=%s\n",runningres[solve->Sres()]);
        return ;
    }
    puts("pass");
    allPass++;
}

static void judgeClient_CPP_test()
{
    judgeClient_AC();
    judgeClient_TLE();
    judgeClient_MLE();
    judgeClient_PE();
    judgeClient_RE();
    judgeClient_WA();
    printf("you all test:%d,you pass:%d\nyou pass rate:%.2lf%%\n",all,allPass,1.0*allPass/all*100);
}


static void judgeClient_PY3_AC(){
    all++;
    printf("Test AC:");
    Solve* solve = nullptr;
    solve =  new Solve("1",1,"a=input()\nls=a.split()\nprint('hello world')",1000,128,PYTHON3);
    judgeClient juc(solve);
    status = OJ_JUDGE;
    juc.judge();
    if(solve->Sres() != OJ_AC){
        printf("Failed res=%s\n",runningres[solve->Sres()]);
        return ;
    }
    puts("pass");
    allPass++;
}

static void judgeClient_PY3_PE(){
    all++;
    printf("Test PE:");
    Solve* solve = nullptr;
    solve  = new Solve("1",1,"print(\"\"\"\nhello world\"\"\")",1000,128,PYTHON3);
    judgeClient juc(solve);
    status = OJ_JUDGE;
    juc.judge();
    if(solve->Sres() != OJ_PE){
        printf("Failed res=%s\n",runningres[solve->Sres()]);
        return ;
    }
    puts("pass");
    allPass++;
    delete solve;
}

static void judgeClient_PY3_RE(){
    all++;
    printf("Test RE:");
    Solve* solve = nullptr;
    solve =  new Solve("1",1,"1/0",1000,128,PYTHON3);
    judgeClient juc(solve);
    status = OJ_JUDGE;
    juc.judge();
    if(solve->Sres() != OJ_RE){
        printf("Failed res=%s\n",runningres[solve->Sres()]);
        return ;
    }
    puts("pass");
    allPass++;
}

static void judgeClient_PY3_WA(){
     all++;
    printf("Test WA:");
    Solve* solve = nullptr;
    solve =  new Solve("1",1,"print(\'2345\')",1000,128,PYTHON3);
    judgeClient juc(solve);
    status = OJ_JUDGE;
    juc.judge();
    if(solve->Sres() != OJ_WA){
        printf("Failed res=%s\n",runningres[solve->Sres()]);
        return ;
    }
    puts("pass");
    allPass++;
}

static void judgeClient_PY3_TLE(){
    all++;
    printf("Test TLE:");
    Solve *solve;
    solve = new Solve("1",1,"while True:\n\tpass",1000,128,PYTHON3);
    judgeClient juc(solve);
    status = OJ_JUDGE;
    juc.judge();
    if(solve->Sres() != OJ_TLE){
        printf("Failed res=%s\n",runningres[solve->Sres()]);
        return ;
    }
    puts("pass");
    allPass++;
}

static void judgeClient_PY3_MLE(){
     all++;
    printf("Test MLE:");
    Solve* solve = nullptr;
    solve =  new Solve("1",1,"ls = [None] * 1024 * 1024 * 1024",1000,128,PYTHON3);
    judgeClient juc(solve);
    status = OJ_JUDGE;
    juc.judge();
    if(solve->Sres() != OJ_MLE){
        printf("Failed res=%s\n",runningres[solve->Sres()]);
        return ;
    }
    puts("pass");
    allPass++;
}


static void judgeClient_PY3_test()
{
    allPass = 0;
    all = 0;
    judgeClient_PY3_AC();
    judgeClient_PY3_TLE();
    judgeClient_PY3_MLE();
    judgeClient_PY3_PE();
    judgeClient_PY3_RE();
    judgeClient_PY3_WA();
    printf("you all py3 test:%d,you pass:%d\nyou pass rate:%.2lf%%\n",all,allPass,1.0*allPass/all*100);
}

int main()
{
    mlog::init("./log");
    judgeClient_CPP_test();
    judgeClient_PY3_test();
    mlog::destory();
    printf("exit!\n");
    return 0;
}