package repository

import (
	"Grinder/server/internal/models"
	"errors"
	"sync"
)

type RoomsMemory struct {
	rooms []models.Room
	mutex sync.Mutex
}

//func NewRooms() Rooms {
//	return &RoomsMemory{}
//}

func (rooms *RoomsMemory) CreateRoom(username string) error {
	return errors.New("Not implemented")
}

func (rooms *RoomsMemory) JoinRoom(username string) error {
	return errors.New("Not implemented")
}
