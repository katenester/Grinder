package service

import (
	"Grinder/server/internal/repository"
	"net"
)

type PlayersService struct {
	repo repository.Players
}

func NewPlayersService(repo repository.Players) *PlayersService {
	return &PlayersService{
		repo: repo,
	}
}

func (p *PlayersService) CreatePlayer(username string, conn net.Conn) error {
	return p.repo.CreatePlayer(username, conn)
}
