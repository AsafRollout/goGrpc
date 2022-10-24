package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"

	"goGrpc/speakers"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9876", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := speakers.NewAudioOutputClient(conn)

	fmt.Print("Welcome to chat")

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		_, err := c.Play(context.Background(), &speakers.Audio{Body: text})
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}
		//log.Printf("Response from server: %s", response.Body)
	}
}
