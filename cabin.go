package main

import (
	"fmt"
)

type command struct {
	name     byte
	floor    int
	up, down bool
}

type Cabin struct {
	lowerFloor, higherFloor, currentFloor int
	opened                                bool
	calls                                 map[int]command
	gos                                   map[int]command
}

const (
	OPEN     = "OPEN"
	CLOSE    = "CLOSE"
	UP       = "UP"
	DOWN     = "DOWN"
	NOTHING  = "NOTHING"
	CMD_CALL = 'c'
	CMD_GO   = 'g'
	CALLUP   = 'u'
	CALLDOWN = 'd'
)

var debug = false

func (c *Cabin) processCommand(cmd *command) string {
	if cmd.floor < c.lowerFloor || cmd.floor > c.higherFloor {
		c.floorProcessed(cmd.floor)
		return NOTHING
	}
	if cmd.floor > c.currentFloor {
		c.currentFloor++
		return UP
	}
	if cmd.floor < c.currentFloor {
		c.currentFloor--
		return DOWN
	}
	// floor == c.currentFloor
	c.opened = true
	c.floorProcessed(cmd.floor)
	return OPEN
}

func (c *Cabin) floorProcessed(floor int) {
	delete(c.gos, floor)
	delete(c.calls, floor)
}

func (c *Cabin) trace(msg string) {
	if debug {
		fmt.Printf("%s:(%d/%d) current=%d\ncalls=%+v\ngos=%+v\n\n", msg, c.lowerFloor, c.higherFloor, c.currentFloor, c.calls, c.gos)
	}
}

func (c *Cabin) NextCommand() (ret string) {
	c.trace("Start NEXT")
	defer func() { c.trace("End NEXT " + ret) }()

	if c.opened {
		c.opened = false
		return CLOSE
	}
	cmd := c.nextGo()
	if cmd == nil {
		cmd = c.nextCall()
	}
	if cmd != nil {
		return c.processCommand(cmd)
	}
	return NOTHING
}

func (c *Cabin) Reset(lowerFloor, higherFloor int) {
	initCabin(c, lowerFloor, higherFloor)
}

func (c *Cabin) Call(floor int, dir byte) {
	if call, ok := c.calls[floor]; ok {
		call.up = true
		call.down = true
	} else {
		c.calls[floor] = command{
			name:  CMD_CALL,
			floor: floor,
			up:    dir == CALLUP,
			down:  dir == CALLDOWN,
		}
	}
}

func (c *Cabin) Go(floor int) {
	if _, ok := c.gos[floor]; !ok {
		c.gos[floor] = command{
			name:  CMD_GO,
			floor: floor,
			up:    floor > c.currentFloor,
			down:  floor < c.currentFloor,
		}
	}
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
	c.calls = make(map[int]command)
	c.gos = make(map[int]command)
	c.opened = false
	c.trace("init")
}

func (c *Cabin) nextCall() *command {
	if len(c.calls) == 0 {
		return nil
	}
	for _, cmd := range c.calls {
		return &cmd
	}
	return nil
}

func (c *Cabin) nextGo() *command {
	if len(c.gos) == 0 {
		return nil
	}
	for _, cmd := range c.gos {
		return &cmd
	}
	return nil
}

func (c *Cabin) goProcessed(floor int) {
}
