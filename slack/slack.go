package slack

import (
	"log"
	"strings"

	"net/http"
)

type Slack interface {
	SendMessage(message string) error
}

type slack struct {
	token   string
	channel string
	text    string
}

func NewSlack(token, channel string) Slack {
	return &slack{
		token:   token,
		channel: channel,
	}
}

func (s *slack) SendMessage(message string) error {
	s.text = message
	return s.sendRequest()
}

func (s *slack) sendRequest() error {
	req, err := http.NewRequest(
		"POST",
		"https://slack.com/api/chat.postMessage",
		strings.NewReader(`{"channel":"`+s.channel+`","text":"`+s.text+`"}`),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.token)
	log.Println("slack", req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
