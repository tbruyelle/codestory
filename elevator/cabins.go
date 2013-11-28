package elevator

import (
	"log"
	"time"
)

type Cabins struct {
	lowerFloor, higherFloor, currentFloor, cabCount int
	cabinSize                                       int
	calls                                           []command
	cabs                                            []*Cabin
}

func (c *Cabins) String() string {
	var s string
	for i := 0; i < c.cabCount; i++ {
		s += c.cabs[i].String()
	}
	return s
}

func NewCabins(lowerFloor, higherFloor, cabinSize, cabinCount int, debug bool) *Cabins {
	c := new(Cabins)
	initCabins(c, lowerFloor, higherFloor, cabinSize, cabinCount, debug)
	return c
}

func initCabins(c *Cabins, lowerFloor, higherFloor, cabinSize, cabinCount int, debug bool) {
	c.lowerFloor = lowerFloor
	c.currentFloor = 0
	c.higherFloor = higherFloor
	c.calls = make([]command, 0)
	c.cabinSize = cabinSize
	c.cabCount = cabinCount
	c.cabs = make([]*Cabin, cabinCount)
	for i := 0; i < cabinCount; i++ {
		c.cabs[i] = NewCabin(lowerFloor, higherFloor, cabinSize, debug)
	}
}

func (c *Cabins) NextCommands() []string {
	r := make([]string, c.cabCount)
	for i, c := range c.cabs {
		r[i] = c.NextCommand()
	}
	return r
}

func (c *Cabins) Call(atFloor int, to string) {
	// Determine the nearest idle elevator
	cabin := -1
	maxFloor := c.higherFloor - c.lowerFloor
	for i := 0; i < len(c.cabs); i++ {
		diff := floorDiff(c.cabs[i].currentFloor, atFloor)
		if diff < maxFloor && c.cabs[i].IsIdle() {
			maxFloor = diff
			cabin = i
		}
	}
	if cabin == -1 {
		// if no idle cabin, found the one in the same direction
		for i := 0; i < len(c.cabs); i++ {
			if c.cabs[i].MatchDirection(atFloor) {
				cabin = i
			}
		}
	}
	if cabin == -1 {
		//if not match choose the first cabin
		cabin = 0
	}
	// call the nearest
	c.cabs[cabin].Call(atFloor, to)
}

func (c *Cabins) Go(floorToGo, cabin int) {
	c.cabs[cabin].Go(floorToGo)
}

func (c *Cabins) Reset(lowerFloor, higherFloor, cabinSize, cabinCount int, cause string) {
	log.Printf("%s ---> Reset requested %d/%d/%d/%d msg=%s\n", time.Now(), lowerFloor, higherFloor, cabinSize, cabinCount, cause)
	initCabins(c, lowerFloor, higherFloor, cabinSize, cabinCount, false)
}

func (c *Cabins) Debug(enabled bool) {
	for _, c := range c.cabs {
		c.Debug(enabled)
	}
}
func (c *Cabins) Ditdlamerde() {
	for _, c := range c.cabs {
		c.Ditdlamerde()
	}
}

func (c *Cabins) UserHasEntered(cabin int) {
	c.cabs[cabin].UserHasEntered()
}

func (c *Cabins) UserHasExited(cabin int) {
	c.cabs[cabin].UserHasExited()
}
