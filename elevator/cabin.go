package elevator

import (
	"fmt"
	"log"
	"time"
)

type command struct {
	name     byte
	Floor    int
	Up, Down bool
	Count    int
}

func (c command) String() string {
	var s string
	if c.name == CMD_CALL {
		s = "CALL"
	} else {
		s = "GO"
	}
	return fmt.Sprintf("(%s floor=%d up=%t down=%t count=%d)", s, c.Floor, c.Up, c.Down, c.Count)
}

type Cabin struct {
	lowerFloor, higherFloor, CurrentFloor int
	Opened                                bool
	Calls                                 []command
	Gos                                   []command
	Direction                             string
	CabinSize, Crew                       int
	debug, ditdlamerde                    bool
}

func (c *Cabin) String() string {
	return fmt.Sprintf("open=%t direction=%s crew=%d/%d (%d/%d) -- floor=%d\ngos=%s\ncalls=%s",
		c.Opened, c.Direction, c.Crew, c.CabinSize,
		c.lowerFloor, c.higherFloor, c.CurrentFloor, c.Gos, c.Calls)
}

func (c *Cabin) IsIdle() bool {
	return len(c.Gos) == 0 && len(c.Calls) == 0 && c.Crew == 0
}

func (c *Cabin) MatchDirection(floor int) bool {
	if c.CurrentFloor == floor {
		return true
	}
	if c.CurrentFloor < floor {
		return c.Direction == UP
	}
	return c.Direction == DOWN
}

const (
	OPEN      = "OPEN"
	OPEN_UP   = "OPEN_UP"
	OPEN_DOWN = "OPEN_DOWN"
	CLOSE     = "CLOSE"
	UP        = "UP"
	DOWN      = "DOWN"
	NOTHING   = "NOTHING"
	POOP      = "POOP"
	CMD_CALL  = 'c'
	CMD_GO    = 'g'
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
			c.Direction = ret
		} else {
			c.Direction = NOTHING
		}
	}()

	if c.ditdlamerde {
		c.ditdlamerde = false
		return POOP
	}

	if c.Opened {
		// before close check is theres a command for currentFloor
		if c.shouldStopAtCurrentFloor(nil) {
			// command for current floor, keep the door opened
			c.floorProcessed(c.CurrentFloor)
			return NOTHING
		}
		c.Opened = false
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
	log.Printf("%s ---> Reset requested %d/%d/%d msg=%s\n", time.Now(), lowerFloor, higherFloor, cabinSize, cause)
	initCabin(c, lowerFloor, higherFloor, cabinSize)
}

func (c *Cabin) Call(floor int, dir string) {
	i := findFloor(c.Calls, floor)
	if i < len(c.Calls) {
		c.Calls[i].Up = c.Calls[i].Up || dir == UP
		c.Calls[i].Down = c.Calls[i].Down || dir == DOWN
		c.Calls[i].Count++
	} else {
		c.Calls = append(c.Calls,
			command{
				name:  CMD_CALL,
				Floor: floor,
				Up:    dir == UP,
				Down:  dir == DOWN,
				Count: 1,
			})
	}
}

func (c *Cabin) Go(floor int) {
	i := findFloor(c.Gos, floor)
	if i < len(c.Gos) {
		c.Gos[i].Count++
	} else {
		c.Gos = append(c.Gos,
			command{
				name:  CMD_GO,
				Floor: floor,
				Up:    floor > c.CurrentFloor,
				Down:  floor < c.CurrentFloor,
				Count: 1,
			})
	}
}

func (c *Cabin) Debug(enabled bool) {
	c.debug = enabled
}

func (c *Cabin) nextCall() *command {
	if len(c.Calls) == 0 {
		return nil
	}
	return &c.Calls[0]
}

func (c *Cabin) nextGo() *command {
	if len(c.Gos) == 0 {
		return nil
	}
	// take first to have a direction
	cmd := c.Gos[0]
	maxdiff, ind := 0, 0
	for i := 0; i < len(c.Gos); i++ {
		if !sameDir(cmd, c.Gos[i]) {
			continue
		}
		diff := floorDiff(c.Gos[i].Floor, c.CurrentFloor)
		if diff > maxdiff {
			maxdiff = diff
			ind = i
		}
	}
	return &c.Gos[ind]
}

func floorDiff(f1, f2 int) int {
	diff := f1 - f2
	if diff < 0 {
		diff = -diff
	}
	return diff
}

func sameDir(cmd1, cmd2 command) bool {
	return cmd1.Up && cmd2.Up || cmd1.Down && cmd2.Down
}

