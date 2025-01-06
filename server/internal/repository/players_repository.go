package repository

import (
	"Grinder/Protocol"
	"Grinder/server/internal/models"
	"encoding/json"
	"errors"
	"log"
	"net"
	"sort"
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
func (p *PlayersMemory) GetTop(conn net.Conn, req Protocol.Request) error {
	// Копируем исходный слайс
	p.mu.Lock()
	sortedPlayers := append([]models.Player(nil), p.players...)
	defer p.mu.Unlock()
	// Сортируем копию слайса
	sort.Slice(sortedPlayers, func(i, j int) bool {
		return sortedPlayers[i].Score > sortedPlayers[j].Score
	})
	// Преобразуем слайс игроков в JSON
	data, err := json.Marshal(sortedPlayers)
	log.Println(sortedPlayers)
	log.Println(data)
	if err != nil {
		return Protocol.Response{Cod: Protocol.StatusInternalServerErrorCode, Message: Protocol.RelateError(Protocol.StatusInternalServerErrorCode)}
	}
	return Protocol.Response{Cod: Protocol.StatusSuccessCode, Message: Protocol.RelateError(Protocol.StatusSuccessCode), Body: data}
}
func (p *PlayersMemory) Exit(conn net.Conn, req Protocol.Request) error {
	// Нужно очистить поле для игроков, закрыть игру(не conn)
	return errors.New("Not Implemented")
}

func (p *PlayersMemory) GetUser(username string) (models.Player, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	for _, player := range p.players {
		if player.Name == username {
			if player.IsConnected {
				return player, nil
			}
			// Если пользователь отключен
			return models.Player{}, Protocol.Response{Cod: Protocol.StatusConflictCode, Message: Protocol.RelateError(Protocol.StatusConflictCode)}
		}
	}
	return models.Player{}, Protocol.Response{Cod: Protocol.StatusNotFoundCode, Message: Protocol.RelateError(Protocol.StatusNotFoundCode)}
}
