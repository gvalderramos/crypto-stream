package main

import (
	"log"

	"github.com/gvalderramos/crypto-stream/internal/crypto_stream_api"
)

func main() {
	conn := crypto_stream_api.NewRabbitMqConn()
	defer conn.Close()

	ch, err := conn.Channel()
	crypto_stream_api.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"crypto-stream-event", // name
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	crypto_stream_api.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	crypto_stream_api.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// TODO save these events
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
