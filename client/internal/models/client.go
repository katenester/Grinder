package models

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"net"
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
	request := fmt.Sprint(viper.GetString("sign"), username)
	_, err := c.conn.Write([]byte(request))
	if err != nil {
		return err
	}

	c.name = username
	return nil
}
func (c *ClientKaT) CreateRoom() error {
	request := fmt.Sprint(viper.GetString("create"))
	_, err := c.conn.Write([]byte(request))
	if err != nil {
		return err
	}
	return nil
}
func (c *ClientKaT) JoinRoom() error {
	request := fmt.Sprint(viper.GetString("join"))
	_, err := c.conn.Write([]byte(request))
	if err != nil {
		return err
	}
	return nil
}
func (c *ClientKaT) StartGame() error {
	request := fmt.Sprint(viper.GetString("start"))
	_, err := c.conn.Write([]byte(request))
	if err != nil {
		return err
	}
	return nil
}
func (c *ClientKaT) Play() error {
	request := fmt.Sprint(viper.GetString("game"))
	_, err := c.conn.Write([]byte(request))
	if err != nil {
		return err
	}
	return nil
}
func (c *ClientKaT) GetTopScores() error {
	request := fmt.Sprint(viper.GetString("top"))
	_, err := c.conn.Write([]byte(request))
	if err != nil {
		return err
	}
	return nil
}
func (c *ClientKaT) Stop() error {
	request := fmt.Sprint(viper.GetString("stop"))
	_, err := c.conn.Write([]byte(request))
	if err != nil {
		return err
	}
	return nil
}
