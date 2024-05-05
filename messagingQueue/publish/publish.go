package publish

import (
	"context"
	"fmt"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/saurabh-sde/otp-authentication-go/utility"
)

func PublishToMQ(mobile string) error {
	utility.Print(nil, "Publish SMS via MQ: ", mobile)

	// conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", os.Getenv("RABBIT_USER"), os.Getenv("RABBIT_PASSWORD"), os.Getenv("RABBIT_HOST"), os.Getenv("RABBIT_PORT"))
	conn, err := amqp.Dial(connStr)
	if err != nil {
		utility.Print(&err, "Failed to connect to RabbitMQ")
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		utility.Print(&err, "Failed to connect to RabbitMQ")
		return err
	}
	defer ch.Close()

	topic := "verification"
	err = ch.ExchangeDeclare(
		topic,    // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		utility.Print(&err, "Failed to declare an exchange")
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		topic,     // exchange
		"SendOTP", // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(mobile),
		})
	if err != nil {
		utility.Print(&err, "Failed to publish a message")
		return err
	}
	utility.Print(nil, "published message to topic: %s  msg: %+v", topic, mobile)
	return nil
}
