package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		timestampBytes, err := ioutil.ReadFile("/usr/src/app/files/timestamp")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		timestamp := string(timestampBytes)
		timestampHash := string(hashBytes(timestampBytes))

		response := fmt.Sprintf("Timestamp: %s, Hash: %s", timestamp, timestampHash)

		if _, err := w.Write([]byte(response)); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func hashBytes(input []byte) []byte {
	h := sha1.New()
	h.Write(input)
	return h.Sum(nil)
}
