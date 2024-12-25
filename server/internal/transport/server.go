package transport

import (
	"Grinder/Protocol"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net"
	"sync"
)

// CommandHandler - Тип для функции-обработчика
type CommandHandler func(conn net.Conn)

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
func (s *Server) HandleCommand(conn net.Conn, command string) {
	if handler, ok := s.handlers[command]; ok {
		handler(conn)
	} else {
		fmt.Fprintf(conn, "Команда %s не найдена\n", command)
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
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err)
			continue
		}
		// Запускаем выполнение команды полученной из запроса клиента
		go s.HandleCommand(conn, s.getCommand(conn))
	}
}

// Получить команду из запроса
func (s *Server) getCommand(conn net.Conn) string {
	var req Protocol.Request
	// Чтение данных с вервера
	decoder := json.NewDecoder(conn)
	err := decoder.Decode(&req)
	// Ошибка при декодировании
	if err != nil {
		log.Print(err.Error(), req)
		return ""
	}
	return req.Command
}

func (s *Server) InitRouter() {
	s.RegisterHandler(viper.GetString("sign"), s.handle.Sign)
	s.RegisterHandler(viper.GetString("game_user"), s.handle.GameUser)
	s.RegisterHandler(viper.GetString("game_server"), s.handle.GameServer)
	s.RegisterHandler(viper.GetString("MakeMove"), s.handle.MakeMove)
	s.RegisterHandler(viper.GetString("TakeChips"), s.handle.TakeChips)
	s.RegisterHandler(viper.GetString("MoveChips"), s.handle.MoveChips)
	s.RegisterHandler(viper.GetString("top"), s.handle.GetTop)
	s.RegisterHandler(viper.GetString("exit"), s.handle.Exit)
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
