#include<rabbitmq.h>
#include<iostream>
#include<string>
using namespace std;
void callback(amqp_envelope_t event){
    printf("recived:%d\n",*(int*)event.message.body.bytes);
    return ;
}
int main(){
    RabbitMQ mq("127.0.0.1",5672,"ahutoj","2019ahut");
    auto read = mq.createConsumer("hello");
    read.consumeMessage(callback);
    return 0;
}