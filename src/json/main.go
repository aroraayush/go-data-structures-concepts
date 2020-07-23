// Go doesn't have the concept of classes
// It uses structs to create objects and do OOP
package main

import (
	"encoding/json"
	"fmt"
)

// Comment required to remove linting issues for exported struct (struct name in Upper Camel Case)
type Book1 struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Comment required to remove linting issues for exported struct (struct name in Upper Camel Case)
type Book2 struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

// Comment required to remove linting issues for exported struct (struct name in Upper Camel Case)
type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {

	// Calling like a constructor
	book1 := Book1{Title: "Harry Potter", Author: "J.K Rowling"}
	fmt.Println("book1: ", book1) // book1:  {Harry Potter J.K Rowling}

	author := Author{Name: "Ayush", Age: 25, Developer: true}
	fmt.Println("author: ", author) // author:  {Ayush 25 true}

	book2 := Book2{Title: "Harry Potter", Author: author}
	fmt.Println("book2: ", book2) // book2:  {Harry Potter {Ayush 25 true}}

	// Marshal returns the JSON encoding of book1
	byteArray, err := json.Marshal(book1)
	if err != nil {
		fmt.Println("err while json marshalling: ", err.Error())
	}
	fmt.Println("book1 json: ", string(byteArray)) // book1 json:  {"title":"Harry Potter","author":"J.K Rowling"}

	// Should re-use old variable, declaring new
	// Linting error : no new variables on left side of :=
	byteArray, err2 := json.Marshal(author)
	if err2 != nil {
		fmt.Println("err while json marshalling: ", err2.Error())
	}
	fmt.Println("author json: ", string(byteArray)) // author json:  {"name":"Ayush","age":25,"is_developer":true}

	// MarshalIndent(v, prefix,indentCount)
	// It is like Marshal but applies Indent to format the output.
	// Each JSON element in the output will begin on a new line beginning
	// with prefix followed by one or more copies of indent according
	// to the indentation nesting.
	byteArray3, err3 := json.MarshalIndent(book2, "", "  ")
	if err3 != nil {
		fmt.Println("err while json marshalling: ", err3.Error())
	}
	fmt.Println("book2 json indented: \n", string(byteArray3))
	// book2 json indented:
	// {
	// "title": "Harry Potter",
	// "author": {
	// 		"name": "Ayush",
	// 		"age": 25,
	// 		"is_developer": true
	// 	  }
	// }

}
