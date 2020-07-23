package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Comment required to remove linting issues for exported struct (struct name in Upper Camel Case)
// Modelling all fields in JSON file
type Person struct {
	Fname  string `json: "fname"`
	Lname  string `json: "lname"`
	Gender string `json: "gender"`
	Height int    `json: "height"`
}

// Modelling few fields in JSON file
type Person2 struct {
	Fname  string `json: "fname"`
	Height int    `json: "height"`
}

func main() {

	content, err := ioutil.ReadFile("json_file.json")
	if err != nil {
		fmt.Println("err while json marshalling: ", err.Error())
	}

	var persons []Person

	err2 := json.Unmarshal(content, &persons)
	if err2 != nil {
		fmt.Println("Err while json unmarshalling: ", err2.Error())
	}
	for _, x := range persons {
		fmt.Printf("First name : %s \n", x.Fname)
		fmt.Printf("Height : %d \n", x.Height)
		fmt.Printf("Person : %s \n\n", x)

		// First name : Ayush
		// Age : 177
		// Person : {Ayush Arora M %!s(int=177)}

	}
	
	content_partial, err3 := ioutil.ReadFile("json_file.json")
	if err3 != nil {
		fmt.Println("err while json marshalling: ", err3.Error())
	}
	var person2 Person2

	err2 := json.Unmarshal(content_partial, &person2)
	if err2 != nil {
		fmt.Println("err2 while json unmarshalling: ", err2.Error())
	}
	fmt.Printf("First name : %s \n", person2.Fname)
	fmt.Printf("Height : %d \n", person2.Height)
	fmt.Printf("Person : %s \n\n", person2)

	// First name : Ayush
	// Height : 177
	// Person : {Ayush %!s(int=177)}
}
