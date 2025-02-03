package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"real-time-chat/hub"
	"real-time-chat/handlers"
)

func main() {
	logChan := make(chan string, 100)
	go logger(logChan)

	go hub.Run()

	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/ws", handlers.HandleWS)
	http.HandleFunc("/history", handlers.HistoryHandler)
	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/status", handlers.StatusHandler)

	addr := ":8080"
	logChan <- fmt.Sprintf("Servidor iniciado em %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func logger(ch chan string) {
	file, err := os.OpenFile("chat.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Erro abrindo arquivo de log:", err)
		return
	}
	defer file.Close()
	for {
		msg := <-ch
		timestamp := time.Now().Format(time.RFC3339)
		logLine := fmt.Sprintf("%s: %s\n", timestamp, msg)
		file.WriteString(logLine)
	}
}
