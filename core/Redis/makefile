include redis.mk
CC = g++
all:test

test:redis.o test.o
	${CC} $^ -o $@ ${LIBS}
clean:
	rm *.o test