package main

import (
	"fmt"
)

// NEED SOME FIXES
func main() {
	var i int64 = -33 //Число, в котором будем менять символ
	var a int64 = 3   //Позиция символа который будем менять. от 1 до 63
	var b int64
	b = 1 << (a - 1)
	fmt.Printf("start number \t%b\n", i) // print standart number in binary
	a1 := i ^ b
	fmt.Printf("result number \t%b\n", a1) // Измененное число
}

//-64
