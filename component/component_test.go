package component

import "testing"

func TestComponent(t *testing.T) {
	life := NewParent("life")
	animal := NewParent("animal")
	animal.Add(NewLeaf("cat"), NewLeaf("dog"))
	plant := NewParent("plant")
	plant.Add(NewLeaf("fruit"), NewLeaf("unknown"))
	life.Add(animal, plant)
	life.Print("")
	animal.Print("")
}
