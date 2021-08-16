package builder

import "testing"

func TestBuilder(t *testing.T) {
	c1 := &Config1{}
	c1.SetMax(2).SetMin(1).SetMust(0)
	c := NewConfig(c1)
	c.builder.Build()
	//c1.SetMax(1).SetMin(2).SetMust(0)
	// c.builder.Build()
}
