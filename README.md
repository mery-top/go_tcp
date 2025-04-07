# Go TCP Server 🌐

This is a simple **TCP server written in Go** that allows communication between the server and a connected client using standard input and output. 

## Usage 🚀

### 1. Clone this repository

```bash
git clone https://github.com/yourusername/go-tcp-server.git
cd go-tcp-server
```

## How It Works ⚙️

- The server listens on TCP port `8080`.
- When a new client connects, the server:
  - Sends a welcome message.
  - Starts two goroutines:
    - One for reading and printing client messages.
    - Another for sending messages typed by the server user.
- The server continues running until the client disconnects.
