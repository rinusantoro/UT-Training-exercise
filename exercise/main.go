package main

import (
	"fmt"
)

func main(){
	x := 9
	y := 8 

	// Manual
	//add := x + y
	//min := x - y
	//multiple := x * y
	//division := x / y

	//fmt.Println("x + y = ", add)
	//fmt.Println("x - y = ", min)
	//fmt.Println("x * y = ", multiple)
	//fmt.Println("x / y = ", division)

	//Menggunakan func
	resultAdd 		:= add(x,y)
	resultMin		:= min(x,y)
	resultMultiple 	:= multiple(x,y)
	resultDivision 	:= division(x,y)

	fmt.Println("x = ",x)
	fmt.Println("y = ",y)
	fmt.Println("x + y = ",resultAdd)
	fmt.Println("x - y = ",resultMin)
	fmt.Println("x * y = ",resultMultiple)
	fmt.Println("x / y = ",resultDivision)

}

func add(x int, y int) (int){
	result := x + y
	return result
}

func min(x int, y int) (int){
	result := x - y
	return result
}

func multiple(x int, y int) (int){
	result := x * y
	return result
}

func division(x int, y int) (int){
	result := x * y
	return result
}