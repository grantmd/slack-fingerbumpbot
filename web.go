package main

// Slack outgoing webhooks are handled here. Requests come in and are run through
// the processor to generate a response, which is sent back to Slack.
//
// Create an outgoing webhook in your Slack here:
// https://my.slack.com/services/new/outgoing-webhook

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type WebhookResponse struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		incomingText := r.PostFormValue("text")
		log.Printf("Handling incoming request: %s", incomingText)

		if incomingText == ":point_right:" {
			time.Sleep(5 * time.Second)

			var response WebhookResponse
			response.Username = botUsername
			response.Text = ":point_left:"
			log.Printf("Sending response: %s", response.Text)

			b, err := json.Marshal(response)
			if err != nil {
				log.Fatal(err)
			}

			w.Write(b)
		}
	})
}

func StartServer(port int) {
	log.Printf("Starting HTTP server on %d", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
