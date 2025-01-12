package channels

import "fmt"

func Buffered() {
	bc := make(chan int, 2)

	fmt.Println("sending values into buffered channel")
	bc <- 1
	bc <- 2

	fmt.Println("receiving values")
	fmt.Println(<-bc)
	fmt.Println(<-bc)
}
