package models

import (
	"encoding/json"
	"log"
	"net"
)

type Player struct {
	Name        string   `json:"name"`         // Игровое имя
	Score       int      `json:"score"`        // Счёт
	Conn        net.Conn `json:"-"`            // Подключение
	IsConnected bool     `json:"is-connected"` // Проверка подключения
}

func (p *Player) SendPlayer(message string) error {
	dat, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = p.Conn.Write(dat)
	if err != nil {
		log.Println(err)
		return err
	}
}

func (p *Player) ReceivePlayer() (string, error) {

}
