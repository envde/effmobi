# Тестовый проект на GO

## Запуск проекта

Притяните git репозиторий. 

Добавте в корень проекта файл .env.

```env
APP_PORT=8080
APP_ENV=dev

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=effmobi
DB_SSLMODE=disable

```

Выполните команду 

docker compose up --build

Проект запустится.
