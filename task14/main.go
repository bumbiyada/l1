package main

import "fmt"

//First variant by using fmt.Sprintf
func type1(obj interface{}) {
	tmp := fmt.Sprintf("%T", obj)
	fmt.Println(tmp)
}

func type2(obj interface{}) {

}
func main() {
	var b []interface{}
	b = append(b, 43, "grg", 545.54, true)
	fmt.Println(b)

	for idx, val := range b {
		fmt.Println(idx, val)
	}
}
