package builder

// Builder ...
type Builder interface {
	Build()
}

// Config ...
type Config struct {
	builder Builder
}

// NewConfig return a *Config with given Builder
func NewConfig(b Builder) *Config {
	return &Config{builder: b}
}

// Config1 ...
type Config1 struct {
	max  int
	min  int
	must interface{}
}

// Build check the parameters when construct object
func (c *Config1) Build() {
	if c.must == nil {
		panic("parameter must should be initial")
	}
	if c.max < c.min {
		panic("max cant smaller than min")
	}
}

func (c *Config1) SetMax(max int) *Config1 {
	c.max = max
	return c
}

func (c *Config1) SetMin(min int) *Config1 {
	c.min = min
	return c
}

func (c *Config1) SetMust(must interface{}) *Config1 {
	c.must = must
	return c
}
