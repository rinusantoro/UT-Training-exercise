package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	Name string `json:"name"`
	Age string `json:"age"`
	Address string `json:"address"`
}

func main(){
	slice := []string{"rinusantoro","sleman","31","sonasa","jogja","25"}

	var jsonData, err = json.Marshal(slice)

	if err != nil{
		fmt.Println("Error marshalling json " + err.Error())
		return
	}
	
	var JsonString = string(jsonData)
	fmt.Println(JsonString)

	//for index, item := range a {
	//	fmt.Printf("%v: %v\n", index, item)
	//}


}