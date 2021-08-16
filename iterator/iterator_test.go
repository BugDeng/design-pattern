package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	it := NewArrayIterator(arr)
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
