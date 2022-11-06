CXXFLAGS += -std=c++11 -g

JudgeObj = judgeClient_test.o\
	../judgeClient.o\
	../mlog.o\
	../Solve.o\
 	../Tpool/cond.o\
  	../Tpool/locker.o\
   	../Tpool/threadpool.o\
    ../Tpool/sem.o

MlogObj = logtest.o\
	../mlog.o\
	../Tpool/cond.o\
	../Tpool/locker.o

judgeClient_test:${JudgeObj}
	g++ $^ -o $@ -pthread -std=c++11

logtest:${MlogObj}
	g++ $^ -o $@ -pthread -std=c++11

clean:
	rm ${JudgeObj}
	rm judgeClient_test