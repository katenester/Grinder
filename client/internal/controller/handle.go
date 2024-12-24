package controller

import (
	"Grinder/client/internal/models"
	"fmt"
)

type Handle interface {
	Sign() error
	CreateRoom() error
	JoinRoom() error
	StartGame() error
	Play() error
	GetTopScores() error
	Stop() error
	Exit() error
}
type HandlerCondole struct {
	client models.Client
}

func NewHandler(client models.Client) Handle {
	return &HandlerCondole{client: client}
}
func (c *HandlerCondole) Sign() error {
	for {
		fmt.Println("Please enter your username:")
		var username string
		fmt.Scanln(&username)
		c.client.Sign(username)
	}
}
func (c *HandlerCondole) CreateRoom() error {

}
func (c *HandlerCondole) JoinRoom() error {

}
func (c *HandlerCondole) StartGame() error {

}
func (c *HandlerCondole) Play() error {

}
func (c *HandlerCondole) GetTopScores() error {

}
func (c *HandlerCondole) Stop() error {

}
func (c *HandlerCondole) Exit() error {

}
