package service

import (
	"Grinder/server/internal/repository"
	"net"
)

type Players interface {
	CreatePlayer(username string, conn net.Conn) error
}
type Service struct {
	Players
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Players: NewPlayersService(repo.Players),
	}
}
