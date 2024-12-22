package transport

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net"
	"sync"
)

// CommandHandler - Тип для функции-обработчика
type CommandHandler func(conn net.Conn, args []string)

// Server - Структура сервера
type Server struct {
	handlers map[string]CommandHandler
	mu       sync.Mutex
}

// NewServer - Конструктор нового сервера
func NewServer() *Server {
	return &Server{
		handlers: make(map[string]CommandHandler),
	}
}

// RegisterHandler - Метод для регистрации обработчика команды
func (s *Server) RegisterHandler(command string, handler CommandHandler) {
	s.mu.Lock()
	s.handlers[command] = handler
	s.mu.Unlock()
}

// HandleCommand - Метод для обработки команды
func (s *Server) HandleCommand(conn net.Conn, command string, args []string) {
	if handler, ok := s.handlers[command]; ok {
		handler(conn, args)
	} else {
		fmt.Fprintf(conn, "Команда %s не найдена\n", command)
	}
}
a
func (s *Server) ListenAndServe() {
	listener, err := net.Listen("tcp", ":"+viper.GetString("port"))
	defer listener.Close()
	if err != nil {
		log.Printf("Error starting server: %v\n", err)
		return
	}
	log.Println("Server listening on port 80...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go handleClient(conn)
	}
}
