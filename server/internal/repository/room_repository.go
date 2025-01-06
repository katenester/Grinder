package repository

import (
	"Grinder/Protocol"
	"Grinder/server/internal/models"
	"sync"
)

type RoomsMemory struct {
	rooms []models.Room
	mutex sync.Mutex
}

//func NewRooms() Rooms {
//	return &RoomsMemory{}
//}

func NewRoomsMemory() *RoomsMemory {
	return &RoomsMemory{
		rooms: make([]models.Room, 0),
	}
}

func (rooms *RoomsMemory) CreateRoom(players []models.Player) error {
	if len(rooms.rooms) > 0 {
		return Protocol.Response{Cod: Protocol.StatusInternalServerErrorCode, Message: Protocol.RelateError(Protocol.StatusInternalServerErrorCode)}
	}
	rooms.mutex.Lock()
	defer rooms.mutex.Unlock()
	rooms.rooms = append(rooms.rooms, models.Room{
		Players:    players,
		IsNetwork:  len(players) > 1, // если игроков больше 1 (иначе игра с сервером)
		Board:      [16]int{},
		MillsBuilt: make(map[int][][]int),
	})
	return Protocol.Response{Cod: Protocol.StatusCreatedSuccessCode, Message: Protocol.RelateError(Protocol.StatusCreatedSuccessCode)}
}
