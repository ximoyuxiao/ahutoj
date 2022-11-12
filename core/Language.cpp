#include "Language.h"
#include "okcalls.h"
Language* Language::SolveToLanguage(Solve* solve){
    switch (solve->Lang())
    {
    case C:
        return new C_Language();
        break;
    case CPP:
    case CPP17:
        return new Cpp_Language(17);
        break;
    case CPP11:
        return new Cpp_Language(11);
        break;
    case PYTHON3:
        return new Python3_Language();
        break;
    default:
        return NULL;
        break;
    }
}


lanuage C_Language::getLanguage(){
    return C;
}

lanuage Cpp_Language::getLanguage(){
    return CPP;
}

lanuage Python3_Language::getLanguage(){
    return PYTHON3;
}

lanuage Java_Language::getLanguage(){
    return JAVA;
}

Cpp_Language::Cpp_Language(int version){
    this->version = version;
}

void C_Language::init_syscalls_limits(int call_counter[]){
    for (int i = 0; i == 0 || LANG_CV[i]; i++)
        call_counter[LANG_CV[i]] = HOJ_MAX_LIMIT;
}

void Cpp_Language::init_syscalls_limits(int call_counter[]){
    for (int i = 0; i == 0 || LANG_CV[i]; i++)
		call_counter[LANG_CV[i]] = HOJ_MAX_LIMIT;
}

void Python3_Language::init_syscalls_limits(int call_counter[]){
    for (int i = 0; i == 0 || LANG_YV[i]; i++)
		call_counter[LANG_YV[i]] = HOJ_MAX_LIMIT;
}

void Java_Language::init_syscalls_limits(int call_counter[]){
    
}

void C_Language::compile(char *dir, int pid,const char *src){
    char comp[MAXBUFF];
    char sourceFile[128];
    sprintf(sourceFile,"%s/%d.c",dir, pid);
    FILE* fp = fopen(sourceFile,"w");
    fprintf(fp,"%s",src);
    fclose(fp);
    sprintf(comp,"gcc %s -o %s/main 2>%s/err.txt",sourceFile,dir,dir);
    system(comp);
}

void Cpp_Language::compile(char *dir, int pid,const char *src){
    char comp[MAXBUFF];
    char sourceFile[128];
    sprintf(sourceFile, "%s/%d.cpp", dir, pid);
    FILE* fp = fopen(sourceFile,"w");
    fprintf(fp,"%s", src);
    fclose(fp);
    sprintf(comp,"g++ %s -o %s/main 2>%s/err.txt",sourceFile,dir,dir);
    system(comp);
}

void Python3_Language::compile(char *dir, int pid,const char *src){
    //Python是及时解释性语言，但是要产生错误文件，防止上层判断报错
    char errFile[MAXBUFF];
    char sourceFile[128];
    sprintf(sourceFile, "%s/main.py", dir);
    sprintf(errFile, "%s/err.txt", dir);
    FILE *fp = fopen(sourceFile, "w");
    fprintf(fp, "%s", src);
    fclose(fp);

    FILE *errfp = fopen(errFile, "w");
    fclose(errfp);
}

void Java_Language::compile(char *dir, int pid,const char *src){
    char comp[MAXBUFF];
    char sourceFile[128];
    sprintf(sourceFile,"%s/Main.java",dir);
    FILE* fp = fopen(sourceFile,"w");
    fprintf(fp,"%s", src);
    fclose(fp);
    sprintf(comp,"javac %s %s/Main.class 2>%s/err.txt",sourceFile,dir,dir);
    system(comp);
}

void C_Language::run(char *dir, char * const envp[5]){
    char path[128]={0};
    sprintf(path,"%s/main",dir);
    execle(path,path,NULL,envp);
}

void Cpp_Language::run(char *dir, char * const envp[5]){
    char path[128]={0};
    sprintf(path,"%s/main",dir);
    execle(path,path,NULL,envp);
}

void Python3_Language::run(char *dir, char * const envp[5]){
    char path[128]={0};
    sprintf(path,"%s/main.py",dir);
    execle("/usr/bin/python3", "/usr/bin/python3", path, NULL,envp);
}

void Java_Language::run(char *dir, char * const envp[5]){
    execle("java", "java", "Main.class", NULL, envp);
}