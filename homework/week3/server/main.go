package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// go serveApp()
	// go serveDebug()

	done := make(chan error, 2)
	stop := make(chan struct{})

	go func() {
		done <- serveDebug(stop)
	}()

	go func() {
		done <- serveApp(stop)
	}()

	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Printf("err: %v\n", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}

func serveApp(stop chan struct{}) error {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(rw, "Hello Go")
	// })
	// log.Printf("start app ...\n")
	// return http.ListenAndServe(":8080", nil)
	return serve(":8080", &Simple{}, stop)
}

type Simple struct{}

func (s *Simple) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello Go")
}

func serveDebug(stop chan struct{}) error {
	// log.Printf("start debug ...\n")
	// return http.ListenAndServe(":8001", http.DefaultServeMux)
	return serve(":8001", &Simple{}, stop)
}

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()
	log.Printf("start server at %s", addr)
	return s.ListenAndServe()
}
