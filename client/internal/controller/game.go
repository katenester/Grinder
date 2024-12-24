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
	var username string
	// Отправка имени
	for {
		fmt.Println("Please enter your username:")
		fmt.Scanln(&username)
		// Создаем команду
		resp := Proto.Request{Command: viper.GetString("sign"), Username: username}
		// Отправляем серверу json
		encoder := json.NewEncoder(conn)
		err := encoder.Encode(resp)
		if err != nil {
			log.Print(err.Error())
			continue
		}
		// Чтение данных с вервера
		var msg Proto.Response
		decoder := json.NewDecoder(conn)
		err = decoder.Decode(&msg)
		// Ошибка при декодировании
		if err != nil {
			log.Print(err.Error())
			continue
		}
		if msg.Cod != 200 {
			fmt.Println(msg, "Попробуйте ещё раз")
		}
	}
}
