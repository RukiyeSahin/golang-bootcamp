package main

import (
	"fmt"
	"sync"
	"time"
)

func basic_channel_imp() {

	// a <<< ==== >>>> b
	//var dataStream chan interface{}
	//dataStream = make(chan interface{})

	// directional channel definations
	//var receiverChan <-chan interface{}
	//var senderChan chan<- interface{}
	//dataStream := make(chan interface{})
	//
	//receiverChan = dataStream
	//senderChan = dataStream
	strStream := make(chan string)
	fmt.Println("waiting...")
	go func() {
		time.Sleep(time.Second * 3)
		//close(strStream)
		//return
		strStream <- "Hello World"
	}()
	fmt.Println("waiting sub goroutine message..")

	//blocking!!!
	str, ok := <-strStream
	if ok {
		fmt.Println("message received...:", str, ok)
	} else {
		fmt.Println("channel closed", ok)
	}
	fmt.Println(str)
}

func channel_range() {
	intStream := make(chan int, 4)

	//sender
	go func() {
		defer close(intStream)
		defer fmt.Println("done\n")
		for i := 0; i < 5; i++ {
			fmt.Printf("Sending %d\n", i)

			intStream <- i
		}
	}()

	for val := range intStream {
		fmt.Printf("Got the value : %d\n", val)
	}

}

func channel_range_nonblocking() {
	fmt.Println("BLOCKING CHANNELS")
	channel_range()
	fmt.Println("NON-BLOCKING CHANNELS")
	begin := make(chan interface{})
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", x)
		}(i)
	}
	fmt.Println("unblocking goroutines")
	close(begin)
	wg.Wait()
}
