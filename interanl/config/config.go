package config

import (
	"fmt"
	"os"
)

// Конфигурация
type Config struct {
	App AppConfig // Настройки приложения
	DB  DBConfig  // Настройки базы данных
}

// Конфигурация приложения
type AppConfig struct {
	Port string // Порт на котором будет запущен сервис
	Env  string
}

// Настройки БД
type DBConfig struct {
	Host     string // Хост
	Port     string // Порт
	User     string // Пользователь
	Password string // Пароль
	Name     string // Название БД
	SSLMode  string // Настройка SSL
}

// Конструирует строку соединения с БД
func (d DBConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		d.User, d.Password, d.Host, d.Port, d.Name, d.SSLMode,
	)
}

// Загружает файл конфигураци из .env
func Load() (*Config, error) {
	cfg := &Config{
		App: AppConfig{
			Port: getEnv("APP_PORT", "8080"),
			Env:  getEnv("APP_ENV", "dev"),
		},
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "effmobi"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}
	return cfg, nil
}

// Возвращает env
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
