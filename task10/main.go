package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -5.2, 6.4, 145.7}
	res := map[int][]float64{}
	for _, val := range arr {
		fmt.Printf("%v, ", val)
		tmp := val / 10
		tmp = math.Trunc(tmp) * 10 //or math.Floor
		res[int(tmp)] = append(res[int(tmp)], val)

	}
	// Print all our groups
	fmt.Println()
	for key, value := range res {
		fmt.Println("Key:", key, "Value:", value)
	}
}
