package recieve

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/saurabh-sde/otp-authentication-go/service"
	"github.com/saurabh-sde/otp-authentication-go/utility"
)

func InitializeMQConsumer() (err error) {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", os.Getenv("RABBIT_USER"), os.Getenv("RABBIT_PASSWORD"), os.Getenv("RABBIT_HOST"), os.Getenv("RABBIT_PORT"))
	conn, err := amqp.Dial(connStr)
	if err != nil {
		utility.Print(&err, "Failed to connect to RabbitMQ")
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		utility.Print(&err, "Failed to open a channel")
		return
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
	}

	q, err := ch.QueueDeclare(
		"",    // name
		true,  // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		utility.Print(&err, "Failed to declare a queue")
	}

	err = ch.QueueBind(
		q.Name,    // queue name
		"SendOTP", // routing key
		topic,     // exchange
		false,
		nil)

	if err != nil {
		utility.Print(&err, "Failed to bind a queue")
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
		utility.Print(&err, "Failed to register a consumer")
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			utility.Print(nil, "Recieved msg on topic[%s]: %s", topic, d.Body)
			// send OTP
			service.TwilioSendOTP(string(d.Body))
		}
	}()

	utility.Print(nil, "Waiting for messages. To exit press CTRL+C")
	<-forever

	return
}
