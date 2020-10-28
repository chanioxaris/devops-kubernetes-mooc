package main

import (
	"log"
	"os"
	"time"
)

func main() {
	for {
		f, err := os.OpenFile("/usr/src/app/files/timestamp", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		if _, err = f.WriteString(time.Now().UTC().String()); err != nil {
			log.Fatal(err)
		}

		if err = f.Close(); err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second * 5)
	}
}
