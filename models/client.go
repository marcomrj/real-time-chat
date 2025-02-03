package models

import "github.com/gorilla/websocket"

type Client struct {
	Conn        *websocket.Conn
	Username    string
	Room        string
	RateLimiter chan struct{}
