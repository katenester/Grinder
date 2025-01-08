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
		Id:         len(rooms.Rooms), // id комнаты (для первой комнаты -0 и т/д)
		Players:    players,
		StatusGame: models.Create,    // статус- игра создана
		IsNetwork:  len(players) > 1, // если игроков больше 1 (иначе игра с сервером)
		Board:      [16]int{},
		MillsBuilt: make(map[int][][]int),
	})
	return Protocol.Response{Cod: Protocol.StatusCreatedSuccessCode, Message: Protocol.RelateError(Protocol.StatusCreatedSuccessCode)}
}

func (rooms *RoomsMemory) GetRoomId(players []models.Player) (int, error) {
	// итерация по комнатам
	for _, room := range rooms.Rooms {
		// Если количество игроков не совпадает, это уже не та комната
		if len(players) != len(room.Players) {
			continue
		}
		// если один игрок
		if len(players) == 1 {
			if players[0] == room.Players[0] {
				return room.Id, nil
			}
		}
		// Если сопадают 2 игрока
		if (players[0] == room.Players[0] || players[0] == room.Players[1]) && (players[1] == room.Players[0] || players[1] == room.Players[1]) {
			return room.Id, nil
		}
	}
	return 0, Protocol.Response{Cod: Protocol.StatusNotFoundCode, Message: Protocol.RelateError(Protocol.StatusNotFoundCode)}
}

// GetBoard - получить доску
func (rooms *RoomsMemory) GetBoard(idRoom int) [16]int {
	return rooms.Rooms[idRoom].Board
}

// SetBoard - установить новую доску
func (rooms *RoomsMemory) SetBoard(idRoom int, newBoard [16]int) {
	rooms.Mutex.Lock()
	rooms.Rooms[idRoom].Board = newBoard
	rooms.Mutex.Unlock()
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
