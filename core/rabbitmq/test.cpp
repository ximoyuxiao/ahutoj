#include<rabbitmq.h>
#include<iostream>
#include<string>
using namespace std;
void callback(amqp_envelope_t event){
    printf("recived:%s\n",(char*)event.message.body.bytes);
    return ;
}
int main(){
    RabbitMQ mq("127.0.0.1",5672,"ahutoj","2019ahut");
    auto send = mq.createProducer();
    auto read = mq.createConsumer("hello");
    send.sendMessage("hello",(void*)"hello wrold",sizeof("hello world"));
    cout<<"sendl"<<endl;
    char buff[1000];
    read.consumeMessage(callback);
    return 0;
}