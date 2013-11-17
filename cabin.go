package main

import (
	"fmt"
)

type call struct {
	atFloor int
	to      byte
}

type Cabin struct {
	lowerFloor, higherFloor, currentFloor int
	calls                                 []call
	opened                                bool
}
var NOCALL   = call{-1, CALLUP}

const (
	OPEN     = "OPEN"
	CLOSE    = "CLOSE"
	UP       = "UP"
	DOWN     = "DOWN"
	NOTHING  = "NOTHING"
	CALLUP   = 'u'
	CALLDOWN = 'd'
)

func (c *Cabin) NextCommand() string {
	if c.opened {
		c.opened = false
		return CLOSE
	}
	call := c.nextCall()
	if call.atFloor >= 0 {
		if call.atFloor > c.higherFloor {
			return NOTHING
		}
		if call.atFloor == c.currentFloor {
			c.opened = true
			c.callProcessed()
			return OPEN
		}
		if call.atFloor > c.currentFloor {
			c.currentFloor++
			return UP
		}
		if call.atFloor < c.currentFloor {
			c.currentFloor--
			return DOWN
		}
	}
	return NOTHING
}

func (c *Cabin) Reset() {
}

func (c *Cabin) Call(atFloor int, to byte) {
	c.calls = append(c.calls, call{atFloor, to})
}

func (c *Cabin) Go(floorToGo int) {
	fmt.Println("go", floorToGo)
}

func (c *Cabin) UserHasEntered() {
	fmt.Println("UserHasEntered")
}

func (c *Cabin) UserHasExited() {
	fmt.Println("UserHasExited")
}

func NewCabin() *Cabin {
	c := new(Cabin)
	c.lowerFloor = 0
	c.currentFloor = 0
	c.higherFloor = 20
	c.calls = make([]call, 0, c.higherFloor)
	return c
}

func (c *Cabin) nextCall() call {
	if len(c.calls) == 0 {
		return NOCALL
	}
	return c.calls[0]
}

func (c *Cabin) callProcessed() {
	c.calls = c.calls[1:]
}
