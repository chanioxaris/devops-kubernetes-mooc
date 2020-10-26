package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
)

func main() {
	http.HandleFunc("/", indexPageHandler())

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	fmt.Println("Server started in port 8080")

	wg.Wait()
}

const indexPage = `
<html>
  <h1>Hello from project server</h1>
</html>
`

func indexPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("foo").Parse(indexPage)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if err = t.Execute(w, nil); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}
