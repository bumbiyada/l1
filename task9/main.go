package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	br1 := make(chan bool)
	br2 := make(chan bool)
	arr := []int{4, 5, 6, 7, 8}
	var wg sync.WaitGroup
	fmt.Println("goroutine started")

	// get data from array and push it to first channel
	wg.Add(1)
	go func() {
		for _, val := range arr {
			fmt.Printf("sending from array \t %v\n", val)
			ch1 <- val
		}
		wg.Done()
		fmt.Println("goroutine finished")
	}()
	fmt.Println("goroutine started")
	t := time.Now()

	// get data from 1-st chan, multiply it and push to 2-nd chan
	wg.Add(1)
	go func() {
		incr := 0
		for {
			select {
			case val := <-ch1:
				res := val * val
				ch2 <- res
			case <-br1:
				wg.Done()
				fmt.Println(incr)
				fmt.Println(time.Since(t).Microseconds())
				fmt.Println("goroutine finished")
			default:
				time.Sleep(time.Millisecond * 20)
				incr++
			}
		}

	}()

	// print values from chan 2 to standard output
	wg.Add(1)
	fmt.Println("goroutine started")
	go func() {
		for {
			select {
			case val := <-ch2:
				fmt.Printf("printing multiply \t%v\n", val)
			case <-br2:
				wg.Done()
				fmt.Println("goroutine finished")
			default:

			}
		}
	}()
	// goroutine that listens os. SIGINT and interrupts all goroutines
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		for {
			select {
			case sig := <-c:
				fmt.Println("got signal to interrupt\t\t", sig)
				br1 <- true
				br2 <- true
			default:
			}
		}
	}()
	fmt.Println("waiting")

	wg.Wait()
	fmt.Println("APP EXITED SUCESSFULY")
}
