package main

import (
	"fmt"
)

type Employee interface {
	GetName() string
	SetName(string)
}
type engineer struct {
	name string
}
func (receiver *engineer) GetName()string  {
	return receiver.name
}
func (receiver *engineer) SetName(s string) {
	receiver.name = s
}
func DoSomething() interface{}{

	return 5
}
func main(){
	 //e :=engineer{name: "Alper"}
	e := CreateEngineer("Alper")
	e.SetName("Muhammed")
	fmt.Printf("%s\n",e.GetName())
	 val := DoSomething()

	 fmt.Println( 5 + val.(int))
}

// factory pattern
func CreateEngineer(name string) Employee{
	instance := &engineer{name: name}
	fmt.Printf("%v\n",instance)
	return instance
}



