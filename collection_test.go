package collection

import (
	"fmt"
	"testing"
)

// doubleInt is used as a UnaryFunction for testing Map and MapX
func doubleInt(i Value) Value {
	return i.(int) * 2
}

// addInt is used as a BinaryFunction for testing Reduce and ReduceFirst
func addInt(a, b Value) Value {
	return a.(int) + b.(int)
}

// isEvenInt is used as a UnaryPredicate for testing Filter and sort
func isEvenInt(a Value) bool {
	return a.(int)%2 == 0
}

func TestVector_Length(t *testing.T) {
	v := NewVector([]Value{1, 2, 3})

	if length := v.Length(); length != 3 {
		t.Errorf("Expected vector.Length() to be %d, got %d", 3, length)
	}
}

func TestVector_Append(t *testing.T) {
	v := NewVector(nil)
	v.Append(1)

	if length := v.Length(); length != 1 {
		t.Errorf("Expected vector.Append(1) to increase length by one")
	}

	if item := v.Item(0); item.(int) != 1 {
		t.Errorf("Expected the first item to be 1, got %v", item)
	}
}

func TestVector_Item(t *testing.T) {
	v := NewVector([]Value{1, 2, 3})

	length := v.Length()
	for index := 0; index < length; index++ {
		if item := v.Item(index); item.(int) != index+1 {
			t.Errorf("Expected vector.Item(%d) to be %d, got %v", index, index+1, item)
		}
	}
}

func TestVector_SetItem(t *testing.T) {
	v := NewVector(nil)
	v.Append(1)
	v.SetItem(0, 2)

	if actual, expected := v.Item(0), 2; actual != expected {
		t.Errorf("Expected v.Item(0) to be %v, got %v", expected, actual)
	}
}

func TestVector_String(t *testing.T) {
	v := NewVector([]Value{1, 2, 3})

	if actual, expected := v.String(), "[1 2 3]"; actual != expected {
		t.Errorf("Expected v.String() to be %v, got %v", expected, actual)
	}
}

func TestMap(t *testing.T) {
	v := NewVector([]Value{1, 2, 3})
	mv := Map(v, doubleInt)
	len := v.Length()

	for index := 0; index < len; index++ {
		if mv.Item(index) != 2*v.Item(index).(int) {
			t.Errorf("Expected mappedVector[%d] to be %d", index, 2*v.Item(index).(int))
		}
	}
}

func ExampleMap() {
	v := NewVector([]Value{1, 2, 3})

	r := Map(v, func(i Value) Value {
		return i.(int) * 2
	})
	fmt.Printf("%v\n", r)
	// Output: [2 4 6]
}

func TestMapX(t *testing.T) {
	v := NewVector([]Value{1, 2, 3})
	mv := MapX(v, doubleInt)
	len := v.Length()

	for index := 0; index < len; index++ {
		if actual, expected := mv.Item(index), v.Item(index); actual != expected {
			t.Errorf("Expected mv.Item(index) to be %v, got %v", expected, actual)
		}
	}
}

func ExampleMapX() {
	v := NewVector([]Value{2, 4, 6})
	r := MapX(v, func(i Value) Value {
		return i.(int) * 2
	})

	fmt.Printf("v: %v\n", v)
	fmt.Printf("r: %v\n", r)

	// Output:
	// v: [4 8 12]
	// r: [4 8 12]
}

func TestReduce(t *testing.T) {
	v := NewVector([]Value{2, 4, 6})
	sum := Reduce(v, addInt, 0)
	if actual, expected := sum, 12; actual != expected {
		t.Errorf("Expected sum to be %v, got %v", expected, actual)
	}
}

func ExampleReduce() {
	r := NewVector([]Value{1, 2, 3})
	s := Reduce(r, func(i, j Value) Value {
		return i.(int) + j.(int)
	}, 0)

	fmt.Printf("Sum: %v\n", s)

	// Output:
	// Sum: 6
}

func TestReduceFirst(t *testing.T) {
	v := NewVector([]Value{1, 2, 3})
	if actual, expected := ReduceFirst(v, addInt), 6; actual != expected {
		t.Errorf("Expected ReduceFirst(v, addInt) to be %v, got %v", expected, actual)
	}
}

func ExampleReduceFirst() {
	r := NewVector([]Value{1, 2, 3})
	s := ReduceFirst(r, func(i, j Value) Value {
		return i.(int) + j.(int)
	})

	fmt.Printf("Sum: %v\n", s)
	// Output:
	// Sum: 6
}

func TestFilter(t *testing.T) {
	v := NewVector([]Value{1, 2, 3})
	filtered := Filter(v, isEvenInt)

	if actual, expected := filtered.Item(0), 2; actual != expected {
		t.Errorf("Expected filtered.Item(0) to be %v, got %v", expected, actual)
	}

	if actual, expected := filtered.Length(), 1; actual != expected {
		t.Errorf("Expected filtered.Length() to be %v, got %v", expected, actual)
	}
}

func ExampleFilter() {
	r := NewVector([]Value{1, 2, 3})

	even := func(i Value) bool {
		return i.(int)%2 == 0
	}

	evenNumbers := Filter(r, even)

	fmt.Printf("Even numbers: %v\n", evenNumbers)

	// Output:
	// Even numbers: [2]
}
