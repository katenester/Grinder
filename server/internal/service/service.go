package service

import (
	"Grinder/Protocol"
	"Grinder/server/internal/models"
	"Grinder/server/internal/repository"
	"net"
)

type Players interface {
	CreatePlayer(username string, conn net.Conn) error
	GetTop(conn net.Conn, req Protocol.Request) error
	Exit(conn net.Conn, req Protocol.Request) error
	GetUser(username string) (models.Player, error)
}

type Game interface {
	CreateRoom(players []models.Player) error
}
type Service struct {
	Players
	Game
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Players: NewPlayersService(repo.Players),
		Game:    NewGameService(repo.Game),
	}
}
