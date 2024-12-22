package models

import (
	"context"
	"net"
)

type Client struct {
	conn           net.Conn //соединение
	name           string   // имя пользователя
	idRoomActivity int      // id активной комнаты
	board          [][]byte
	ctx            context.Context
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn, ctx: context.Background()}
}
