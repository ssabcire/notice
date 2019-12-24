package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/slack/line", lineHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Print("error to ListenAndServe.", err)
	}
	log.Printf("サーバーを起動しました")
}

func lineHandler(w http.ResponseWriter, r *http.Request) {
	kind := "Slack"
	name := "Line"
	dsText, err := dsGet(kind, name)
	if err != nil {
		log.Print("Error to dsGet. ", err)
		return
	}
	scText, err := line_scraping()
	if err != nil || scText == "" {
		log.Print("Error to line_scraping. ", err)
		return
	}
	if scText != dsText {
		err := dsPut(kind, name, scText)
		if err != nil {
			log.Print("Error to Datastore Put. ", err)
			return
		}
		err = send(scText, mustGetenv("LINE_TOPIC"))
		if err != nil {
			log.Print("Error to send. ", err)
			return
		}
		log.Print("Send to message.")
	}
}
