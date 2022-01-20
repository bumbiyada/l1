package main

import (
	"fmt"
	"sync"
	"time"
)

// first variant make my own structure with map and mutex
//struct that has mutex and map, for example [int]int
type m struct {
	mu sync.Mutex
	m  map[int]int
}

// function to add value in map
func (v *m) add(i int) {
	v.mu.Lock()
	v.m[i] = i
	v.mu.Unlock()
}

// function to init map
func (v *m) init() {
	v.m = map[int]int{}
}

func main() {
	var m m

	// second variant to use existing sync.Map that has allready mutex and some methods
	var m2 sync.Map
	m.init()
	go func() {
		for i := 0; i < 10; i++ {
			m.add(i)
			m2.Store(i, i)
		}
	}()
	go func() {
		for i := 10; i < 20; i++ {
			m.add(i)
			m2.Store(i, i)
		}
	}()
	time.Sleep(3)
	for i := 0; i < 20; i++ {
		fmt.Println(m.m[i])
		fmt.Println(m2.Load(i))
	}
}
