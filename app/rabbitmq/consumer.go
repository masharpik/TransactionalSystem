package rabbitmq

import (
	"fmt"
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
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	return
}

func (consumer *Consumer) Listen() {
	go func() {
		for d := range consumer.msgs {
			taskID := string(d.Body)
			task, ok := taskMap.Load(taskID)
			if ok {
				taskFunc, ok := task.(func())
				if ok {
					taskFunc()
					taskMap.Delete(taskID)
				} else {
					logger.LogOperationError(fmt.Errorf("Задача с ID %s не является функцией", taskID))
				}
			} else {
				logger.LogOperationError(fmt.Errorf("Задача с ID %s не найдена", taskID))
			}
		}
	}()
}

