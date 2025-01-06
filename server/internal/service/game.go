package service

import (
	"Grinder/server/internal/models"
	"Grinder/server/internal/repository"
)

type GameService struct {
	repo repository.Game
}

func NewGameService(repo repository.Game) *GameService {
	return &GameService{
		repo: repo,
	}
}

func (g *GameService) CreateRoom(players []models.Player) error {
	return g.repo.CreateRoom(players)
}
