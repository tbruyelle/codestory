package elevator

import "strings"
import "testing"

var cs *Cabins

func setupCs() {
	cs = NewCabins(0, 20, 10, 2, false)
}

func TestCabinsNextCommands(t *testing.T) {
	setupCs()

	c := cs.NextCommands()

	assert(t, c[0], NOTHING)
	assert(t, c[1], NOTHING)
}

func TestCabinsGo(t *testing.T) {
	setupCs()

	cs.Go(1, 0)
	cs.Go(2, 1)

	assertInt(t, cs.Cabs[0].Gos[0].Floor, 1)
	assertInt(t, cs.Cabs[1].Gos[0].Floor, 2)
}

func TestCabinsUserHasEntered(t *testing.T) {
	setupCs()

	cs.UserHasEntered(0)
	cs.UserHasEntered(0)
	cs.UserHasEntered(1)

	assertInt(t, cs.Cabs[0].Crew, 2)
	assertInt(t, cs.Cabs[1].Crew, 1)
}

func TestCabinsUserHasExited(t *testing.T) {
	setupCs()
	cs.UserHasEntered(0)
	cs.UserHasEntered(0)
	cs.UserHasEntered(0)
	cs.UserHasEntered(1)
	cs.UserHasEntered(1)

	cs.UserHasExited(0)
	cs.UserHasExited(0)
	cs.UserHasExited(1)

	assertInt(t, cs.Cabs[0].Crew, 1)
	assertInt(t, cs.Cabs[1].Crew, 1)
}

func TestCabinsCallNearestIdleCabin(t *testing.T) {
	setupCs()
	cs.Cabs[0].CurrentFloor = 5
	cs.Cabs[1].CurrentFloor = 3

	cs.Call(0, UP)
	c := nextCommandss(cs)

	assert(t, c[0], NOTHING)
	assert(t, c[1], DOWN+DOWN+DOWN+OPEN+CLOSE+NOTHING)
}

func TestCabinsCallSameDistanceChooseFirstIdleCabin(t *testing.T) {
	setupCs()
	cs.Cabs[0].CurrentFloor = 3
	cs.Cabs[1].CurrentFloor = 1

	cs.Call(2, UP)
	c := nextCommandss(cs)

	assert(t, c[0], DOWN+OPEN+CLOSE+NOTHING)
	assert(t, c[1], NOTHING)
}

