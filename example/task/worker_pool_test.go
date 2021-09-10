package task

import (
	"fmt"
	"testing"
	"time"
)

func Test_worker_pool(t *testing.T) {
	var worker = func(id int, jobs, results chan int) {
		for job := range jobs {
			fmt.Println("worker", id, "started job ", job)
			time.Sleep(time.Second)
			fmt.Println("worker", id, "finished job", job)
			results <- job * 2
		}
	}

	const numOfJobs = 5
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	for i := 0; i < 3; i++ {
		go worker(i+1, jobs, results)
	}

	for i := 0; i < numOfJobs; i++ {
		jobs <- i + 1
	}
	close(jobs)

	for i := 0; i < numOfJobs; i++ {
		<-results
	}
}
