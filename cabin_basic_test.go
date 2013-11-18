package main

import (
	"testing"
)

func TestBasicGoCurrentFloor(t *testing.T) {
	setup()
	e.Go(0)

	c := nextCommands(e)

	assert(t, c, OPEN+CLOSE+NOTHING)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 0)
	assertDoorClosed(t, e)
}

func TestBasicGoUp(t *testing.T) {
	setup()
	e.Go(2)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 2)
	assertDoorClosed(t, e)
}

func TestBasicGoDown(t *testing.T) {
	setup()
	e.currentFloor = 2
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

func TestBasicCallCurrentFloor(t *testing.T) {
	setup()
	e.Call(0, CALLUP)

	c := nextCommands(e)

	assert(t, c, OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 0)
	assertDoorClosed(t, e)
}

func TestBasicCallTooLow(t *testing.T) {
	setup()
	e.Call(-1, CALLUP)

	c := e.NextCommand()

	assert(t, c, NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 0)
	assertDoorClosed(t, e)
}

func TestBasicCallTooHigh(t *testing.T) {
	setup()
	e.Call(21, CALLUP)

	c := e.NextCommand()

	assert(t, c, NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 0)
	assertDoorClosed(t, e)
}

func TestBasicCallUp(t *testing.T) {
	setup()
	e.Call(2, CALLUP)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 2)
	assertDoorClosed(t, e)
}

func TestBasicCallDown(t *testing.T) {
	setup()
	e.currentFloor = 3
	e.Call(1, CALLUP)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 1)
	assertDoorClosed(t, e)
}

func TestBasicCalls(t *testing.T) {
	setup()
	e.Call(2, CALLUP)
	e.Call(3, CALLUP)
	e.Call(1, CALLUP)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+UP+OPEN+CLOSE+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 1)
	assertDoorClosed(t, e)
}

func TestBasicGos(t *testing.T) {
	setup()
	e.Go(2)
	e.Go(3)
	e.Go(1)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+UP+OPEN+CLOSE+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 1)
	assertDoorClosed(t, e)
}

func TestCallNegativeFloors(t *testing.T) {
	setup()
	e.lowerFloor = -3
	e.Call(2, CALLDOWN)
	e.Call(-3, CALLUP)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+DOWN+DOWN+DOWN+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, -3)
	assertNoMoreCall(t, e)
	assertDoorClosed(t, e)
}

func TestGoNegativeFloors(t *testing.T) {
	setup()
	e.lowerFloor = -3
	e.Go(2)
	e.Go(-3)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+DOWN+DOWN+DOWN+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, -3)
	assertNoMoreGo(t, e)
	assertDoorClosed(t, e)
}
