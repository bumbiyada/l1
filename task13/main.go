package main

import "fmt"

// first variant by using Python-like syntax
func swap1(a, b *int) {
	*a, *b = *b, *a
}

// second variant by using XOR
func swap2(A, B *int) {
	*A ^= *B
	*B ^= *A
	*A ^= *B
}

//third variant by arithmetic MAGIC
func swap3(A, B *int) {
	*A = *A + *B
	*B = *A - *B
	*A = *A - *B
}
func main() {
	var A = 64
	var B = 128
	swap1(&A, &B)
	fmt.Println(A, B)

	swap2(&A, &B)
	fmt.Println(A, B)
	swap3(&A, &B)
	fmt.Println(A, B)
}
