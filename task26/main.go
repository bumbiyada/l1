package main

import (
	"fmt"
	"strings"
)

func Uniq(str string) bool {
	str = strings.ToLower(str)
	var tmp []rune
	buff := map[string]bool{}
	tmp = []rune(str)
	for _, val := range tmp {
		//if buff[string(val)] == nil; buff[string(val)] = true
		fmt.Println(val)
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
	var str string = "aabcd"
	res := Uniq(str)
	fmt.Println(res)
}
