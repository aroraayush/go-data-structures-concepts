package main

import (
	"fmt"
	"net/http"
)

// const not in upper case becuase we don't
// want to expose it ot other packages
const port int = 8000
const path string = "/hello"

func main() {

	http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("<html><body><center><h1>Hello World</h1></center></body></html>"))
	})

	portString := fmt.Sprintf(":%d", port)

	fmt.Printf("Starting server at port %d listening at path %s", port, path)

	err := http.ListenAndServe(portString, nil)

	if err != nil {
		panic(err)
	}

}
