package main

import (
	"context"
	"fmt"

	"log"

	ts "grpc-mongo/grpc/profile"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := ts.NewProfileServiceClient(conn)

	response, err := client.CreateProfile(context.Background(), &ts.Profile{Name: "Subasri"})
	if err != nil {
		log.Fatalf("Failed to call CreateProfile: %v", err)
	}

	fmt.Printf("Response: %s\n", response.Name)
}