func TestCabinsCallChooseFirstCabinNoSameDirection(t *testing.T) {
	setupCs()
	cs.Cabs[0].CurrentFloor = 3
	cs.Go(4, 0) // first cab goes up
	cs.Cabs[1].CurrentFloor = 1
	cs.Go(0, 1)              // second cab goes down
	tmp := cs.NextCommands() // start moving

	cs.Call(2, UP)
	c := nextCommandss(cs)
	c[0] = tmp[0] + c[0]
	c[1] = tmp[1] + c[1]

	assert(t, c[0], UP+OPEN+CLOSE+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assert(t, c[1], DOWN+OPEN+CLOSE+NOTHING)
}

func TestCabinsCallChooseCabinSameDirectionUp(t *testing.T) {
	setupCs()
	cs.Cabs[0].CurrentFloor = 12
	cs.Go(10, 0) // give a down direction to cabin 0
	cs.Cabs[1].CurrentFloor = 9
	cs.Go(14, 1)             // give a up direction to cabin 1
	tmp := cs.NextCommands() // start moving

	cs.Call(15, UP)
	c := nextCommandss(cs)
	c[0] = tmp[0] + c[0]
	c[1] = tmp[1] + c[1]

	assert(t, c[0], DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assert(t, c[1], UP+UP+UP+UP+UP+OPEN+CLOSE+UP+OPEN+CLOSE+NOTHING)
}

func TestCabinsCallChooseNearestCabinInSameDirection(t *testing.T) {
	setupCs()
	cs.Cabs[0].CurrentFloor = 2
	cs.Go(6, 0)
	cs.Cabs[1].CurrentFloor = 4
	cs.Go(6, 1)
	tmp := cs.NextCommands() // start moving

	cs.Call(5, UP)
	c := nextCommandss(cs)
	c[0] = tmp[0] + c[0]
	c[1] = tmp[1] + c[1]

	assert(t, c[0], UP+UP+UP+UP+OPEN+CLOSE+NOTHING)
	assert(t, c[1], UP+OPEN+CLOSE+UP+OPEN+CLOSE+NOTHING)
}

func TestCabinsCallOtherDirectionChooseNearestCabinInSameDirection(t *testing.T) {
	setupCs()
	cs.Cabs[0].CurrentFloor = 2
	cs.Go(6, 0)
	cs.Cabs[1].CurrentFloor = 4
	cs.Go(6, 1)
	tmp := cs.NextCommands() // start moving

	cs.Call(5, DOWN)
	c := nextCommandss(cs)
	c[0] = tmp[0] + c[0]
	c[1] = tmp[1] + c[1]

	assert(t, c[0], UP+UP+UP+UP+OPEN+CLOSE+NOTHING)
	assert(t, c[1], UP+UP+OPEN+CLOSE+DOWN+OPEN+CLOSE+NOTHING)
}

func TestCabinsCallSkipFullCabin(t *testing.T) {
	setupCs()
	cs.Cabs[0].CurrentFloor = 2
	cs.Cabs[0].Crew = cs.CabinSize
	cs.Go(5, 0)
	cs.Cabs[1].CurrentFloor = 1
	cs.Go(5, 1)
	tmp := cs.NextCommands() // start moving

	cs.Call(3, UP)
	c := nextCommandss(cs)
	c[0] = tmp[0] + c[0]
	c[1] = tmp[1] + c[1]

	assert(t, c[0], UP+UP+UP+OPEN+CLOSE+NOTHING)
	assert(t, c[1], UP+UP+OPEN+CLOSE+UP+UP+OPEN+CLOSE+NOTHING)
}

func TestCabinsCallBothFullChooseNearestCabin(t *testing.T) {
	setupCs()
	cs.Cabs[0].CurrentFloor = 2
	cs.Cabs[0].Crew = cs.CabinSize
	cs.Go(5, 0)
	cs.Cabs[1].CurrentFloor = 1
	cs.Cabs[1].Crew = cs.CabinSize
	cs.Go(5, 1)
	tmp := cs.NextCommands() // start moving

	cs.Call(3, UP)
	c := nextCommandss(cs)
	c[0] = tmp[0] + c[0]
	c[1] = tmp[1] + c[1]

	assert(t, c[0], UP+UP+UP+OPEN+CLOSE+DOWN+DOWN+OPEN+CLOSE+NOTHING)
	assert(t, c[1], UP+UP+UP+UP+OPEN+CLOSE+NOTHING)
}

func TestCabinsCallChooseLessLoadedCabin(t *testing.T) {
	setupCs()
	cs.Cabs[0].CurrentFloor = 2
	cs.Cabs[0].Crew = 4

	cs.Call(3, UP)
	c := nextCommandss(cs)

	assert(t, c[0], NOTHING)
	assert(t, c[1], UP+UP+UP+OPEN+CLOSE+NOTHING)
}

func nextCommandss(cs *Cabins) []string {
	cmds := make([]string, 2)
	for i := 0; i < 100; i++ {
		cmd := cs.NextCommands()
		doBreak := true
		if cmd[0] != NOTHING || !strings.HasSuffix(cmds[0], NOTHING) {
			doBreak = false
			cmds[0] += cmd[0]
		}
		if cmd[1] != NOTHING || !strings.HasSuffix(cmds[1], NOTHING) {
			doBreak = false
			cmds[1] += cmd[1]
		}
		if doBreak {
			break
		}
	}
	return cmds
}

//func newMockCabin(id, l, h, c, cc int) *Cabin {
//	m := &mockCabin{id: id, LowerFloor: l,
//		higherFloor: h, cabinSize: c, cabinCount: cc}
//	return m
//}
//
//type mockCabin struct {
//	Cabin
//	id int
//}
//
//func (m *mockCabin) NextCommand() string {
//	return "MOCK" + m.id
//}
