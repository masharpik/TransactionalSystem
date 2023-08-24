package logger

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/masharpik/TransactionalSystem/utils/literals"
)

var DebugOn bool

func InitLogger() {
	if !DebugOn {
		return
	}

	file, err := os.OpenFile("file.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Warn(literals.LogOpenLogFileError, err)
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
			ForceColors:   true,
		})
		return
	}

	loggerConsole := logrus.New()
	loggerFile := logrus.New()

	loggerConsole.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
	loggerFile.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	loggerConsole.SetOutput(os.Stderr)
	loggerFile.SetOutput(file)

	logrus.SetOutput(io.MultiWriter(loggerConsole.Writer(), loggerFile.Writer()))
}

func getRequestInfo(r *http.Request) map[string]interface{} {
	return map[string]interface{}{
		"method":     r.Method,
		"path":       r.URL.Path,
		"query":      r.URL.Query(),
		"header":     r.Header,
		"body":       r.Body,
		"remoteAddr": r.RemoteAddr,
	}
}

func LogRequestSuccess(r *http.Request, statusCode int) {
	if !DebugOn {
		return
	}

	info := getRequestInfo(r)
	info["status"] = statusCode
	logrus.WithFields(logrus.Fields(info)).Info(literals.LogRequestSuccess)
}

func LogRequestError(r *http.Request, statusCode int, err error) {
	if !DebugOn {
		return
	}

	info := getRequestInfo(r)
	info["error"] = err
	info["status"] = statusCode
	logrus.WithFields(logrus.Fields(info)).Error(literals.LogRequestError)
}

func LogOperationSuccess(operations ...string) {
	if !DebugOn {
		return
	}

	logrus.Info(strings.Join(operations, " "))
}

func LogOperationError(err error) {
	if !DebugOn {
		return
	}

	logrus.Error(err)
}

func LogOperationFatal(err error) {
	logrus.Fatal(err)
}
