#include <string>
#include <vector>
#include <amqp_tcp_socket.h>
#include <amqp.h>
#include <amqp_framing.h>
#include <cstring>
#include <rabbitmq.h>
#include<../mlog.h>
using namespace my;
RabbitMQ::RabbitMQ(std::string host, int port, std::string user, std::string password) :
    m_host(host), m_port(port), m_user(user), m_password(password), m_poolSize(10)
{
    m_connectionPool.resize(m_poolSize);
    for (int i = 0; i < m_poolSize; i++) {
        m_connectionPool[i] = nullptr;
    }
}

RabbitMQ::~RabbitMQ() {
    for (auto &conn : m_connectionPool) {
        if (conn != nullptr) {
            amqp_destroy_connection(conn);
            conn = nullptr;
        }
    }
}

Producer RabbitMQ::createProducer() {
    return Producer(this);
}

Consumer RabbitMQ::createConsumer(std::string queueName) {
    return Consumer(this, queueName);
}

bool isConnectionOpen(amqp_connection_state_t conn) {
    amqp_channel_t channel = 1;
    amqp_channel_open_ok_t* res = amqp_channel_open(conn, channel);
    if (res == nullptr) {
        return false;
    }
    amqp_channel_close(conn, channel, AMQP_REPLY_SUCCESS);
    return true;
}

amqp_connection_state_t RabbitMQ::getConnection() {
    amqp_connection_state_t conn = nullptr;
    poolLocker.lock();
    for (int i = 0; i < m_poolSize; i++) {
        if (m_connectionPool[i] != nullptr && isConnectionOpen(m_connectionPool[i])) {
            conn = m_connectionPool[i];
            m_connectionPool[i] = nullptr;
            break;
        }
    }
    poolLocker.unlock();
    if (conn == nullptr) {
        printf("amqp://%s:%s@oj-rabbitmq", m_user.c_str(), m_password.c_str());
        conn = amqp_new_connection();
        if(!conn){
            DLOG( "amqp_socket_open failed");
            return nullptr;
        }
        amqp_socket_t* socket = amqp_tcp_socket_new(conn);
        auto ret = amqp_socket_open(socket, m_host.c_str(), m_port);
        if (ret !=0){
           DLOG( "amqp_socket_open failed,err:%s",amqp_error_string2(ret));
           return nullptr;
        }
        auto rep = amqp_login(conn, "/", 0, 131072, 0, AMQP_SASL_METHOD_PLAIN, m_user.c_str(), m_password.c_str());
        if (rep.reply_type != AMQP_RESPONSE_NORMAL) {
            DLOG( "amqp_login failed,err:%s",amqp_error_string2(rep.library_error));
            return nullptr;
        }
    }
    return conn;
}

void RabbitMQ::releaseConnection(amqp_connection_state_t conn) {
    bool ret = false;
    poolLocker.lock();
    for (int i = 0; i < m_poolSize; i++) {
        if (m_connectionPool[i] == nullptr) {
            m_connectionPool[i] = conn;
            ret  = true;
            break;
        }
    }
    poolLocker.unlock();
    if(!ret){
        amqp_connection_close(conn, AMQP_REPLY_SUCCESS);
        amqp_destroy_connection(conn);
        DLOG( "releaseConnetion\n");
    }
    return ;
}

Producer::Producer(RabbitMQ* rmq) : m_rmq(rmq) {}

bool Producer::sendMessage(std::string queueName, void* messageBody, size_t messageSize) {
    amqp_connection_state_t conn = m_rmq->getConnection();
    if (conn == nullptr) {
        DLOG( "Error get connection\n");
        return false;
    }

    amqp_channel_t channel = 1;
    amqp_basic_properties_t props;
    amqp_bytes_t messageBytes = amqp_bytes_malloc(messageSize);
    memcpy(messageBytes.bytes, messageBody, messageSize);
    if (messageBytes.bytes == nullptr) {
        DLOG( "Error malloc amqp_bytes\n");
        m_rmq->releaseConnection(conn);
        return false;
    }

    amqp_channel_open_ok_t *channel_open = amqp_channel_open(conn, channel);
    if (channel_open == nullptr) {
        DLOG( "Error opening channel\n");
        amqp_channel_close(conn, channel, AMQP_REPLY_SUCCESS);
        m_rmq->releaseConnection(conn);
        return false;
    }
    amqp_queue_declare_ok_t* queue = amqp_queue_declare(conn, channel, amqp_cstring_bytes(queueName.c_str()), false, false, false, false, amqp_empty_table);
    if(queue == nullptr){
        DLOG( "Error declare amqp_queue\n");
        amqp_bytes_free(messageBytes);
        amqp_channel_close(conn, channel, AMQP_REPLY_SUCCESS);
        m_rmq->releaseConnection(conn);
        return false;
    }
    amqp_bytes_t queueBytes = amqp_bytes_malloc_dup(queue->queue);
    if (queueBytes.bytes == nullptr) {
        DLOG( "Error malloc amqp_bytes\n");
        amqp_bytes_free(messageBytes);
        amqp_channel_close(conn, channel, AMQP_REPLY_SUCCESS);
        m_rmq->releaseConnection(conn);
        return false;
    }

    props._flags = AMQP_BASIC_CONTENT_TYPE_FLAG | AMQP_BASIC_DELIVERY_MODE_FLAG;
    props.content_type = amqp_cstring_bytes("application/json");
    props.delivery_mode = 2; // 持久化消息

    int ret = amqp_basic_publish(conn, channel, amqp_empty_bytes, queueBytes, 0, 0, &props, messageBytes);
    amqp_bytes_free(queueBytes);
    amqp_bytes_free(messageBytes);
    amqp_channel_close(conn, channel, AMQP_REPLY_SUCCESS);
    m_rmq->releaseConnection(conn);
    return ret == AMQP_STATUS_OK;
}


