package domain

import (
	"time"

	"github.com/google/uuid"
)

// Модель подписки
type Subscription struct {
	ID          int64      // Идентификатор подписки
	ServiceName string     // Имя сервиса
	Price       int32      // Стоимость подписки
	UserID      uuid.UUID  // Идентификатор пользователя
	StartDate   time.Time  // Время начала подписки
	EndDate     *time.Time // Время окончания подписки
	CreatedAt   time.Time  // Дата создания подписки
	UpdatedAt   time.Time  // Дата обновления подписки
}
