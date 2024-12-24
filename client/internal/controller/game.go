package controller

import (
	"Grinder/Proto"
	"encoding/json"
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/spf13/viper"
	"log"
	"net"
)

func StartGame(conn net.Conn) {
	// Логинимся
	username := ChooseUsername(conn)
	// Выбираем стратегию
	ChooseStrategy(conn, username)
	// Ход игры
	ch := make(chan struct{})
	go Exit(ch)
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
func Accept(conn net.Conn) (Proto.Response, error) {
	var resp Proto.Response
	// Чтение данных с вервера
	decoder := json.NewDecoder(conn)
	err := decoder.Decode(&resp)
	// Ошибка при декодировании
	if err != nil || resp.Cod != 200 {
		log.Print(err.Error(), resp)
		return Proto.Response{}, err
	}
	return resp, nil
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
		_, err = Accept(conn)
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
			_, err = Accept(conn)
			if err != nil {
				break
			}
			return
		case 2:
			err = Send(conn, Proto.Request{Command: viper.GetString("game_server"), Username: username})
			if err != nil {
				break
			}
			_, err = Accept(conn)
			if err != nil {
				break
			}
			return
		default:
			fmt.Println("Выбор не распознан.Нужно выбрать пункт 1 или 2.")
		}
	}
}

func MakeMove(conn net.Conn, username string) {
	err := Send(conn, Proto.Request{Command: viper.GetString("MakeMove"), Username: username})
	if err != nil {
		return
	}
	var resp Proto.Response
	resp, err = Accept(conn)
	if err != nil {
		return
	}
	fmt.Println(resp)
	return
}
func TakeChips(conn net.Conn, username string) {
	err := Send(conn, Proto.Request{Command: viper.GetString("TakeChips"), Username: username})
	if err != nil {
		return
	}
	var resp Proto.Response
	resp, err = Accept(conn)
	if err != nil {
		return
	}
	fmt.Println(resp)
	return
}
func MoveChips(conn net.Conn, username string) {
	err := Send(conn, Proto.Request{Command: viper.GetString("MoveChips"), Username: username})
	if err != nil {
		return
	}
	var resp Proto.Response
	resp, err = Accept(conn)
	if err != nil {
		return
	}
	fmt.Println(resp)
	return
}
func GetTopScores(conn net.Conn, username string) {
	err := Send(conn, Proto.Request{Command: viper.GetString("top"), Username: username})
	if err != nil {
		return
	}
	var resp Proto.Response
	resp, err = Accept(conn)
	if err != nil {
		return
	}
	fmt.Println(resp)
	return
}
func Exit(ch chan struct{}) {
	// Открываем клавиатуру
	if err := keyboard.Open(); err != nil {
		fmt.Println("Ошибка при открытии клавиатуры:", err)
		return
	}
	defer keyboard.Close()

	for {
		// Чтение нажатой клавиши
		_, key, err := keyboard.GetKey()о
		if err != nil {
			fmt.Println("Ошибка при чтении клавиши:", err)
			return
		}

		// Если нажата клавиша F1
		if key == keyboard.KeyF1 {
			// Отправляем пустую структуру в канал(завершение игры)
			ch <- struct{}{}
			close(ch)
			fmt.Println("F1 нажата, отправлено сообщение в канал.")
		}
	}
}
