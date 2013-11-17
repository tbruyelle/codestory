package main

import (
	"strings"
	"testing"
)

func assert(t *testing.T, value string, want string) {
	if value != want {
		t.Errorf("expected %s, returned %s", want, value)
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
}

func TestBasicCallCurrentFloor(t *testing.T) {
	e := NewCabin()
	e.Call(0, CALLUP)

	c := nextCommands(e)

	assert(t, c, OPEN+CLOSE+NOTHING)
}

func TestBasicCallTooLow(t *testing.T) {
	e := NewCabin()
	e.Call(-1, CALLUP)

	c := e.NextCommand()

	assert(t, c, NOTHING)
}

func TestBasicCallTooHigh(t *testing.T) {
	e := NewCabin()
	e.Call(21, CALLUP)

	c := e.NextCommand()

	assert(t, c, NOTHING)
}

func TestBasicCallUp(t *testing.T) {
	e := NewCabin()
	e.Call(2, CALLUP)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING)
}

func TestBasicCallDown(t *testing.T) {
	e := NewCabin()
	e.currentFloor = 2
	e.Call(0, CALLUP)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+OPEN+CLOSE+NOTHING)
}

func TestBasicCalls(t *testing.T) {
	e := NewCabin()
	e.Call(2, CALLUP)
	e.Call(3, CALLUP)
	e.Call(1, CALLUP)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+UP+OPEN+CLOSE+DOWN+DOWN+OPEN+CLOSE+NOTHING)
}
