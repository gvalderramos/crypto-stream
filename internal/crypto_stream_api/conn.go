package crypto_stream_api

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func NewWebSocketConn(wsUrl *url.URL) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(wsUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Connected to '%s' WebSocket path.\n", wsUrl.Hostname())
	return conn, err
}

func NewRabbitMqConn() *amqp.Connection {
	host := os.Getenv("EVENT_BROKER_HOST")
	user := os.Getenv("EVENT_BROKER_USER")
	pass := os.Getenv("EVENT_BROKER_PASS")
	eventUrl := fmt.Sprintf("amqp://%s:%s@%s:5672", user, pass, host)
	conn, err := amqp.Dial(eventUrl)
	FailOnError(err, "Failed to connect with RabbitMQ.")
	return conn
}

func PostNewEvent(cryptoEvent *CryptoStreamEvent) {
	conn := NewRabbitMqConn()
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to create the event channel.")

	q, err := ch.QueueDeclare(
		"crypto-stream-event", // name
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	FailOnError(err, "Failed to declare the queue context")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msg, err := cryptoEvent.String()
	FailOnError(err, "Failed to encode the message.")

	err = ch.PublishWithContext(
		ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
	FailOnError(err, "Failed to send message")
}

func databaseConn() *sql.DB {
	host := os.Getenv("CRYPTO_STREAM_DATABASE_HOST")
	const (
		port     = 5432
		user     = "guest"
		password = "guest"
		dbname   = "crypto_stream"
	)
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	db, err := sql.Open("postgres", psqlInfo)
	FailOnError(err, "Unable to connect with the database.")
	return db
}

func SaveEvent(cryptoEvent *CryptoStreamEvent) {
	db := databaseConn()
	defer db.Close()

	err := db.Ping()
	FailOnError(err, "Fail checking database's connection healthy.")

	fmt.Println("Successfully connected!")
}
