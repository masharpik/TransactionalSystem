package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/masharpik/TransactionalSystem/app"
	"github.com/masharpik/TransactionalSystem/app/rabbitmq"
	"github.com/masharpik/TransactionalSystem/utils/literals"
	"github.com/masharpik/TransactionalSystem/utils/logger"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.LogOperationFatal(fmt.Errorf(literals.LogEnvFileNotFound))
	}

	debugOnStr := os.Getenv("DEBUG_ON")
	if debugOnStr == "" {
		logger.LogOperationFatal(fmt.Errorf(literals.LogEnvVarIsNil, "DEBUG_ON"))
	}

	logger.DebugOn, err = strconv.ParseBool(debugOnStr)
	if err != nil {
		logger.LogOperationFatal(err)
	}

	logger.InitLogger()
}

func main() {
	url := rabbitmq.GetConnUrl()
	conn, err := rabbitmq.GetConn(url)
	defer conn.Close()
	if err != nil {
		logger.LogOperationFatal(err)
	}

	ch, err := rabbitmq.GetChannel(conn)
	defer ch.Close()
	if err != nil {
		logger.LogOperationFatal(fmt.Errorf(literals.LogOpenningRabbitMQCChannelError, err))
	}

	sender, err := rabbitmq.CreateSender(ch)
	if err != nil {
		logger.LogOperationFatal(fmt.Errorf(literals.LogDeclarationRabbitMQQueueError, err))
	}

	consumer, err := rabbitmq.CreateConsumer(ch)
	if err != nil {
		logger.LogOperationFatal(fmt.Errorf(literals.LogDeclarationRabbitMQQueueError, err))
	}
	consumer.Listen()

	r, err := app.RegisterUrls(sender)
	if err != nil {
		logger.LogOperationFatal(err)
	}

	err = app.StartServer(r)
	if err != nil {
		logger.LogOperationFatal(err)
	}
}
