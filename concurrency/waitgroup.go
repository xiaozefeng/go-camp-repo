package concurrency

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type httpPkg struct{}

func (httpPkg) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var h httpPkg

func BatchGet(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		url := url
		go func() {
			log.Printf("start request: %s", url)
			defer wg.Done()
			if content, err := h.Get(url); err != nil {
				log.Printf("get %s failed, error: %v", url, err)
			} else {
				log.Printf("request %s, response: %v", url, string(content))
			}
		}()
	}
	wg.Wait()
}
