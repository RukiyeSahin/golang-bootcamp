package main

import (
	"fmt"
	"net/http"
)

func main() {

	type Result struct {
		Error    error
		Response *http.Response
		Url      string
	}

	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		responses := make(chan Result)

		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				result := Result{Error: err, Response: resp, Url: url}
				select {
				case <-done:
					return
				case responses <- result:
				}
			}
		}()

		return responses
	}

	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://www.trendyol.com", "https://badhostname"}
	//go func() {
	//	time.Sleep(time.Millisecond*500)
	//	fmt.Println("terminating")
	//	done<- true
	//}()
	for response := range checkStatus(done, urls...) {

		if response.Error != nil {
			fmt.Printf("Url:%s, Error: %v\n", response.Url, response.Error)
			continue
		}
		fmt.Printf("Url:%s,Response: %v\n", response.Url, response.Response.StatusCode)
	}

}
