package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("This is a test log generated by go program every 30 seconds")
		time.Sleep(30 * time.Second)
	}
}