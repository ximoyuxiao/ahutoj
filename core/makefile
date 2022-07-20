

judgeClient:judgeClient.o mlog.o Solve.o Tpool/cond.o Tpool/locker.o Tpool/threadpool.o Tpool/sem.o
	g++ $^ -o $@ -pthread -std=c++11

clean:
	rm *.o
	rm Tpool/*.o
	rm logtest
	rm judgeClient