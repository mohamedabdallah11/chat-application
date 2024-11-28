package commons

type ParticipantInfo struct {
	ConnectionPort  int    
	ParticipantName string 
}


type CommunicationPayload struct {
	MessageContent string 
	ConnectionPort int    
}
