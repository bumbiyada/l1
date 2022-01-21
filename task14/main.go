package main

import (
	"fmt"
	"reflect"
)

//First variant by using fmt.Sprintf
func type1(obj interface{}) {
	tmp := fmt.Sprintf("%T", obj)
	fmt.Println(tmp)
}

// Second variant by reflect
func type2(obj interface{}) {
	tmp := reflect.TypeOf(obj)
	fmt.Println(tmp)
}

// Third variant by switch case and Type Assertion
func type3(obj interface{}) {
	switch v := obj.(type) {
	case int:
		fmt.Println("int", v)
	case float64:
		fmt.Println("float64", v)
	case float32:
		fmt.Println("float32", v)
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("bool", v)
	}
}
func main() {
	var b []interface{}
	b = append(b, 43, "grg", 545.54, true)
	fmt.Println(b)

	for idx, val := range b {
		fmt.Println(idx, val)
		type1(val)
		type2(val)
		type3(val)
	}
}
