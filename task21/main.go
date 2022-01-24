package main

import (
	"fmt"
	"strings"
)

// type client struct {
// }

// func (c *client) insertLightningConnectorIntoComputer(com computer) {
// 	fmt.Println("Client inserts Lightning connector into computer.")
// 	com.insertIntoLightningPort()
// }

// type computer interface {
// 	insertIntoLightningPort()
// }

// type mac struct {
// }

// func (m *mac) insertIntoLightningPort() {
// 	fmt.Println("Lightning connector is plugged into mac machine.")
// }

// type windows struct{}

// func (w *windows) insertIntoUSBPort() {
// 	fmt.Println("USB connector is plugged into windows machine.")
// }

// type windowsAdapter struct {
// 	windowMachine *windows
// }

// func (w *windowsAdapter) insertIntoLightningPort() {
// 	fmt.Println("Adapter converts Lightning signal to USB.")
// 	w.windowMachine.insertIntoUSBPort()
// }

// func main() {

// 	client := &client{}
// 	mac := &mac{}

// 	client.insertLightningConnectorIntoComputer(mac)

// 	windowsMachine := &windows{}
// 	windowsMachineAdapter := &windowsAdapter{
// 		windowMachine: windowsMachine,
// 	}

// 	client.insertLightningConnectorIntoComputer(windowsMachineAdapter)
// }

type client struct {
}

func (c *client) operation(com enter) {
	fmt.Println("Client make some work.")
	com.Upper()
}

type enter interface {
	Upper()
}

type str struct {
	str string
}

func (s *str) Upper() {
	fmt.Println("Doing our string operations ", s.str)
	str := strings.ToUpper(s.str)
	fmt.Println(str)
}

type runi struct {
	val []rune
}

func (w *runi) Convert(a []rune) (res string) {
	fmt.Println("Converting to string")
	res = string(w.val)
	return res
}

type runeAdapter struct {
	Runeval *runi
}

func (a *runeAdapter) Upper() {
	fmt.Println("Adapter converts rune to str")
	str := a.Runeval.Convert(a.Runeval.val)
	str = strings.ToUpper(str)
	fmt.Println(str)
}

func main() {

	client := &client{}
	str := &str{"aboba"}

	client.operation(str)

	Runeval := &runi{[]rune("abiba")}
	Adapter := &runeAdapter{
		Runeval: Runeval,
	}

	client.operation(Adapter)
}