func (c *Cabin) UserHasEntered() {
	if c.Crew >= c.CabinSize {
		log.Println("OUps cabin size exceeded !")
		return
	}
	c.Crew++
}

func (c *Cabin) UserHasExited() {
	if c.Crew <= 0 {
		log.Println("OUps cabin is empty")
		return
	}
	c.Crew--
}

func NewCabin(lowerFloor, higherFloor, cabinSize int, d bool) *Cabin {
	c := new(Cabin)
	initCabin(c, lowerFloor, higherFloor, cabinSize)
	c.debug = d
	return c
}

func initCabin(c *Cabin, lowerFloor, higherFloor, cabinSize int) {
	c.lowerFloor = lowerFloor
	c.CurrentFloor = 0
	c.higherFloor = higherFloor
	c.Calls = make([]command, 0)
	c.Gos = make([]command, 0)
	c.Opened = false
	c.CabinSize = cabinSize
	c.Crew = 0
	c.trace("INIT")
}

func (c *Cabin) isFull() bool {
	return c.Crew >= c.CabinSize
}

func (c *Cabin) processCommand(cmd *command) string {
	if cmd.Floor < c.lowerFloor || cmd.Floor > c.higherFloor {
		c.floorProcessed(cmd.Floor)
		return NOTHING
	}
	if cmd.Floor > c.CurrentFloor {
		if c.shouldStopAtCurrentFloor(cmd) {
			return c.processCmdCurrentFloor()
		}
		c.CurrentFloor++
		return UP
	}
	if cmd.Floor < c.CurrentFloor {
		if c.shouldStopAtCurrentFloor(cmd) {
			return c.processCmdCurrentFloor()
		}
		c.CurrentFloor--
		return DOWN
	}
	// floor == c.CurrentFloor
	return c.processCmdCurrentFloor()
}

func (c *Cabin) floorProcessed(floor int) {
	c.deleteGo(floor)
	c.deleteCall(floor)
}

func (c *Cabin) processCmdCurrentFloor() string {
	c.Opened = true
	c.floorProcessed(c.CurrentFloor)
	switch c.Direction {
		case UP:
		return OPEN_UP
		case DOWN:
		return OPEN_DOWN
	}
	return OPEN
}

func (c *Cabin) shouldStopAtCurrentFloor(currentCmd *command) bool {
	if hasFloor(c.Gos, c.CurrentFloor) {
		return true
	}
	if c.isFull() {
		// never stop if cabin is full
		return false
	}
	i := findFloor(c.Calls, c.CurrentFloor)
	if i < len(c.Calls) {
		// found a call for current floor
		if c.Calls[i].Count+c.Crew > c.CabinSize {
			// cabin will be full dont stop
			return false
		}
		if currentCmd == nil {
			return true
		}
		// check if current direction matches with call direction
		switch c.Direction {
		case UP:
			if currentCmd.Floor == c.higherFloor {
				// the destination is the higher floor,
				// so stop if CALL UP
				return c.Calls[i].Up
			}
			return c.Calls[i].Up && currentCmd.Up
		case DOWN:
			if currentCmd.Floor == c.lowerFloor {
				// the destination is the lower floor,
				// so stop if CALL down
				return c.Calls[i].Down
			}
			return c.Calls[i].Down && currentCmd.Down
		default:
			// the cabin is idle here, check match direction with current command
			return sameDir(c.Calls[i], *currentCmd)
		}
	}
	return false
}

func findFloor(cmds []command, floor int) int {
	for i := 0; i < len(cmds); i++ {
		if cmds[i].Floor == floor {
			return i
		}
	}
	return len(cmds)
}

func hasFloor(cmds []command, floor int) bool {
	return findFloor(cmds, floor) < len(cmds)
}

func (c *Cabin) deleteGo(floor int) {
	i := findFloor(c.Gos, floor)
	//fmt.Printf("delete floor %d GOS %+v\nfound %d\n", floor, c.Gos, i)
	if i < len(c.Gos) {
		c.Gos = c.Gos[:i+copy(c.Gos[i:], c.Gos[i+1:])]
		//cmds[i], cmds = cmds[len(cmds)-1], cmds[:len(cmds)-1]
	}
}

func (c *Cabin) deleteCall(floor int) {
	i := findFloor(c.Calls, floor)
	if i < len(c.Calls) {
		c.Calls = c.Calls[:i+copy(c.Calls[i:], c.Calls[i+1:])]
		//cmds[i], cmds = cmds[len(cmds)-1], cmds[:len(cmds)-1]
	}
}

func (c *Cabin) trace(msg string) {
	if c.debug {
		log.Printf("%s:\n%s\n=================\n", msg, c)
	}
}
