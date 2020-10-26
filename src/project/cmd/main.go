package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	fmt.Println("Server started in port 8080")

	wg.Wait()
}
