package rabbitmq

import (
	"context"
	"encoding/json"
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

func NewWorker(amqpUrl string, queueName string, repository interfaces.Repository, logger *zap.Logger) *ReminderWorker {
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
		jsonBody, err := json.Marshal(event)
		if err != nil {
			return err
		}
		err = w.ch.Publish(
			"",           // exchange
			w.queue.Name, // routing key
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        jsonBody,
			})
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *ReminderWorker) Run() {
	forever := make(chan bool)
	go func() {
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
	}()
	w.logger.Info(" [*] Reminder worker started.")
	<-forever
}
