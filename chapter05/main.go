package main

import "fmt"

//var name string
//func init(){
//	fmt.Println("init:init")
//	name ="Test"
//}
//
//func main(){
//	fmt.Println("init:main")
//	fmt.Println(name)
//}
var whatisTeh = AnswerMe()
func AnswerMe() int  {
	fmt.Println("called,answer me")
	return 5
}
func init(){
	fmt.Println("called,init")
	whatisTeh = 0
}
func main(){
	fmt.Println("called,main")
	if whatisTeh ==0{
		fmt.Println("here we go!",whatisTeh)
	}
}

