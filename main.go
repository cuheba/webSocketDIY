package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)



var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	// allowed all connections
	CheckOrigin: func(r *http.Request) bool { return true },
}


func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		log.Printf("Recieved: %s\n", msg)

		if err := SaveMessage(string(msg)); err != nil {
			log.Println("Redis save error:", err)
		}

		// echo msg back ro client
		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func handleHistory(w http.ResponseWriter, r *http.Request) {
	// Получаем историю через функцию из data.go
	messages, err := GetHistory()
	if err != nil {
		http.Error(w, "Failed to fetch history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Отправляем сырые данные построчно
	for _, msg := range messages {
		fmt.Fprintf(w, "%s\n", msg)
	}
}



func main() {
	InitRedis()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Only serve the exact root path; prevent directory listing
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/history", handleHistory)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

