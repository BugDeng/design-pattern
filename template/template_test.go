package template

import "testing"

func TestTemplate(t *testing.T) {
	p := NewProject("www.project.com", "/root")
	p.Deploy()
}
