package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	fileServer := http.FileServer(http.Dir("/usr/src/app/files"))

	http.HandleFunc("/", indexPageHandler())
	http.Handle("/usr/src/app/files/", http.StripPrefix("/usr/src/app/files", fileServer))

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := http.ListenAndServe(":3000", nil); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Server started in port 3000")

	wg.Wait()
}

const indexPage = `
<html>
	<h1>Hello from project server</h1>

	<img src="/usr/src/app/files/image.jpg" width="500" height="500"/>

	<br><br>	

	<form>
  		<input type="text" name="todo" maxlength="140"><br><br>
  		<input type="submit" value="Send">
	</form>

	<br><br>

	<h2>List of todos</h2>
	<ul>
	  <li>Fix todo form</li>
	  <li>Retrieve todos list</li>
	</ul>
</html>
`

func indexPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := downloadImage(); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

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

func downloadImage() error {
	if fileExists("/usr/src/app/files/image.jpg") {
		return nil
	}

	response, err := http.Get("https://picsum.photos/1200")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	f, err := os.Create("/usr/src/app/files/image.jpg")
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = io.Copy(f, response.Body); err != nil {
		return err
	}

	return nil
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}
