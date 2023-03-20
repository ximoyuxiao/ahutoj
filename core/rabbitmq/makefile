CC=g++
CXXFLAGS=-std=c++11 -g -Wall -I/usr/local/include/amqp -I.
LDFLAGS=-L/usr/local/lib -L/usr/local/lib/x86_64-linux-gnu -lrabbitmq
all:rabitmq test server client

rabitmq:rabbitmq.o

test:rabbitmq.o test.o
	${CC} $^ -o $@ $(LDFLAGS)

server:server.o rabbitmq.o
	${CC} $^ -o $@ $(LDFLAGS)
	
client:client.o rabbitmq.o
	${CC} $^ -o $@ $(LDFLAGS)