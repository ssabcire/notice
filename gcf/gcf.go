package slack

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

type payload struct {
	Text string `json:"text"`
}

func send(webhookURL string, text string) (err error) {
	p, err := json.Marshal(payload{Text: text})
	if err != nil {
		return err
	}
	resp, err := http.PostForm(webhookURL, url.Values{"payload": {string(p)}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Pubsubからメッセージを受け取ってSlackにメッセージを送る
func Do(ctx context.Context, m PubSubMessage) error {
	text := string(m.Data)
	if text == "" {
		log.Print("テキストが空でした")
	}
	webhookURL := os.Getenv("WEBHOOK_URL")
	err := send(webhookURL, text)
	if err != nil {
		log.Print("send error. err: ", err)
	}
	log.Printf("Slackへ送信が完了しました。text:%s", text)
	return nil
}
