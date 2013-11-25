package main

import (
	"fmt"
)

type command struct {
	name     byte
	floor    int
	up, down bool
	count    int
}

func (c command) String() string {
	var s string
	if c.name == CMD_CALL {
		s = "CALL"
	} else {
		s = "GO"
	}
	return fmt.Sprintf("(%s floor=%d up=%t down=%t count=%d)", s, c.floor, c.up, c.down, c.count)
}

type Cabin struct {
	lowerFloor, higherFloor, currentFloor int
	opened                                bool
	calls                                 []command
	gos                                   []command
	direction                             string
	cabinSize, crew                       int
	debug, ditdlamerde                    bool
}

func (c *Cabin) String() string {
	return fmt.Sprintf("open=%t direction=%s crew=%d/%d (%d/%d) -- floor=%d\ngos=%s\ncalls=%s",
		c.opened, c.direction, c.crew, c.cabinSize,
		c.lowerFloor, c.higherFloor, c.currentFloor, c.gos, c.calls)
}

const (
	OPEN     = "OPEN"
	CLOSE    = "CLOSE"
	UP       = "UP"
	DOWN     = "DOWN"
	NOTHING  = "NOTHING"
	POOP     = "POOP"
	CMD_CALL = 'c'
	CMD_GO   = 'g'
)

func (c *Cabin) Ditdlamerde() {
	c.ditdlamerde = true
}

func (c *Cabin) NextCommand() (ret string) {
	var cmd *command
	c.trace("\nNEXT")
	defer func() {
		if c.debug {
			s := "RETURN " + ret
			if cmd != nil {
				s += fmt.Sprintf(" cmd=%s", cmd)
			}
			c.trace(s)
		}
	}()

	defer func() {
		// remind the elevator direction
		if ret == UP || ret == DOWN {
			c.direction = ret
		} else {
			c.direction = NOTHING
		}
	}()

	if c.ditdlamerde {
		c.ditdlamerde = false
		return POOP
	}

	if c.opened {
		// before close check is theres a command for currentFloor
		if c.shouldStopAtCurrentFloor(nil) {
			// command for current floor, keep the door opened
			c.floorProcessed(c.currentFloor)
			return NOTHING
		}
		c.opened = false
		return CLOSE
	}
	cmd = c.nextGo()
	if cmd == nil {
		cmd = c.nextCall()
	}
	if cmd != nil {
		return c.processCommand(cmd)
	}
	return NOTHING
}

func (c *Cabin) Reset(lowerFloor, higherFloor, cabinSize int, cause string) {
	fmt.Printf("---> Reset requested %d/%d/%d msg=%s\n", lowerFloor, higherFloor, cabinSize, cause)
	initCabin(c, lowerFloor, higherFloor, cabinSize)
}

func (c *Cabin) Call(floor int, dir string) {
	i := findFloor(c.calls, floor)
	if i < len(c.calls) {
		c.calls[i].up = c.calls[i].up||dir==UP
		c.calls[i].down = c.calls[i].down||dir==DOWN
		c.calls[i].count++
	} else {
		c.calls = append(c.calls,
			command{
				name:  CMD_CALL,
				floor: floor,
				up:    dir == UP,
				down:  dir == DOWN,
				count: 1,
			})
	}
}

func (c *Cabin) Go(floor int) {
	i := findFloor(c.gos, floor)
	if i < len(c.gos) {
		c.gos[i].count++
	} else {
		c.gos = append(c.gos,
			command{
				name:  CMD_GO,
				floor: floor,
				up:    floor > c.currentFloor,
				down:  floor < c.currentFloor,
				count: 1,
			})
	}
}

func (c *Cabin) Debug(enabled bool) {
	c.debug = enabled
}

func (c *Cabin) nextCall() *command {
	if len(c.calls) == 0 {
		return nil
	}
	return &c.calls[0]
}

func (c *Cabin) nextGo() *command {
	if len(c.gos) == 0 {
		return nil
	}
	return &c.gos[0]
}

func (c *Cabin) UserHasEntered() {
	if c.crew >= c.cabinSize {
		fmt.Println("OUps cabin size exceeded !")
		return
	}
	c.crew++
}

