package hub

import (
	"sync"

	"real-time-chat/models"
)

var (
	Clients = make(map[*models.Client]bool)
	Broadcast = make(chan models.Message, 100)
	History = make(map[string][]models.Message)
	mu      sync.Mutex
)

func Run() {
	for {
		msg := <-Broadcast

		if msg.Type == "chat" || msg.Type == "system" {
			mu.Lock()
			History[msg.Room] = append(History[msg.Room], msg)
			if len(History[msg.Room]) > 50 {
				History[msg.Room] = History[msg.Room][len(History[msg.Room])-50:]
			}
			mu.Unlock()
		}

		mu.Lock()
		for client := range Clients {
			if client.Room == msg.Room {
				if msg.Type == "private" {
					continue
				}
				client.Conn.WriteJSON(msg)
			}
		}
		mu.Unlock()
	}
}

func AddClient(client *models.Client) {
	mu.Lock()
	Clients[client] = true
	mu.Unlock()
}

func RemoveClient(client *models.Client) {
	mu.Lock()
	delete(Clients, client)
	mu.Unlock()
}

func GetUsers(room string) []string {
	mu.Lock()
	defer mu.Unlock()
	var users []string
	for client := range Clients {
		if client.Room == room {
			users = append(users, client.Username)
		}
	}
	return users
}

func GetHistory(room string) []models.Message {
	mu.Lock()
	defer mu.Unlock()
	return History[room]
}
