package main

import (
	"fmt"
	"strings"
)

//First variant
func swap1(str string) {
	tmp := []rune(str)
	res := []string{}
	buf := []rune{}
	for _, val := range tmp {
		if val == 32 {
			res = append(res, string(buf))
			buf = nil
		} else {
			buf = append(buf, val)
		}
	}
	res = append(res, string(buf))
	for i := (len(res) - 1); i >= 0; i-- {
		fmt.Printf("%s ", res[i])
	}
	var b strings.Builder
	for i := (len(res) - 1); i >= 0; i-- {
		//b.WriteString(res[i])
		fmt.Fprintf(&b, "%s ", res[i])
	}
	r := b.String()
	fmt.Println(r)
}

//Second Variant
func swap2(str string) {
	tmp := []rune(str)
	res := []string{}
	buf := []rune{}
	var p1 int = 0
	var p2 int = 0
	for idx, val := range tmp {
		if val == 32 {
			p1 = p2
			p2 = idx
			buf = tmp[p1:p2]
			res = append(res, string(buf))
			p2++
			buf = nil
		}
	}
	p1 = p2
	p2 = len(tmp)
	buf = tmp[p1:p2]
	res = append(res, string(buf))
	for i := (len(res) - 1); i >= 0; i-- {
		fmt.Printf("%s ", res[i])
	}
}
func main() {
	var str string = `snow dog sun`
	swap2(str)
	fmt.Println()
	swap1(str)
}
