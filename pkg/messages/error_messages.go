package messages

const (
	EmptyAuthHeader   = "пустой заголовок авторизации"
	InvalidAuthHeader = "неверный формат заголовка авторизации"

	ReadConfigError = "ошибка считывания конфигов"
	ReadEnvError    = "ошибка считывания env файла"

	DBConnectionError = "не удалось подключиться к БД"

	StartHTTPServerError = "возникла ошибка запуска https сервера"

	InvalidSigningMethod = "невалидный метод подписи"
	UnknownTokenClaims   = "тип token claims должны быть типа *tokenClaims"
	InvalidTypeOfUserId  = "невалидный тип айди пользователя"
)
