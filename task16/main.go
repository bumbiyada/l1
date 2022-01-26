package main

import (
	"fmt"
	"time"
)

var id int = 0

func sort_node(arr []int, low int, high int) {
	if low < high {
		id++
		p := partition_hoar(arr, low, high)
		//fmt.Println("partition", p)
		time.Sleep(20 * time.Millisecond)
		//fmt.Println("left node ", low, p-1)
		sort_node(arr, low, p)
		//fmt.Println("right node ", p+1, high)
		sort_node(arr, p+1, high)
	}

}
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

// Hoar partition
func partition_hoar(arr []int, low int, high int) (p int) {
	//p = arr[(low+high)/2]
	p = arr[low]
	fmt.Println(arr, "\t", p, "\t", low, "\t", high)

	i := low - 1
	j := high + 1
	for {
		for {
			j--
			if arr[j] >= p {
				break
			}
		}

		for {
			i++
			if arr[i] <= p {
				break
			}
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			return j
		}

	}
}

// just starts first node
func quicksort(arr []int) {

	sort_node(arr, 0, len(arr)-1)

}
func main() {
	arr := []int{4, 5, 3, 6, 3, 2, 65, 3, 4, 6, 8, 4, 8, 9, 56, 75, 78, 5, 5, 45, 535}
	quicksort(arr)
	fmt.Println(arr)
	fmt.Print(id)
}
