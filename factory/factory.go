package main

import "fmt"

type Factory interface {
	Print()
}

func NewFactory(t int) Factory {
	switch t {
	case 0:
		return &AFactory{}
	case 1:
		return &BFactory{}
	}

	return nil
}

type AFactory struct{}

func (A *AFactory) Print() {
	fmt.Println("I'm A")
}

type BFactory struct{}

func (B *BFactory) Print() {
	fmt.Println("I'm B")
}
