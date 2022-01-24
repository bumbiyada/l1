package main

import "fmt"

func delete(slice []int, n int) (res []int) {
	if n >= len(slice) {
		fmt.Println("too high value")
		return nil
	} else {
		res := []int{}
		res = slice[0:n]
		c := slice[n+1:]
		res = append(res, c...)
		return res
	}

}
func main() {
	a := []int{4, 6, 2, 7, 4, 3}
	a = delete(a, 6)

	fmt.Println(a)
}
