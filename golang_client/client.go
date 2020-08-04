package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/coreyjs/grpc-cache-example/cache"
	"google.golang.org/grpc"
)

func main() {
	var (
		writing = true
		buf     []byte
		n       int
		status  *cache.UploadStatus
	)

	actionPtr := flag.String("a", "", "action to preform.  status, get, store")
	fileInputPtr := flag.String("f", "", "file name")

	flag.Parse()

	if actionPtr == nil {
		fmt.Println("-a is empty")
		return
	}

	if *actionPtr != "status" && *fileInputPtr == "" {
		fmt.Println("-f filename is required for any file related action")
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer conn.Close()

	client := cache.NewCacheHubClient(conn)

	if *actionPtr == "status" {
		fmt.Println("Getting Status ...")
		statusRequest := cache.StatusRequest{
			Name: "Hello from golang client",
		}

		response, err := client.GetStatus(context.Background(), &statusRequest)
		if err != nil {
			log.Fatalf("Error when calling GetStatus: %s", err)
		}

		log.Printf("Response from server: %s", response.Message)
	} else if *actionPtr == "store" {
		// ensure file exists
		if !fileExists(*fileInputPtr) {
			fmt.Printf("File can not be found: %s\n", *fileInputPtr)
		}

		file, err := os.Open(*fileInputPtr)
		if err != nil {
			log.Fatalf("Error opening file: %s", err)
		}
		defer file.Close()

		// open a stream-based connection with the
		// gRPC server
		var stats cache.Stats
		stream, err := client.Upload(context.Background())
		if err != nil {
			log.Fatalf("Failed to upload file: %s", err)
		}
		defer stream.CloseSend()
		stats.StartedAt = time.Now()
		buf = make([]byte, 32) // 1 = chunk size
		for writing {
			n, err = file.Read(buf)
			if err != nil {
				if err == io.EOF {
					writing = false
					err = nil
					continue
				}
				log.Fatalf("Error reading chunk: %s", err)
			}
			_, fileIdentifer := filepath.Split(file.Name())
			err = stream.Send(&cache.Chunk{
				Content:    buf[:n],
				Identifier: fileIdentifer,
			})
			if err != nil {
				if err == io.EOF {
					writing = false
					err = nil
				} else {
					log.Fatalf("error sending chunk: %s", err)
				}

			}
		}

		stats.FinishedAt = time.Now()
		status, err = stream.CloseAndRecv()

		if err != nil {
			log.Fatalf("Error on closing stream: %s", err)
		}

		if status.Code != cache.UploadStatusCode_Ok {
			fmt.Printf("Status code not valid: %s", status.Code)
		}

	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
