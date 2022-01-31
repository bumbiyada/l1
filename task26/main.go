package main

import (
	"fmt"
	"strings"
)

// find out if all chars in str are unique
func Uniq(str string) bool {
	str = strings.ToLower(str)
	var tmp []rune
	buff := map[string]bool{}
	tmp = []rune(str)
	for _, val := range tmp {
		b := buff[string(val)]
		if b == false {
			buff[string(val)] = true
		} else if b == true {
			return false
		}

	}
	return true
}

func main() {
	var str string = "abcd"
	res := Uniq(str)
	fmt.Println(res)
}
