package repository

import (
	"Grinder/Protocol"
	"Grinder/server/internal/models"
	"net"
)

type Players interface {
	CreatePlayer(username string, conn net.Conn) error
	GetTop(conn net.Conn, req Protocol.Request) error
	Exit(conn net.Conn, req Protocol.Request) error
	GetUser(username string) (models.Player, error)
}
type Game interface {
	CreateRoom(players []models.Player) (int, error)
	// GetRoomId(players []models.Player) (int, error)
	GetBoard(idRoom int) [16]int
	SetBoard(idRoom int, newBoard [16]int)
	GetMillsBuilt(idRoom int, currentPlayer int) [][]int
	AppendMillsBuilt(idRoom int, currentPlayer int, mill []int)
	GetPlayer(idRoom int) []models.Player
	//GetModeInit(user models.Player)(int,error)
}
type Repository struct {
	Players
	Game
}

func NewRepository() *Repository {
	return &Repository{
		Players: NewPlayersMemory(),
		Game:    NewRoomsMemory(),
	}
}
