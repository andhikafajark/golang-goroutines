package golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(1 * time.Second)
		channel <- "Master"
		fmt.Println("Finish Send Data to Channel")
	}()

	data := <-channel

	fmt.Println(data)
	time.Sleep(1 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "Master"
	fmt.Println("Finish Send Data to Channel")
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel

	fmt.Println(data)
	time.Sleep(1 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "Master"
	fmt.Println("Finish Send Data to Channel")
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
	fmt.Println("Finish Get Data from Channel")
}

func TestChannelInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Master"
	channel <- "Zero"
	channel <- "One"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Finish")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Looping - " + strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println("Get data", data)
	}

	fmt.Println("Finish")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from Channnel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from Channnel 2", data)
			counter++
		default:
			fmt.Println("Waiting")
		}
		if counter == 2 {
			break
		}
	}
}
