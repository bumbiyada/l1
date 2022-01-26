package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

// First variant
func Worker(number int, ch chan int, close chan bool) {
	//defer WG.Done()
	for {
		select {
		case <-close:
			fmt.Printf("Worker %v GOT SIGINT AND NEED TO BE CLOSED\n", number)
			return
		case val := <-ch:
			fmt.Printf("Worker %v printing value = %v\n", number, val)
		default:
			time.Sleep(200 * time.Millisecond)
		}
	}
}
func main() {
	flag.Parse()
	args := flag.Args()
	n, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println("got wrong arg")
	}
	var WG sync.WaitGroup
	var c = make(chan int, 5)
	var close = make(chan bool)

	WG.Add(1)
	go func(c chan int, close chan bool) {
		defer WG.Done()
		for {
			select {
			case <-close:
				fmt.Printf("Generator GOT SIGINT AND NEED TO BE CLOSED\n")
				return
			default:
				val := rand.Int()
				c <- val
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(c, close)
	for i := 1; i <= n; i++ {
		WG.Add(1)
		go func(i int, c chan int, close chan bool) {
			defer WG.Done()
			Worker(i, c, close)

		}(i, c, close)

	}
	WG.Add(1)
	go func(close chan bool) {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		for {
			select {
			case <-sigint:
				fmt.Println("got signal to interrupt\t\t")
				for i := 0; i < n+1; i++ {
					close <- true
				}

				WG.Done()
				return
			default:
				time.Sleep(50 * time.Millisecond)
			}
		}
	}(close)
	WG.Wait()
	log.Println("APP finished sucessfully")
}
