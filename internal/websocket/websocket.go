package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel
var mutex = &sync.Mutex{}

// Message object
type Message struct {
	Type string `json:"type"`
	Body string `json:"body"`
}

// HandleConnections handles incoming WebSocket connections
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Failed to upgrade connection: %v", err)
		return
	}
	defer ws.Close()

	mutex.Lock()
	clients[ws] = true
	mutex.Unlock()

	for {
		var msg Message
		// Read new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			deleteClient(ws)
			break
		}
		// Send new message to the broadcast channel
		broadcast <- msg
	}
}

// deleteClient safely removes a client from the clients map
func deleteClient(ws *websocket.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(clients, ws)
}

// HandleMessages listens for messages on the broadcast channel and forwards them to all clients
func HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it to every connected client
		mutex.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error writing message: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}
