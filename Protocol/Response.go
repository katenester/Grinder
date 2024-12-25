package Protocol

const (
	StatusSuccessCode             = 200 // Успех
	StatusCreatedSuccessCode      = 201
	StatusBadRequestCode          = 400 // Некорректный формат
	StatusNotFoundCode            = 404 // Неизвестная команда
	StatusUnauthorizedCode        = 403 // Неавторизованное действие
	StatusConflictCode            = 409 // Пользователь уже существует
	StatusInternalServerErrorCode = 500 // Ошибка на сервере
)
const (
	StatusSuccess             = "OK"
	StatusCreatedSuccess      = "CRETE"                 // Успех
	StatusBadRequest          = "BAD REQUEST"           // Некорректный формат
	StatusNotFound            = "NOT FOUND"             // Неизвестная команда
	StatusUnauthorized        = "UNAUTHORIZED"          // Неавторизованное действие
	StatusConflict            = "CONFLICT"              // Пользователь уже существует
	StatusInternalServerError = "INTERNAL SERVER ERROR" // Ошибка на сервере
)

type Response struct {
	Cod     int    `json:"cod" binding:"required"`
	Message string `json:"message" binding:"required"`
	Body    string `json:"body"` // доска пока так но возможно будет массив
}

func (res Response) Error() string {
	return res.Message
}
func RelateError(code int) string {
	switch code {
	case StatusSuccessCode:
		return StatusSuccess
	case StatusCreatedSuccessCode:
		return StatusCreatedSuccess
	case StatusBadRequestCode:
		return StatusBadRequest
	case StatusNotFoundCode:
		return StatusNotFound
	case StatusUnauthorizedCode:
		return StatusUnauthorized
	case StatusConflictCode:
		return StatusConflict
	default:
		return StatusInternalServerError
	}
}
