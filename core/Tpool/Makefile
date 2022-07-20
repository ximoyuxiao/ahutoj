LIBS += -pthread

CXXFLAGS += -std=c++11 -Wall -g

TARGET = test

OBJECT =test.o\
cond.o\
locker.o\
sem.o\
threadpool.o
all:test


${TARGET}:${OBJECT}
	g++ $^ -o $@ ${LIBS}

clear:
	rm *.o test