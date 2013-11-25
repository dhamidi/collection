package collection

import (
	"fmt"
)

type Collection interface {
	Item(index int) interface{}
	SetItem(index int, value interface{}) interface{}
	Length() int
	Append(item interface{}) Collection
}

type Vector struct {
	data []interface{}
}

type Value interface{}
type UnaryFunction func (Value) Value
type UnaryPredicate func (interface{}) bool
type BinaryFunction func (Value, Value) Value

func NewVector(items []interface{}) *Vector {
	data := items

	if items == nil {
		data = []interface{}{}
	}

	return &Vector{
		data: data,
	}
}

func NewVectorInt(items []int) *Vector {
	result := &Vector{data: []interface{}{}}

	if items == nil {
		return result
	}

	for _, item := range items {
		result.Append(item)
	}

	return result
}

func (self *Vector) String() string {
	return fmt.Sprintf("%v", self.data)
}

func (self *Vector) Length() int {
	return len(self.data)
}

func (self *Vector) Item(index int) interface{} {
	return self.data[index]
}

func (self *Vector) SetItem(index int, value interface{}) interface{} {
	oldValue := self.Item(index)
	self.data[index] = value
	return oldValue
}

func (self *Vector) Append(item interface{}) Collection {
	self.data = append(self.data, item)
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

func Reduce(c Collection, f BinaryFunction) interface{} {
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
	var item interface{}

	for index := 0; index < len; index++ {
		item = c.Item(index)
		if p(item) {
			result.Append(item)
		}
	}

	return result
}
