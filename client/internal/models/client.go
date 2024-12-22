package models

import (
	"context"
	"fmt"
	"net"
)

// Запросы для сервера
const (
	sign         = "SIGN %s"
	createRoom   = "CreateRoom"
	joinRoom     = ""
	startGame    = "StartGame"
	play         = "Play"
	getTopScores = "GetTopScores"
	stop         = "Stop"
)

type Client interface {
	Sign(username string) error
	CreateRoom() error
	JoinRoom() error
	StartGame() error
	Play() error
	GetTopScores() error
	Stop() error
}

type ClientKaT struct {
	conn           net.Conn //соединение
	name           string   // имя пользователя
	idRoomActivity int      // id активной комнаты
	board          [][]byte // текущее представление доски
	ctx            context.Context
}

func NewClientKat(conn net.Conn) Client {
	return &ClientKaT{conn: conn, ctx: context.Background()}
}

func (c *ClientKaT) Sign(username string) error {
	request := fmt.Sprintf(sign, username)
	_, err := c.conn.Write([]byte(request))
	if err != nil {
		return err
	}
	c.name = username
	return nil
}
func (c *ClientKaT) CreateRoom() error {
	request := fmt.Sprintf(createRoom)
	_, err := c.conn.Write([]byte(request))
	if err != nil {
		return err
	}
	return nil
}
func (c *ClientKaT) JoinRoom() error {
	request := fmt.Sprintf(joinRoom)
	_, err := c.conn.Write([]byte(request))
	if err != nil {
		return err
	}
	return nil
}
func (c *ClientKaT) StartGame() error {
	request := fmt.Sprintf(startGame)
	_, err := c.conn.Write([]byte(request))
	if err != nil {
		return err
	}
	return nil
}
func (c *ClientKaT) Play() error {
	request := fmt.Sprintf(play)
	_, err := c.conn.Write([]byte(request))
	if err != nil {
		return err
	}
	return nil
}
func (c *ClientKaT) GetTopScores() error {
	request := fmt.Sprintf(getTopScores)
	_, err := c.conn.Write([]byte(request))
	if err != nil {
		return err
	}
	return nil
}
func (c *ClientKaT) Stop() error {
	request := fmt.Sprintf(stop)
	_, err := c.conn.Write([]byte(request))
	if err != nil {
		return err
	}
	return nil
}
