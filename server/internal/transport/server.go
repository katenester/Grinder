package transport

import (
	"Grinder/Protocol"
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"net"
	"sync"
)

// CommandHandler - Тип для функции-обработчика
type CommandHandler func(conn net.Conn, req Protocol.Request)

// Server - Структура сервера
type Server struct {
	handlers map[string]CommandHandler
	handle   *Handler // обработчик путей
	mu       sync.Mutex
}

// NewServer - Конструктор нового сервера
func NewServer(handle *Handler) *Server {
	return &Server{
		handlers: make(map[string]CommandHandler),
		handle:   handle,
	}
}

// RegisterHandler - Метод для регистрации обработчика команды
func (s *Server) RegisterHandler(command string, handler CommandHandler) {
	s.mu.Lock()
	s.handlers[command] = handler
	s.mu.Unlock()
}

// HandleCommand - Метод для обработки команды
func (s *Server) HandleCommand(conn net.Conn, req Protocol.Request, command string) {
	if handler, ok := s.handlers[command]; ok {
		handler(conn, req)
	} else {
		s.handle.sendResponse(conn, Protocol.Response{Cod: Protocol.StatusNotFoundCode, Message: Protocol.RelateError(Protocol.StatusNotFoundCode)})
	}
}
func (s *Server) Run() {
	listener, err := net.Listen("tcp", ":"+viper.GetString("port"))
	defer listener.Close()
	if err != nil {
		log.Printf("Error starting server: %v\n", err)
		return
	}
	log.Println("Server listening on port 80...")
	// Регистрируем пути
	s.InitRouter()
	for {
		// Слушаем соединения
		conn, err := listener.Accept()
		//log.Print("conn: ", conn, "com:", s.getCommand(conn))
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err)
			continue
		}
		//var req Protocol.Request
		//// Чтение данных с вервера
		//decoder := json.NewDecoder(conn)
		//_ = decoder.Decode(&req)
		//log.Print("conn: ", conn, "com2:", req)
		// Запускаем выполнение команды полученной из запроса клиента
		go func(conn net.Conn) {
			for {
				req, err := s.getRequest(conn)
				if err != nil {
					s.handle.sendResponse(conn, err.(Protocol.Response))
				}
				go s.HandleCommand(conn, req, req.Command)
			}
		}(conn)
	}
}

//func (h *Handler) sendResponse(conn net.Conn, resp Protocol.Response) {
//	log.Print("sendResponse: ", resp)
//	// Отправляем клиенту
//	encoder := json.NewEncoder(conn)
//	errNew := encoder.Encode(resp)
//	if errNew != nil {
//		log.Println(errNew.Error())
//	}
//}

// Получить запрос
func (s *Server) getRequest(conn net.Conn) (Protocol.Request, error) {
	var req Protocol.Request
	// Чтение данных с вервера
	decoder := json.NewDecoder(conn)
	err := decoder.Decode(&req)
	log.Println("getRequest: ", req)
	// Ошибка при декодировании
	if err != nil {
		log.Print(err.Error(), req)
		return Protocol.Request{}, Protocol.Response{Cod: Protocol.StatusBadRequestCode, Message: Protocol.RelateError(Protocol.StatusBadRequestCode)}
	}
	return req, nil
}

func (s *Server) InitRouter() {
	s.RegisterHandler(viper.GetString("command.sign"), s.handle.Sign)
	s.RegisterHandler(viper.GetString("command.game_user"), s.handle.GameUser)
	s.RegisterHandler(viper.GetString("command.game_server"), s.handle.GameServer)
	s.RegisterHandler(viper.GetString("command.MakeMove"), s.handle.MakeMove)
	s.RegisterHandler(viper.GetString("command.TakeChips"), s.handle.TakeChips)
	s.RegisterHandler(viper.GetString("command.MoveChips"), s.handle.MoveChips)
	s.RegisterHandler(viper.GetString("command.top"), s.handle.GetTop)
	s.RegisterHandler(viper.GetString("command.exit"), s.handle.Exit)
}

//func (s *Server) send(req Protocol.Request) error {
//	// Отправляем серверу json
//	encoder := json.NewEncoder(c.conn)
//	err := encoder.Encode(req)
//	if err != nil {
//		log.Print(err.Error())
//	}
//	return err
//}
//func (s *Server) accept() (Protocol.Response, error) {
//	var resp Protocol.Response
//	// Чтение данных с вервера
//	decoder := json.NewDecoder(c.conn)
//	err := decoder.Decode(&resp)
//	// Ошибка при декодировании
//	if err != nil || resp.Cod != 200 {
//		log.Print(err.Error(), resp)
//		return Protocol.Response{}, err
//	}
//	return resp, nil
//}
