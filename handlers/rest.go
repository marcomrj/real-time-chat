package handlers

import (
	"encoding/json"
	"net/http"

	"real-time-chat/hub"
)

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	room := r.URL.Query().Get("room")
	if room == "" {
		room = "default"
	}
	msgs := hub.GetHistory(room)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msgs)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	room := r.URL.Query().Get("room")
	if room == "" {
		room = "default"
	}
	users := hub.GetUsers(room)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	total := len(hub.Clients)
	status := map[string]interface{}{
		"total_clients": total,
		"timestamp":     "placeholder",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
