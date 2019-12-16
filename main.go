package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age string `json:"age"`
}

func main() {
	var JsonString = `{"name": "rinu", "age": "31"}`
	var JsonData = []byte(JsonString)

	var data User

	var err = json.Unmarshal(JsonData, &data)

	if err != nil{
		fmt.Println("Error Unmarshalling json " + err.Error())
		return
	}

	fmt.Println(data)

	fmt.Println("My name is " + data.Name)
	fmt.Println("My age is " + data.Age)
}
