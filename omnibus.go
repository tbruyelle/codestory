package main

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

func (o *Omnibus) NextCommand() string {
	i := o.counter % len(commands)
	o.counter = o.counter + 1
	return commands[i]
}

func (o *Omnibus) Reset() {
o.counter=0}

func NewOmnibus() *Omnibus {
	return new(Omnibus)
}
