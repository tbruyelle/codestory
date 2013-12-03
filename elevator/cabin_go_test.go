package elevator

import (
	"testing"
)

func TestNextGoUpChooseFarestFloor(t *testing.T) {
	setup()
	e.Go(1)
	e.Go(3)
	e.Go(5)

	c := e.nextGo()

	if c.Floor != 5 {
		t.Errorf("incorrect next GO, expected floor 5 but was floor %d", c.Floor)
	}
}

func TestNextGoDownChooseFarestFloor(t *testing.T) {
	setup()
	e.CurrentFloor = 5
	e.Go(4)
	e.Go(3)
	e.Go(1)

	c := e.nextGo()

	if c.Floor != 1 {
		t.Errorf("incorrect next GO, expected floor 1 but was floor %d", c.Floor)
	}
}

func TestNextGoUpDownChooseFarestFloor(t *testing.T) {
	setup()
	e.higherFloor = 20
	e.CurrentFloor = 10
	e.Go(11)
	e.Go(3)
	e.Go(1)
	e.Go(19)

	c := e.nextGo()

	if c.Floor != 19 {
		t.Errorf("incorrect next GO, expected floor 19 but was floor %d", c.Floor)
	}
}

func TestNextGoDownUpChooseFarestFloor(t *testing.T) {
	setup()
	e.higherFloor = 20
	e.CurrentFloor = 10
	e.Go(9)
	e.Go(3)
	e.Go(1)
	e.Go(19)

	c := e.nextGo()

	if c.Floor != 1 {
		t.Errorf("incorrect next GO, expected floor 1 but was floor %d", c.Floor)
	}
}

func TestGosSameFloor(t *testing.T) {
	setup()

	e.Go(1)
	e.Go(1)

	assertNbGo(t, e, 1)
}

func TestGoCurrentFloor(t *testing.T) {
	setup()
	e.Go(0)

	c := nextCommands(e)

	assert(t, c, OPEN+CLOSE+NOTHING)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 0)
	assertDoorClosed(t, e)
}

func TestGoUp(t *testing.T) {
	setup()
	e.Go(2)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 2)
	assertDoorClosed(t, e)
}

func TestGoDown(t *testing.T) {
	setup()
	e.CurrentFloor = 2
	e.Go(0)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 0)
	assertNoMoreGo(t, e)
	assertDoorClosed(t, e)
}

func TestGoTooLow(t *testing.T) {
	setup()
	e.Go(-1)

	c := nextCommands(e)

	assert(t, c, NOTHING)
	assertFloor(t, e, 0)
	assertNoMoreGo(t, e)
	assertDoorClosed(t, e)
}
func TestGoTooHigh(t *testing.T) {
	setup()
	e.Go(100)

	c := nextCommands(e)

	assert(t, c, NOTHING)
	assertFloor(t, e, 0)
	assertNoMoreGo(t, e)
	assertDoorClosed(t, e)
}

func TestGos(t *testing.T) {
	setup()

	e.Go(2)
	c := nextCommands(e)
	e.Go(3)
	c += nextCommands(e)
	e.Go(1)
	c += nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING+UP+OPEN+CLOSE+NOTHING+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 1)
	assertDoorClosed(t, e)
}

func TestGoNegativeFloors(t *testing.T) {
	setup()
	e.lowerFloor = -3

	e.Go(2)
	c := nextCommands(e)
	e.Go(-3)
	c += nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING+DOWN+DOWN+DOWN+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, -3)
	assertNoMoreGo(t, e)
	assertDoorClosed(t, e)
}

func TestGoSameDirectionUp(t *testing.T) {
	setup()
	e.Go(2)
	e.Go(3)
	e.Go(5)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+UP+OPEN+CLOSE+UP+UP+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 5)
	assertNoMoreGo(t, e)
	assertDoorClosed(t, e)
}

func TestGoSameDirectionDown(t *testing.T) {
	setup()
	e.CurrentFloor = 5
	e.Go(4)
	e.Go(2)
	e.Go(1)

	c := nextCommands(e)

	assert(t, c, DOWN+OPEN+CLOSE+DOWN+DOWN+OPEN+CLOSE+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 1)
	assertNoMoreGo(t, e)
	assertDoorClosed(t, e)
}

func TestGosUpChooseNearest(t *testing.T) {
	setup()
	e.Go(4)
	e.Go(2)
	e.Go(3)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+UP+OPEN+CLOSE+UP+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 4)
	assertDoorClosed(t, e)
	assertNoMoreGo(t, e)
}

func TestGosDownChooseNearest(t *testing.T) {
	setup()
	e.CurrentFloor = 5
	e.Go(1)
	e.Go(3)
	e.Go(2)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+OPEN+CLOSE+DOWN+OPEN+CLOSE+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 1)
	assertNoMoreGo(t, e)
	assertDoorClosed(t, e)
}

func TestGoUpOpen(t *testing.T) {
	setup()
	e.Go(2)

	c := nextCommands(e)

	assertReal(t, c, UP+UP+OPEN+CLOSE+NOTHING)
}

func TestGoDownOpen(t *testing.T) {
	setup()
	e.CurrentFloor=3
	e.Go(1)

	c := nextCommands(e)

	assertReal(t, c, DOWN+DOWN+OPEN+CLOSE+NOTHING)
}

func TestGoUpDownOpenDown(t *testing.T) {
setup()
e.CurrentFloor=1
e.Go(2)
e.Go(0)

c:=nextCommands(e)

assertReal(t, c, UP+OPEN_DOWN+DOWN+DOWN+OPEN+CLOSE+NOTHING)
}
