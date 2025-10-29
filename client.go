package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
	"time"
)

type Message struct {
	Author string
	Body   string
	Time   string
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer client.Close()

	fmt.Println("Connected to chat server")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("\nWelcome,", name, "! Type your messages below to chat.")
	fmt.Println("(Type 'exit' to leave the chat)\n")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if strings.ToLower(text) == "exit" {
			fmt.Println("Goodbye")
			break
		}

		msg := Message{Author: name, Body: text, Time: time.Now().Format("2006-01-02 15:04:05")}
		var history []Message

		err = client.Call("Chat.SendMessage", msg, &history)
		if err != nil {
			log.Println("Error sending message:", err)
			continue
		}

		fmt.Println("\n--- Chat History ---")
		for _, m := range history {
			fmt.Printf("[%s] %s: %s\n", m.Time, m.Author, m.Body)
		}
		fmt.Println("--------------------\n")
	}
}
