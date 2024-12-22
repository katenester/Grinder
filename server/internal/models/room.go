package models

import "time"

type Room struct {
	id           int
	Players      [2]Player // Хранение игроков
	LastActivity time.Time // Последняя активность
	isNetwork    bool      // Сетевая ли игра (с другом ли)
	// Далее поля в зависимости от типа игры
}
