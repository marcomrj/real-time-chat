package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"real-time-chat/hub"
	"real-time-chat/models"
	"real-time-chat/utils"
	"real-time-chat/utils/ratelimiter"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWS(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Erro ao estabelecer conexão: %v", err)
		return
	}

	room := r.URL.Query().Get("room")
	if room == "" {
		room = "default"
	}
	username := r.URL.Query().Get("username")
	if username == "" {
		username = "Anônimo"
	}

	client := &models.Client{
		Conn:        ws,
		Username:    username,
		Room:        room,
		RateLimiter: make(chan struct{}, 5),
	}
	go ratelimiter.StartRateLimiter(client.RateLimiter, 5, time.Second)

	hub.AddClient(client)

	sysMsg := models.Message{
		Username: "Sistema",
		Message:  fmt.Sprintf("%s entrou na sala.", username),
		Room:     room,
		Type:     "system",
		Time:     time.Now().Format("15:04:05"),
	}
	hub.Broadcast <- sysMsg

	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Erro ao ler JSON: %v", err)
			hub.RemoveClient(client)
			leftMsg := models.Message{
				Username: "Sistema",
				Message:  fmt.Sprintf("%s saiu da sala.", username),
				Room:     room,
				Type:     "system",
				Time:     time.Now().Format("15:04:05"),
			}
			hub.Broadcast <- leftMsg
			break
		}

		msg.Time = time.Now().Format("15:04:05")

		select {
		case <-client.RateLimiter:
		default:
			warning := models.Message{
				Username: "Sistema",
				Message:  "Rate limit excedido. Aguarde um momento.",
				Room:     room,
				Type:     "system",
				Time:     time.Now().Format("15:04:05"),
			}
			ws.WriteJSON(warning)
			continue
		}

		if msg.Type == "typing" {
			msg.Room = client.Room
			hub.Broadcast <- msg
			continue
		}

		if utils.ProcessCommand(client, msg) {
			continue
		}

		msg.Type = "chat"
		msg.Room = client.Room
		msg.Username = client.Username

		hub.Broadcast <- msg
	}
}
