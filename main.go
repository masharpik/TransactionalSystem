package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/masharpik/TransactionalSystem/utils/literals"
	"github.com/masharpik/TransactionalSystem/utils/logger"
	"github.com/masharpik/TransactionalSystem/app"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.LogOperationFatal(errors.New(literals.LogEnvFileNotFound))
	}

	debugOnStr := os.Getenv("DEBUG_ON")
	if debugOnStr == "" {
		logger.LogOperationFatal(errors.New(fmt.Sprintf(literals.LogEnvVarIsNil, "DEBUG_ON")))
	}

	logger.DebugOn, err = strconv.ParseBool(os.Getenv("DEBUG_ON"))
	if err != nil {
		logger.LogOperationFatal(err)
	}

	logger.InitLogger()
}

func main() {
	r, err := app.RegisterUrls()
	if err != nil {
		logger.LogOperationFatal(err)
	}

	err = app.StartServer(r)
	if err != nil {
		logger.LogOperationFatal(err)
	}
}
