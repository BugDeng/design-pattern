package chain

import (
	"fmt"
	"time"
)

// Handler ...
type Handler interface {
	DoHandle() bool
}

// HandlerChain implements the Handler interface and defines the successor
type HandlerChain struct {
	Handler
	successor *HandlerChain
}

// SetSuccessor set the successor handler
func (h *HandlerChain) SetSuccessor(successor *HandlerChain) {
	h.successor = successor
}

// Handle if do not handle the event then call the children's handler
func (h *HandlerChain) Handle() {
	handled := h.DoHandle()
	if h.successor != nil && !handled {
		h.successor.Handle()
	}
}

// DoHandle return true if handled the event
func (h *HandlerChain) DoHandle() bool {
	if h.Handler == nil {
		fmt.Println("default chain")
		return true
	}
	return h.Handler.DoHandle()
}

// HandlerA ...
type HandlerA struct{}

// NewHandlerA ...
func NewHandlerA() *HandlerChain {
	return &HandlerChain{Handler: &HandlerA{}}
}

// DoHandle ...
func (h *HandlerA) DoHandle() bool {
	handled := time.Now().Unix()%2 == 0
	fmt.Println("handler A", handled)
	if handled {
		return true
	}
	return false
}

// HandlerB ...
type HandlerB struct{}

// NewHandlerB ...
func NewHandlerB() *HandlerChain {
	return &HandlerChain{Handler: &HandlerB{}}
}

// DoHandle ...
func (h *HandlerB) DoHandle() bool {
	fmt.Println("handler B")
	return true
}
