package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Comment required to remove linting issues for exported struct (struct name in Upper Camel Case)

type Person struct {
	Fname  string `json: "fname"`
	Lname  string `json: "lname"`
	Gender string `json: "gender"`
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
		fmt.Printf("Age : %d \n", x.Height)
		fmt.Printf("Person : %s \n\n", x)

		// First name : Ayush
		// Age : 177
		// Person : {Ayush Arora M %!s(int=177)}

	}
}
