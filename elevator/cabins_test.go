package elevator

import "testing"

var cs *Cabins

func setupCs() {
	cs = NewCabins(0, 5, 10, 2, false)
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

	assertInt(t, cs.cabs[0].gos[0].floor, 1)
	assertInt(t, cs.cabs[1].gos[0].floor, 2)
}

func TestCabinsUserHasEntered(t *testing.T) {
	setupCs()

	cs.UserHasEntered(0)
	cs.UserHasEntered(0)
	cs.UserHasEntered(1)

	assertInt(t, cs.cabs[0].crew, 2)
	assertInt(t, cs.cabs[1].crew, 1)
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

	assertInt(t, cs.cabs[0].crew, 1)
	assertInt(t, cs.cabs[1].crew, 1)
}

func TestCabinsCallNearestCabin(t *testing.T) {
	setupCs()
	cs.cabs[0].currentFloor = 5
	cs.cabs[1].currentFloor = 3

	cs.Call(0, UP)
	c := cs.NextCommands()

	assert(t, c[0], NOTHING)
	assert(t, c[1], DOWN)
}

//func newMockCabin(id, l, h, c, cc int) *Cabin {
//	m := &mockCabin{id: id, lowerFloor: l,
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
