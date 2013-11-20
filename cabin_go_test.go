package main

import (
	"testing"
)

func TestGoCalledFloor(t *testing.T) {
	setup()
	e.Go(4)
	e.Call(4, UP)

	c := nextCommands(e)

	assert(t, c, UP+UP+UP+UP+OPEN+CLOSE+NOTHING)
	assertDoorClosed(t, e)
	assertNoMoreGo(t, e)
	assertNoMoreCall(t, e)
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
	e.currentFloor = 5
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
	e.currentFloor = 5
	e.Go(1)
	e.Go(3)
	e.Go(2)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+OPEN+CLOSE+DOWN+OPEN+CLOSE+DOWN+OPEN+CLOSE+NOTHING)
	assertFloor(t, e, 1)
	assertNoMoreGo(t, e)
	assertDoorClosed(t, e)
}
