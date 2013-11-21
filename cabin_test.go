package main

import (
	"strings"
	"testing"
)

var e *Cabin

func setup() {
	e = NewCabin(0, 5, 10, false)
}

func assert(t *testing.T, value string, want string) {
	if value != want {
		t.Errorf("expected %s but was %s", want, value)
	}
}

func assertDoorClosed(t *testing.T, e *Cabin) {
	if e.opened {
		t.Error("exected door closed but was opened")
	}
}

func assertNoMoreGo(t *testing.T, e *Cabin) {
	assertNbGo(t, e, 0)
}

func assertNbGo(t *testing.T, e *Cabin, nbGos int) {
	if len(e.gos) != nbGos {
		t.Errorf("expected %d GO but was %d", nbGos, len(e.gos))
	}
}

func assertNoMoreCall(t *testing.T, e *Cabin) {
	assertNbCall(t, e, 0)
}

func assertNbCall(t *testing.T, e *Cabin, nbCalls int) {
	if len(e.calls) != nbCalls {
		t.Errorf("expected %d CALL but was %d", nbCalls, len(e.calls))
	}
}

func assertFloor(t *testing.T, c *Cabin, floor int) {
	if c.currentFloor != floor {
		t.Errorf("expected floor %d but was %d", floor, c.currentFloor)
	}
}

func nextCommands(e Elevator) string {
	var s string
	for i := 0; i < 100; i++ {
		c := e.NextCommand()
		if c == NOTHING && strings.HasSuffix(s, NOTHING) {
			// ends where there is 2 following NOTHINGs
			break
		}
		s += c
	}
	return s
}

func TestIdle(t *testing.T) {
	setup()

	c := e.NextCommand()

	assert(t, c, NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 0)
	assertDoorClosed(t, e)
}

func TestOpenedDoor(t *testing.T) {
	setup()
	e.opened = true

	c := nextCommands(e)

	assert(t, c, CLOSE+NOTHING)
	assertDoorClosed(t, e)
}

func TestReset(t *testing.T) {
	setup()
	e.Call(2, UP)
	e.Call(3, DOWN)
	e.Go(5)
	e.UserHasEntered()
	nextCommands(e)
	e.opened = true

	e.Reset(-1, 50, 500)

	assertFloor(t, e, 0)
	if e.lowerFloor != -1 {
		t.Errorf("bad lower floor")
	}
	if e.higherFloor != 50 {
		t.Errorf("bad higher floor")
	}
	assertNoMoreCall(t, e)
	assertNoMoreGo(t, e)
	assertDoorClosed(t, e)
	if e.cabinSize!=500{
	t.Errorf("bad cabinsize")}
}
