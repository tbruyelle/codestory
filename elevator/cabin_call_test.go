package elevator

import (
	"testing"
)

func TestCallDirection(t *testing.T) {
	setup()

	e.Call(1, UP)
	e.Call(1, UP)
	e.Call(2, UP)
	e.Call(2, DOWN)
	e.Call(3, DOWN)
	e.Call(3, DOWN)

	if !e.Calls[0].Up || e.Calls[0].Down {
		t.Errorf("incorrect direction, should up=true, down=false but was up=%t, down=%t", e.Calls[0].Up, e.Calls[0].Down)
	}
	if !e.Calls[1].Up || !e.Calls[1].Down {
		t.Errorf("incorrect direction, should up=true, down=true but was up=%t, down=%t", e.Calls[1].Up, e.Calls[1].Down)
	}
	if e.Calls[2].Up || !e.Calls[2].Down {
		t.Errorf("incorrect direction, should up=false, down=true but was up=%t, down=%t", e.Calls[2].Up, e.Calls[2].Down)
	}
}

func TestCallDownAtMaxFloorStopAtCallUp(t *testing.T) {
	setup()
	e.Call(5, DOWN)
	e.Call(4, UP)
	e.Call(1, UP)

	c := nextCommands(e)

	assert(t, c, UP+OPEN+CLOSE+UP+UP+UP+OPEN+CLOSE+UP+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 5)
	assertDoorClosed(t, e)
	assertNoMoreCall(t, e)
}

func TestCallUpAtMinFloorStopAtCallDown(t *testing.T) {
	setup()
	e.CurrentFloor = 5
	e.Call(0, UP)
	e.Call(4, DOWN)
	e.Call(1, DOWN)

	c := nextCommands(e)

	assert(t, c, DOWN+OPEN+CLOSE+DOWN+DOWN+DOWN+OPEN+CLOSE+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 0)
	assertDoorClosed(t, e)
	assertNoMoreCall(t, e)
}

func TestDownCallDownSkipCallUp(t *testing.T) {
	setup()
	e.CurrentFloor = 5
	e.Call(3, UP)
	e.Call(4, DOWN)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+OPEN+CLOSE+UP+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 4)
	assertDoorClosed(t, e)
	assertNoMoreCall(t, e)
}

func TestUpCallUpSkipCallDown(t *testing.T) {
	setup()
	e.Call(3, DOWN)
	e.Call(2, UP)

	c := nextCommands(e)

	assert(t, c, UP+UP+UP+OPEN+CLOSE+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 2)
	assertDoorClosed(t, e)
	assertNoMoreCall(t, e)
}

func TestUpCallDownSkipCallUp(t *testing.T) {
	setup()
	e.Call(4, DOWN)
	e.Call(3, UP)

	c := nextCommands(e)

	assert(t, c, UP+UP+UP+UP+OPEN+CLOSE+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 3)
	assertDoorClosed(t, e)
	assertNoMoreCall(t, e)
}

func TestDownCallUpSkipCallDown(t *testing.T) {
	setup()
	e.CurrentFloor = 5
	e.Call(2, UP)
	e.Call(3, DOWN)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+DOWN+OPEN+CLOSE+UP+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 3)
	assertDoorClosed(t, e)
	assertNoMoreCall(t, e)
}

func TestDownCallUpSkipCallUp(t *testing.T) {
	setup()
	e.CurrentFloor = 5
	e.Call(3, UP)
	e.Call(4, UP)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+OPEN+CLOSE+UP+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 4)
	assertDoorClosed(t, e)
	assertNoMoreCall(t, e)
}

func TestUpCallDownSkipCallDown(t *testing.T) {
	setup()
	e.Call(3, DOWN)
	e.Call(2, DOWN)

	c := nextCommands(e)

	assert(t, c, UP+UP+UP+OPEN+CLOSE+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 2)
	assertDoorClosed(t, e)
	assertNoMoreCall(t, e)
}

