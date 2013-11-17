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

func TestWhenCallFloorTooLowThenReturnNOTHING(t *testing.T) {
	e := NewCabin()
	e.Call(-1)

	c := e.NextCommand()

	assert(t, c, NOTHING)
}

func TestWhenCallFloorTooHighThenReturnNOTHING(t *testing.T) {
	e := NewCabin()
	e.Call(21)

	c := e.NextCommand()

	assert(t, c, NOTHING)
}

func TestWhenCallFloorUpThenReturnUP(t *testing.T) {
	e := NewCabin()
	e.Call(1)

	c := e.NextCommand()

	assert(t, c, UP)
}

func TestWhenCallFloorDownThenReturnDOWN(t *testing.T) {
	e := NewCabin()
	e.currentFloor = 2
	e.Call(1)

	c := e.NextCommand()

	assert(t, c, DOWN)
}

func TestWhenCall1FloorUpThenReturnUPNOTHING(t *testing.T) {
	e := NewCabin()
	e.Call(1)

	c := e.NextCommand() + e.NextCommand()

	assert(t, c, UP+NOTHING)
}

func TestWhenCall1FloorDownThenReturnDOWNNOTHING(t *testing.T) {
	e := NewCabin()
	e.currentFloor = 2
	e.Call(1)

	c := e.NextCommand() + e.NextCommand()

	assert(t, c, DOWN+NOTHING)
}

func TestWhenCall3FloorUpThenReturnUPUPUPNOTHING(t *testing.T) {
	e := NewCabin()
	e.Call(3)

	c := e.NextCommand() + e.NextCommand() + e.NextCommand() + e.NextCommand()

	assert(t, c, UP+UP+UP+NOTHING)
}

func TestWhenCall3FloorDownThenReturnDOWNDOWNDOWNNOTHING(t *testing.T) {
	e := NewCabin()
	e.currentFloor = 3
	e.Call(0)

	c := e.NextCommand() + e.NextCommand() + e.NextCommand() + e.NextCommand()

	assert(t, c, DOWN+DOWN+DOWN+NOTHING)
}
