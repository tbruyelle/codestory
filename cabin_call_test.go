package main

import (
	"testing"
)

func TestCallCurrentFloor(t *testing.T) {
	setup()
	e.Call(0, CALLUP)

	c := nextCommands(e)

	assert(t, c, OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 0)
	assertDoorClosed(t, e)
}

func TestCallTooLow(t *testing.T) {
	setup()
	e.Call(-1, CALLUP)

	c := e.NextCommand()

	assert(t, c, NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 0)
	assertDoorClosed(t, e)
}

func TestCallTooHigh(t *testing.T) {
	setup()
	e.Call(21, CALLUP)

	c := e.NextCommand()

	assert(t, c, NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 0)
	assertDoorClosed(t, e)
}

func TestCallUp(t *testing.T) {
	setup()
	e.Call(2, CALLUP)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 2)
	assertDoorClosed(t, e)
}

func TestCallDown(t *testing.T) {
	setup()
	e.currentFloor = 3
	e.Call(1, CALLUP)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 1)
	assertDoorClosed(t, e)
}

func TestCalls(t *testing.T) {
	setup()

	e.Call(2, CALLUP)
	c := nextCommands(e)
	e.Call(3, CALLUP)
	c += nextCommands(e)
	e.Call(1, CALLUP)
	c += nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING+UP+OPEN+CLOSE+NOTHING+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 1)
	assertDoorClosed(t, e)
}

func TestCallNegativeFloors(t *testing.T) {
	setup()
	e.lowerFloor = -3

	e.Call(2, CALLDOWN)
	c := nextCommands(e)
	e.Call(-3, CALLUP)
	c += nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING+DOWN+DOWN+DOWN+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, -3)
	assertNoMoreCall(t, e)
	assertDoorClosed(t, e)
}
