### Go Code (main.go)

1. **Imports and Upgrader Definition:**
   ```go
   import (
       "fmt"
       "net/http"
       "github.com/gorilla/websocket"
       "sync"
   )

   var upgrader = websocket.Upgrader{
       ReadBufferSize:  1024,
       WriteBufferSize: 1024,
       CheckOrigin: func(r *http.Request) bool {
           return true
       },
   }
   ```
   - Imports necessary packages including `fmt`, `net/http`, `github.com/gorilla/websocket`, and `sync`.
   - Defines an upgrader to upgrade HTTP connections to WebSocket connections with specified buffer sizes and a function to allow all origins.

2. **Client and Hub Structures:**
   ```go
   type Client struct {
       conn *websocket.Conn
   }

   type Hub struct {
       clients    map[*Client]bool
       broadcast  chan []byte
       register   chan *Client
       unregister chan *Client
   }
   ```
   - `Client` represents a WebSocket connection.
   - `Hub` manages connected clients and handles broadcasting messages.

3. **Global Variables:**
   ```go
   var hub = Hub{
       clients:    make(map[*Client]bool),
       broadcast:  make(chan []byte),
       register:   make(chan *Client),
       unregister: make(chan *Client),
   }

   var mu sync.Mutex
   ```
   - Initializes a global `hub` instance and a mutex `mu` for synchronization.

4. **Hub Run Method:**
   ```go
   func (h *Hub) run() {
       for {
           select {
           case client := <-h.register:
               h.clients[client] = true
           case client := <-h.unregister:
               if _, ok := h.clients[client]; ok {
                   delete(h.clients, client)
                   client.conn.Close()
               }
           case message := <-h.broadcast:
               for client := range h.clients {
                   mu.Lock()
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
   ```
   - Handles client registration, unregistration, and broadcasting messages to all connected clients.

5. **Echo Handler:**
   ```go
   func echo(w http.ResponseWriter, r *http.Request) {
       conn, err := upgrader.Upgrade(w, r, nil)
       if err != nil {
           fmt.Println(err)
           return
       }
       client := &Client{conn: conn}
       hub.register <- client

       defer func() {
           hub.unregister <- client
       }()

       for {
           _, message, err := conn.ReadMessage()
           if err != nil {
               fmt.Println(err)
               return
           }
           hub.broadcast <- message
       }
   }
   ```
   - Upgrades HTTP requests to WebSocket connections and manages communication by reading and broadcasting messages.

6. **Main Function:**
   ```go
   func main() {
       http.Handle("/", http.FileServer(http.Dir("./static")))
       http.HandleFunc("/ws", echo)

       go hub.run()

       fmt.Println("Server started at :8080")
       if err := http.ListenAndServe(":8080", nil); err != nil {
           fmt.Println("ListenAndServe:", err)
       }
   }
   ```
   - Serves static files and handles WebSocket requests at the `/ws` endpoint.
   - Starts the hub in a separate goroutine and the HTTP server on port 8080.

### HTML Code (index.html)

```html
<!DOCTYPE html>
<html>
<head>
    <title>WebSocket Chat</title>
</head>
<body>
    <h1>WebSocket Chat</h1>
    <input type="text" id="message" placeholder="Enter message">
    <button onclick="sendMessage()">Send</button>
    <div id="chat"></div>

    <script>
        let ws = new WebSocket("ws://localhost:8080/ws");

        ws.onmessage = function(event) {
            let chat = document.getElementById("chat");
            let message = document.createElement("div");
            message.innerText = event.data;
            chat.appendChild(message);
        };

        function sendMessage() {
            let message = document.getElementById("message").value;
            ws.send(message);
            document.getElementById("message").value = '';
        }
    </script>
</body>
</html>
```

1. **HTML Structure:**
   - Contains an input field for typing messages, a button to send messages, and a div to display chat messages.

2. **WebSocket Connection:**
   ```javascript
   let ws = new WebSocket("ws://localhost:8080/ws");
   ```
   - Establishes a WebSocket connection to the server at `ws://localhost:8080/ws`.

3. **Handling Incoming Messages:**
   ```javascript
   ws.onmessage = function(event) {
       let chat = document.getElementById("chat");
       let message = document.createElement("div");
       message.innerText = event.data;
       chat.appendChild(message);
   };
   ```
   - Appends incoming messages to the chat div.

4. **Sending Messages:**
   ```javascript
   function sendMessage() {
       let message = document.getElementById("message").value;
       ws.send(message);
       document.getElementById("message").value = '';
   }
   ```
   - Sends the typed message to the WebSocket server and clears the input field.

### Flow of the Application

1. **Client Connection:**
   - When the HTML page loads, it creates a WebSocket connection to the server.
   - The server upgrades the HTTP connection to a WebSocket connection and registers the new client.

2. **Message Sending:**
   - When a user types a message and clicks the "Send" button, the message is sent to the server via WebSocket.
   - The server receives the message and broadcasts it to all connected clients.

3. **Message Receiving:**
   - All connected clients receive the broadcasted message and display it in their chat window.

This setup allows multiple clients to connect to the server and exchange messages in real-time via WebSocket.