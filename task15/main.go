package main

import "fmt"

// troubles because UTF-8 , ASCII symbols take 1 byte, other unicode symbols may take up to 4 bytes, that`s why
// take slice of string is not a good idea because we itterate on bytes, not on symbols
// var justString string
// func someFunc() {
// v := createHugeString(1 << 10)
// justString = v[:100]
// }

// func main() {
// someFunc()git
func main() {
	//example of crash
	var a = `превед медвед`
	var b = a[:5]
	fmt.Println(b)
	// correct example, now while using []rune we itterate by symbol, not by bytes
	var c = []rune(a)
	d := c[:5]
	fmt.Printf("%s", string(d))
}
