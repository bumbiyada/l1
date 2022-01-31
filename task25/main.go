package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	sleep(4)
	fmt.Println(time.Now())
	fmt.Print("HELLO WORLD")
}

// aprooved function of sleep, simple cycle is in task5
func sleep(s int) {
	<-time.After(time.Second * time.Duration(s))
}
