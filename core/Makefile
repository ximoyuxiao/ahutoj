include Redis/redis.mk
include Tpool/tpool.mk
CXXFLAGS +=-I/usr/local/mysql/include/mysql -I/usr/include/mysql -I/usr/local/include/amqp -I./rabbitmq -g -Wall -std=c++11
LDFLAGS += -L/usr/local/mysql/lib/mysql -L/usr/lib/mysql  -L/usr/lib64/mysql -L/usr/local/lib -L/usr/local/lib/x86_64-linux-gnu 
LIBS += -lmysqlclient -lrabbitmq
TARGET = judged
SOURCE = main.o\
judgeClient.o\
Solution.o\
Solve.o\
mlog.o\
mydb.o\
Redis/redis.o\
Tpool/locker.o\
Tpool/cond.o\
Tpool/sem.o\
Tpool/threadpool.o\
Conf.o\
readConfig.o\
Language.o\
rabbitmq/rabbitmq.o
all:${TARGET}

${TARGET}:${SOURCE}
	g++ $^ -o $@ ${LDFLAGS} ${LIBS}

clean:
	rm ${TARGET} ${SOURCE}