package main

import (
	"log"
	"net/http"
	"os"
)

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
}

func dsHandler(w http.ResponseWriter, r *http.Request) {
	kind := "Task"
	name := "a"
	dsText, err := dsGet(kind, name)
	if err != nil {
		log.Print(err)
	}
	scText, err := line_scraping()
	if err != nil || scText == "" {
		log.Print(err)
		os.Exit(1)
	}
	if scText != dsText {
		err := dsPut(kind, name, scText)
		if err != nil {
			log.Print(err)
		}
		log.Print("Send to message.")
	}
}
