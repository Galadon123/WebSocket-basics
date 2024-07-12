# WebSocket Echo Server Project

This document provides an understanding of the WebSocket Echo Server implemented in Go, with a corresponding HTML client for testing.

## Project Structure

```
project-root/
├── main.go
└── static/
    └── index.html
```

- **main.go**: Contains the Go code for the WebSocket server.
- **static/index.html**: Contains the HTML code for the WebSocket client.

## WebSocket Server (main.go)

### Imports

```go
import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)
```

- **fmt**: For formatted I/O operations.
- **net/http**: For HTTP client and server implementations.
- **github.com/gorilla/websocket**: For WebSocket implementation.

### WebSocket Upgrader

```go
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
```

- **upgrader**: Configures the WebSocket upgrade with buffer sizes and a function to check the origin of requests. The `CheckOrigin` function is set to always return `true` to allow connections from any origin.

### Echo Handler

```go
func echo(w http.ResponseWriter, r *http.Request) {
	// Upgrade the connection to a WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		// Read message from WebSocket
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print received message
		fmt.Printf("Received: %s\n", message)

		// Write message back to WebSocket
		if err := conn.WriteMessage(messageType, message); err != nil {
			fmt.Println(err)
			return
		}
	}
}
```

- **conn, err := upgrader.Upgrade(w, r, nil)**: Upgrades the HTTP connection to a WebSocket connection. If an error occurs, it prints the error and returns.
- **defer conn.Close()**: Ensures that the WebSocket connection is closed when the function exits.
- **for loop**: Continuously reads messages from the WebSocket and echoes them back to the client.
  - **conn.ReadMessage()**: Reads a message from the WebSocket.
  - **conn.WriteMessage(messageType, message)**: Sends the received message back to the WebSocket.

### Main Function

```go
func main() {
	// Serve static files from the "static" directory
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ws", echo) // Handle WebSocket requests

	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
```

- **http.Handle("/", http.FileServer(http.Dir("./static")))**: Serves static files from the "static" directory.
- **http.HandleFunc("/ws", echo)**: Registers the `echo` handler for WebSocket requests at the `/ws` endpoint.
- **http.ListenAndServe(":8080", nil)**: Starts the HTTP server on port 8080. If an error occurs, it prints the error.

## WebSocket Client (index.html)

```html
<!DOCTYPE html>
<html>
<head>
	<title>WebSocket Echo Test</title>
</head>
<body>
	<h1>WebSocket Echo server Test</h1>
	<input type="text" id="message" placeholder="Enter message">
	<button onclick="sendMessage()">Send</button>
	<p id="response"></p>

	<script>
		let ws = new WebSocket("ws://localhost:8080/ws");

        ws.onopen = () => {
            console.log("Connected to the WebSocket server");
        };
        
		ws.onmessage = function(event) {
			document.getElementById("response").innerText = "Received: " + event.data;
		};

		function sendMessage() {
			let message = document.getElementById("message").value;
			ws.send(message);
		}
	</script>
</body>
</html>
```

### Explanation

- **WebSocket Initialization**
  ```javascript
  let ws = new WebSocket("ws://localhost:8080/ws");
  ```

  Creates a new WebSocket connection to the server at `ws://localhost:8080/ws`.

- **WebSocket Events**
  ```javascript
  ws.onopen = () => {
      console.log("Connected to the WebSocket server");
  };
  
  ws.onmessage = function(event) {
      document.getElementById("response").innerText = "Received: " + event.data;
  };
  ```

  - **ws.onopen**: Logs a message to the console when the WebSocket connection is established.
  - **ws.onmessage**: Updates the `response` paragraph with the received message data.

- **Sending Messages**
  ```javascript
  function sendMessage() {
      let message = document.getElementById("message").value;
      ws.send(message);
  }
  ```

  Reads the message from the input field and sends it to the WebSocket server.

### Running the Project

1. **Start the Go Server**:
   ```sh
   go run main.go
   ```
   This starts the WebSocket server on `http://localhost:8080`.

2. **Open the WebSocket Client**:
   - Navigate to `http://localhost:8080` in your web browser.
   - Enter a message in the input field and click "Send".
   - The server will echo the message back, and it will be displayed in the `response` paragraph.

## Summary

This project demonstrates a basic WebSocket echo server using Go and a simple HTML client to interact with the server. The server listens for WebSocket connections, echoes received messages back to the client, and serves the HTML client. The client connects to the server, sends messages, and displays the echoed messages.