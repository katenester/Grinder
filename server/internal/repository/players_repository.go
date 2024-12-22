package repository

import (
	"Grinder/server/internal/models"
	"errors"
	"sync"
)

type PlayersMemory struct {
	players []models.Player
	mu      sync.RWMutex
}

func NewPlayers() Players {
	return &PlayersMemory{}
}

func (p *PlayersMemory) CreatePlayer(username string) error {
	return errors.New("Not implemented")
}
func (p *PlayersMemory) GetPlayer(username string) error {
	return errors.New("Not implemented")
}
