package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
	"time"
)

// Message represents a chat message
type Message struct {
	Author string
	Body   string
	Time   string
}

// Chat provides RPC methods
type Chat struct {
	mu       sync.Mutex
	messages []Message
}

// SendMessage appends a message and returns the whole history
func (c *Chat) SendMessage(msg Message, reply *[]Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	msg.Time = time.Now().Format("2006-01-02 15:04:05")
	c.messages = append(c.messages, msg)
	*reply = make([]Message, len(c.messages))
	copy(*reply, c.messages)
	return nil
}

// GetHistory returns the whole chat history
func (c *Chat) GetHistory(_ struct{}, reply *[]Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	*reply = make([]Message, len(c.messages))
	copy(*reply, c.messages)
	return nil
}

func main() {
	chat := &Chat{}
	err := rpc.Register(chat)
	if err != nil {
		log.Fatalf("rpc register error: %v", err)
	}

	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}
	fmt.Println("Chat server listening on :1234")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
