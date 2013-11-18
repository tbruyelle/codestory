package main

import (
	"strings"
	"testing"
)

func TestBasicGoCurrentFloor(t *testing.T) {
	e := newElevator()
	e.Go(0)

	c := nextCommands(e)

	assert(t, c, OPEN+CLOSE+NOTHING)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 0)
}

func TestBasicGoUp(t *testing.T) {
	e := newElevator()
	e.Go(2)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 2)
}

func TestBasicGoDown(t *testing.T) {
	e := newElevator()
	e.currentFloor = 2
	e.Go(0)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 0)
	assertNoMoreGo(t, e)
}

func TestGoTooLow(t *testing.T) {
	e := newElevator()
	e.Go(-1)

	c := nextCommands(e)

	assert(t, c, NOTHING)
	assertFloor(t, e, 0)
	assertNoMoreGo(t, e)
}
func TestGoTooHigh(t *testing.T) {
	e := newElevator()
	e.Go(100)

	c := nextCommands(e)

	assert(t, c, NOTHING)
	assertFloor(t, e, 0)
	assertNoMoreGo(t, e)
}

func newElevator() *Cabin {
	return NewCabin(0, 5)
}

func assert(t *testing.T, value string, want string) {
	if value != want {
		t.Errorf("expected %s but was %s", want, value)
	}
}

func assertNoMoreGo(t *testing.T, e *Cabin) {
	if len(e.gos) > 0 {
		t.Errorf("expected no more GO but still %d", len(e.gos))
	}
}
func assertNoMoreCall(t *testing.T, e *Cabin) {
	if len(e.calls) > 0 {
		t.Errorf("expected no more CALL but still %d", len(e.calls))
	}
}
func assertFloor(t *testing.T, c *Cabin, floor int) {
	if c.currentFloor != floor {
		t.Errorf("expected floor %d but was %d", floor, c.currentFloor)
	}
}

func nextCommands(e Elevator) string {
	var s string
	for !strings.Contains(s, NOTHING) {
		s += e.NextCommand()
	}
	return s
}

func TestWhenIdleReturnNOTHING(t *testing.T) {
	e := newElevator()

	c := e.NextCommand()

	assert(t, c, NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 0)
}

func TestBasicCallCurrentFloor(t *testing.T) {
	e := newElevator()
	e.Call(0, CALLUP)

	c := nextCommands(e)

	assert(t, c, OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 0)
}

func TestBasicCallTooLow(t *testing.T) {
	e := newElevator()
	e.Call(-1, CALLUP)

	c := e.NextCommand()

	assert(t, c, NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 0)
}

func TestBasicCallTooHigh(t *testing.T) {
	e := newElevator()
	e.Call(21, CALLUP)

	c := e.NextCommand()

	assert(t, c, NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 0)
}

func TestBasicCallUp(t *testing.T) {
	e := newElevator()
	e.Call(2, CALLUP)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 2)
}

func TestBasicCallDown(t *testing.T) {
	e := newElevator()
	e.currentFloor = 3
	e.Call(1, CALLUP)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 1)
}

func TestBasicCalls(t *testing.T) {
	e := newElevator()
	e.Call(2, CALLUP)
	e.Call(3, CALLUP)
	e.Call(1, CALLUP)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+UP+OPEN+CLOSE+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 1)
}

func TestReset(t *testing.T) {
	e := newElevator()
	e.Call(2, CALLUP)
	e.Call(3, CALLDOWN)
	e.Go(5)
	e.UserHasEntered()
	nextCommands(e)

	e.Reset(-1, 50)

	assertFloor(t, e, 0)
	if e.lowerFloor != -1 {
		t.Errorf("bad lower floor")
	}
	if e.higherFloor != 50 {
		t.Errorf("bad higher floor")
	}
	assertNoMoreCall(t, e)
	assertNoMoreGo(t, e)
}

func TestCallNegativeFloors(t *testing.T) {
	e := newElevator()
	e.lowerFloor = -3
	e.Call(2, CALLDOWN)
	e.Call(-3, CALLUP)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+DOWN+DOWN+DOWN+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, -3)
	assertNoMoreCall(t, e)
}
