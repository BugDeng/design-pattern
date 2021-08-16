package proxy

import "testing"

func TestProxy(t *testing.T) {
	ci := &ControllerImpl{name: "rabbit"}
	var ctrl Controller = &ControllerImplProxy{ci: ci}
	ctrl.Method()
}
