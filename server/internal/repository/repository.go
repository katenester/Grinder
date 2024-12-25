package repository

import "net"

type Players interface {
	CreatePlayer(username string, conn net.Conn) error
}

type Repository struct {
	Players
}

func NewRepository() *Repository {
	return &Repository{
		Players: NewPlayersMemory(),
	}
}
