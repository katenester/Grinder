package transport

import (
	"github.com/spf13/viper"
	"net"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRouter(srv Server) {
	srv.RegisterHandler(viper.GetString("sign"), h.Sign)
	srv.RegisterHandler(viper.GetString("create"), h.CreateRoom)
	srv.RegisterHandler(viper.GetString("join"), h.JoinRoom)
	srv.RegisterHandler(viper.GetString("start"), h.StartGame)
	srv.RegisterHandler(viper.GetString("game"), h.Play)
	srv.RegisterHandler(viper.GetString("top"), h.GetTopScores)
	srv.RegisterHandler(viper.GetString("stop"), h.Stop)
	srv.RegisterHandler(viper.GetString("exit"), h.Exit)
}

func (h *Handler) Sign(conn net.Conn, args []string) {

}
func (h *Handler) CreateRoom(conn net.Conn, args []string) {

}
func (h *Handler) JoinRoom(conn net.Conn, args []string) {

}
func (h *Handler) StartGame(conn net.Conn, args []string) {

}
func (h *Handler) Play(conn net.Conn, args []string) {

}
func (h *Handler) GetTopScores(conn net.Conn, args []string) {

}
func (h *Handler) Stop(conn net.Conn, args []string) {

}

func (h *Handler) Exit(conn net.Conn, args []string) {

}
