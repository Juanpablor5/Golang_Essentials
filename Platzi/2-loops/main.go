package main

import "fmt"

func main() {
	//For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("==")

	// For while
	cnt := 0
	for cnt < 10 {
		fmt.Println(cnt)
		cnt++
	}
}
