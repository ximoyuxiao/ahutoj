#include<rabbitmq.h>
#include<iostream>
#include<string>
using namespace std;
void callback(amqp_envelope_t event){
    printf("recived:%s\n",(char*)event.message.body.bytes);
    return ;
}
int main(){
    RabbitMQ mq("rabbitmq",5672,"ahutoj","123456");
    auto send = mq.createProducer();
    auto read = mq.createConsumer("hello");
    send.sendMessage("hello",(void*)"hello wrold",sizeof("hello world"));
    cout<<"sendl"<<endl;
    char buff[1000];
    read.consumeMessage(callback);
    return 0;
}