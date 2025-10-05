package ws

import (
	"encoding/json"
)

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

var currentHub *Hub

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					delete(h.clients, client)
					close(client.send)

				}
			}
		}
	}
}

func SetHub(h *Hub) {
	currentHub = h
}

func GetHub() *Hub { return currentHub }

type Event struct {
	Resource string      `json:"resource"`
	Action   string      `json:"action"`
	Data     interface{} `json:"data"`
}

func Emit(resource, action string, data interface{}) {
	if currentHub == nil {
		return
	}
	payload, err := json.Marshal(Event{Resource: resource, Action: action, Data: data})
	if err != nil {
		return
	}
	currentHub.broadcast <- payload
}
