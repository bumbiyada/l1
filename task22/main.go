package main

import (
	"fmt"
	"math/big"
)

type mystruct struct {
	val *big.Float
}

func (m *mystruct) init(val float64) {
	m.val = big.NewFloat(val)
}
func (m *mystruct) plus(arg *mystruct) mystruct {
	var z big.Float
	z.Add(m.val, arg.val)
	var res mystruct
	res.val = &z
	return res
}

func (m *mystruct) minus(arg *mystruct) mystruct {
	var z big.Float
	z.Sub(m.val, arg.val)
	var res mystruct
	res.val = &z
	return res
}

func (m *mystruct) multiply(arg *mystruct) mystruct {
	var z big.Float
	z.Mul(m.val, arg.val)
	var res mystruct
	res.val = &z
	return res
}

func (m *mystruct) division(arg *mystruct) mystruct {
	var z big.Float
	z.Quo(m.val, arg.val)
	var res mystruct
	res.val = &z
	return res
}

func (m *mystruct) print() {
	fmt.Println(m.val.String())
}
func main() {
	var a, b, c mystruct
	a.init(2)
	b.init(3)
	fmt.Printf("result of plus ")
	c = a.plus(&b)
	c.print()
	fmt.Printf("result of minus ")
	c = a.minus(&b)
	c.print()
	fmt.Printf("result of multiply ")
	c = a.multiply(&b)
	c.print()
	fmt.Printf("result of division ")
	c = a.division(&b)
	c.print()
}
