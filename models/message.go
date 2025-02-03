package models

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Room     string `json:"room"`
	Type     string `json:"type"`
	Target   string `json:"target,omitempty"`
	Time     string `json:"time,omitempty"`
}
