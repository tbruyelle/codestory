package main

import (
	"testing"
)

func TestMixedPreventStopIfCallCountExceedCabinSize(t *testing.T) {
	setup()
	e.crew = e.cabinSize - 3
	e.Go(2)
	e.Call(1, UP)
	e.Call(1, UP)
	e.Call(1, UP)
	e.Call(1, UP)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+DOWN+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 1)
	assertDoorClosed(t, e)
}

func TestMixedPreventStopCallIfFull(t *testing.T) {
	setup()
	e.crew = e.cabinSize
	e.Go(2)
	e.Call(1, UP)

	c := nextCommands(e)

	assert(t, c, UP+UP+OPEN+CLOSE+DOWN+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 1)
	assertDoorClosed(t, e)
}

func TestMixedGoDownCallUpCurrentFloor(t *testing.T) {
	setup()
	e.currentFloor = 3
	e.Go(2)
	e.Call(3, UP)

	c := nextCommands(e)

	assert(t, c, DOWN+OPEN+CLOSE+UP+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 3)
	assertDoorClosed(t, e)
}

func TestMixedGoUpCallDownCurrentFloor(t *testing.T) {
	setup()
	e.currentFloor = 3
	e.Go(4)
	e.Call(3, DOWN)

	c := nextCommands(e)

	assert(t, c, UP+OPEN+CLOSE+DOWN+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 3)
	assertDoorClosed(t, e)
}

func TestMixedGoDownCallDownCurrentFloor(t *testing.T) {
	setup()
	e.currentFloor = 5
	e.Go(4)
	e.Call(5, DOWN)

	c := nextCommands(e)

	assert(t, c, OPEN+CLOSE+DOWN+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 4)
	assertDoorClosed(t, e)
}

func TestMixedGoUpCallUpCurrentFloor(t *testing.T) {
	setup()
	e.Go(1)
	e.Call(0, UP)

	c := nextCommands(e)

	assert(t, c, OPEN+CLOSE+UP+OPEN+CLOSE+NOTHING)
	assertNoMoreCall(t, e)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 1)
	assertDoorClosed(t, e)
}

func TestMixedGoUpSkipCallDown(t *testing.T) {
	setup()
	e.Go(4)
	e.Call(3, DOWN)

	c := nextCommands(e)

	assert(t, c, UP+UP+UP+UP+OPEN+CLOSE+DOWN+OPEN+CLOSE+NOTHING)
	assertDoorClosed(t, e)
	assertNoMoreCall(t, e)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 3)
}

func TestMixedGoDownSkipCallUp(t *testing.T) {
	setup()
	e.currentFloor = 5
	e.Go(1)
	e.Call(3, UP)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+DOWN+DOWN+OPEN+CLOSE+UP+UP+OPEN+CLOSE+NOTHING)
	assertDoorClosed(t, e)
	assertNoMoreCall(t, e)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 3)
}

func TestMixedGoUpStopAtCallUp(t *testing.T) {
	setup()
	e.Go(4)
	e.Call(3, UP)

	c := nextCommands(e)

	assert(t, c, UP+UP+UP+OPEN+CLOSE+UP+OPEN+CLOSE+NOTHING)
	assertDoorClosed(t, e)
	assertNoMoreCall(t, e)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 4)
}

func TestMixedGoDownStopAtCallDown(t *testing.T) {
	setup()
	e.currentFloor = 5
	e.Go(1)
	e.Call(3, DOWN)

	c := nextCommands(e)

	assert(t, c, DOWN+DOWN+OPEN+CLOSE+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assertDoorClosed(t, e)
	assertNoMoreCall(t, e)
	assertNoMoreGo(t, e)
	assertFloor(t, e, 1)
}

func TestMixedGoCalledFloor(t *testing.T) {
	setup()
	e.Go(4)
	e.Call(4, UP)

	c := nextCommands(e)

	assert(t, c, UP+UP+UP+UP+OPEN+CLOSE+NOTHING)
	assertDoorClosed(t, e)
	assertNoMoreGo(t, e)
	assertNoMoreCall(t, e)
}
