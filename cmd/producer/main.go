package main

import "github.com/TSRangel/RabbitMQ_intro/pkg/rabbitmq"



func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()

	rabbitmq.ExchangeDeclare(ch, rabbitmq.ExName, rabbitmq.Kind)
	rabbitmq.QueueBind(ch, rabbitmq.Queue, rabbitmq.Key, rabbitmq.ExName)

	err := rabbitmq.Publish(ch, "Hello Future!", rabbitmq.ExName, rabbitmq.Key)
	if err != nil {
		panic(err)
	}
}
