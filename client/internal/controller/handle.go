package controller

import (
	"Grinder/client/internal/models"
	"fmt"
)

type Handle interface {
	Sign()
	CreateRoom()
	JoinRoom()
	StartGame()
	Play()
	GetTopScores()
	Stop()
	Exit()
}
type HandlerCondole struct {
	client models.Client
}

func NewHandler(client models.Client) Handle {
	return &HandlerCondole{client: client}
}
func (c *HandlerCondole) Sign() {
	for {
		fmt.Println("Please enter your username:")
		var username string
		fmt.Scanln(&username)
		c.client.Sign(username)
	}
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
