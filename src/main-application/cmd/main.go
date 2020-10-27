package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

func main() {
	message := uuid.New()

	wg := &sync.WaitGroup{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := fmt.Sprintf("%v: %v\n", time.Now().UTC(), message)

		if _, err := w.Write([]byte(response)); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := http.ListenAndServe(":3000", nil); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			fmt.Printf("%v: %v\n", time.Now().UTC(), message)
			time.Sleep(time.Second * 5)
		}
	}()

	wg.Wait()
}
