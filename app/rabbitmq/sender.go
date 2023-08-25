package rabbitmq

import (
	"os"

	"github.com/google/uuid"
	"github.com/masharpik/TransactionalSystem/utils/logger"
	"github.com/streadway/amqp"
)

type Sender struct {
	ch         *amqp.Channel
	queue_name string
}

func CreateSender(ch *amqp.Channel) (sender *Sender, err error) {
	sender = &Sender{
		ch:         ch,
		queue_name: os.Getenv("RABBITMQ_QUEUE_NAME"),
	}

	err = sender.getQueue()

	return
}

func (sender *Sender) getQueue() (err error) {
	_, err = sender.ch.QueueDeclare(
		sender.queue_name, // name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)

	return
}

func (sender *Sender) PushTask(callback func()) (err error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		logger.LogOperationError(err)
	}
	taskId := uuid.String()
	taskMap.Store(taskId, callback)

	err = sender.ch.Publish(
		"",                // exchange
		sender.queue_name, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(taskId),
		})
	return
}
