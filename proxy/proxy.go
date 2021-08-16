package proxy

import "fmt"

type Controller interface {
	Method()
}

type ControllerImpl struct {
	name string
}

func (c *ControllerImpl) Method() {
	fmt.Println(c.name)
}

type ControllerImplProxy struct {
	ci *ControllerImpl
}

// Method proxy extends original Method
func (cp *ControllerImplProxy) Method() {
	fmt.Println("pre")
	cp.ci.Method()
	fmt.Println("post")
}
