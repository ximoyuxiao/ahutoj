package middlewares

import (
	"ahutoj/web/utils"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"time"
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
	if rabbitmq != nil {
		return rabbitmq, nil
	}
	// utils.GetLogInstance().Debug("NewRabbitMQ")
	uri := fmt.Sprintf("amqp://%v:%v@rabbitmq", User, Password)
	rabbitmq, _ = newRabbitMQ(uri, poolSize) //强制重新连接
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
	for len(pool) < poolSize {
		connection, err := amqp.Dial(uri)
		if err != nil {
			logger.Errorf("call Dial failed,now Redialing, len(pool)=%v, err=%s", len(pool), err.Error())
			time.Sleep(5 * time.Second)
			continue
		}
		pool <- connection
	}
	logger.Info("RabbitMQ connection pool create success")
	return &RabbitMQ{ConnectionPool: pool}, nil
}

func isConnectionOpen(conn *amqp.Connection) bool {
	ch, err := conn.Channel()
	if err != nil || ch == nil {
		return false
	}
	defer ch.Close()
	return true
}

func (r *RabbitMQ) GetConnection() (*amqp.Connection, error) {
	for {
		select {
		case conn := <-r.ConnectionPool:
			if !isConnectionOpen(conn) {
				utils.GetLogInstance().Infof("Connection failure,now RegetConnection, len(r.ConnectionPool)=%v", len(r.ConnectionPool))
				continue
			}
			return conn, nil
		default:
			uri := fmt.Sprintf("amqp://%v:%v@rabbitmq", r.User, r.Password)
			// conn,err:=Re(func()(*amqp.Connection,error){return amqp.Dial(uri)}, 3, 5)
			conn, err := amqp.Dial(uri)
			if err != nil {
				utils.GetLogInstance().Errorf("call Dial failed, err=%s", err.Error())
				time.Sleep(5 * time.Second)
				continue
			}
			return conn, nil
		}
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

	ch, _ := conn.Channel()
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
		logger.Errorf("call Channel failed, err=%s", err.Error())
		return nil, err
	}
	// defer func() {
	// 	if err := ch.Close(); err != nil {
	// 		logger.Errorf("call Channel Close failed, err=%s", err.Error())
	// 	}
	// }()我超，不能关
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
