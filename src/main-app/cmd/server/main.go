package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	message := uuid.New().String()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		timestampBytes, err := ioutil.ReadFile("/usr/src/app/files/timestamp")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		pingPongBytes, err := ioutil.ReadFile("/usr/src/app/files/pingpong")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		response := fmt.Sprintf("%s: %s\nPing / Pongs: %s", string(timestampBytes), message, string(pingPongBytes))

		if _, err := w.Write([]byte(response)); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
