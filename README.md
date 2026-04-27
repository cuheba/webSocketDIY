# WebSocket Echo & Redis History

<img width="834" height="665" alt="image" src="https://github.com/user-attachments/assets/cd55a5a5-c0ee-423f-a776-a563f8fcda76" /> 

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

## *Ввод терминала
<img width="693" height="165" alt="image" src="https://github.com/user-attachments/assets/a7293893-c5c6-40c2-9d3e-1467738e252a" />

