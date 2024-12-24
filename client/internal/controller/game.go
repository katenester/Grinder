package controller

import (
	"Grinder/client/internal/models"
	"fmt"
	"github.com/eiannone/keyboard"
)

func StartGame(client *models.Client) {
	// Логинимся
	client.ChooseUsername()
	// Выбираем стратегию
	client.ChooseStrategy()
	// Ход игры
	ch := make(chan struct{})
	go Exit(ch, client)
}
func Exit(ch chan struct{}, client *models.Client) {
	// Открываем клавиатуру
	if err := keyboard.Open(); err != nil {
		fmt.Println("Ошибка при открытии клавиатуры:", err)
		return
	}
	defer keyboard.Close()

	for {
		// Чтение нажатой клавиши
		_, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Ошибка при чтении клавиши:", err)
			return
		}

		// Если нажата клавиша F1
		if key == keyboard.KeyF1 {
			client.Exit()
			// Отправляем пустую структуру в канал(завершение игры)
			ch <- struct{}{}
			close(ch)
			fmt.Println("F1 нажата, отправлено сообщение в канал. Конец игры. Т	Ы СДАЛСЯ")
		}
	}
}
