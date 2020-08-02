package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
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
	fmt.Println("Welcome to the CacheHub interactive shell!")
	fmt.Println("Pick on of the following ...")
	fmt.Println("   1 - Get Status Update of gRPC Server")
	fmt.Println("   2 - Upload a file to the cache")
	fmt.Println("   3 - Retrieve a file from the cache")
	fmt.Println("   4 - Exit")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer conn.Close()

	client := cache.NewCacheHubClient(conn)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		char, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println(err)
			break
		}

		if char == '1' {
			fmt.Println("Getting Status ...")
			statusRequest := cache.StatusRequest{
				Name: "Hello from golang client",
			}

			response, err := client.GetStatus(context.Background(), &statusRequest)
			if err != nil {
				log.Fatalf("Error when calling GetStatus: %s", err)
			}

			log.Printf("Response from server: %s", response.Message)
		} else if char == '2' {
			fmt.Println("Enter file name")
			//fileName = strings.Replace(fileName, "\n", "", -1)
			var fileName string
			fmt.Scanf("%s", &fileName)
			// ensure file exists
			if !fileExists(fileName) {
				fmt.Printf("File can not be found: %s\n", fileName)
			}

			file, err := os.Open(fileName)
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
			buf = make([]byte, 1) // 1 = chunk size
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
				err = stream.Send(&cache.Chunk{
					Content: buf[:n],
				})
				if err != nil {
					log.Fatalf("error sending chunk: %s", err)
				}
			}

			stats.FinishedAt = time.Now()
			status, err = stream.CloseAndRecv()

			if err != nil {

			}

			if status.Code != cache.UploadStatusCode_Ok {
				// log
			}

		} else if char == '4' {
			break
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
