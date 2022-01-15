package main

import (
	"fmt"
	"log"
)

type Human struct {
	Name   string
	Age    int
	Height int
	Weight int
}

type Action struct {
	Human Human
	Name  string
}

func (a *Action) New(n string) {
	a.Name = n
	log.Println(a.Name)
}
func (h *Human) Birthday() {
	h.Age = h.Age + 1
	log.Println("HOOORAY IT`S BIRTHDAY ! ", h.Name)
}
func main() {
	fmt.Println("Hello world")
	Maxim := Human{"Maxim", 24, 186, 72}
	Maxim.Age = 24
	Maxim.Birthday()
	log.Println("Maxim`s age = ", Maxim.Age)
	a := Action{}
	a.New("aboba")
}
