package main

import (
	"fmt"
)

// return 2 ^ n number if n = 1 -> return
func multiply(n int64) int64 {
	if (n < 0) || (n > 64) {
		return -1
	} else if n == 0 {
		return 0
	} else {
		var res int64 = 1
		for i := 1; i < int(n); i++ {
			res = res * 2
		}
		return res
	}

}
func main() {
	var i int64 = 64                     //Число, в котором будем менять символ
	var a int64 = 7                      //Позиция символа который будем менять. от 1 до 63
	b := multiply(a)                     // Получаем
	fmt.Printf("start number \t%b\n", i) // print standart number in binary
	a1 := i ^ b
	fmt.Printf("result number \t%b\n", a1) // Измененное число
}

//-64
