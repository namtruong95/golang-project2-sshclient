package main

import (
	"fmt"
	"log"

	"github.com/melbahja/goph"
)

func main() {
	client := initClient()

	// Defer closing the network connection.
	defer client.Close()

	// Execute your command.
	out, err := client.Run("ls -la ~/.ssh")

	if err != nil {
		log.Fatal(err)
	}

	// Get your output as []byte.
	fmt.Println(string(out))
}

func initClient() *goph.Client {
	user := "hung-pc"
	host := "192.168.1.108"
	auth, _ := goph.Key("/Users/dsoft/.ssh/audio-book/audio-book-dev.pem", "")
	// Start new ssh connection with private key.
	client, err := goph.New(user, host, auth)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
