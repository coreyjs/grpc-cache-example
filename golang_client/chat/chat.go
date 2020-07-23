package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {

}

func (s *Server) SayHello(ctx context.Context, message *Message)(*Message, error) {
	log.Printf("Recieved message body from client: %s", message.Body)
	return &Message{Body: "Hello from server!"}, nil
}