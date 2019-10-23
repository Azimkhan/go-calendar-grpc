package rabbitmq

import (
	"context"
	"github.com/Azimkhan/go-calendar-grpc/internal/domain/interfaces"
	"github.com/Azimkhan/go-calendar-grpc/internal/domain/models"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"time"
)

type ReminderWorker struct {
	repo   interfaces.Repository
	ch     *amqp.Channel
	queue  *amqp.Queue
	logger *zap.Logger
	conn   *amqp.Connection
}

func failOnError(err error, msg string, logger *zap.Logger) {
	if err != nil {
		logger.Fatal(msg, zap.Error(err))
	}
}

func NewWorker(amqpUri string, queueName string, repository interfaces.Repository, logger *zap.Logger) *ReminderWorker {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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

	worker := &ReminderWorker{
		conn:   conn,
		repo:   repository,
		ch:     ch,
		queue:  &q,
		logger: logger,
	}
	return worker
}

func (w *ReminderWorker) Close() {
	err := w.ch.Close()
	if err != nil {
		w.logger.Warn("Failed to close AMQP channel", zap.Error(err))
	}

	err = w.conn.Close()
	if err != nil {
		w.logger.Warn("Failed to close AMQP connection", zap.Error(err))
	}
}
func (w *ReminderWorker) fetchEvents(from time.Time, to time.Time) ([]*models.CalendarEvent, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	events, err := w.repo.FetchByDateRange(from, to, ctx)
	if err != nil {
		return nil, err
	}
	return events, nil
}
func (w *ReminderWorker) sendReminders(events []*models.CalendarEvent) error {
	for _, event := range events {
		err := w.ch.Publish(
			"",           // exchange
			w.queue.Name, // routing key
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(event.Name),
			})
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *ReminderWorker) Work() {
	for {
		// send reminders for every minute
		from := time.Now().UTC()
		to := from.Add(1 * time.Minute)

		events, err := w.fetchEvents(from, to)
		if err != nil {
			failOnError(err, "Failed to fetch from repository", w.logger)
		}

		err = w.sendReminders(events)
		if err != nil {
			failOnError(err, "Failed to send reminders", w.logger)
		}
		// Sleep until next period
		dur := to.Sub(time.Now())
		time.Sleep(dur)
	}
}
