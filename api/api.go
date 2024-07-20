package api

import (
	"net/http"
	"strings"
)

type Api interface {
	SendMessage(message string) error
}

type api struct {
	endpoint string
	key      string
	text     string
}

func NewApiServer(endpoint, api_key string) Api {

	return &api{
		endpoint: endpoint,
		key:      api_key,
	}
}

func (s *api) SendMessage(message string) error {
	s.text = message
	return s.sendRequest()
}

func (s *api) sendRequest() error {
	req, err := http.NewRequest(
		"POST",
		s.endpoint,
		strings.NewReader(`{"text":"`+s.text+`"}`),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
