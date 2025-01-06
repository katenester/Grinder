package service

import (
	"Grinder/Protocol"
	"Grinder/server/internal/models"
	"Grinder/server/internal/repository"
	"net"
)

type PlayersService struct {
	repo repository.Players
}

func NewPlayersService(repo repository.Players) *PlayersService {
	return &PlayersService{
		repo: repo,
	}
}

func (p *PlayersService) CreatePlayer(username string, conn net.Conn) error {
	return p.repo.CreatePlayer(username, conn)
}
func (p *PlayersService) GetTop(conn net.Conn, req Protocol.Request) error {
	return p.repo.GetTop(conn, req)
	//h.sendResponse(conn, h.service.Players.CreatePlayer(req.Username, conn).(Protocol.Response))
}

func (p *PlayersService) Exit(conn net.Conn, req Protocol.Request) error {
	return p.repo.Exit(conn, req)
}

func (p *PlayersService) GetUser(username string) (models.Player, error) {
	return p.repo.GetUser(username)
}
