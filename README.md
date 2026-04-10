# GREEN-API Test Page

Go веб-приложение для тестирования методов GREEN-API.

## Функционал

Приложение поддерживает следующие методы GREEN-API:

- **getSettings** - получение настроек инстанса
- **getStateInstance** - получение состояния инстанса
- **sendMessage** - отправка текстового сообщения
- **sendFileByUrl** - отправка файла по URL

## Структура проекта

```
greenapitest/
├── cmd/
│   └── server/
│       └── main.go           # Точка входа приложения
├── internal/
│   ├── api/
│   │   └── client.go         # HTTP клиент для GREEN-API
│   ├── handlers/
│   │   └── handlers.go       # HTTP обработчики
│   └── models/
│       └── models.go         # Структуры данных
├── web/
│   ├── static/
│   │   ├── css/
│   │   │   └── style.css     # Стили
│   │   └── js/
│   │       └── app.js        # JavaScript
│   └── templates/
│       └── index.html        # HTML шаблон
├── go.mod
└── README.md
```

## Запуск

```bash
# Клонировать репозиторий
git clone https://github.com/username/greenapitest.git
cd greenapitest

# Запустить сервер
go run cmd/server/main.go

# Или собрать и запустить
go build -o server cmd/server/main.go
./server
```

Сервер запустится на `http://localhost:8080`

## Использование

1. Создайте инстанс в [личном кабинете GREEN-API](https://green-api.com/)
2. Подключите WhatsApp, отсканировав QR-код
3. Откройте `http://localhost:8080` в браузере
4. Введите `idInstance` и `ApiTokenInstance`
5. Используйте кнопки для вызова методов API

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/getSettings` | POST | Получить настройки инстанса |
| `/api/getStateInstance` | POST | Получить состояние инстанса |
| `/api/sendMessage` | POST | Отправить сообщение |
| `/api/sendFileByUrl` | POST | Отправить файл по URL |

## Переменные окружения

- `PORT` - порт сервера (по умолчанию: 8080)

## Документация

- [GREEN-API Documentation](https://green-api.com/docs/)
