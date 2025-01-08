package repository

import (
	"Grinder/Protocol"
	"Grinder/server/internal/models"
	"sync"
)

type RoomsMemory struct {
	Rooms []models.Room
	Mutex sync.Mutex
}

//func NewRooms() Rooms {
//	return &RoomsMemory{}
//}

func NewRoomsMemory() *RoomsMemory {
	return &RoomsMemory{
		Rooms: make([]models.Room, 0),
	}
}

func (rooms *RoomsMemory) CreateRoom(players []models.Player) error {
	if len(rooms.Rooms) > 0 {
		return Protocol.Response{Cod: Protocol.StatusInternalServerErrorCode, Message: Protocol.RelateError(Protocol.StatusInternalServerErrorCode)}
	}
	rooms.Mutex.Lock()
	defer rooms.Mutex.Unlock()
	rooms.Rooms = append(rooms.Rooms, models.Room{
		Players:    players,
		StatusGame: models.Create,    // статус- игра создана
		IsNetwork:  len(players) > 1, // если игроков больше 1 (иначе игра с сервером)
		Board:      [16]int{},
		MillsBuilt: make(map[int][][]int),
	})
	return Protocol.Response{Cod: Protocol.StatusCreatedSuccessCode, Message: Protocol.RelateError(Protocol.StatusCreatedSuccessCode)}
}

// GetModeInit - получение режима игры (сетевая или нет)
//func (rooms *RoomsMemory) GetModeInit(user models.Player)(int,error){
//	// итерация по комнатам
//	for _, room := range rooms.Rooms {
//		// итерация по игрокам
//		for _, player := range room.Players {
//			// Если нашли нужного игрока
//			if player==user{
//				// если сетевая игра
//				if len(room.Players)>1{
//					return 1,nil
//				}
//				// если игра с сервером
//				return 2,nil
//			}
//		}
//	}
//	return 0, Protocol.Response{Cod: Protocol.StatusNotFoundCode, Message: Protocol.RelateError(Protocol.StatusNotFoundCode)}
//}
