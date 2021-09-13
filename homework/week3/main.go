package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {

	srv := &http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello Go")
	})

	ctx, cancel := context.WithCancel(context.Background())
	g, errCtx := errgroup.WithContext(ctx)

	// deal shutdown
	g.Go(func() error {
		<-errCtx.Done()
		log.Println("stopping http server")
		//TODO: clean other resources
		return srv.Shutdown(errCtx)
	})

	// start http server
	g.Go(func() error {
		log.Print("staring http server at address :8080\n")
		return srv.ListenAndServe()
	})

	// deal linux system signal
	g.Go(func() error {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-done:
			log.Println("sig:", sig)
			cancel()
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Printf("group error: %v\n", err)
	}
	log.Println("all group done!")
}
