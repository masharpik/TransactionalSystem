package rabbitmq

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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

func (consumer *Consumer) sendMessage(d amqp.Delivery, userId string, newAmount float64, link string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", link, nil)

	logger.LogOperationSuccess(fmt.Sprintf("Снятие с баланса.\nТекущее состояние пользователя: %s\nБаланс: %f\n", userId, newAmount))
	req.Header.Set("Custom-Status", "200")

	_, err := client.Do(req)
	if err != nil {
		logger.LogOperationError(fmt.Errorf("Произошла ошибка при попытке отослать простой GET-запрос: %w", err))
	}

	err = d.Ack(false)
	if err != nil {
		logger.LogOperationError(fmt.Errorf("Произошла ошибка при попытке подтверждения выполненного сообщения: %w", err))
	}
}

func (consumer *Consumer) Listen() {
	go func() {
		for d := range consumer.msgs {
			var data taskData
			err := json.Unmarshal(d.Body, &data)
			if err != nil {
				logger.LogOperationError(fmt.Errorf("Произошла ошибка при получении данных брокером: %w", err))
				continue
			}

			userId := data.UserID
			newAmount := data.NewAmount
			link := data.Link

			go consumer.sendMessage(d, userId, newAmount, link)
		}
	}()
}
