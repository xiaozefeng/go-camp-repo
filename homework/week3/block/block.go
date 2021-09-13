package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Hello, GopherCon SG")
	})

	go func() {
		log.Print("server started at 8080\n")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()


	select{}
}