func (c *Cabin) UserHasExited() {
	if c.crew <= 0 {
		fmt.Println("OUps cabin is empty")
		return
	}
	c.crew--
}

func NewCabin(lowerFloor, higherFloor, cabinSize int, d bool) *Cabin {
	c := new(Cabin)
	initCabin(c, lowerFloor, higherFloor, cabinSize)
	c.debug = d
	return c
}

func initCabin(c *Cabin, lowerFloor, higherFloor, cabinSize int) {
	c.lowerFloor = lowerFloor
	c.currentFloor = 0
	c.higherFloor = higherFloor
	c.calls = make([]command, 0)
	c.gos = make([]command, 0)
	c.opened = false
	c.cabinSize = cabinSize
	c.crew = 0
	c.trace("INIT")
}

func (c *Cabin) isFull() bool {
	return c.crew >= c.cabinSize
}

func (c *Cabin) processCommand(cmd *command) string {
	if cmd.floor < c.lowerFloor || cmd.floor > c.higherFloor {
		c.floorProcessed(cmd.floor)
		return NOTHING
	}
	if cmd.floor > c.currentFloor {
		if c.shouldStopAtCurrentFloor(cmd) {
			return c.processCmdCurrentFloor()
		}
		c.currentFloor++
		return UP
	}
	if cmd.floor < c.currentFloor {
		if c.shouldStopAtCurrentFloor(cmd) {
			return c.processCmdCurrentFloor()
		}
		c.currentFloor--
		return DOWN
	}
	// floor == c.currentFloor
	return c.processCmdCurrentFloor()
}

func (c *Cabin) floorProcessed(floor int) {
	c.deleteGo(floor)
	c.deleteCall(floor)
}

func (c *Cabin) processCmdCurrentFloor() string {
	c.opened = true
	c.floorProcessed(c.currentFloor)
	return OPEN
}

func (c *Cabin) shouldStopAtCurrentFloor(currentCmd *command) bool {
	if hasFloor(c.gos, c.currentFloor) {
		return true
	}
	if c.isFull() {
		// never stop if cabin is full
		return false
	}
	i := findFloor(c.calls, c.currentFloor)
	if i < len(c.calls) {
		// found a call for current floor
		if currentCmd == nil {
			return true
		}
		// check if current direction matches with call direction
		switch c.direction {
		case UP:
			if currentCmd.floor == c.higherFloor {
				// the destination is the higher floor,
				// so stop if CALL UP
				return c.calls[i].up
			}
			return c.calls[i].up && currentCmd.up
		case DOWN:
			if currentCmd.floor == c.lowerFloor {
				// the destination is the lower floor,
				// so stop if CALL down
				return c.calls[i].down
			}
			return c.calls[i].down && currentCmd.down
		default:
			// the cabin is idle here, check match direction with current command
			return c.calls[i].up && currentCmd.up || c.calls[i].down && currentCmd.down
		}
	}
	return false
}

func findFloor(cmds []command, floor int) int {
	for i := 0; i < len(cmds); i++ {
		if cmds[i].floor == floor {
			return i
		}
	}
	return len(cmds)
}

func hasFloor(cmds []command, floor int) bool {
	return findFloor(cmds, floor) < len(cmds)
}

func (c *Cabin) deleteGo(floor int) {
	i := findFloor(c.gos, floor)
	//fmt.Printf("delete floor %d GOS %+v\nfound %d\n", floor, c.gos, i)
	if i < len(c.gos) {
		c.gos = c.gos[:i+copy(c.gos[i:], c.gos[i+1:])]
		//cmds[i], cmds = cmds[len(cmds)-1], cmds[:len(cmds)-1]
	}
}

func (c *Cabin) deleteCall(floor int) {
	i := findFloor(c.calls, floor)
	if i < len(c.calls) {
		c.calls = c.calls[:i+copy(c.calls[i:], c.calls[i+1:])]
		//cmds[i], cmds = cmds[len(cmds)-1], cmds[:len(cmds)-1]
	}
}

func (c *Cabin) trace(msg string) {
	if c.debug {
		fmt.Printf("%s:\n%s\n=================\n", msg, c)
	}
}
