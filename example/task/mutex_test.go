package task

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type SafeMap struct {
	m    map[int]int
	lock *sync.Mutex
}

func newSafeMap() *SafeMap {
	return &SafeMap{
		m:    map[int]int{},
		lock: &sync.Mutex{},
	}
}

func (s *SafeMap) Set(key, value int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m[key] = value
}

func (s *SafeMap) Get(key int) int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.m[key]
}

func (s *SafeMap) print() {
	s.lock.Lock()
	s.lock.Unlock()
	fmt.Println("state", s.m)
}

func Test_mutex(t *testing.T) {
	var readOps uint64
	var writeOps uint64

	var sm = newSafeMap()

	for i := 0; i < 100; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				total += sm.Get(key)
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				sm.Set(key, val)
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps", writeOpsFinal)

	sm.print()

}

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func Test_stateful_goroutines(t *testing.T) {

	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp

				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Microsecond)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Microsecond)
			}
		}()

	}

	time.Sleep(time.Second)
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps", readOpsFinal)

	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps", writeOpsFinal)

}
