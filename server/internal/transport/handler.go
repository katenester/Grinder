package transport

import (
	"Grinder/Protocol"
	"Grinder/server/internal/models"
	"Grinder/server/internal/service"
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"
)

type Handler struct {
	service *service.Service
	mu      sync.Mutex
	queue   map[models.Player]chan struct{}
	ch      chan models.Player
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
		queue:   make(map[models.Player]chan struct{}),
		ch:      make(chan models.Player),
	}
}

//func (h *Handler) InitRouter() {
//	srv.RegisterHandler(viper.GetString("sign"), h.Sign)
//	srv.RegisterHandler(viper.GetString("game_user"), h.GameUser)
//	srv.RegisterHandler(viper.GetString("game_server"), h.GameServer)
//	srv.RegisterHandler(viper.GetString("MakeMove"), h.MakeMove)
//	srv.RegisterHandler(viper.GetString("TakeChips"), h.TakeChips)
//	srv.RegisterHandler(viper.GetString("MoveChips"), h.MoveChips)
//	srv.RegisterHandler(viper.GetString("top"), h.GetTop)
//	srv.RegisterHandler(viper.GetString("exit"), h.Exit)
//}

//	func sendToClient(conn net.Conn){
//		log.Print(err.Error(), req)
//		encoder := json.NewEncoder(conn)
//		err := encoder.Encode(req)
//	}
//func (h *Handler) sendErrorResponse(conn net.Conn, code int) {
//	// Собираем ответ
//	resp := Protocol.Response{Cod: code, Message: Errors.RelateError(code)}
//	log.Print("sendResponse: ", resp)
//	// Отправляем клиенту
//	encoder := json.NewEncoder(conn)
//	errNew := encoder.Encode(resp)
//	if errNew != nil {
//		log.Println(errNew.Error())
//	}
//}

func (h *Handler) sendResponse(conn net.Conn, resp Protocol.Response) {
	log.Println("sendResponse: ", resp.Cod, resp.Message)
	// Отправляем клиенту
	//encoder := json.NewEncoder(conn)
	//errNew := encoder.Encode(resp)
	//if errNew != nil {
	//	log.Println(errNew.Error())
	//}
	dat, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Write(dat)
	if err != nil {
		log.Println(err)
	}
	// Отправка body (полезной нагрузки)
	if resp.Body != nil {
		data := resp.Body
		log.Println(data)
		// Записываем тело
		_, err = conn.Write(data)
	}
}
func (h *Handler) Sign(conn net.Conn, req Protocol.Request) {
	// Выполянем создание/ вход пользователя и отправляем response
	h.sendResponse(conn, h.service.Players.CreatePlayer(req.Username, conn).(Protocol.Response))
}
func (h *Handler) GameUser(conn net.Conn, req Protocol.Request) {
	// Нужен пользователь
	user, err := h.service.GetUser(req.Username)
	// Если пользователь найден или неактивен
	if err != nil {
		h.sendResponse(conn, err.(Protocol.Response))
	}
	// Добавляем пользователя в очередь
	h.mu.Lock()
	h.queue[user] = make(chan struct{})
	h.mu.Unlock()
	quit := make(chan struct{})
	select {
	// Если нашли игрока
	case players := <-h.findUser(user, quit):
		log.Println("Игрока нашли для игры")
		h.sendResponse(conn, h.service.Game.CreateRoom([]models.Player{user, players}).(Protocol.Response))
	// Если дедлайн
	case <-time.After(timeFindUser):
		quit <- struct{}{}
		// Удаляем пользователя из очереди
		h.mu.Lock()
		delete(h.queue, user)
		h.mu.Unlock()
		h.sendResponse(conn, Protocol.Response{Cod: Protocol.StatusTimeOutCode, Message: Protocol.StatusTimeOut})
	// Если текущего пользователя забрал другой пользователь
	case <-h.queue[user]:
		quit <- struct{}{}
		log.Println("Игрока взяли в игру")
		// Игра началась
		//h.sendResponse(conn, h.service.Game.CreateRoom([]models.Player{user, players}).(Protocol.Response))
	}

	//// Канал для отмены дедлайна
	//removeUser := make(chan struct{})
	//// Добавляем пользователя в очередь и запускаем таймер для удаления, если его не взяли
	//go func() {
	//	select {
	//	// Если прошло время ожидания в очереди
	//	case <-time.After(timeFindUser):
	//		fmt.Println("User removed from queue due to timeout:", user.Name)
	//		<-h.ch // Убираем его из канала, если за 3 минуты его никто не взял
	//	case <-removeUser:
	//		// Если пользователь был взят из канала до истечения времени
	//		fmt.Println("User was taken from queue:", user.Name)
	//	}
	//}()
	//for{
	//	select {
	//	// либо ждём записи в канал (то есть ставим текущего чела в очередь) (и подождём пока его возьмут)
	//	case h.ch <- user:
	//		fmt.Println("User placed in queue:", user.Name)
	//	// Либо ждём пока появится чел в очереди
	//	case user2 := <-h.ch:
	//		if user.Name == user2.Name {
	//			// Если это тот же пользователь, продолжаем ожидание
	//			fmt.Println("User tried to take themselves from queue, waiting for another:", user.Name)
	//			continue
	//		}
	//		// Если это другой пользователь => создаём комнату
	//		h.sendResponse(conn, h.service.Game.CreateRoom([]models.Player{user, user2}).(Protocol.Response))
	//		break
	//	// Либо таймаут ожидания
	//	case <-time.After(timeFindUser):
	//		// Сообщаем, что пользователь не был взят
	//		removeUser <- struct{}{}
	//		h.sendResponse(conn, Protocol.Response{Cod: Protocol.StatusTimeOutCode, Message: Protocol.StatusTimeOut})
	//		break
	//	}
	//}

}
func (h *Handler) findUser(user models.Player, quit chan struct{}) chan models.Player {
	result := make(chan models.Player)
	go func() {
		for {
			for player, _ := range h.queue {
				select {
				case <-quit:
					// Кидаем дефолтного игрока для освобождения ресурсов
					result <- models.Player{}
					return
				default:
					// Если нашли готового игрока
					if user != player {
						h.mu.Lock()
						// Делаем оповещение для другого пользователя
						h.queue[player] <- struct{}{}
						// Удаляем игроков из очереди
						delete(h.queue, user)
						delete(h.queue, player)
						h.mu.Unlock()
						result <- player
						return
					}
				}
			}
		}
	}()
	return result
}
func (h *Handler) GameServer(conn net.Conn, req Protocol.Request) {
	user, err := h.service.GetUser(req.Username)
	// Если пользователь найден или неактивен
	if err != nil {
		h.sendResponse(conn, err.(Protocol.Response))
	}
	h.sendResponse(conn, h.service.Game.CreateRoom([]models.Player{user}).(Protocol.Response))
}
func (h *Handler) MakeMove(conn net.Conn, req Protocol.Request) {

}
func (h *Handler) TakeChips(conn net.Conn, req Protocol.Request) {

}
func (h *Handler) MoveChips(conn net.Conn, req Protocol.Request) {

}
func (h *Handler) GetTop(conn net.Conn, req Protocol.Request) {
	//h.sendResponse(conn, h.service.Players.CreatePlayer(req.Username, conn).(Protocol.Response))
	h.sendResponse(conn, h.service.Players.GetTop(conn, req).(Protocol.Response))
}

func (h *Handler) Exit(conn net.Conn, req Protocol.Request) {
	h.sendResponse(conn, h.service.Players.Exit(conn, req).(Protocol.Response))
}
