package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 5)
	for i := 0; i < 5; i++ {
		go sendMessage(ch, i)
	}

	printMessages(ch)
}

func sendMessage(out chan<- string, n int) {
	for {
		date := time.Now().Format("15:04:05")
		out <- fmt.Sprintf("Message at: %s", date)
		fmt.Println("Sending message at:", date)
	}
}

func printMessages(in <-chan string) {
	for m := range in {
		fmt.Println(m)
		time.Sleep(1 * time.Second)
	}
}
