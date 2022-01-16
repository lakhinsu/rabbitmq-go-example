package utils

import (
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

type RMQConsumer struct {
	Queue            string
	ConnectionString string
	MsgHandler       func(queue string, msg amqp.Delivery, err error)
}

func (x RMQConsumer) OnError(err error, msg string) {
	if err != nil {
		x.MsgHandler(x.Queue, amqp.Delivery{}, err)
	}
}

func (x RMQConsumer) Consume() {
	conn, err := amqp.Dial(x.ConnectionString)
	x.OnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	x.OnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		x.Queue, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	x.OnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	x.OnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			x.MsgHandler(x.Queue, d, nil)
		}
	}()
	log.Info().Msgf("Started listening for messages on '%s' queue", x.Queue)
	<-forever
}
