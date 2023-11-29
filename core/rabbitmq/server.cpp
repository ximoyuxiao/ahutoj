#include<rabbitmq.h>
#include<iostream>
#include<string>
using namespace std;
#include<unistd.h>
int main(){
    int i = 0;
    for(;;){
        RabbitMQ mq("rabbitmq",5672,"ahutoj","123456");
        auto send = mq.createProducer();
        send.sendMessage("hello",(void*)&i,sizeof(int));
        usleep(1e3);
        i++;
    }
    return 0;
}