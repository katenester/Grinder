package controller

import "Grinder/client/internal/models"

type Handle interface {
	Sig()
	CreateRoom()
	JoinRoom()
	StartGame()
	Play()
	GetTopScores()
	Stop()
	Exit()
}
type HandlerCondole struct {
	client *models.Client
}

func NewHandler(client *models.Client) Handle {
	return &HandlerCondole{client: client}
}
func (c *HandlerCondole) Sig() {

}
func (c *HandlerCondole) CreateRoom() {

}
func (c *HandlerCondole) JoinRoom() {

}
func (c *HandlerCondole) StartGame() {

}
func (c *HandlerCondole) Play() {

}
func (c *HandlerCondole) GetTopScores() {

}
func (c *HandlerCondole) Stop() {

}
func (c *HandlerCondole) Exit() {

}
