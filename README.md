# Simple Chatroom using Go RPC

This project is a simple **chatroom** application implemented in Go using **RPC (Remote Procedure Call)**.

## Project Description

The project simulates a basic distributed chat system where multiple clients can send and receive messages through a central server.

### Server

- Listens for incoming RPC connections on port `1234`.
- Stores all messages in memory (a long list).
- Returns the chat history to any client that sends a message.

### Client

- Connects to the RPC server.
- Sends messages to the server.
- Fetches the entire chat history from the server after sending each message.
- Keeps running until the user types `exit`.

## How to Run

1. Run the server:
   ```bash
   go run server.go
   ```
2. go run client.go and enter your name
3. Type messages and press Enter.
4. Type exit to quit.

## Demo Video

[Demo Video](https://drive.google.com/file/d/1VswRrWTrpVR6Ka484fERKTFHfr_qk_jR/view?usp=sharing)

## Documentation

This project contains two main Go files:

- **server.go**:  
  Handles incoming RPC requests from clients.  
  Stores chat messages in memory and returns the full chat history to each client.

- **client.go**:  
  Connects to the server and lets the user send messages.  
  After each message, it fetches and displays the updated chat history.

## How It Works

1. The **server** starts and listens on port 1234.
2. Each **client** connects to the server via RPC.
3. When a user types a message, the client sends it to the server.
4. The server stores it and returns the full message list to all connected clients.
5. The client displays the messages on screen.
