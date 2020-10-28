package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var counter int

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter += 1

		f, err := os.OpenFile("/usr/src/app/files/pingpong", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if _, err = f.WriteString(strconv.Itoa(counter)); err != nil {
			log.Fatal(err)
		}

		if _, err := w.Write([]byte(strconv.Itoa(counter))); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
