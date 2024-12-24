package Errors

const (
	StatusSuccessCode             = 200 // Успех
	StatusBadRequestCode          = 400 // Некорректный формат
	StatusNotFoundCode            = 404 // Неизвестная команда
	StatusUnauthorizedCode        = 403 // Неавторизованное действие
	StatusConflictCode            = 409 // Пользователь уже существует
	StatusInternalServerErrorCode = 500 // Ошибка на сервере
)
const (
	StatusSuccess             = "StatusSuccess"             // Успех
	StatusBadRequest          = "StatusBadRequest"          // Некорректный формат
	StatusNotFound            = "StatusNotFound"            // Неизвестная команда
	StatusUnauthorized        = "StatusUnauthorized"        // Неавторизованное действие
	StatusConflict            = "StatusConflict"            // Пользователь уже существует
	StatusInternalServerError = "StatusInternalServerError" // Ошибка на сервере
)

type MyError struct {
	status  int
	message string
}

func (m *MyError) Error() string {
	return m.message
}
