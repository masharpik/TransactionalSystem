package rabbitmq

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/masharpik/TransactionalSystem/app/transaction/utils"
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

func (sender *Sender) PushTask(status utils.StatusTransaction) (err error) {
	jsonData, err := json.Marshal(status)
	if err != nil {
		return fmt.Errorf("Произошла ошибка при создании таска (Marshal'инге данных): %w", err)
	}

	err = sender.ch.Publish(
		"",                // exchange
		sender.queue_name, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		},
	)
	if err != nil {
		return fmt.Errorf("Произошла ошибка при пуше таска: %w", err)
	}

	return
}
