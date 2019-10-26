package rabbitmq

import (
	"encoding/json"
	"github.com/Azimkhan/go-calendar-grpc/internal/domain/models"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type ReminderConsumer struct {
	ch     *amqp.Channel
	queue  *amqp.Queue
	conn   *amqp.Connection
	logger *zap.Logger
}

func NewConsumer(amqpUrl string, queueName string, logger *zap.Logger) *ReminderConsumer {
	conn, err := amqp.Dial(amqpUrl)
	failOnError(err, "Failed to connect to RabbitMQ", logger)

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel", logger)

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	failOnError(err, "Failed to declare a queue", logger)

	consumer := &ReminderConsumer{
		conn:   conn,
		ch:     ch,
		queue:  &q,
		logger: logger,
	}

	return consumer
}

func (c *ReminderConsumer) Run() {
	msgs, err := c.ch.Consume(
		c.queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	failOnError(err, "Failed to declare a queue", c.logger)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			event := models.CalendarEvent{}
			err := json.Unmarshal(d.Body, &event)
			if err != nil {
				c.logger.Error("Unmarshal error", zap.Error(err))
				continue
			}
			c.logger.Info("Received reminder", zap.Any("event", event))
		}
	}()

	c.logger.Info(" [*] Waiting for messages. To exit press CTRL+C.")
	<-forever
}
