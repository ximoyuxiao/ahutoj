include Redis/redis.mk
include Tpool/tpool.mk
CXXFLAGS +=-I/usr/local/mysql/include/mysql -I/usr/include/mysql -g -Wall -std=c++11
LDFLAGS += -L/usr/local/mysql/lib/mysql -L/usr/lib/mysql  -L/usr/lib64/mysql
LIBS += -lmysqlclient
TARGET = judged
SOURCE = main.o\
judgeClient.o\
SolutionDb.o\
Solve.o\
mlog.o\
mydb.o\
Redis/redis.o\
Tpool/locker.o\
Tpool/cond.o\
Tpool/sem.o\
Tpool/threadpool.o\
Conf.o\
readConfig.o
all:${TARGET}

${TARGET}:${SOURCE}
	g++ $^ -o $@ ${LIBS}

clean:
	rm ${TARGET} ${SOURCE}