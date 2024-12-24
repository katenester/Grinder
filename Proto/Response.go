package Proto

type Response struct {
	Cod     int    `json:"cod" binding:"required"`
	Message string `json:"message" binding:"required"`
	Body    string `json:"body"` // доска пока так но возможно будет массив
}
