package main

import (
	"fmt"
)

type Human struct {
	Name   string
	Age    int
	Height int
	Weight int
}

type Action struct {
	Human
	Act string
}

func (a *Action) New(name string, age int, height int, weight int) {
	a.Name = name
	a.Age = age
	a.Height = height
	a.Weight = weight
	fmt.Printf("Created new human, name= %s, age= %v, height= %v, weight= %v\n", a.Name, a.Age, a.Height, a.Weight)
}
func (h *Human) Birthday() {
	h.Age = h.Age + 1
	fmt.Println("HOOORAY IT`S BIRTHDAY ! ", h.Name)
}
func (h *Human) WeightMyself() {
	if h.Weight < 55 {
		fmt.Println("Your weight less than normal ", h.Weight)
	} else if h.Weight > 100 {
		fmt.Println("Your weight is above normal ", h.Weight)
	} else {
		fmt.Println("Well Done, your weight is OK ", h.Weight)
	}

}

func main() {

	// call method of Human through Action
	a := Action{}
	a.New("Nicolay", 25, 180, 77)
	a.Birthday()
	a.WeightMyself()
}
