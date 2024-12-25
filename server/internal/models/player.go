package models

import "net"

type Player struct {
	Name        string   // Игровое имя
	Score       int      // Счёт
	Conn        net.Conn // Подключение
	IsConnected bool     // Проверка подключения
}
