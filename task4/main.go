package main

import (
	"flag"
	"fmt"
	"sync"
)

// First variant
func Square(n int) (res int) {
	return (n * n)
}
func main() {
	var WG sync.WaitGroup
	flag.Parse()
	// making new channel with buffer = 5, it`s important
	var c = make(chan int, 5)
	var summ int = 0
	// starting goroutines in cycle
	for i := 2; i <= 10; i += 2 {
		WG.Add(1)
		go func(i int, WG *sync.WaitGroup, c chan int) {
			res := Square(i)
			c <- res
			WG.Done()
		}(i, &WG, c)

	}
	WG.Wait()

	// reading values from channel
	for i := 0; i < 5; i++ {
		tmp := <-c
		summ = summ + tmp
	}
	close(c)
	// Printing result
	fmt.Println(summ)
}
