// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ws

import (
	"encoding/json"
	"log"
)

var hub *Hub

type datamsg struct {
	Event string                 `json:"event"`
	Data  map[string]interface{} `json:"data"`
}

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func Init() {
	hub = newHub()
	go hub.Run()
}

func GetHub() *Hub {
	return hub
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			log.Printf("New ws client\n")
			h.clients[client] = true
		case client := <-h.unregister:
			log.Printf("Disconnected ws client\n")
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			var d datamsg
			json.Unmarshal(message.Message, &d)

			log.Printf("New event: %s\n", d.Event)

			if (d.Event == "identify") {
				ident := string(d.Data["identity"])
				message.Client.Identity = ident
				log.Printf("Client identified as %s\n", ident)
				continue
			}

			for client := range h.clients {
				select {
				case client.send <- message.Message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *Hub) Broadcast(msg []byte) {
	for client := range h.clients {
		client.send <- msg
	}
}

func (h *Hub) BroadcastEvent(evt string, data interface{}) {
	msg, _ := json.Marshal(map[string]interface{}{
		"event": evt,
		"data":  data,
	})

	h.Broadcast(msg)
}
