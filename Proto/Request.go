package Proto

type Request struct {
	Command  string `json:"command" binding:"required"`
	Username string `json:"username" binding:"required"`
	Stage    string `json:"stage"`
	Body     string `json:"body"` // доска пока так но возможно будет массив
}
