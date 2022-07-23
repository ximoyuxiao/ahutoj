include redis.mk
CC = g++
test:redis.o test.o
	${CC} $^ -o $@ ${LIBS}
clean:
	rm *.o test