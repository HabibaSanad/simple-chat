package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
	"time"
)

// must match server.Message
type Message struct {
	Author string
	Body   string
	Time   string
}

func main() {
	name := flag.String("name", "Anonymous", "your display name")
	addr := flag.String("addr", "localhost:1234", "server address")
	flag.Parse()

	client, err := rpc.Dial("tcp", *addr)
	if err != nil {
		log.Fatalf("cannot connect to server: %v", err)
	}
	defer client.Close()

	// fetch history at start
	var history []Message
	err = client.Call("Chat.GetHistory", struct{}{}, &history)
	if err == nil {
		printHistory(history)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type messages and press Enter. Type 'exit' to quit.")
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		if strings.ToLower(text) == "exit" {
			fmt.Println("Exiting...")
			return
		}

		msg := Message{Author: *name, Body: text, Time: time.Now().Format("2006-01-02 15:04:05")}
		var newHistory []Message
		err = client.Call("Chat.SendMessage", msg, &newHistory)
		if err != nil {
			log.Println("Error sending message:", err)
			// try to reconnect once
			client.Close()
			client, err = rpc.Dial("tcp", *addr)
			if err != nil {
				log.Println("Reconnect failed:", err)
				continue
			}
			// retry send
			err = client.Call("Chat.SendMessage", msg, &newHistory)
			if err != nil {
				log.Println("Send failed after reconnect:", err)
				continue
			}
		}
		printHistory(newHistory)
	}
}

func printHistory(h []Message) {
	fmt.Println("--- chat history ---")
	for _, m := range h {
		fmt.Printf("[%s] %s: %s\n", m.Time, m.Author, m.Body)
	}
	fmt.Println("--------------------")
}
