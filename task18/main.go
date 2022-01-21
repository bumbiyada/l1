package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// First variant - using my own struct with mutex
type incrementor struct {
	i  uint64
	mu sync.Mutex
}

// Second variant - using atomic

func (ic *incrementor) increment() {
	ic.mu.Lock()
	ic.i++
	ic.mu.Unlock()
}
func (ic *incrementor) print() {
	fmt.Printf("value of inceremtor = %v\n", ic.i)
}
func main() {
	var wg sync.WaitGroup
	var incr incrementor
	var atom int64
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			incr.increment()
			atomic.AddInt64(&atom, 1)
		}

	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			incr.increment()
			atomic.AddInt64(&atom, 1)
		}
	}()
	wg.Wait()
	incr.print()
	fmt.Println(atom)
}
