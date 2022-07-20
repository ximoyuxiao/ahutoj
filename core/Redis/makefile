CC = g++
LIBS += -lhiredis
CXXFLAGS += -I/usr/include/hiredis/
LDFLAGS += -L/usr/local/lib/
test:redis.o test.o
	${CC} $^ -o $@ ${LIBS}
clear:
	rm *.o test