package minitask

import (
	"fmt"
)

type sender struct {
	Name    string
	Message string
}

var messages = []sender{
	{Name: "Alfan", Message: "Hello"},
	{Name: "Given", Message: "Test"},
}

func RunBoard() {
	message := make(chan sender)

	go whiteboard(message)

	for _, v := range messages {
		sendMessage(message, v)
	}
}

func sendMessage(mchan chan sender, messages sender) {
	mchan <- messages
}

func whiteboard(mchan chan sender) {
	for data := range mchan {
		fmt.Printf("From: %s, Message: %s \n", data.Name, data.Message)
	}
}
