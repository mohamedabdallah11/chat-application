package main

import (
	"bufio"
	"chat-application/commons"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"strings"
)

type Client struct {
	Name string
	Port int
}

func (c *Client) ReceiveMessage(payload commons.CommunicationPayload, ack *string) error {
	fmt.Printf("[%d]: %s\n", payload.ConnectionPort, payload.MessageContent)
	*ack = "Message received"
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Error allocating port: %v\n", err)
	}
	defer listener.Close()
	port := listener.Addr().(*net.TCPAddr).Port

	client := &Client{Name: name, Port: port}

	err = rpc.Register(client)
	if err != nil {
		log.Fatalf("Error registering client: %v\n", err)
	}

	serverConn, err := rpc.Dial("tcp", "localhost:42586")
	if err != nil {
		log.Fatalf("Error connecting to server: %v\n", err)
	}
	defer serverConn.Close()

	participantInfo := commons.ParticipantInfo{
		ConnectionPort:  port,
		ParticipantName: name,
	}
	var reply string
	err = serverConn.Call("Server.RegisterParticipant", participantInfo, &reply)
	if err != nil {
		log.Fatalf("Error registering with server: %v\n", err)
	}
	fmt.Println(reply)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Connection error: %v\n", err)
				continue
			}
			go rpc.ServeConn(conn)
		}
	}()

	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "" {
			continue
		}

		payload := commons.CommunicationPayload{
			MessageContent: message,
			ConnectionPort: port,
		}
		err := serverConn.Call("Server.BroadcastMessage", payload, &reply)
		if err != nil {
			log.Printf("Error sending message: %v\n", err)
		}
	}
}
