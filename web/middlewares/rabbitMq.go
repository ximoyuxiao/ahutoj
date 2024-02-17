package middlewares

import (
	"ahutoj/web/utils"
	"encoding/json"
	"fmt"
	"github.com/bytedance/gopkg/util/logger"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Host           string
	Port           int
	User           string
	Password       string
	ConnectionPool chan *amqp.Connection
}

var rabbitmq *RabbitMQ = nil

func NewRabbitMQ(Host string, Port int, User string, Password string, poolSize int) (*RabbitMQ, error) {
	var err error
	if rabbitmq != nil {
		return rabbitmq, nil
	}
	uri := fmt.Sprintf("amqp://%v:%v@%v", User, Password, Host)
	rabbitmq, err = newRabbitMQ(uri, poolSize)
	if err != nil {
		return nil, err
	}
	rabbitmq.Host = Host
	rabbitmq.Password = Password
	rabbitmq.User = User
	rabbitmq.Port = Port
	return rabbitmq, nil
}

func GetRabbitMq() *RabbitMQ {
	return rabbitmq
}
func newRabbitMQ(uri string, poolSize int) (*RabbitMQ, error) {
	pool := make(chan *amqp.Connection, poolSize)
	for i := 0; i < poolSize; i++ {
		connection, err := amqp.Dial(uri)
		if err != nil {
			return nil, err
		}
		pool <- connection
	}
	test := &RabbitMQ{ConnectionPool: pool}
	logger := utils.GetLogInstance()
	logger.Debug("rabbitmq:", test)
	return &RabbitMQ{ConnectionPool: pool}, nil
}

func (r *RabbitMQ) GetConnection() (*amqp.Connection, error) {
	select {
	case conn := <-r.ConnectionPool:
		return conn, nil
	default:
		uri := fmt.Sprintf("amqp://%v:%v@%v", r.User, r.Password, r.Host)
		conn, err := amqp.Dial(uri)
		logger := utils.GetLogInstance()
		logger.Debug("URI", uri)
		if err != nil {
			logger.Errorf("call Dial failed, conn=%v, err=%s", conn, err.Error())
			return nil, err
		}
		return conn, nil
	}
}

func (r *RabbitMQ) ReleaseConnection(conn *amqp.Connection) {
	r.ConnectionPool <- conn
}

type Producer struct {
	RabbitMQ *RabbitMQ
}

func NewProducer(rmq *RabbitMQ) *Producer {
	return &Producer{RabbitMQ: rmq}
}

func (p *Producer) SendMessage(queueName string, messageBody interface{}) error {
	conn, err := p.RabbitMQ.GetConnection()
	if err != nil {
		logger.Errorf("call SendGetConnection failed, submit=%v, err=%s", conn, err.Error())
		return err
	}
	defer p.RabbitMQ.ReleaseConnection(conn)

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // queue name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		logger.Errorf("call SendQueueDeclare failed, queue=%v, err=%s", q, err.Error())
		return err
	}
	data, _ := json.Marshal(messageBody)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
	if err != nil {
		logger.Errorf("call DoPublish failed, data=%v, err=%s", data, err.Error())
		return err
	}

	return nil
}

type Consumer struct {
	RabbitMQ  *RabbitMQ
	QueueName string
}

func NewConsumer(rmq *RabbitMQ, queueName string) *Consumer {
	return &Consumer{RabbitMQ: rmq, QueueName: queueName}
}

func (c *Consumer) ConsumeMessage() (<-chan amqp.Delivery, error) {
	conn, err := c.RabbitMQ.GetConnection()
	if err != nil {
		logger.Errorf("call ConsumeGetConnection failed, conn=%v, err=%s", conn, err.Error())
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {

		return nil, err
	}

	q, err := ch.QueueDeclare(
		c.QueueName, // queue name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		logger.Errorf("call ConsumeQueueDeclare failed, queue=%v, err=%s", q, err.Error())
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil {
		logger.Errorf("call DoConsume failed, consume=%v, err=%s", msgs, err.Error())
		return nil, err
	}

	return msgs, nil
}
