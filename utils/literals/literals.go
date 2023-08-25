package literals

const (
	LogOpenLogFileError                         = "Не удалось открыть файл или директорию для логирования, будет использоваться дефолтное os.Stderr"
	LogRequestSuccess                           = "Запрос выполнился успешно"
	LogRequestError                             = "Запрос выполнился неуспешно"
	LogEnvFileNotFound                          = "Не найден .env файл"
	LogEnvVarIsNil                              = "Переменная \"%s\" окружения не обнаружена"
	LogErrorOccurredBeforeResponseWriterMethods = "Возникла ошибка до исполнения методов http.ResponseWriter: :%w"
	LogServerWasStarted                         = "Сервер был запущен по адресу %s"
	LogPanicOccured                             = "В процессе обработки запроса возникла паника: %w"
	LogConnDBFailed                             = "Не удалось подключиться к БД: %w"
	LogConnDBTimeout                            = "Время на попытку подключиться к БД вышло"
	LogConnDBSuccess                            = "Подключение к бд произошло успешно: %s"
	CodeUniqConflict                            = "23505"
	LogOpenningRabbitMQConnError                = "Не удалось открыть соединение с брокером сообщений"
	LogOpenningRabbitMQCChannelError            = "Не удалось открыть канал с брокером сообщений: %w"
	LogDeclarationRabbitMQQueueError            = "Ошибка в создании очереди в брокере сообщений: %w"
	LogDeclarationRabbitMQConsumerError         = "Ошибка в создании обработчике в брокере сообщений: %w"
	LogConnRabbitMQSuccess                      = "Подключение к брокеру сообщений успешно"
)
