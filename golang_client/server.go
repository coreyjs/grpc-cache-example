package main

// import (
// 	"log"
// 	"net"

// 	"github.com/coreyjs/grpc-cache-example/chat"
// 	"google.golang.org/grpc"
// )

// func main() {
// 	lis, err := net.Listen("tcp", ":9000")
// 	if err != nil {
// 		log.Fatalf("Failed to listen to port 9000: %v", err)
// 	}

// 	s := chat.Server{}
// 	grpcServer := grpc.NewServer()

// 	chat.RegisterChatServiceServer(grpcServer, &s)

// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
// 	}

// }
