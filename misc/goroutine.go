package misc

import (
	"fmt"
	"time"
)

// CheckChannel passes a goroutine operation via channel
func CheckChannel() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

// CheckBufferedChannel checks a buffered channel operation
func CheckBufferedChannel() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// CheckChannelSynchronization function
func CheckChannelSynchronization(done chan bool) {
	fmt.Printf("Working . . .")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

// Synchronization func
func Synchronization() {
	done := make(chan bool, 1)
	go CheckChannelSynchronization(done)

	<-done
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// ChannelDirections func
func ChannelDirections() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "Passed Message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
