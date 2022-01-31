package main

import (
	"fmt"
	"sync"
)

// First variant. multiply values and printing them
func Square(n int) {
	fmt.Println(n * n)
}

func main() {
	var WG sync.WaitGroup
	// starting goroutines in cycle
	for i := 2; i <= 10; i += 2 {
		WG.Add(1)
		go func(i int, WG *sync.WaitGroup) {
			Square(i)
			WG.Done()
		}(i, &WG)
	}
	WG.Wait()
}
