package visitor

import "fmt"

// Visitor decoupling the resource and operation
type Visitor interface {
	Visit(Resource)
}

// Resource accept an operation
type Resource interface {
	Accept(Visitor)
}

// PDF ...
type PDF struct {
	name string
}

// NewPDF ...
func NewPDF() Resource {
	return &PDF{name: "pdf"}
}

// Accept call the given visitor's operation
func (pdf *PDF) Accept(v Visitor) {
	v.Visit(pdf)
}

// ZIP ...
type ZIP struct {
	name string
}

// NewZIP ...
func NewZIP() Resource {
	return &ZIP{name: "zip"}
}

// Accept call the given visitor's operation
func (zip *ZIP) Accept(v Visitor) {
	v.Visit(zip)
}

// Exactor exact the resource
type Exactor struct{}

// NewExactor ...
func NewExactor() Visitor {
	return &Exactor{}
}

// Visit ...
func (e *Exactor) Visit(r Resource) {
	switch rt := r.(type) {
	case *PDF:
		fmt.Println("exact ", rt.name)
	case *ZIP:
		fmt.Println("exact ", rt.name)
	}
}

// Compressor compress the resource
type Compressor struct{}

// NewCompressor ...
func NewCompressor() Visitor {
	return &Compressor{}
}

// Visit ...
func (c *Compressor) Visit(r Resource) {
	switch rt := r.(type) {
	case *PDF:
		fmt.Println("compress ", rt.name)
	case *ZIP:
		fmt.Println("compress ", rt.name)
	}
}
