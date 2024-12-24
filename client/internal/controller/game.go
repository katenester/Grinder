package controller

import (
	"Grinder/Proto"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net"
)

func StartGame(ch chan struct{}, conn net.Conn) {
	// Логинимся
	username := ChooseUsername(conn)
	// Выбираем имя
	ChooseStrategy(conn, username)
}
func Send(conn net.Conn, req Proto.Request) error {
	// Отправляем серверу json
	encoder := json.NewEncoder(conn)
	err := encoder.Encode(req)
	if err != nil {
		log.Print(err.Error())
	}
	return err
}
func Accept(conn net.Conn) error {
	var resp Proto.Response
	// Чтение данных с вервера
	decoder := json.NewDecoder(conn)
	err := decoder.Decode(&resp)
	// Ошибка при декодировании
	if err != nil || resp.Cod != 200 {
		log.Print(err.Error(), resp)
		return err
	}
	return nil
}
func ChooseUsername(conn net.Conn) string {
	var username string
	var err error
	for {
		fmt.Println("Please enter your username:")
		fmt.Scanln(&username)
		err = Send(conn, Proto.Request{Command: viper.GetString("sign"), Username: username})
		if err != nil {
			continue
		}
		err = Accept(conn)
		if err != nil {
			continue
		}
		// Всё норм, имя прошло
		return username
	}
}
func ChooseStrategy(conn net.Conn, username string) {
	var err error
	// Выбираем стратегию
	var choose int
	for {
		fmt.Println("Выберите: \n 1. Сетевая игра \n Одиночная игра")
		fmt.Scanln(&choose)
		switch choose {
		case 1:
			err = Send(conn, Proto.Request{Command: viper.GetString("game_user"), Username: username})
			if err != nil {
				break
			}
			err = Accept(conn)
			if err != nil {
				break
			}
			return
		case 2:
			err = Send(conn, Proto.Request{Command: viper.GetString("game_server"), Username: username})
			if err != nil {
				break
			}
			err = Accept(conn)
			if err != nil {
				break
			}
			return
		default:
			fmt.Println("Выбор не распознан.Нужно выбрать пункт 1 или 2.")
		}
	}
}
