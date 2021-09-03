package concurrency

import "testing"

func TestBatchGet(t *testing.T) {
	var urls = []string{
		"http://127.0.0.1:8080/hello",
		"http://127.0.0.1:8080/hello",
		"http://127.0.0.1:8080/hello",
	}
	BatchGet(urls)
}
