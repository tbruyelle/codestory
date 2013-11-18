package main

import (
	"strings"
	"testing"
)

var e *Cabin

func setup() {
	e = NewCabin(0, 5)
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
	if len(e.gos) > 0 {
		t.Errorf("expected no more GO but still %d", len(e.gos))
	}
}

func assertNoMoreCall(t *testing.T, e *Cabin) {
	if len(e.calls) > 0 {
		t.Errorf("expected no more CALL but still %d", len(e.calls))
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

func TestIdle(t *testing.T) {
	setup()

	c := e.NextCommand()

	assert(t, c, NOTHING)
	assertNoMoreCall(t, e)
	assertFloor(t, e, 0)
	assertDoorClosed(t, e)
}

func TestReset(t *testing.T) {
	setup()
	e.Call(2, CALLUP)
	e.Call(3, CALLDOWN)
	e.Go(5)
	e.UserHasEntered()
	nextCommands(e)
	e.opened = true

	e.Reset(-1, 50)

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
}
