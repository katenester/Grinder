package models

import "net"

type Player struct {
	name        string   // Игровое имя
	Score       int      // Счёт
	conn        net.Conn // Подключение
	isConnected bool     // Проверка подключения
}
