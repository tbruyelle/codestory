package main

import (
	"strings"
	"testing"
)

func assert(t *testing.T, value string, want string) {
	if value != want {
		t.Errorf("expected %s but was %s", want, value)
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
	e := NewCabin()

	c := e.NextCommand()

	assert(t, c, NOTHING)
	assertFloor(t, e, 0)
}

func TestBasicCallCurrentFloor(t *testing.T) {
	e := NewCabin()
	e.Call(0, CALLUP)

	c := nextCommands(e)

	assert(t, c, OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 0)
}

func TestBasicCallTooLow(t *testing.T) {
	e := NewCabin()
	e.Call(-1, CALLUP)

	c := e.NextCommand()

	assert(t, c, NOTHING)
	assertFloor(t, e, 0)
}

func TestBasicCallTooHigh(t *testing.T) {
	e := NewCabin()
	e.Call(21, CALLUP)

	c := e.NextCommand()

	assert(t, c, NOTHING)
	assertFloor(t, e, 0)
}

func TestBasicCallUp(t *testing.T) {
	e := NewCabin()
	e.Call(2, CALLUP)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 2)
}

func TestBasicCallDown(t *testing.T) {
	e := NewCabin()
	e.currentFloor = 2
	e.Call(0, CALLUP)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 0)
}

func TestBasicCalls(t *testing.T) {
	e := NewCabin()
	e.Call(2, CALLUP)
	e.Call(3, CALLUP)
	e.Call(1, CALLUP)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+UP+OPEN+CLOSE+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 1)
}

func TestReset(t *testing.T) {
	e := NewCabin()
	e.Call(2, CALLUP)
	e.Call(3, CALLDOWN)
	e.Go(5)
	e.UserHasEntered()

	e.Reset()

	assertFloor(t, e, 0)
}
