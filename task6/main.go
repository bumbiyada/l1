package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx1, cancel1 := context.WithCancel(ctx)
	ctx2, cancel2 := context.WithTimeout(ctx, time.Second*3)
	defer cancel2()
	var wg sync.WaitGroup
	var close = make(chan bool)
	// first Case simple channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-close:
				log.Println("1 Goroutine closed by channel")
				return
			default:
				fmt.Println("1 Goroutine DOSOMETHING")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	// second case ctx.WithCancel
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				log.Println("2 Goroutine closed by context.WithCancel")
				return
			default:
				fmt.Println("2 Goroutine DOSOMETHING")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx1)
	// third case WithTimeout
	// ctx.WithDeadline is similar
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				log.Println("3 Goroutine closed by context.WithTimeout")
				return
			default:
				fmt.Println("3 Goroutine DOSOMETHING")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx2)
	// sigint listener
	wg.Add(1)
	go func(cancel1, cancel2 context.CancelFunc) {
		defer wg.Done()
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		for {
			select {
			case <-sigint:
				log.Println("got sigint")
				cancel1()
				cancel2()
				close <- true
				return
			default:
				time.Sleep(20 * time.Millisecond)
			}
		}
	}(cancel1, cancel2)
	wg.Wait()
}
