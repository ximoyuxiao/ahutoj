#ifndef LANGUAGE_H_
#define LANGUAGE_H_

#include <string>
#include <unistd.h>
#include "result.h"

#include "Solve.h"

const int call_array_size = 512;
const int MAXBUFF = 1024;


class Language{
public:
    static Language* SolveToLanguage(Solve* solve);
    virtual lanuage getLanguage() = 0;
    virtual void init_syscalls_limits(int call_counter[]) = 0;
    virtual void compile(char *dir, int pid, const char *src) = 0;
    virtual void run(char *dir, char * const envp[5]) = 0;
};

class C_Language : public Language{
public:
    lanuage getLanguage();
    void init_syscalls_limits(int call_counter[]);
    void compile(char *dir, int pid, const char *src);
    void run(char *dir, char * const envp[5]);
};
class Cpp_Language : public Language{
private:
    int version;
public:
    Cpp_Language(int version);
    lanuage getLanguage();
    void init_syscalls_limits(int call_counter[]);
    void compile(char *dir, int pid, const char *src);
    void run(char *dir, char * const envp[5]);
};
class Python3_Language : public Language{
public:
    lanuage getLanguage();
    void init_syscalls_limits(int call_counter[]);
    void compile(char *dir, int pid, const char *src);
    void run(char *dir, char * const envp[5]);
};
class Java_Language : public Language{
public:
    lanuage getLanguage();
    void init_syscalls_limits(int call_counter[]);
    void compile(char *dir, int pid, const char *src);
    void run(char *dir, char * const envp[5]);
};

#endif