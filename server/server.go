package main

import (
	"chat-application/commons"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

type Server struct {
	mutex        sync.Mutex
	participants map[int]commons.ParticipantInfo
}

func (s *Server) RegisterParticipant(participant commons.ParticipantInfo, reply *string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.participants[participant.ConnectionPort] = participant
	*reply = fmt.Sprintf("Welcome %s! You have been registered.", participant.ParticipantName)
	log.Printf("Participant Registered: %v\n", participant)
	return nil
}

func (s *Server) BroadcastMessage(payload commons.CommunicationPayload, reply *string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, participant := range s.participants {
		if participant.ConnectionPort != payload.ConnectionPort {
			client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%d", participant.ConnectionPort))
			if err != nil {
				log.Printf("Failed to send message to %s: %v\n", participant.ParticipantName, err)
				continue
			}
			defer client.Close()

			var ack string
			err = client.Call("Client.ReceiveMessage", payload, &ack)
			if err != nil {
				log.Printf("Error calling ReceiveMessage: %v\n", err)
			}
		}
	}
	*reply = "Message broadcasted successfully."
	return nil
}

func main() {
	server := &Server{
		participants: make(map[int]commons.ParticipantInfo),
	}

	err := rpc.Register(server)
	if err != nil {
		log.Fatalf("Error registering server: %v\n", err)
	}

	listener, err := net.Listen("tcp", ":42586")
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
	defer listener.Close()
	log.Println("Server started on port 42586")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection error: %v\n", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
