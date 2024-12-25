package repository

import (
	"Grinder/Protocol"
	"net"
)

type Players interface {
	CreatePlayer(username string, conn net.Conn) error
	GetTop(conn net.Conn, req Protocol.Request) error
	Exit(conn net.Conn, req Protocol.Request) error
}

type Repository struct {
	Players
}

func NewRepository() *Repository {
	return &Repository{
		Players: NewPlayersMemory(),
	}
}
