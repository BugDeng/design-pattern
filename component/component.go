package component

import "fmt"

// Node is a interface of tree like node
type Node interface {
	Parent() Node
	Name() string
	Print(pre string)
}

// LeafNode is a node without child
type LeafNode struct {
	Node
	name string
}

// NewLeaf returns a LeafNode pointer with given name
func NewLeaf(name string) *LeafNode {
	return &LeafNode{name: name}
}

// Print prints the leaf name
func (leaf *LeafNode) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, leaf.name)
}

// ParentNode is a node with child node
type ParentNode struct {
	Node
	childs []Node
	name   string
}

// NewParent returns a node with given name
func NewParent(name string) *ParentNode {
	return &ParentNode{childs: make([]Node, 0, 5), name: name}
}

// Print prints the parent node and all children node
func (parent *ParentNode) Print(pre string) {
	fmt.Printf("%s+%s\n", pre, parent.name)
	pre += "  "
	for _, node := range parent.childs {
		node.Print(pre)
	}
}

// Add adds child node
func (parent *ParentNode) Add(child ...Node) {
	parent.childs = append(parent.childs, child...)
}