func TestUpCallUpStopAtCallUp(t *testing.T) {
	setup()
	e.Call(4, UP)
	e.Call(3, UP)

	c := nextCommands(e)

	assert(t, c, UP+UP+UP+OPEN+CLOSE+UP+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 4)
	assertDoorClosed(t, e)
	assertNoMoreCall(t, e)
}

func TestDownCallDownStopAtCallDown(t *testing.T) {
	setup()
	e.CurrentFloor = 5
	e.Call(2, DOWN)
	e.Call(3, DOWN)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+OPEN+CLOSE+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 2)
	assertDoorClosed(t, e)
	assertNoMoreCall(t, e)
}

func TestCallCurrentFloorOpenedDoor(t *testing.T) {
	setup()
	e.Opened = true
	e.Call(0, UP)

	c := nextCommands(e)

	assert(t, c, NOTHING+CLOSE+NOTHING)
}

func TestCallSameDirectionUp(t *testing.T) {
	setup()
	e.Call(2, DOWN)
	e.Call(3, DOWN)
	e.Call(5, DOWN)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+UP+OPEN+CLOSE+UP+UP+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 5)
	assertNoMoreCall(t, e)
	assertDoorClosed(t, e)
}
func TestCallSameDirectionDown(t *testing.T) {
	setup()
	e.CurrentFloor = 5
	e.Call(4, DOWN)
	e.Call(2, DOWN)
	e.Call(1, DOWN)

	c := nextCommands(e)

	assert(t, c, DOWN+OPEN+CLOSE+DOWN+DOWN+OPEN+CLOSE+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 1)
	assertNoMoreCall(t, e)
	assertDoorClosed(t, e)
}

func TestCallsSameFloor(t *testing.T) {
	setup()

	e.Call(1, UP)
	e.Call(1, UP)

	assertNbCall(t, e, 1)
}

func TestCallCurrentFloor(t *testing.T) {
	setup()
	e.Call(0, UP)

	c := nextCommands(e)

	assert(t, c, OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 0)
	assertDoorClosed(t, e)
}

func TestCallTooLow(t *testing.T) {
	setup()
	e.Call(-1, UP)

	c := e.NextCommand()

	assert(t, c, NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 0)
	assertDoorClosed(t, e)
}

func TestCallTooHigh(t *testing.T) {
	setup()
	e.Call(21, UP)

	c := e.NextCommand()

	assert(t, c, NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 0)
	assertDoorClosed(t, e)
}

func TestCallUp(t *testing.T) {
	setup()
	e.Call(2, UP)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 2)
	assertDoorClosed(t, e)
}

func TestCallDown(t *testing.T) {
	setup()
	e.CurrentFloor = 3
	e.Call(1, UP)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 1)
	assertDoorClosed(t, e)
}

func TestCalls(t *testing.T) {
	setup()

	e.Call(2, UP)
	c := nextCommands(e)
	e.Call(3, UP)
	c += nextCommands(e)
	e.Call(1, UP)
	c += nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING+UP+OPEN+CLOSE+NOTHING+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 1)
	assertDoorClosed(t, e)
}

func TestCallNegativeFloors(t *testing.T) {
	setup()
	e.lowerFloor = -3

	e.Call(2, DOWN)
	c := nextCommands(e)
	e.Call(-3, UP)
	c += nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING+DOWN+DOWN+DOWN+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, -3)
	assertNoMoreCall(t, e)
	assertDoorClosed(t, e)
}

func TestCallOpenUp(t *testing.T) {
	setup()
	e.Call(2, UP)

	c := nextCommands(e)

	assertReal(t, c, UP+UP+OPEN_UP+CLOSE+NOTHING)
}

func TestCallOpenDown(t *testing.T) {
	setup()
	e.CurrentFloor = 3
	e.Call(1, DOWN)

	c := nextCommands(e)

	assertReal(t, c, DOWN+DOWN+OPEN_DOWN+CLOSE+NOTHING)
}
