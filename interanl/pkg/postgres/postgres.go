package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Возвращает соединение из пула. Это работает эффективнее, чем создавать новое
// соединение с БД каждый запрос.
func NewPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	// Получаем пул , который управляет соединениями с БД
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New: %w", err)
	}
	// Проверяем работает ли соединение
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pool.Ping: %w", err)
	}
	return pool, nil
}
