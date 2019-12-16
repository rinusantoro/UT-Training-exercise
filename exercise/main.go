package main

import (
	"fmt"
	"errors"
)

func main(){
	//x := 9.0
	//y := 8.0
	var x,y float64 	
	var operator string

	fmt.Println("Masukkan nilai x : ")
	fmt.Scan(&x)
	fmt.Println("Masukkan nilai y : ")
	fmt.Scan(&y)
	fmt.Println("Masukkan operator : ")
	fmt.Scan(&operator)

	//Exercise 1
	//Menggunakan func
	//resultAdd 		:= add(x,y)
	//resultMin		:= min(x,y)
	//resultMultiple 	:= multiple(x,y)
	//resultDivision 	:= division(x,y)

	//fmt.Println("y = ",y)
	//fmt.Println("x + y = ",resultAdd)
	//fmt.Println("x - y = ",resultMin)
	//fmt.Println("x * y = ",resultMultiple)
	//fmt.Println("x / y = ",resultDivision)

	result,err := Arithmatic(x,y, operator)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("Arithmatic = ",result,err)

}

/*Exercise 2
func add(x float64, y float64) (float64){
	result := x + y
	return result
}

func min(x float64, y float64) (float64){
	result := x - y
	return result
}

func multiple(x float64, y float64) (float64){
	result := x * y
	return result
}

func division(x float64, y float64) (float64){
	result := x / y
	return result
}
*/

func Arithmatic(x float64, y float64, s string) (float64, error){
	switch s {
	case "+" :
		val := x + y
		//fmt.Println("x + y = ",val)
		return val,nil
	case "-":
		val := x - y
		//fmt.Println("x - y = ",val)
		return val,nil
	case "*":
		val := x * y
		//fmt.Println("x * y = ",val)
		return val,nil
	case "/":
		val := x / y
		//fmt.Println("x / y = ",val)
		return val,nil
	default:

		return 0, errors.New("Input tanda Salah")
	
	}
	return 0,nil
}