package main

import "fmt"

// iterate by symbols []rune
func main() {
	var a = `главрыба`
	b := []rune(a)
	l := len(b)
	for i := 0; i < (l / 2); i++ {
		b[i], b[l-i-1] = b[l-i-1], b[i]
	}
	fmt.Println(string(b))
}
