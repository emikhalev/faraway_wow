package main

import (
	"context"
	"flag"
	"github.com/emikhalev/faraway_wow/pkg/client"
	"log"
	"net"
)

//nolint:gochecknoglobals
var (
	addr = "localhost:52345"
)

//nolint:gochecknoinits
func init() {
	flag.StringVar(&addr, "addr", "", "path to config file")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	conn, err := net.Dial("tcp", addr)
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
