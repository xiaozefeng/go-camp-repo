package ratelimiter

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

type RateLimiter interface {
	Take() error
}

var _ RateLimiter = &SlidingWindowCounter{}

var once sync.Once

type SlidingWindowCounter struct {
	incurRequests    int32
	durationRequests chan int32
	accuracy         time.Duration
	snippet          time.Duration
	currentRequests  int32
	allowRquests     int32
}

var ErrLimitExceeded = errors.New("too many requests, exceeded the limit.")

func (swc *SlidingWindowCounter) Take() error {
	once.Do(func() {
		go sliding(swc)
		go calculate(swc)
	})
	currentRequests := atomic.LoadInt32(&swc.currentRequests)
	if currentRequests >= swc.currentRequests {
		return ErrLimitExceeded
	}
	if !atomic.CompareAndSwapInt32(&swc.currentRequests, currentRequests, currentRequests+1) {
		return ErrLimitExceeded
	}

	atomic.AddInt32(&swc.currentRequests, 1)

	return nil
}

func sliding(swc *SlidingWindowCounter) {
	for {
		select {
		case <-time.After(swc.accuracy):
			t := atomic.SwapInt32(&swc.currentRequests, 0)
			swc.durationRequests <- t
		}
	}
}

func calculate(swc *SlidingWindowCounter) {
	for {
		<-time.After(swc.accuracy)
		if len(swc.durationRequests) == cap(swc.durationRequests) {
			break
		}
	}

	for {
		<-time.After(swc.accuracy)
		t := <-swc.durationRequests
		if t != 0 {
			atomic.AddInt32(&swc.currentRequests, -t)
		}
	}
}

func NewSlidingWindowCounter(accuracy, snippet time.Duration, allowRequests int32) *SlidingWindowCounter {
	return &SlidingWindowCounter{
		durationRequests: make(chan int32, snippet/accuracy/1000),
		accuracy:         accuracy,
		allowRquests:     allowRequests,
		snippet:          snippet,
	}
}
