package main

import (
	"fmt"
	"time"
)

func select_statement_with_channel() {

	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(time.Second * 5)
		close(c)
	}()

	fmt.Println("blocking on read..")

	workCounter := 0
loop:
	for true {
		select {
		case <-c:
			break loop
		default:
		}
		workCounter++
		time.Sleep(time.Second * 1)
	}
	fmt.Printf("Kanal kapanana kadar %d kadar bekledi. Süre: %v", workCounter, time.Since(start))
}
