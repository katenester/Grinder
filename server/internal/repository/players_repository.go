package repository

import (
	"Grinder/Protocol"
	"Grinder/server/internal/models"
	"net"
	"sync"
)

// PlayersMemory - Хранение в оперативке игроков
type PlayersMemory struct {
	players []models.Player
	mu      sync.RWMutex
}

func NewPlayersMemory() *PlayersMemory {
	return &PlayersMemory{
		players: make([]models.Player, 0),
	}
}

func (p *PlayersMemory) CreatePlayer(username string, conn net.Conn) error {
	for id, player := range p.players {
		// если игрок найден
		if username == player.Name {
			// Проверка на подключение к серверу
			switch player.IsConnected {
			// Если подключен
			case true:
				return Protocol.Response{Cod: Protocol.StatusConflictCode, Message: Protocol.RelateError(Protocol.StatusConflictCode)}
			case false:
				p.players[id].IsConnected = true
				return Protocol.Response{Cod: Protocol.StatusSuccessCode, Message: Protocol.RelateError(Protocol.StatusSuccessCode)}
			}
		}
	}
	// Если игрок не найден => добавляем
	p.mu.Lock()
	p.players = append(p.players, models.Player{Name: username, Conn: conn, IsConnected: true})
	defer p.mu.Unlock()
	return Protocol.Response{Cod: Protocol.StatusCreatedSuccessCode, Message: Protocol.RelateError(Protocol.StatusCreatedSuccessCode)}
}
