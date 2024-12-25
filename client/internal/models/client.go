package models

import (
	"Grinder/Protocol"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net"
)

type Client struct {
	conn           net.Conn //соединение
	name           string   // имя пользователя
	idRoomActivity int      // id активной комнаты
	board          [][]byte // текущее представление доски
	ctx            context.Context
}

func NewClientKat(conn net.Conn) *Client {
	return &Client{conn: conn, ctx: context.Background()}
}

func (c *Client) ChooseUsername() {
	var username string
	for {
		fmt.Println("Please enter your username:")
		fmt.Scanln(&username)
		err := c.send(Protocol.Request{Command: viper.GetString("command.sign"), Username: username})
		if err != nil {
			continue
		}
		resp := c.accept()
		fmt.Println("Ответ от сервера:", resp)
		if resp.Cod == 200 || resp.Cod == 201 {
			// Всё норм, имя прошло
			c.name = username
			return
		}
	}
}
func (c *Client) ChooseStrategy() {
	var err error
	// Выбираем стратегию
	var choose int
	for {
		fmt.Println("Выберите: \n 1. Сетевая игра \n Одиночная игра")
		fmt.Scanln(&choose)
		switch choose {
		case 1:
			err = c.send(Protocol.Request{Command: viper.GetString("command.game_user"), Username: c.name})
			if err != nil {
				break
			}
			err = c.accept()
			if err != nil {
				break
			}
			return
		case 2:
			err = c.send(Protocol.Request{Command: viper.GetString("command.game_server"), Username: c.name})
			if err != nil {
				break
			}
			err = c.accept()
			if err != nil {
				break
			}
			return
		default:
			fmt.Println("Выбор не распознан.Нужно выбрать пункт 1 или 2.")
		}
	}
}

func (c *Client) MakeMove() {
	err := c.send(Protocol.Request{Command: viper.GetString("command.MakeMove"), Username: c.name})
	if err != nil {
		return
	}
	var resp Protocol.Response
	resp = c.accept()
	if err != nil {
		return
	}
	fmt.Println(resp)
	return
}
func (c *Client) TakeChips() {
	err := c.send(Protocol.Request{Command: viper.GetString("command.TakeChips"), Username: c.name})
	if err != nil {
		return
	}
	var resp Protocol.Response
	resp = c.accept()
	if err != nil {
		return
	}
	fmt.Println(resp)
	return
}
func (c *Client) MoveChips() {
	err := c.send(Protocol.Request{Command: viper.GetString("command.MoveChips"), Username: c.name})
	if err != nil {
		return
	}
	var resp Protocol.Response
	resp = c.accept()
	if err != nil {
		return
	}
	fmt.Println(resp)
	return
}
func (c *Client) GetTopScores() {
	err := c.send(Protocol.Request{Command: viper.GetString("command.top"), Username: c.name})
	if err != nil {
		return
	}
	var resp Protocol.Response
	resp = c.accept()
	fmt.Println("Ответ от сервера:", resp)
	return
}
func (c *Client) Exit() {
	err := c.send(Protocol.Request{Command: viper.GetString("command.exit"), Username: c.name})
	if err != nil {
		return
	}
	var resp Protocol.Response
	resp = c.accept()
	if err != nil {
		return
	}
	fmt.Println(resp)
	// Очистка доски
	c.board = make([][]byte, 0)
	return
}

func (c *Client) send(req Protocol.Request) error {
	// Отправляем серверу json
	encoder := json.NewEncoder(c.conn)
	err := encoder.Encode(req)
	fmt.Println("req", req)
	if err != nil {
		log.Print(err.Error())
	}
	return err
}
func (c *Client) accept() Protocol.Response {
	resp := Protocol.Response{}
	// Чтение данных с вервера
	decoder := json.NewDecoder(c.conn)
	err := decoder.Decode(&resp)
	if err != nil {
		resp = Protocol.Response{Cod: 500, Message: Protocol.RelateError(500)}
	}
	return resp
}
