package service

import (
	"Grinder/Protocol"
	"Grinder/server/internal/repository"
	"net"
)

type Players interface {
	CreatePlayer(username string, conn net.Conn) error
	GetTop(conn net.Conn, req Protocol.Request) error
	Exit(conn net.Conn, req Protocol.Request) error
}
type Service struct {
	Players
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Players: NewPlayersService(repo.Players),
	}
}
