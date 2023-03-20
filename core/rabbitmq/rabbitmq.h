#ifndef RABBIT_MQ_H__
#define RABBIT_MQ_H__
#include <string>
#include <vector>
#include <amqp_tcp_socket.h>
#include <amqp.h>
#include <amqp_framing.h>
#include <cstring>
class Producer;
class Consumer;
class RabbitMQ {
public:
    RabbitMQ(std::string host, int port, std::string user, std::string password);
    ~RabbitMQ();
    Producer createProducer();
    Consumer createConsumer(std::string queueName);
    amqp_connection_state_t getConnection();
    void releaseConnection(amqp_connection_state_t conn);
private:
    std::string m_host;
    int m_port;
    std::string m_user;
    std::string m_password;
    int m_poolSize;
    std::vector<amqp_connection_state_t> m_connectionPool;
};

class Producer {
public:
    Producer(RabbitMQ& rmq);

    bool sendMessage(std::string queueName, void* messageBody, size_t messageSize);
private:
    RabbitMQ& m_rmq;
};

class Consumer {
public:
    Consumer(RabbitMQ& rmq, std::string queueName);
    int consumeMessage(void (*callback)(amqp_envelope_t));
private:
    RabbitMQ& m_rmq;
    std::string m_queueName;
    void (*m_callback)(amqp_envelope_t);
};
#endif