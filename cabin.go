package main

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
	CALLNO   = 'n'
)

func (c *Cabin) NextCommand() string {
	if c.opened {
		c.opened = false
		return CLOSE
	}
	if len(c.gos) > 0 {
		floor := c.nextGo()
		if floor < c.lowerFloor || floor > c.higherFloor {
			c.goProcessed()
			return NOTHING
		}
		if floor == c.currentFloor {
			c.opened = true
			c.goProcessed()
			return OPEN
		}
		if floor > c.currentFloor {
			c.currentFloor++
			return UP
		}
		if floor < c.currentFloor {
			c.currentFloor--
			return DOWN
		}
	}
	call := c.nextCall()
	if call.to != CALLNO {
		if call.atFloor < c.lowerFloor ||call.atFloor > c.higherFloor {
			c.callProcessed()
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

func (c *Cabin) Reset(lowerFloor, higherFloor int) {
	initCabin(c, lowerFloor, higherFloor)
}

func (c *Cabin) Call(atFloor int, to byte) {
	c.calls = append(c.calls, call{atFloor, to})
}

func (c *Cabin) Go(floorToGo int) {
	c.gos = append(c.gos, floorToGo)
}

func (c *Cabin) UserHasEntered() {
}

func (c *Cabin) UserHasExited() {
}

func NewCabin(lowerFloor, higherFloor int) *Cabin {
	c := new(Cabin)
	initCabin(c, lowerFloor, higherFloor)
	return c
}

func initCabin(c *Cabin, lowerFloor, higherFloor int) {
	c.lowerFloor = lowerFloor
	c.currentFloor = 0
	c.higherFloor = higherFloor
	c.calls = make([]call, 0, c.higherFloor-c.lowerFloor)
	c.gos = make([]int, 0, c.higherFloor-c.lowerFloor)
}

func (c *Cabin) nextCall() call {
	if len(c.calls) == 0 {
		return call{to: CALLNO}
	}
	return c.calls[0]
}

func (c *Cabin) callProcessed() {
	c.calls = c.calls[1:]
}

func (c *Cabin) nextGo() int {
	return c.gos[0]
}

func (c *Cabin) goProcessed() {
	c.gos = c.gos[1:]
}
