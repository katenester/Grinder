package models

import (
	"net"
)

type Player struct {
	Name        string   `json:"name"`         // Игровое имя
	Score       int      `json:"score"`        // Счёт
	Conn        net.Conn `json:"-"`            // Подключение
	IsConnected bool     `json:"is-connected"` // Проверка подключения
}
