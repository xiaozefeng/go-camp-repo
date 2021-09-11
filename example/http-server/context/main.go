package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello",  hello)
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, r *http.Request){
	ctx:= r.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello hadnler ended")
	select {
	case <-time.After(10 *time.Second):
		fmt.Fprintf(w, "hello\n")
		
	}
}


