package main

import "fmt"

// first realistaion by Map[string]bool
func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	n := map[string]bool{}
	for _, val := range arr {
		if n[val] == false {
			n[val] = true
			fmt.Println("Added value in our map", val)
		} else if n[val] == true {
			fmt.Println("value is allready in our map", val)
		}
	}
	for key, val := range n {
		fmt.Printf("key = %s , val = %v\n", key, val)
	}
}
