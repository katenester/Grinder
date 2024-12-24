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
	var req Proto.Request
	var resp Proto.Response
	// Логинимся
	var username string
	// Отправка имени
	for {
		fmt.Println("Please enter your username:")
		fmt.Scanln(&username)
		// Создаем команду
		req = Proto.Request{Command: viper.GetString("sign"), Username: username}
		// Отправляем серверу json
		encoder := json.NewEncoder(conn)
		err := encoder.Encode(req)
		if err != nil {
			log.Print(err.Error())
			continue
		}
		// Чтение данных с вервера
		decoder := json.NewDecoder(conn)
		err = decoder.Decode(&resp)
		// Ошибка при декодировании
		if err != nil {
			log.Print(err.Error())
			continue
		}
		if resp.Cod != 200 {
			fmt.Println(resp, "Попробуйте ещё раз")
		}
		// Всё норм, имя прошло
		break
	}

	// Выбираем стратегию
	for {
	var choose int
Loop:
	for {
		fmt.Println("Выберите: \n 1. Сетевая игра \n Одиночная игра")
		fmt.Scanln(&choose)
		switch choose {
		case 1:
			req = Proto.Request{Command: viper.GetString("game_user"), Username: username}
			break Loop
		case 2:
			req = Proto.Request{Command: viper.GetString("game_server"), Username: username}
			break Loop
		default:
			fmt.Println("Выбор не распознан.Нужно выбрать пункт 1 или 2.")
		}
	}
	// Отправляем серверу json
	encoder := json.NewEncoder(conn)
	err := encoder.Encode(req)
	if err != nil {
		log.Print(err.Error())

	}
	// Чтение данных с вервера
	decoder := json.NewDecoder(conn)
	err = decoder.Decode(&resp)
	// Ошибка при декодировании
	if err != nil {
		log.Print(err.Error())
		continue
	}
	if resp.Cod != 200 {
		fmt.Println(resp, "Попробуйте ещё раз")
	}
	// Всё норм, дейсвие прошло
	break
}
