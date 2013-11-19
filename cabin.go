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
	gos                                   []int
}

const (
	OPEN     = "OPEN"
	CLOSE    = "CLOSE"
	UP       = "UP"
	DOWN     = "DOWN"
	NOTHING  = "NOTHING"
	CALLUP   = 'u'
	CALLDOWN = 'd'
)

var debug = false

func processCommand(c *Cabin, floor int, processed func()) string {
	if floor < c.lowerFloor || floor > c.higherFloor {
		processed()
		return NOTHING
	}
	if floor > c.currentFloor {
		c.currentFloor++
		return UP
	}
	if floor < c.currentFloor {
		c.currentFloor--
		return DOWN
	}
	// floor == c.currentFloor
	c.opened = true
	processed()
	return OPEN
}

func (c *Cabin) trace(msg string) {
	if debug {
		fmt.Printf("%s: current=%d\ncalls=%+v\ngos=%+v\n\n", msg, c.currentFloor, c.calls, c.gos)
	}
}

func (c *Cabin) NextCommand() string {
	c.trace("Start NEXT")
	defer c.trace("End NEXT")

	if c.opened {
		c.opened = false
		return CLOSE
	}
	floor := c.nextGo()
	if floor != nil {
		return processCommand(c, *floor, c.goProcessed)
	}
	call := c.nextCall()
	if call != nil {
		return processCommand(c, call.atFloor, c.callProcessed)
	}
	return NOTHING
}

func (c *Cabin) Reset(lowerFloor, higherFloor int) {
	initCabin(c, lowerFloor, higherFloor)
}

func (c *Cabin) Call(atFloor int, to byte) {
	for i := 0; i < len(c.calls); i++ {
		if atFloor == c.calls[i].atFloor {
			// ignore already registerd call
			return
		}
	}
	c.calls = append(c.calls, call{atFloor, to})
}

func (c *Cabin) Go(floorToGo int) {
	for i := 0; i < len(c.gos); i++ {
		if floorToGo == c.gos[i] {
			// ignore alread registered go
			return
		}
	}
	c.gos = append(c.gos, floorToGo)
}

func (c *Cabin) UserHasEntered() {
}

func (c *Cabin) UserHasExited() {
}

func NewCabin(lowerFloor, higherFloor int, d bool) *Cabin {
	c := new(Cabin)
	initCabin(c, lowerFloor, higherFloor)
	debug = d
	return c
}

func initCabin(c *Cabin, lowerFloor, higherFloor int) {
	c.lowerFloor = lowerFloor
	c.currentFloor = 0
	c.higherFloor = higherFloor
	c.calls = make([]call, 0, c.higherFloor-c.lowerFloor)
	c.gos = make([]int, 0, c.higherFloor-c.lowerFloor)
	c.opened = false
	c.trace("init")
}

func (c *Cabin) nextCall() *call {
	if len(c.calls) == 0 {
		return nil
	}
	return &c.calls[0]
}

func (c *Cabin) callProcessed() {
	c.calls = c.calls[1:]
}

func (c *Cabin) nextGo() *int {
	if len(c.gos) == 0 {
		return nil
	}
	return &c.gos[0]
}

func (c *Cabin) goProcessed() {
	c.gos = c.gos[1:]
}
