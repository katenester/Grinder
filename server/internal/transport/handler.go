package transport

import (
	"Grinder/Protocol"
	"Grinder/server/internal/service"
	"encoding/json"
	"log"
	"net"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
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
	encoder := json.NewEncoder(conn)
	errNew := encoder.Encode(resp)
	if errNew != nil {
		log.Println(errNew.Error())
	}
}
func (h *Handler) Sign(conn net.Conn, req Protocol.Request) {
	// Выполянем создание/ вход пользователя и отправляем response
	h.sendResponse(conn, h.service.Players.CreatePlayer(req.Username, conn).(Protocol.Response))
}
func (h *Handler) GameUser(conn net.Conn, req Protocol.Request) {

}
func (h *Handler) GameServer(conn net.Conn, req Protocol.Request) {

}
func (h *Handler) MakeMove(conn net.Conn, req Protocol.Request) {

}
func (h *Handler) TakeChips(conn net.Conn, req Protocol.Request) {

}
func (h *Handler) MoveChips(conn net.Conn, req Protocol.Request) {

}
func (h *Handler) GetTop(conn net.Conn, req Protocol.Request) {
	//h.sendResponse(conn, h.service.Players.CreatePlayer(req.Username, conn).(Protocol.Response))
	log.Println("metka")
	h.sendResponse(conn, h.service.Players.GetTop(conn, req).(Protocol.Response))
}

func (h *Handler) Exit(conn net.Conn, req Protocol.Request) {
	h.sendResponse(conn, h.service.Players.Exit(conn, req).(Protocol.Response))
}
