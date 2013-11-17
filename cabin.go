package main

import (
	"fmt"
)

type Cabin struct {
	lowerFloor, higherFloor, currentFloor int
	call                                  int
	opened                                bool
}

const (
	OPEN    = "OPEN"
	CLOSE   = "CLOSE"
	UP      = "UP"
	DOWN    = "DOWN"
	NOTHING = "NOTHING"
)

func (c *Cabin) NextCommand() string {
	if c.opened {
		c.opened = false
		c.call = -1
		return CLOSE
	}
	if c.call >= 0 {
		if c.call > c.higherFloor {
			return NOTHING
		}
		if c.call == c.currentFloor {
			c.opened = true
			return OPEN
		}
		if c.call > c.currentFloor {
			c.currentFloor++
			return UP
		}
		if c.call < c.currentFloor {
			c.currentFloor--
			return DOWN
		}
	}
	return NOTHING
}

func (c *Cabin) Reset() {
}

func (c *Cabin) Call(atFloor int) {
	c.call = atFloor
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
	c.call = -1
	return c
}
