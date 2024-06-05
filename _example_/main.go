package main

import (
	"fmt"

	amqplib "github.com/root27/go-rabbit"
)

func main() {

	conn, err := amqplib.Connect("amqp://guest:guest@localhost:5672/")

	if err != nil {

		fmt.Println("Error connection rabbit: ", err)

		return

	}

	defer conn.Close()

	fmt.Printf("connected to RabbitMQ\n")

	// Send a message
	msg := "Error log example"

	_, err = amqplib.Send(conn, "testing", []byte(msg), "text/plain")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("message sent: %s\n", msg)

}
