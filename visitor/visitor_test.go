package visitor

import "testing"

func TestVisitor(t *testing.T) {
	pdf := NewPDF()
	zip := NewZIP()
	e := NewExactor()
	c := NewCompressor()
	pdf.Accept(e)
	pdf.Accept(c)
	zip.Accept(e)
	zip.Accept(c)
}
