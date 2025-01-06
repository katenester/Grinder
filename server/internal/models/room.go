package models

import "time"

type Room struct {
	Players      []Player  // Хранение игроков
	LastActivity time.Time // Последняя активность
	IsNetwork    bool      // Сетевая ли игра (с другом ли)
	// Далее поля в зависимости от типа игры
	Board      [16]int         // игровое поле
	MillsBuilt map[int][][]int // Хранит список построенных мельниц для каждого игрока
}
