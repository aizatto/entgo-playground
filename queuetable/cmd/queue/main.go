package main

import (
	"context"
	"log"
	"os"
	"queuetable/internal/clients"
)

func main() {
	err := run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	entClient, err := clients.EntClient(ctx)
	if err != nil {
		return err
	}
	defer entClient.Close()

	// lock the row
	entClient.Tx()
}
