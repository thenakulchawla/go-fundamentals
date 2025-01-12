package channels

import (
	"fmt"
	"time"
)

func Unbuffured() {

	uc := make(chan int)
	go func() {
		fmt.Println("starting goroutine with unbuffered channel")
		uc <- 1
		uc <- 2
		fmt.Println("message sent to unbuffered channel")
	}()

	fmt.Println("sleep for 1 second")
	time.Sleep(1 * time.Second)
	fmt.Println("receiving from unbuffered channel")
	fmt.Println(<-uc)
	fmt.Println(<-uc)
}

// UnbufferedBlocked in an unbuffered channel,
// if we don't receive first, it gets blocked
// this will result in a fatal error
func UnbufferedBlocked() {

	uc := make(chan int)
	fmt.Println("receiving values")

	fmt.Println(<-uc)
	fmt.Println(<-uc)
	fmt.Println("sending after receiving")
	go func() {
		fmt.Println("starting goroutine with unbuffered channel")
		uc <- 1
		uc <- 2
		fmt.Println("message sent to unbuffered channel")
	}()

}

// UnbufferedGoroutine as long as the code to receive is before or runnable,
func UnbufferedGoroutine() {
	ch := make(chan string) // Unbuffered channel

	go func() {
		fmt.Println("sleeping")
		time.Sleep(1 * time.Second)
		fmt.Println("Sending...")
		ch <- "Hello from sender"
		fmt.Println("Message sent")
	}()

	fmt.Println("Receiving...")
	fmt.Println(<-ch) // This will block until the sender sends a value
	fmt.Println("Message received")
}
