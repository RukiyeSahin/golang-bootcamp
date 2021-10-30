package main

import "fmt"

const currency = "TL"
var Dolar,Euro = "$","€"

func struct_func(){

	type Player struct {
		Name	string
		Age		int
	}
	alper := Player{
		Name: "Alper",
		Age:  40,
	}
	alper.Age=42
	fmt.Printf("%v\n",alper)
}
func nesned_struct_func(){

	type Player struct {
		Name string
		Age	int
	}
	type Team struct {
		Name string
		Players [2]Player
	}
	myTeam := Team{
		Name:    "Devnot",
	}
	players := [...]Player{
		{			Name: "Ahmet",			Age:  20,		},
		{			Name: "Ali",			Age:  20,		},
	}
	myTeam.Players = players
	fmt.Printf("MyTeam: %v",myTeam)
}




func main()  {
	//struct_func()
	//nesned_struct_func()
}
