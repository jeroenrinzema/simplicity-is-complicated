package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Message represents a Slack API message
type Message struct {
	Text string `json:"text"`
}

func manipulate(message *Message) {
	message.Text = message.Text + " Never gonna let you down."
}

func main() {
	webhook := os.Getenv("SLACK_WEBHOOK")
	message := Message{
		Text: "Never gonna give you up!",
	}

	manipulate(&message)

	bb, _ := json.Marshal(message)
	http.Post(webhook, "application/json", bytes.NewBuffer(bb))

	fmt.Println("success!")
}
