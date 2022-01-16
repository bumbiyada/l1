package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World")
	var i int64 = 63
	fmt.Println(i)
	fmt.Println(strconv.FormatInt(i, 2)) // print standart number in 101
	var c int64 = 4
	fmt.Println(strconv.FormatInt(c, 2)) // print c number in 101
	a1 := i ^ c
	fmt.Println(strconv.FormatInt(a1, 2)) // print i ^ c number in 101
	a2 := i & c
	fmt.Println(strconv.FormatInt(a2, 2)) // print i & c numnber in 101
}
