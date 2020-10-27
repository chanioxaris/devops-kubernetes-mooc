package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	var counter int

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(strconv.Itoa(counter))); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}

		counter += 1
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
