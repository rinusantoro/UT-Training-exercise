package main

import "fmt"

func main(){
	var a = []string{"foo1","Bar1","foo2","bar2"}

	for index, item := range a {
		fmt.Printf("%v: %v\n", index, item)
	}


}