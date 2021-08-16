package observer

import "fmt"

// Observer is implemented by the observer who should listen the subject's action
type Observer interface {
	Update(name string)
}

// Monster implemented as an Observer
type Monster struct {
	name string
}

// Update the action
func (m *Monster) Update(name string) {
	fmt.Println(m.name, "attack", name)
}

// Subject is implemeneted by the one who should notify the observers
type Subject interface {
	Register(name string, ob Observer)
	Remove(name string)
	Notify()
}

// Hero implemented the Subject interface
type Hero struct {
	obs  map[string]Observer
	name string
}

// NewHero return a Subject with given name
func NewHero(name string) Subject {
	return &Hero{obs: make(map[string]Observer), name: name}
}

// Register an Observer with given name
func (h *Hero) Register(name string, ob Observer) {
	h.obs[name] = ob
}

// Remove an Observer of given name
func (h *Hero) Remove(name string) {
	delete(h.obs, name)
}

// Notify all the observers
func (h *Hero) Notify() {
	for _, ob := range h.obs {
		ob.Update(h.name)
	}
}
