package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/root27/slacknotifyMQ/controllers"
	"os"
	"os/signal"
	"syscall"
)

var (
	welcomeMessage = `
	
	
   _____ _            _    _   _       _   _  __       __  __  ____  
  / ____| |          | |  | \ | |     | | (_)/ _|     |  \/  |/ __ \ 
 | (___ | | __ _  ___| | _|  \| | ___ | |_ _| |_ _   _| \  / | |  | |
  \___ \| |/   /|/ __| |/ / .   |/ _ \| __| |  _| | | | |\/| | |  | |
  ____) | | (_| | (__|   <| |\  | (_) | |_| | | | |_| | |  | | |__| |
 |_____/|_|\__,_|\___|_|\_\_| \_|\___/ \__|_|_|  \__, |_|  |_|\___\_\
                                                  __/ |              
                                                 |___/               


	SlackNotifyMQ - Bringing your RabbitMQ alerts to Slack!
`
)

func main() {

	fmt.Println(welcomeMessage)

	var url string

	fmt.Print(color.YellowString("Enter the rabbitMQ host address (default: amqp://guest:guest@localhost:5672) : "))

	fmt.Scanln(&url)

	if url == "" {

		url = "amqp://guest:guest@localhost:5672/"

	}

	var queueName string

	fmt.Print(color.YellowString("Enter the rabbitMQ queue name: "))

	fmt.Scanln(&queueName)

	var slackToken string

	fmt.Print(color.YellowString("Enter the token of your slack: "))

	fmt.Scanln(&slackToken)

	var slackChannel string

	fmt.Print(color.YellowString("Enter the channel ID of slack that notifications will be sent to (e.g. #general) : "))

	fmt.Scanln(&slackChannel)

	s := lib.SlackNotify{
		RabbitHost:   url,
		RabbitQueue:  queueName,
		SlackToken:   slackToken,
		SlackChannel: slackChannel,
	}

	stopChan := make(chan os.Signal, 1)

	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stopChan
		fmt.Println(color.RedString("Notifier Stopped"))
		os.Exit(0)
	}()

	s.HandleRabbit()

}
