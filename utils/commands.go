package utils

import (
	"fmt"
	"strings"
	"time"

	"real-time-chat/hub"
	"real-time-chat/models"
)


func ProcessCommand(client *models.Client, msg models.Message) bool {
	if len(msg.Message) == 0 || msg.Message[0] != '/' {
		return false
	}
	parts := strings.SplitN(msg.Message, " ", 3)
	command := parts[0]
	switch command {
	case "/users":
		users := hub.GetUsers(client.Room)
		sysMsg := models.Message{
			Username: "Sistema",
			Message:  "Usuários na sala: " + strings.Join(users, ", "),
			Room:     client.Room,
			Type:     "system",
			Time:     time.Now().Format("15:04:05"),
		}
		sendPrivate(client, sysMsg)
		return true

	case "/pm":
		if len(parts) < 3 {
			errMsg := models.Message{
				Username: "Sistema",
				Message:  "Uso: /pm <usuário> <mensagem>",
				Room:     client.Room,
				Type:     "system",
				Time:     time.Now().Format("15:04:05"),
			}
			sendPrivate(client, errMsg)
			return true
		}
		target := parts[1]
		privateMsg := parts[2]
		pm := models.Message{
			Username: client.Username,
			Message:  privateMsg,
			Room:     client.Room,
			Type:     "private",
			Target:   target,
			Time:     time.Now().Format("15:04:05"),
		}
		sent := sendPrivateByUsername(client.Room, target, pm)
		if !sent {
			errMsg := models.Message{
				Username: "Sistema",
				Message:  fmt.Sprintf("Usuário %s não encontrado na sala.", target),
				Room:     client.Room,
				Type:     "system",
				Time:     time.Now().Format("15:04:05"),
			}
			sendPrivate(client, errMsg)
		}
		return true

	case "/kick":
		if client.Username != "admin" {
			errMsg := models.Message{
				Username: "Sistema",
				Message:  "Você não tem permissão para usar esse comando.",
				Room:     client.Room,
				Type:     "system",
				Time:     time.Now().Format("15:04:05"),
			}
			sendPrivate(client, errMsg)
			return true
		}
		if len(parts) < 2 {
			errMsg := models.Message{
				Username: "Sistema",
				Message:  "Uso: /kick <usuário>",
				Room:     client.Room,
				Type:     "system",
				Time:     time.Now().Format("15:04:05"),
			}
			sendPrivate(client, errMsg)
			return true
		}
		target := parts[1]
		kickUser(client.Room, target)
		return true

	default:
		errMsg := models.Message{
			Username: "Sistema",
			Message:  "Comando desconhecido.",
			Room:     client.Room,
			Type:     "system",
			Time:     time.Now().Format("15:04:05"),
		}
		sendPrivate(client, errMsg)
		return true
	}
}

func sendPrivate(client *models.Client, msg models.Message) {
	client.Conn.WriteJSON(msg)
}

func sendPrivateByUsername(room, username string, msg models.Message) bool {
	found := false
	for client := range hub.Clients {
		if client.Room == room && client.Username == username {
			client.Conn.WriteJSON(msg)
			found = true
		}
	}
	return found
}

func kickUser(room, target string) {
	for client := range hub.Clients {
		if client.Room == room && client.Username == target {
			msg := models.Message{
				Username: "Sistema",
				Message:  "Você foi expulso da sala por um administrador.",
				Room:     client.Room,
				Type:     "system",
				Time:     time.Now().Format("15:04:05"),
			}
			client.Conn.WriteJSON(msg)
			client.Conn.Close()
		}
	}
}
