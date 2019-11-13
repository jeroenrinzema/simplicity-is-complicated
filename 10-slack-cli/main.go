package main

import (
	"fmt"
	"bytes"
	"os"
	"net/http"
	"encoding/json"
)

// Message represents a Slack API message
type Message struct {
	Text string `json:"text"`	
}

func main() {
	webhook := os.Getenv("SLACK_WEBHOOK")
	message := Message{
		Text: "Never gonna give you up!",
	}

	bb, _ := json.Marshal(&message)
	http.Post(webhook, "application/json", bytes.NewBuffer(bb))

	fmt.Println("success!")
}