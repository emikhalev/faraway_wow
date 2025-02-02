package main

import (
	"context"
	"log"
	"net"

	"github.com/emikhalev/faraway_wow/pkg/client"
)

const serverAddress = "localhost:52345"

func main() {
	ctx := context.Background()

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer conn.Close()

	quote, err := client.WoW(ctx, conn)
	if err != nil {
		log.Fatal("cannot get quote from server: %w", err)
	}

	log.Printf("quote from server: %s", quote)
}
