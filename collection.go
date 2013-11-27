// Package collection provides common operations on generic containers.
package collection

import (
	"fmt"
)

// Value is an alias for the empty interface type.  This alias is
// provided to make function signatures more readable.
type Value interface{}

// UnaryFunction encodes the type of callbacks expected by Map and MapX
type UnaryFunction func(Value) Value

// UnaryPredicate encodes the type of callbacks expected by Filter
type UnaryPredicate func(Value) bool

// BinaryFunction encodes the type of callbacks expected by Reduce and ReduceFirst
type BinaryFunction func(Value, Value) Value

// Collection is the interface that wraps a generic container with
// indexed access to each element.
//
// Item returns a single element of the collection.  The element
// is identified by `index'.
//
// SetItem sets a single element of the collection, identified
// by `index', to a new value, `value'.  It returns the value of
// the element previously identified by `index'.
//
// Length returns the number of elements in a collection.
//
// Append adds one element, `item', to the end of the
// collection.  It returns a collection containing all elements
// before the call to `Append' and `item'.  Whether the returned
// collection is newly allocated or the same reference as the
// method receiver is left up to implementation.
//
// Empty returns a new, empty collection of the same type.  This is
// required for supporting `Filter' and a non-destructive `Map'.
type Collection interface {
	Item(index int) Value
	SetItem(index int, value Value) Value
	Length() int
	Append(item Value) Collection
	Empty() Collection
}

// Vector is a collection based on Go's builtin slice type.
type Vector []Value

// NewVector creates a new vector initially containing the elements
// passed in `items'.
//
// The value of `items' may be `nil', in which case a new slice of
// length 0 is allocated and used instead.
func NewVector(items []Value) *Vector {
	data := Vector(items)

	if items == nil {
		data = Vector(make([]Value, 0))
	}

	return &data
}

// String returns the string representation of a vector.  This matches
// Go's `v' format specifier of the underlying slice.
func (self *Vector) String() string {
	return fmt.Sprintf("%v", *self)
}

// Length returns the length of the vector.  This is equal to calling
// Go's built-in `len' function on the underlying slice.
func (self *Vector) Length() int {
	return len(*self)
}

// Item returns the element located at position `index'.  This is equal
// to indexing the underlying slice using brackets.  If the index is out
// of bounds, this function panics.
func (self *Vector) Item(index int) Value {
	return (*self)[index]
}

// SetItem sets the element at position `index' to `value' and returns
// the element that was at position `index' before the assignment.
func (self *Vector) SetItem(index int, value Value) Value {
	oldValue := self.Item(index)

	(*self)[index] = value

	return oldValue
}

// Append appends `item' to the end of the underlying slice and returns
// the method receiver.
func (self *Vector) Append(item Value) Collection {
	*self = append(*self, item)

	return self
}

// Empty returns a new, empty vector.  The capacity of the vector is
// initialized to the `Length' of the method receiver.
func (self *Vector) Empty() Collection {
	return NewVector(make([]Value, 0, self.Length()))
}

// Map determines the length of the collection `c' exactly once and then
// iterates over the range [0, length).
//
// It returns a new collection that contains the results of applying `f'
// to each element in that range of the original collection.
func Map(c Collection, f UnaryFunction) Collection {
	results := c.Empty()
	len := c.Length()

	for index := 0; index < len; index++ {
		results.Append(f(c.Item(index)))
	}

	return results
}

// MapX works like `Map', except that it modifies and returns the
// original collection (`c').
func MapX(c Collection, f UnaryFunction) Collection {
	len := c.Length()

	for index := 0; index < len; index++ {
		item := c.Item(index)
		c.SetItem(index, f(item))
	}

	return c
}

// Reduce applies a function of two arguments `f' to an accumulator
// value and successive elements of `c'.  The length of `c' is
// determined once using `Length'.  It returns the value of the
// accumulator at the end of the iteration.
//
// The first argument to `f' is the accumulator (initialized to
// `initial'), the second argument is the current element of the
// collection.
func Reduce(c Collection, f BinaryFunction, initial Value) Value {
	result := initial
	len := c.Length()

	for index := 0; index < len; index++ {
		result = f(result, c.Item(index))
	}

	return result
}

// ReduceFirst works like `Reduce', except that it uses the first
// element of the collection as the initial value.
func ReduceFirst(c Collection, f BinaryFunction) Value {
	result := c.Item(0)
	len := c.Length()

	for index := 1; index < len; index++ {
		result = f(result, c.Item(index))
	}

	return result
}

// Filter returns a new collection containing only the elements of `c'
// for which `p' returns `true'.
func Filter(c Collection, p UnaryPredicate) Collection {
	result := c.Empty()
	len := c.Length()
	var item Value

	for index := 0; index < len; index++ {
		item = c.Item(index)
		if p(item) {
			result.Append(item)
		}
	}

	return result
}
