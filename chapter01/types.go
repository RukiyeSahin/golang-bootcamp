package main

import (
	"fmt"
	"strconv"
)

func prim_types(){


	var myint16 int16
	var myint32 int32

	myint := 83

	fmt.Printf("%d( %T / %v) %d %d", myint,myint,myint,myint16,myint32)

	//myBool, yourBool, ourBool:=true,true,true
	//str := "aa"
}
func conversion_data_types(){

	var men uint8
	men =5
	var women int16
	women = 6

	var people int

	people = int(men) + int(women)
	fmt.Println(people)

	myFloat := float64(15.5)
	fmt.Println(myFloat)

	sstr := "15"

	myint,_ := strconv.Atoi(sstr)
	fmt.Println(myint)



}
func arrays(){
	var days []string
	days = []string{"Monday","Thueday","Wednesday","Thursday","Friday","Saturday","Sunday",}
	days[5] ="Cumartesi"
	//months :=[]string{}

	fmt.Printf("\t%V\n", days)
	fmt.Printf("\t%v\n", days[5])
	fmt.Printf("Len: %d\n", len(days))
}

func slides()  {
	days := [...]string{"Monday","Thueday","Wednesday","Thursday","Friday","Saturday","Sunday",}
	workdays := days[0:5]
	weekends := days[5:7]

	fmt.Println("workdays:",workdays, "weekends:",weekends)
}
func maps(){

	cities := map[string]int{
		"Ankara":6,
		"İstanbul":34,
		"İzmir": 35,
	}
	if val,ok := cities["Ankara"]; ok==true {
		fmt.Println(ok,val)
	}
}
func main(){
	//prim_types()
	///conversion_data_types()
	//arrays()
	//slides()
	//maps()
}
