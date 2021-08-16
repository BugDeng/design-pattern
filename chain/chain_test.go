package chain

import "testing"

func TestHandlerChain_Handle(t *testing.T) {
	type fields struct {
		Handler   Handler
		successor *HandlerChain
	}
	chain := NewHandlerA()
	handlerA := NewHandlerA()
	handlerB := NewHandlerB()
	chain.SetSuccessor(handlerA)
	handlerA.SetSuccessor(handlerB)
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{"case1", fields{chain, handlerA}},
		{"case2", fields{handlerA, handlerB}},
		{"case3", fields{handlerB, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerChain{
				Handler:   tt.fields.Handler,
				successor: tt.fields.successor,
			}
			h.Handle()
		})
	}
}
