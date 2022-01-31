package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {
	var WG sync.WaitGroup
	// making new channel with buffer = 5, it`s important
	var ch = make(chan int, 1)
	var c = make(chan bool)
	var N int = 10

	// Realisation of timer
	WG.Add(1)
	go func() {
		for i := 0; i < N; i++ {
			time.Sleep(1 * time.Second)
		}
		c <- true
		WG.Done()
	}()
	func() {
		log.Println("Starting ...")
		for {
			select {
			case <-c:
				fmt.Println("GOT value from chanel that tell`s me to stop program")
				os.Exit(0)
			case tmp := <-ch:
				fmt.Println("got ", tmp)
				time.Sleep(1 * time.Second)
			default:
				tmp := rand.Intn(100)
				ch <- tmp
				fmt.Println("posted ", tmp)
				time.Sleep(1 * time.Second)
			}

		}
	}()

	WG.Wait()
}
