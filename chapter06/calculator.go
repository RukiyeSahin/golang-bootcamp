package main

import "errors"

func Calculator(x,y int) (int,error){

	if x< 0{
		return 0,errors.New("NOT COOL")
	}
	if y< 0{
		return 0,errors.New("NOT COOL")
	}

	return x+y,nil
}
