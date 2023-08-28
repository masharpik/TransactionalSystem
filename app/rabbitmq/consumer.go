package rabbitmq

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/masharpik/TransactionalSystem/app/transaction/utils"
	"github.com/masharpik/TransactionalSystem/utils/logger"
	"github.com/streadway/amqp"
)

type Consumer struct {
	ch         *amqp.Channel
	queue_name string
	msgs       <-chan amqp.Delivery
}

func CreateConsumer(ch *amqp.Channel) (consumer *Consumer, err error) {
	consumer = &Consumer{
		ch:         ch,
		queue_name: os.Getenv("RABBITMQ_QUEUE_NAME"),
	}

	consumer.msgs, err = consumer.getConsume()
	if err != nil {
		return
	}

	return
}

func (consumer *Consumer) getConsume() (msgs <-chan amqp.Delivery, err error) {
	msgs, err = consumer.ch.Consume(
		consumer.queue_name, // queue
		"",                  // consumer
		false,               // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	return
}

func (consumer *Consumer) sendMessage(d amqp.Delivery, status string, userId string, balance float64, link string) {
	var err error
	client := &http.Client{}
	req, _ := http.NewRequest("GET", link, nil)

	req.Header.Set("Custom-Code", "200")
	req.Header.Set("Custom-Status", status)
	req.Header.Set("Custom-UserId", userId)
	req.Header.Set("Custom-Balance", fmt.Sprint(balance))

	if _, err = client.Do(req); err != nil {
		logger.LogOperationError(fmt.Errorf("Произошла ошибка при попытке отослать простой GET-запрос: %w", err))
	}

	if err = d.Ack(false); err != nil {
		logger.LogOperationError(fmt.Errorf("Произошла ошибка при попытке подтверждения выполненного сообщения: %w", err))
	}
}

func (consumer *Consumer) Listen() {
	go func() {
		for d := range consumer.msgs {
			var data utils.StatusTransaction
			err := json.Unmarshal(d.Body, &data)
			if err != nil {
				logger.LogOperationError(fmt.Errorf("Произошла ошибка при получении данных брокером: %w", err))
				continue
			}

			userId := data.UserID
			status := data.Status
			balance := data.Balance
			link := data.Destination

			go consumer.sendMessage(d, userId, status, balance, link)
		}
	}()
}
