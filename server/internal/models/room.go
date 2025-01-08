package models

import "time"

type Room struct {
	Id           int
	Players      []Player  // Хранение игроков
	LastActivity time.Time // Последняя активность
	StatusGame   Status
	IsNetwork    bool // Сетевая ли игра (с другом ли)
	// Далее поля в зависимости от типа игры
	Board      [16]int         // игровое поле
	MillsBuilt map[int][][]int // Хранит список построенных мельниц для каждого игрока
}

// Status Статус игры
type Status int

// Константы, использующие iota для автоматической генерации значений
const (
	Create    Status = iota // 0
	InProcess               // 1
	WinUser1                // 2
	WinUser2                // 3
	Broken
)
