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

// Timeouts func
func Timeouts() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result_1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result_2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout 2")
	}
}
