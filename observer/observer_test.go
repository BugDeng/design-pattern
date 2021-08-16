package observer

import "testing"

func TestObsever(t *testing.T) {
	h := NewHero("hero")
	monster1 := &Monster{"monster1"}
	monster2 := &Monster{"monster2"}
	h.Register(monster1.name, monster1)
	h.Register(monster2.name, monster2)
	h.Notify()
}
