package collection

import (
	"fmt"
)

type Value interface{}
type UnaryFunction func(Value) Value
type UnaryPredicate func(Value) bool
type BinaryFunction func(Value, Value) Value

type Collection interface {
	Item(index int) Value
	SetItem(index int, value Value) Value
	Length() int
	Append(item Value) Collection
}

type Vector []Value

func NewVector(items []Value) Vector {
	data := items

	if items == nil {
		data = []Value{}
	}

	return data
}

func NewVectorInt(items []int) Vector {
	result := Vector{}

	if items == nil {
		return result
	}

	for _, item := range items {
		result.Append(item)
	}

	return result
}

func (self Vector) String() string {
	return fmt.Sprintf("%v", self)
}

func (self Vector) Length() int {
	return len(self)
}

func (self Vector) Item(index int) Value {
	return self[index]
}

func (self Vector) SetItem(index int, value Value) Value {
	oldValue := self.Item(index)
	self[index] = value
	return oldValue
}

func (self Vector) Append(item Value) Collection {
	self = append(self, item)
	return self
}

func Map(c Collection, f UnaryFunction) Collection {
	results := NewVector(nil)
	len := c.Length()
	for index := 0; index < len; index++ {
		results.Append(f(c.Item(index)))
	}

	return results
}

func MapX(c Collection, f UnaryFunction) Collection {
	len := c.Length()

	for index := 0; index < len; index++ {
		item := c.Item(index)
		c.SetItem(index, f(item))
	}

	return c
}

func Reduce(c Collection, f BinaryFunction) Value {
	result := c.Item(0)
	len := c.Length()
	for index := 1; index < len; index++ {
		result = f(result, c.Item(index))
	}
	return result
}

func Filter(c Collection, p UnaryPredicate) Collection {
	result := NewVector(nil)
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
