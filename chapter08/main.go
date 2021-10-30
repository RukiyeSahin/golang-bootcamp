/*
Concurrency is not Parallelism by Rob Pike
https://www.youtube.com/watch?v=oV9rvDllKEg

GopherCon 2017: Kavya Joshi - Understanding Channels
https://www.youtube.com/watch?v=KBZlN0izeiY

GopherCon 2018: Kavya Joshi - The Scheduler Saga
https://www.youtube.com/watch?v=YHRO5WQGh0k

Extra: ITT 2019 - Kavya Joshi - Let's talk locks!
https://www.youtube.com/watch?v=tjpncm3xTTc

*/

package main

import (
	"fmt"
	"time"
)

func compute(value int){

	for i := 0; i < value; i++ {
		time.Sleep(time.Millisecond * 250)
		fmt.Println(time.Now(),i)
	}
}
func main(){
	//fmt.Println("lets count")
	//go compute(5)
	//go compute(5)
	//fmt.Scanln()

	go func() {
		fmt.Println("never hit me!")
	}()
	fmt.Scanln()

}
