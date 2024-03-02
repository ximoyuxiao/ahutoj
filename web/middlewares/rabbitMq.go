package middlewares

import (
	"ahutoj/web/utils"
	"encoding/json"
	"errors"
	"fmt"
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
	// utils.GetLogInstance().Debug("NewRabbitMQ")
	uri := fmt.Sprintf("amqp://%v:%v@rabbitmq", User, Password)
	rabbitmq, err = newRabbitMQ(uri, poolSize)
	if err != nil {
		utils.GetLogInstance().Errorf("call newRabbitMQ failed,  err=%s", err.Error())
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
	logger := utils.GetLogInstance()
	pool := make(chan *amqp.Connection, poolSize)
	for i := 0; i < poolSize; i++ {
		connection, err := amqp.Dial(uri)
		if err != nil {
			logger.Errorf("call Dial failed, conn=%v, err=%s", connection, err.Error())
			return nil, err
		}
		pool <- connection
	}
	return &RabbitMQ{ConnectionPool: pool}, nil
}

func (r *RabbitMQ) GetConnection() (*amqp.Connection, error) {
	select {
	case conn := <-r.ConnectionPool:
		if conn == nil {
			err := errors.New("received nil connection from ConnectionPool")
			utils.GetLogInstance().Errorf("call conn:=<-r.ConnectionPoll failed,error =%v", err.Error())
			return nil, err
		}
		// utils.GetLogInstance().Debugf("now len(r.ConnectionPool):%v", len(r.ConnectionPool))
		// time.Sleep(100 * time.Second)
		return conn, nil
	default:
		uri := fmt.Sprintf("amqp://%v:%v@rabbitmq", r.User, r.Password)
		conn, err := amqp.Dial(uri)
		logger := utils.GetLogInstance()
		logger.Debug("URI", uri)
		if err != nil {
			logger.Errorf("call Dial failed, conn=%v, err=%s", conn, err.Error())
			return nil, err
		}
		// utils.GetLogInstance().Debugf("now len(r.ConnectionPool):%v", len(r.ConnectionPool))
		return conn, nil
	}
}

func (r *RabbitMQ) ReleaseConnection(conn *amqp.Connection) {
	select {
	case r.ConnectionPool <- conn:
		return
	default:
		utils.GetLogInstance().Info("ConnectionPool is full, closing connection.")
		conn.Close()
	}
}

type Producer struct {
	RabbitMQ *RabbitMQ
}

func NewProducer(rmq *RabbitMQ) *Producer {
	return &Producer{RabbitMQ: rmq}
}

func (p *Producer) SendMessage(queueName string, messageBody interface{}) error {
	logger := utils.GetLogInstance()
	conn, err := p.RabbitMQ.GetConnection()
	if err != nil {
		logger.Errorf("call SendGetConnection failed, submit=%v, err=%s", conn, err.Error())
		return err
	}
	defer p.RabbitMQ.ReleaseConnection(conn)

	ch, err := conn.Channel()
	if err != nil {
		logger.Errorf("call Channel failed, err=%s", err.Error())
		return err
	}
	defer func() {
		if err := ch.Close(); err != nil {
			logger.Errorf("call Channel Close failed, err=%s", err.Error())
		}
	}()
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
	logger := utils.GetLogInstance()
	conn, err := c.RabbitMQ.GetConnection()
	if err != nil {
		logger.Errorf("call ConsumeGetConnection failed, conn=%v, err=%s", conn, err.Error())
		return nil, err
	}
	defer c.RabbitMQ.ReleaseConnection(conn)
	ch, err := conn.Channel()
	if err != nil {
		logger.Errorf("call ConsumeQueueDeclare failed, channel=%v, err=%s", ch, err.Error())
		return nil, err
	}
	defer func() {
		if err := ch.Close(); err != nil {
			logger.Errorf("call Channel Close failed, err=%s", err.Error())
		}
	}()
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
