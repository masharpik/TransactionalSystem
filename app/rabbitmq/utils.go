package rabbitmq

import (
	"fmt"
	"os"
	"time"

	"github.com/masharpik/TransactionalSystem/utils/literals"
	"github.com/masharpik/TransactionalSystem/utils/logger"
	"github.com/streadway/amqp"
)

type taskData struct {
	UserID    string  `json:"userId"`
	NewAmount float64 `json:"newAmount"`
	Link      string  `json:"link"`
}

func failOnError(err error, msg string) {
	if err != nil {
		logger.LogOperationFatal(fmt.Errorf("%s: %s", msg, err))
	}
}

func GetConnUrl() string {
	name := os.Getenv("RABBITMQ_USER")
	pass := os.Getenv("RABBITMQ_PASS")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", name, pass, host, port)
}

func GetConn(url string) (conn *amqp.Connection, err error) {
	ticker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(2 * time.Minute)
	for {
		select {
		case <-timer.C:
			ticker.Stop()
			err = fmt.Errorf(literals.LogOpenningRabbitMQConnError)
			return
		case <-ticker.C:
			conn, err = amqp.Dial(url)
			if err == nil {
				ticker.Stop()
				timer.Stop()
				logger.LogOperationSuccess(literals.LogConnRabbitMQSuccess)
				return
			}
		}
	}
}

func GetChannel(conn *amqp.Connection) (ch *amqp.Channel, err error) {
	ch, err = conn.Channel()
	return
}
