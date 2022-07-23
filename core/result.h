#ifndef __RESULT_H__
#define __RESULT_H__
typedef enum{
    OJ_AC = 1,
    OJ_WA,
    OJ_TLE,
    OJ_MLE,
    OJ_RE,
    OJ_PE,
    OJ_OLE,
    OJ_CE,
    OJ_JUDGE
}SubRes;
static const char *runningres[] = {"","Accept","Wrong Answer","Time Limit Error","Mermory Limit Error",
                                "Runtime Error","Presentation Error","Output Len Error","Compile Error"};
typedef enum{
    C = 1,
    CPP,
    CPP11,
    CPP17,
    JAVA,
    PYTHON3
}lanuage;
#define DATAPATH    "./Data/"  /*测试样例目录*/
#define DEC         ".des"
#define LOGPATH     "./log"
#define CONF        "./config.conf" /*配置文件目录*/
#define COMPDIR     "./run%d"
#define IPC_PATH    "./judge"

#endif