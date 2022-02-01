package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"

	fmt.Println(len(c), cap(c))

	// Range y close
	close(c) // Cierra el canal y no permite recibir mas datos, independientemente de que tenga capacidad

	//c <- "Mensaje3"
	for message := range c {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)
	go message("mensaje1", email1)
	go message("mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}

func say(text string, c chan<- string) {
	c <- text
}

func main2() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)

	fmt.Println(<-c)
}
