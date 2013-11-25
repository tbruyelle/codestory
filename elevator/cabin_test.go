package elevator

import (
	"strings"
	"testing"
)

var e *Cabin

func setup() {
	e = NewCabin(0, 5, 10, false)
}

func userEntered(e *Cabin, nb int) {
	for i := 0; i < nb; i++ {
		e.UserHasEntered()
	}
}

func userExited(e *Cabin, nb int) {
	for i := 0; i < nb; i++ {
		e.UserHasExited()
	}
}

func assert(t *testing.T, value string, want string) {
	if value != want {
		t.Errorf("expected %s but was %s", want, value)
	}
}

func assertInt(t *testing.T, value int, want int) {
	if value != want {
		t.Errorf("expected %d but was %d", want, value)
	}
}


func assertCrew(t *testing.T, e *Cabin, crew int) {
	if e.crew != crew {
		t.Errorf("expected crew %d but was %d", crew, e.crew)
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

	e.Reset(-1, 50, 500, "yeah")

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
	if e.cabinSize != 500 {
		t.Errorf("bad cabinsize")
	}
	assertCrew(t, e, 0)
}

func TestUserHasEntered(t *testing.T) {
	setup()

	e.UserHasEntered()

	assertCrew(t, e, 1)
}

func TestUserHasExited(t *testing.T) {
	setup()
	e.crew = 1

	e.UserHasExited()

	assertCrew(t, e, 0)
}

func TestUsersEnterAndExit(t *testing.T) {
	setup()

	userEntered(e, 3)
	userExited(e, 2)

	assertCrew(t, e, 1)
}

func TestUserCannotEnterIfFull(t *testing.T) {
	setup()
	e.crew = e.cabinSize

	e.UserHasEntered()

	assertCrew(t, e, e.cabinSize)
}

func TestUserCannotExitIfEmpty(t *testing.T) {
	setup()

	e.UserHasExited()

	assertCrew(t, e, 0)
}

func TestDitdlamerde(t *testing.T) {
	setup()
	e.Ditdlamerde()

	c := nextCommands(e)

	assert(t, c, POOP+NOTHING)
}

func TestDebugEnabled(t *testing.T) {
	setup()

	e.Debug(true)

	if !e.debug {
		t.Errorf("debug should be enabled")
	}
}
func TestDebugDisabled(t *testing.T) {
	setup()
	e.debug = true

	e.Debug(false)

	if e.debug {
		t.Errorf("debug should be disabled")
	}
}

func TestCommandCount(t *testing.T) {
	setup()

	e.Call(1, UP)
	e.Call(1, DOWN)
	e.Call(1, DOWN)
	e.Call(2, UP)
	e.Go(1)
	e.Go(3)
	e.Go(3)
	e.Go(3)
	e.Go(3)

	assertNbCall(t, e, 2)
	if e.calls[0].count != 3 {
		t.Errorf("Inccorrect call count, expected 3 but was %d", e.calls[0].count)
	}
	if e.calls[1].count != 1 {
		t.Errorf("Inccorrect call count, expected 1 but was %d", e.calls[1].count)
	}
	assertNbGo(t, e, 2)
	if e.gos[0].count != 1 {
		t.Errorf("Inccorrect go count, expected 1 but was %d", e.gos[0].count)
	}
	if e.gos[1].count != 4 {
		t.Errorf("Inccorrect go count, expected 4 but was %d", e.gos[1].count)
	}
}
