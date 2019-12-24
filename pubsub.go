package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"sync"
)

// PubSub用
var messagesMu sync.Mutex

// TopicにPublishする
func send(message string, topicName string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, mustGetenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		return err
	}
	topic := client.Topic(topicName)
	exists, err := topic.Exists(ctx)
	if err != nil || !exists {
		return err
	}
	msg := &pubsub.Message{
		Data: []byte(message),
	}
	messagesMu.Lock()
	defer messagesMu.Unlock()
	if _, err := topic.Publish(ctx, msg).Get(ctx); err != nil {
		return err
	}
	return nil
}
