package main

import (
	"cloud.google.com/go/datastore"
	"context"
)

type Entity struct {
	Title string
}

// Datastoreの前処理
func initialize() (ctx context.Context, client *datastore.Client, err error) {
	ctx = context.Background()
	client, err = datastore.NewClient(ctx, mustGetenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		return nil, nil, err
	}
	return ctx, client, nil
}

// Datastoreからデータを取得
func dsGet(kind string, name string) (string, error) {
	ctx, client, err := initialize()
	if err != nil {
		return "", err
	}
	k := datastore.NameKey(kind, name, nil)
	e := new(Entity)
	if err := client.Get(ctx, k, e); err != nil {
		return "", err
	}
	return e.Title, nil
}

// Datastoreに引数sを保存
func dsPut(kind string, name string, s string) error {
	ctx, client, err := initialize()
	if err != nil {
		return err
	}
	k := datastore.NameKey(kind, name, nil)
	e := &Entity{Title: s}
	if _, err := client.Put(ctx, k, e); err != nil {
		return err
	}
	return nil
}
