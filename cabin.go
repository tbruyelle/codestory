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
	calls                                 []command
	gos                                   []command
	direction                             string
}

const (
	OPEN     = "OPEN"
	CLOSE    = "CLOSE"
	UP       = "UP"
	DOWN     = "DOWN"
	NOTHING  = "NOTHING"
	CMD_CALL = 'c'
	CMD_GO   = 'g'
)

var debug = false

func (c *Cabin) processCommand(cmd *command) string {
	if cmd.floor < c.lowerFloor || cmd.floor > c.higherFloor {
		c.floorProcessed(cmd.floor)
		return NOTHING
	}
	if cmd.floor > c.currentFloor {
		if c.hasCmdCurrentFloor(cmd) {
			return c.processCmdCurrentFloor()
		}
		c.currentFloor++
		return UP
	}
	if cmd.floor < c.currentFloor {
		if c.hasCmdCurrentFloor(cmd) {
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

func (c *Cabin) hasCmdCurrentFloor(currentCmd *command) bool {
	if hasFloor(c.gos, c.currentFloor) {
		return true
	}
	i := findFloor(c.calls, c.currentFloor)
	if i < len(c.calls) {
		// found a call for current floor
		if currentCmd == nil {
			return true
		}
		// check if current command direction matches with call direction
		switch currentCmd.name {
		case CMD_GO:
			// GO command: directions match if they are identicals
			return c.calls[i].up && currentCmd.up || c.calls[i].down && currentCmd.down
		case CMD_CALL:
			// CALL command: directions match if they are identicals and
			// match the elevator current direction
			switch c.direction {
			case UP:
				if currentCmd.floor==c.higherFloor{
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
				fmt.Println("What to do here ?", c.calls[i], currentCmd, c.direction)
			}
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
	if debug {
		fmt.Printf("%s:(%d/%d) opened=%t current=%d\nCALLS===%+v\nGOS===%+v\n\n", msg, c.lowerFloor, c.higherFloor, c.opened, c.currentFloor, c.calls, c.gos)
	}
}

func (c *Cabin) NextCommand() (ret string) {
	c.trace("Start NEXT")
	defer func() { c.trace("RETURN " + ret) }()
	defer func() {
		// remind the elevator direction
		if ret == UP || ret == DOWN {
			c.direction = ret
		} else {
			c.direction = NOTHING
		}
	}()

	if c.opened {
		// before close check is theres a command for currentFloor
		if c.hasCmdCurrentFloor(nil) {
			// command for current floor, keep the door opened
			c.floorProcessed(c.currentFloor)
			return NOTHING
		}
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

func (c *Cabin) Call(floor int, dir string) {
	ind := findFloor(c.calls, floor)
	if ind < len(c.calls) {
		call := c.calls[ind]
		call.up = true
		call.down = true
	} else {
		c.calls = append(c.calls,
			command{
				name:  CMD_CALL,
				floor: floor,
				up:    dir == UP,
				down:  dir == DOWN,
			})
	}
}

func (c *Cabin) Go(floor int) {
	if !hasFloor(c.gos, floor) {
		c.gos = append(c.gos,
			command{
				name:  CMD_GO,
				floor: floor,
				up:    floor > c.currentFloor,
				down:  floor < c.currentFloor,
			})
	}
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
	c.calls = make([]command, 0)
	c.gos = make([]command, 0)
	c.opened = false
	c.trace("init")
}
