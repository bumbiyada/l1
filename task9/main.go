package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	br1 := make(chan bool)
	br2 := make(chan bool)
	arr := []int{4, 5, 6, 7, 8}
	var wg sync.WaitGroup
	fmt.Println("goroutine started")
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
	wg.Add(1)
	go func() {
		for {
			select {
			case val := <-ch1:
				res := val * val
				ch2 <- res
			case <-br1:
				wg.Done()
				fmt.Println("goroutine finished")
			default:
			}
		}

	}()
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
