package repository

type Players interface {
	CreatePlayer(username string) error
	GetPlayer(username string) error
}

type Rooms interface {
	CreateRoom(username string) error
	JoinRoom(username string) error
}
