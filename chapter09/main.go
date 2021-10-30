package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main1() {

	wg := sync.WaitGroup{}
	wg.Add(5)
	ping := func() {
		defer wg.Done()
		fmt.Print("Ping ")
	}
	for i := 0; i < 5; i++ {
		go ping()
	}
	wg.Wait()
	fmt.Println(",Pong")
}

func main2() {

	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		fmt.Printf("ölo ")
		wg.Done()
		<-c
	}
	const howmanyroutine = 100000
	wg.Add(howmanyroutine)
	before := memConsumed()
	for i := 0; i < howmanyroutine; i++ {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3f kb\n", float64(after-before)/howmanyroutine/1024)
}

func hello(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Printf("Hello from %d\n", i)
}

func main4() {

	wg := sync.WaitGroup{}
	const numberOfGreeters = 5
	wg.Add(numberOfGreeters)
	for i := 0; i < numberOfGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}

func main5() {

	wg := sync.WaitGroup{}
	const numberOfGreeters = 5
	wg.Add(numberOfGreeters)

	for i := 0; i < numberOfGreeters; i++ {
		go func(k int) {
			defer wg.Done()
			fmt.Printf("Hello from %d\n", k)
		}(i)
	}
	wg.Wait()
}
