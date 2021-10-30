package main

import (
	"errors"
	"fmt"
)

func myFunction(firstName, lastName string) (fullname string,err error) {

	fullname = fmt.Sprintf("%s %s",firstName,lastName)
	err = errors.New("error occured")
	return
}

func addOne() func() int{
	var c int
	return func() int{
		c++
		return c+1
	}
}

func main()  {
	//p1,err := myFunction("Alper","Hankendi")
	//
	//fmt.Printf(p1,err)
	var myFunc = addOne()
	var myFunc2 = addOne()
	fmt.Println(myFunc()) //2
	fmt.Println(myFunc()) //3
	fmt.Println(myFunc())//4
	fmt.Println(myFunc2()) // 2 ? 5
	fmt.Println(myFunc2()) // 3 ? 6
	fmt.Println(myFunc2()) // 4 ? 7


}

