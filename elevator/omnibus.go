package elevator

import (
	"fmt"
)

type Omnibus struct {
	counter int
}

var commands = []string{
	"OPEN", "CLOSE", "UP",
	"OPEN", "CLOSE", "UP",
	"OPEN", "CLOSE", "UP",
	"OPEN", "CLOSE", "UP",
	"OPEN", "CLOSE", "UP",
	"OPEN", "CLOSE", "DOWN",
	"OPEN", "CLOSE", "DOWN",
	"OPEN", "CLOSE", "DOWN",
	"OPEN", "CLOSE", "DOWN",
	"OPEN", "CLOSE", "DOWN",
}

func (o *Omnibus) Debug(enabled bool) {
}

func (o *Omnibus) Ditdlamerde() {
}

func (o *Omnibus) NextCommand() string {
	i := o.counter % len(commands)
	o.counter = o.counter + 1
	return commands[i]
}

func (o *Omnibus) Reset(l,h,c int,s string) {
	o.counter = 0
}

func (o *Omnibus) Call(atFloor int, to string) {
	fmt.Println("call", atFloor)
}

func (o *Omnibus) Go(floorToGo int) {
	fmt.Println("go", floorToGo)
}

func (o *Omnibus) UserHasEntered() {
	fmt.Println("UserHasEntered")
}

func (o *Omnibus) UserHasExited() {
	fmt.Println("UserHasExited")
}

func NewOmnibus() *Omnibus {
	return new(Omnibus)
}
