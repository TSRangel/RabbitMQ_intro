package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

var(
	Queue = "minha_fila"
	ExName = "minha_exchange"
	Kind = "direct"
	Key = "chave_de_roteamento"
)

func OpenChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch
}

func QueueDeclare(ch *amqp.Channel, queue string) {
	_, err := ch.QueueDeclare(queue, false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
}

func ExchangeDeclare(ch *amqp.Channel, exName string, kind string) {
	err := ch.ExchangeDeclare(exName, kind, false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
}

func QueueBind(ch *amqp.Channel, queue string, key string, exName string) {
	err := ch.QueueBind(queue, key, exName, false, nil)
	if err != nil {
		panic(err)
	}
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery, queue string) {
	msgs, err := ch.Consume(
		queue,
		"Go_Consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	for msg := range msgs {
		out <- msg
	}
}

func Publish(ch *amqp.Channel, body string, exName string, key string) error {
	err := ch.PublishWithContext(
		context.TODO(),
		exName,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
