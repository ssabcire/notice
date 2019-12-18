package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
	"net/http"
	"os"
)

type Entity struct {
	Aaa string
}

func main() {
	http.HandleFunc("/datastore", dsHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Print("error to ListenAndServe.", err)
	}
	log.Print("ok!")
}

func dsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, mustGetenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Print("new client err. ", err)
		os.Exit(1)
	}
	k := datastore.NameKey("Task", "a", nil)
	e := new(Entity)
	if err := dsClient.Get(ctx, k, e); err != nil {
		log.Print("dsClient error. ", err)
		os.Exit(1)
	}
	e.Aaa, err = scraping()
	if err != nil || e.Aaa == "" {
		log.Print(err)
		e.Aaa = "scraping sippai"
	}
	if _, err := dsClient.Put(ctx, k, e); err != nil {
		log.Print("dsClient.Put error. ", err)
		os.Exit(1)
	}
	log.Print("Send to message.")
}
