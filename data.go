package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

// InitRedis инициализирует подключение к Redis
func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // без пароля
		DB:       0,  // дефолтная БД
	})

	// Проверка подключения
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

// SaveMessage добавляет сообщение в конец списка chat_history
func SaveMessage(msg string) error {
	return rdb.RPush(ctx, "chat_history", msg).Err()
}

// GetHistory возвращает все сообщения из списка chat_history
func GetHistory() ([]string, error) {
	return rdb.LRange(ctx, "chat_history", 0, -1).Result()
}

