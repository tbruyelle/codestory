package elevator

import (
	"log"
	"sort"
	"time"
)

type Cabins struct {
	LowerFloor, HigherFloor, CabCount int
	CabinSize                         int
	Cabs                              []*Cabin
}

func (c *Cabins) String() string {
	var s string
	for i := 0; i < c.CabCount; i++ {
		s += c.Cabs[i].String()
	}
	return s
}

func NewCabins(LowerFloor, higherFloor, cabinSize, cabinCount int, debug bool) *Cabins {
	c := new(Cabins)
	initCabins(c, LowerFloor, higherFloor, cabinSize, cabinCount, debug)
	return c
}

func initCabins(c *Cabins, lowerFloor, higherFloor, cabinSize, cabinCount int, debug bool) {
	c.LowerFloor = lowerFloor
	c.HigherFloor = higherFloor
	c.CabinSize = cabinSize
	c.CabCount = cabinCount
	c.Cabs = make([]*Cabin, cabinCount)
	for i := 0; i < cabinCount; i++ {
		c.Cabs[i] = NewCabin(lowerFloor, higherFloor, cabinSize, debug)
	}
}

func (c *Cabins) NextCommands() []string {
	r := make([]string, c.CabCount)
	for i, c := range c.Cabs {
		r[i] = c.NextCommand()
	}
	return r
}

func (c *Cabins) Call(atFloor int, to string) {
	// order the cabs by their close to the call
	var cabClose []int
	cabMap := make(map[int][]int)
	for i := 0; i < c.CabCount; i++ {
		diff := floorDiff(c.Cabs[i].CurrentFloor, atFloor)
		if c.Cabs[i].isFull() {
		// if cabin full increase the diff
			diff += 1000
		}
		inds := cabMap[diff]
		inds = append(inds, i)
		sort.Ints(inds)
		cabMap[diff] = inds
		cabClose = append(cabClose, diff)
	}
	sort.Ints(cabClose)
	//fmt.Println(cabClose, cabMap)

	// Determine the nearest idle elevator
	cabin := chooseCab(c, cabClose, cabMap, func(i int) bool {
		return c.Cabs[i].IsIdle()
	})
	if cabin == -1 {
		// if no idle cabin, found the one in the same direction
		cabin = chooseCab(c, cabClose, cabMap, func(i int) bool {
			return c.Cabs[i].MatchDirection(atFloor)
		})
	}
	if cabin == -1 {
		//if not match choose the first cabin
		cabin = 0
	}
	// call the nearest
	c.Cabs[cabin].Call(atFloor, to)
}

func chooseCab(c *Cabins, cabClose []int, cabMap map[int][]int, condition func(i int) bool) int {
	for _, k := range cabClose {
		inds := cabMap[k]
		for _, i := range inds {
			if condition(i) {
				//fmt.Println("found", i)
				return i
			}
		}
	}
	//fmt.Println("not found")
	return -1
}

func (c *Cabins) Go(floorToGo, cabin int) {
	c.Cabs[cabin].Go(floorToGo)
}

func (c *Cabins) Reset(lowerFloor, higherFloor, cabinSize, cabinCount int, cause string) {
	log.Printf("%s ---> Reset requested %d/%d/%d/%d msg=%s\n", time.Now(), lowerFloor, higherFloor, cabinSize, cabinCount, cause)
	initCabins(c, lowerFloor, higherFloor, cabinSize, cabinCount, false)
}

func (c *Cabins) Debug(enabled bool) {
	for _, c := range c.Cabs {
		c.Debug(enabled)
	}
}
func (c *Cabins) Ditdlamerde() {
	for _, c := range c.Cabs {
		c.Ditdlamerde()
	}
}

func (c *Cabins) UserHasEntered(cabin int) {
	c.Cabs[cabin].UserHasEntered()
}

func (c *Cabins) UserHasExited(cabin int) {
	c.Cabs[cabin].UserHasExited()
}
