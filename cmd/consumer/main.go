package main

import (
	"fmt"

	"github.com/TSRangel/RabbitMQ_intro/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()

	rabbitmq.QueueDeclare(ch, rabbitmq.Queue)

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, rabbitmq.Queue)
	
	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}