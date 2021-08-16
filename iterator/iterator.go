package iterator

// Iterator is a iterator interface which can range the container
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// ArrayIterator ...
type ArrayIterator struct {
	cursor int
	size   int
	array  []int
}

// NewArrayIterator return a Iterator with given parameters
func NewArrayIterator(array []int) Iterator {
	return &ArrayIterator{size: len(array), array: array}
}

// HasNext return true if next item exists
func (arr *ArrayIterator) HasNext() bool {
	if arr.cursor < arr.size {
		return true
	}
	return false
}

// Next return the current item and move the cursor to the next item
func (arr *ArrayIterator) Next() interface{} {
	if arr.cursor < 0 || arr.cursor >= arr.size {
		panic("out index of iterator")
	}
	val := arr.array[arr.cursor]
	arr.cursor++
	return val
}
