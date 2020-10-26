package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	message := uuid.New()

	for {
		fmt.Printf("%v: %v\n", time.Now().UTC(), message)
		time.Sleep(time.Second * 5)
	}
}
