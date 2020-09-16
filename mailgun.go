package main

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go"
	"log"
	"time"
)

func main8() {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun("sandboxabf655fdbf5a48a1bea93d1d3c25c85d.mailgun.org", "123")

	sender := "joan.santiago.cabezas@gmail.com"
	subject := "Fancy subject!"
	body := "Hello from Mailgun Go!"
	recipient := "joan.santiago.cabezas@gmail.com"

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
