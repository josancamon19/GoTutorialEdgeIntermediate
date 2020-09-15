package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main()  {
	conn, err := amqp.Dial("amqps://xbkpuufb:VhwrEQe-R5p_vfngK_lasP8eMTrQnPgP@grouse.rmq.cloudamqp.com/xbkpuufb")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("Successfully Connected to our RabbitMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	messages, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	//
	forever := make(chan bool)
	go func() {
		for d := range messages {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever

}