Consumer::Consumer(RabbitMQ* rmq, std::string queueName) :
    m_rmq(rmq), m_queueName(queueName), m_callback(nullptr)
{}

int Consumer::consumeMessage(void (*callback)(amqp_envelope_t)) {
    m_callback = callback;
    amqp_connection_state_t conn = m_rmq->getConnection();
    if (conn == nullptr) {
        DLOG( "Error get Connection\n");
        return 1;
    }
    amqp_channel_t channel = 1; // initialize channel to a non-zero value
    int ret = 0;

    amqp_channel_open_ok_t *channel_open = amqp_channel_open(conn, channel);
    if (channel_open == nullptr) {
        DLOG( "Error opening channel\n");
        amqp_channel_close(conn, channel, AMQP_REPLY_SUCCESS);
        m_rmq->releaseConnection(conn);
        return 1;
    }
    amqp_queue_declare_ok_t* queue = amqp_queue_declare(conn, channel, amqp_cstring_bytes(m_queueName.c_str()), false, false, false, false, amqp_empty_table);
    if (queue == nullptr) {
        DLOG( "Error declare amqp_queue\n");
        amqp_channel_close(conn, channel, AMQP_REPLY_SUCCESS);
        m_rmq->releaseConnection(conn);
        return 1;
    }

    amqp_bytes_t queueBytes = amqp_bytes_malloc_dup(queue->queue);
    if (queueBytes.bytes == nullptr) {
        DLOG( "Error malloc amqp_bytes\n");
        amqp_channel_close(conn, channel, AMQP_REPLY_SUCCESS);
        m_rmq->releaseConnection(conn);
        return 1;
    }

    amqp_basic_consume_ok_t *consume_ok = amqp_basic_consume(conn, channel, queueBytes, amqp_empty_bytes, 0, true, false, amqp_empty_table);
    if (consume_ok == nullptr) {
        DLOG( "Error consuming message\n");
        amqp_channel_close(conn, channel, AMQP_REPLY_SUCCESS);
        m_rmq->releaseConnection(conn);
        return 1;
    }

    amqp_bytes_t consumerTag = amqp_bytes_malloc_dup(consume_ok->consumer_tag);
    if (consumerTag.bytes == nullptr) {
        DLOG( "Error malloc amqp_bytes\n");
        amqp_channel_close(conn, channel, AMQP_REPLY_SUCCESS);
        m_rmq->releaseConnection(conn);
        return 1;
    }

    while (true) {
        amqp_envelope_t envelope = { 0 };
        amqp_maybe_release_buffers(conn);

        auto reply = amqp_consume_message(conn, &envelope, NULL, 0);
        if (reply.reply_type == AMQP_RESPONSE_NORMAL) {
            if (m_callback != nullptr) {
                m_callback(envelope);
            }
            amqp_basic_ack(conn, channel, envelope.delivery_tag, false);
            amqp_destroy_envelope(&envelope);
        } else {
            if (reply.library_error < 0) {
                // if (reply.library_error != AMQP_STATUS_TIMEOUT) {
                //     fprintf(stderr, "Error consuming message: %s\n", amqp_error_string2(reply.library_error));
                // }else {
                //     fprintf(stderr, "Timeout consuming message\n");//没有设置超时
                // }
                fprintf(stderr, "Error consuming message: %s\n", amqp_error_string2(reply.library_error));
                ret = 1;
            }
            break;
        }
    }

    amqp_bytes_free(queueBytes);
    amqp_bytes_free(consumerTag);
    amqp_channel_close(conn, channel, AMQP_REPLY_SUCCESS);
    m_rmq->releaseConnection(conn);
    return ret;
}