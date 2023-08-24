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
)
