package repository

import (
	"Grinder/Protocol"
	"Grinder/server/internal/models"
	"fmt"
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

func (rooms *RoomsMemory) CreateRoom(players []models.Player) (int, error) {
	//if len(rooms.Rooms) > 0 {
	//	return Protocol.Response{Cod: Protocol.StatusInternalServerErrorCode, Message: Protocol.RelateError(Protocol.StatusInternalServerErrorCode)}
	//}
	rooms.Mutex.Lock()
	defer rooms.Mutex.Unlock()
	// Если игрок 1 то создаём игрока - сервер
	if len(players) == 1 {
		players = append(players, models.Player{Name: "server" + fmt.Sprintf("%d", len(rooms.Rooms))})
	}
	rooms.Rooms = append(rooms.Rooms, models.Room{
		Id:         len(rooms.Rooms), // id комнаты (для первой комнаты -0 и т/д)
		Players:    players,
		StatusGame: models.Create, // статус- игра создана
		Board:      [16]int{},
		MillsBuilt: make(map[int][][]int),
	})
	// Возвращаем id
	return len(rooms.Rooms), Protocol.Response{Cod: Protocol.StatusCreatedSuccessCode, Message: Protocol.RelateError(Protocol.StatusCreatedSuccessCode)}
}

//func (rooms *RoomsMemory) GetRoomId(players []models.Player) (int, error) {
//	// итерация по комнатам
//	for _, room := range rooms.Rooms {
//		// если один игрок
//		if len(players) == 1 {
//			if players[0] == room.Players[0] {
//				return room.Id, nil
//			}
//		}
//		// Если сопадают 2 игрока
//		if (players[0] == room.Players[0] || players[0] == room.Players[1]) && (players[1] == room.Players[0] || players[1] == room.Players[1]) {
//			return room.Id, nil
//		}
//	}
//	return 0, Protocol.Response{Cod: Protocol.StatusNotFoundCode, Message: Protocol.RelateError(Protocol.StatusNotFoundCode)}
//}

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

func (rooms *RoomsMemory) GetMillsBuilt(idRoom int, currentPlayer int) [][]int {
	return rooms.Rooms[idRoom].MillsBuilt[currentPlayer]
}

func (rooms *RoomsMemory) AppendMillsBuilt(idRoom int, currentPlayer int, mill []int) {
	rooms.Mutex.Lock()
	defer rooms.Mutex.Unlock()
	rooms.Rooms[idRoom].MillsBuilt[currentPlayer] = append(rooms.Rooms[idRoom].MillsBuilt[currentPlayer], mill) // Добавляем построенную мельницу в список
}

func (rooms *RoomsMemory) GetPlayer(idRoom int) []models.Player {
	return rooms.Rooms[idRoom].Players
}

//func (rooms *RoomsMemory) GetRoom(idRoom int) *[16]int {
//	return &rooms.Rooms[idRoom].Board
//}
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
