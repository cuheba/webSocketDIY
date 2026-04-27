# WebSocket Echo & Redis History

Простой веб-сервер на Go, демонстрирующий работу с WebSockets в реальном времени и использование Redis для сохранения истории сообщений.

## 🚀 Особенности
- **WebSocket Echo**: Мгновенная отправка и получение сообщений через веб-сокеты.
- **Redis Persistence**: Каждое полученное сообщение автоматически сохраняется в список Redis (`chat_history`).
- **History Viewer**: Отдельный эндпоинт и кнопка для получения "сырых" данных из базы.

## 🛠 Стек технологий
- **Язык**: Go (Golang)
- **Библиотеки**: 
  - `github.com/gorilla/websocket` — работа с сокетами.
  - `github.com/redis/go-redis/v9` — клиент для Redis.
- **База данных**: Redis (v9)
- **Frontend**: HTML5, Vanilla JavaScript, CSS (Flexbox)



