package main

import (
	"testing"
)

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

func TestGos(t *testing.T) {
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
