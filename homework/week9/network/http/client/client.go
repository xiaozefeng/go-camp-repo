package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   1 * time.Second,
	}
	resp, err := client.Get("http://127.0.0.1:8080")
	if err != nil {
		log.Printf("err = %+v\n", err)
		return
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("err = %+v\n", err)
		return
	}
	log.Printf("content = %s\n", content)

}
