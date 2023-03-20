package middlewares_test

import (
	"ahutoj/web/middlewares"
	"ahutoj/web/utils"
	"fmt"
	"log"
	"sync"
	"testing"
)

func TestRabitMQ(t *testing.T) {
	utils.ConfigInit("/home/moyu/vscode/ahutoj/config.yaml")
	var wg sync.WaitGroup
	mqcfg := utils.GetConfInstance().RabbitMQ
	rabbitmq, err := middlewares.NewRabbitMQ(mqcfg.Host, mqcfg.Port, mqcfg.Username, mqcfg.Password, 10)
	if err != nil {
		log.Fatalf("failed to create rabbitmq object: %v", err)
	}

	producer := middlewares.NewProducer(rabbitmq)

	wg.Add(1)
	go func() {
		defer wg.Done()

		message := []byte("Hello World!")
		err := producer.SendMessage("hello", message)
		if err != nil {
			log.Fatalf("failed to publish message: %v", err)
		}
	}()
	consumer := middlewares.NewConsumer(rabbitmq, "hello")

	wg.Add(1)
	go func() {
		defer wg.Done()

		messages, err := consumer.ConsumeMessage()
		if err != nil {
			log.Fatalf("failed to register consumer: %v", err)
		}

		for msg := range messages {
			fmt.Printf("Received message: %s\n", string(msg.Body))
		}
	}()

	wg.Wait()
}
