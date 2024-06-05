package lib

import (
	"fmt"
	amqplib "github.com/root27/go-rabbit"
	"github.com/slack-go/slack"
)

type SlackNotify struct {
	RabbitHost   string
	RabbitQueue  string
	SlackToken   string
	SlackChannel string
}

func (s *SlackNotify) HandleRabbit() {

	conn, err := amqplib.Connect(s.RabbitHost)

	if err != nil {

		fmt.Printf("Error connecting rabbitmq:%s\n", err.Error())

		return
	}

	fmt.Println("Connected rabbitmq")

	fmt.Println("[X] Messages waiting...")

	msgs, err := amqplib.Receive(conn, s.RabbitQueue)

	if err != nil {

		fmt.Println("Error receiving messages: ", err)

		return
	}

	for d := range msgs {

		dataByte, ok := d.([]byte)

		if !ok {

			fmt.Println("Type assertion error: ", err)

			return

		}

		fmt.Println("Received message: ", string(dataByte))

		s.SlackNotify(string(dataByte))
	}

}

func (s *SlackNotify) SlackNotify(msg string) {

	api := slack.New(s.SlackToken)

	_, _, err := api.PostMessage(s.SlackChannel, slack.MsgOptionText(msg, false))

	if err != nil {

		fmt.Println("Error sending message to slack channel: ", err)

		return

	}

	fmt.Printf("Message sent to slack channel %s --> %s\n", s.SlackChannel, msg)
}
