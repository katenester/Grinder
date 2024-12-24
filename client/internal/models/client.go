package models

import (
	"Grinder/Proto"
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
		err := c.send(Proto.Request{Command: viper.GetString("sign"), Username: username})
		if err != nil {
			continue
		}
		_, err = c.accept()
		if err != nil {
			continue
		}
		// Всё норм, имя прошло
		c.name = username
		return
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
			err = c.send(Proto.Request{Command: viper.GetString("game_user"), Username: c.name})
			if err != nil {
				break
			}
			_, err = c.accept()
			if err != nil {
				break
			}
			return
		case 2:
			err = c.send(Proto.Request{Command: viper.GetString("game_server"), Username: c.name})
			if err != nil {
				break
			}
			_, err = c.accept()
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
	err := c.send(Proto.Request{Command: viper.GetString("MakeMove"), Username: c.name})
	if err != nil {
		return
	}
	var resp Proto.Response
	resp, err = c.accept()
	if err != nil {
		return
	}
	fmt.Println(resp)
	return
}
func (c *Client) TakeChips() {
	err := c.send(Proto.Request{Command: viper.GetString("TakeChips"), Username: c.name})
	if err != nil {
		return
	}
	var resp Proto.Response
	resp, err = c.accept()
	if err != nil {
		return
	}
	fmt.Println(resp)
	return
}
func (c *Client) MoveChips() {
	err := c.send(Proto.Request{Command: viper.GetString("MoveChips"), Username: c.name})
	if err != nil {
		return
	}
	var resp Proto.Response
	resp, err = c.accept()
	if err != nil {
		return
	}
	fmt.Println(resp)
	return
}
func (c *Client) GetTopScores() {
	err := c.send(Proto.Request{Command: viper.GetString("top"), Username: c.name})
	if err != nil {
		return
	}
	var resp Proto.Response
	resp, err = c.accept()
	if err != nil {
		return
	}
	fmt.Println(resp)
	return
}
func (c *Client) Exit() {
	err := c.send(Proto.Request{Command: viper.GetString("exit"), Username: c.name})
	if err != nil {
		return
	}
	var resp Proto.Response
	resp, err = c.accept()
	if err != nil {
		return
	}
	fmt.Println(resp)
	// Очистка доски
	c.board = make([][]byte, 0)
	return
}

func (c *Client) send(req Proto.Request) error {
	// Отправляем серверу json
	encoder := json.NewEncoder(c.conn)
	err := encoder.Encode(req)
	if err != nil {
		log.Print(err.Error())
	}
	return err
}
func (c *Client) accept() (Proto.Response, error) {
	var resp Proto.Response
	// Чтение данных с вервера
	decoder := json.NewDecoder(c.conn)
	err := decoder.Decode(&resp)
	// Ошибка при декодировании
	if err != nil || resp.Cod != 200 {
		log.Print(err.Error(), resp)
		return Proto.Response{}, err
	}
	return resp, nil
}
