

package main

import (

	"io"
	"net/http"
)

func handler(writer http.ResponseWriter, req *http.Request) {

	// curr_string String := "Hello, " + req.URL.Path[1:]
	// curr_string String := req.URL.Path[1:]

	curr_string string := "Hello "

	io.WriteString(writer, curr_string )
}

func main() {

	http.HandleFunc("/", handler)

	http.ListenAndServ(":8888", nil)
}