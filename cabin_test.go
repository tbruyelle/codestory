package main

import (
	"testing"
)

func assert(t *testing.T, value string, want string) {
	if value != want {
		t.Errorf("expected %s, returned %s", want, value)
	}
}

func nextCommands(e Elevator, time int) string {
	var s string
	for i := 0; i < time; i++ {
		s += e.NextCommand()
	}
	return s
}

func TestWhenIdleReturnNOTHING(t *testing.T) {
	e := NewCabin()

	c := e.NextCommand()

	assert(t, c, NOTHING)
}

func TestCallCurrentFloor(t *testing.T) {
	e := NewCabin()
	e.Call(0)

	c := nextCommands(e, 3)

	assert(t, c, OPEN+CLOSE+NOTHING)
}

func TestCallFloorTooLow(t *testing.T) {
	e := NewCabin()
	e.Call(-1)

	c := e.NextCommand()

	assert(t, c, NOTHING)
}

func TestCallFloorTooHigh(t *testing.T) {
	e := NewCabin()
	e.Call(21)

	c := e.NextCommand()

	assert(t, c, NOTHING)
}

func TestCallFloorUp(t *testing.T) {
	e := NewCabin()
	e.Call(2)

	c := nextCommands(e, 5)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING)
}

func TestCallFloorDown(t *testing.T) {
	e := NewCabin()
	e.currentFloor = 2
	e.Call(0)

	c := nextCommands(e, 5)

	assert(t, c, DOWN+DOWN+OPEN+CLOSE+NOTHING)
}
