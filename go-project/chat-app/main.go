package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	"sync"
)

// WebSocket upgrader to upgrade HTTP connection to WebSocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client represents a single WebSocket connection
//For each client connection, we create a new Client object and register it with the hub.
type Client struct {
	conn *websocket.Conn
}

// Hub maintains the set of active clients and broadcasts messages to them
type Hub struct {
	clients map[*Client]bool  // Registered clients
	broadcast chan []byte  // Inbound messages to broadcast to clients
	register chan *Client   // Register requests from the clients
	unregister chan *Client  // Unregister requests from clients
}

var hub = Hub{
	clients: make(map[*Client]bool), // Create a map of clients
	broadcast: make(chan []byte),// Create a channel to broadcast messages
	register: make(chan *Client),// Create a channel to register clients
	unregister: make(chan *Client),// Create a channel to unregister clients
}



var mu sync.Mutex  
//The `var mu sync.Mutex` declares a mutex named `mu` for synchronizing access to shared resources. It prevents race conditions by ensuring that only one goroutine can access the critical section of code at a time. 

// run method for the Hub to handle client registration, unregistration, and broadcasting messages
func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:  // Register a new client connection with the hub 
			h.clients[client] = true 
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client) // Remove the client from the map
				client.conn.Close()
			}
		case message := <-h.broadcast:  // Broadcast message to all clients
			for client := range h.clients {
				mu.Lock() // Lock the mutex before writing to the client connection
				err := client.conn.WriteMessage(websocket.TextMessage, message)
				mu.Unlock()
				if err != nil {
					client.conn.Close()
					delete(h.clients, client)
				}
			}
		}
	}
}

// echo handles WebSocket requests from clients
func echo(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Create a new client for this connection
	client := &Client{conn: conn}
	// Register the client with the hub
	hub.register <- client

	defer func() {
		// Unregister the client when done
		hub.unregister <- client
	}()

	for {
		// Read message from WebSocket
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		// Send message to the hub to broadcast to all clients
		hub.broadcast <- message
	}
}

func main() {
	// Serve static files from the "static" directory
	http.Handle("/", http.FileServer(http.Dir("./static")))
	// Handle WebSocket requests at "/ws" endpoint
	http.HandleFunc("/ws", echo)

	// Run the hub in a separate goroutine
	go hub.run()

	fmt.Println("Server started at :8080")
	// Start the HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
