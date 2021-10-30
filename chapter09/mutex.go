package main

import (
	"fmt"
	"sync"
)

func mutex_sample() {

	var count int = 0
	var lock sync.Mutex
	increment := func() {
		defer lock.Unlock()
		lock.Lock()
		count++
		fmt.Printf("incrementing : %d\n", count)
	}
	decrement := func() {
		defer lock.Unlock()
		lock.Lock()
		count--
		fmt.Printf("decrementing : %d\n", count)
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}
	wg.Wait()
	fmt.Printf("Current count value : %d\n", count)
}

func main() {

	var lock sync.Mutex
	var count int

	inc := func() {
		lock.Lock()
		count++
		lock.Unlock()
	}
	once := sync.Once{}
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(inc)
			inc()
		}()
	}
	wg.Wait()
	fmt.Printf("Current count: %d\n", count)
}
