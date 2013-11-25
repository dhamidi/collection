package collection

import (
	"testing"
)

func ExampleMap() {
	v := NewVectorInt([]int{2, 4, 7})

	r := Map(v, func(i Value) Value {
		return i.(int) * 2
	})

	fmt.Printf("Numbers: %v\n", r)

	// Output:
	// Numbers: [4 8 14]
}

func ExampleMapX() {
	v := NewVectorInt([]int{2,4,6})
	r := MapX(v, func(i Value) Value {
		return i.(int) * 2
	})

	fmt.Printf("v: %v\n", v)
	fmt.Printf("r: %v\n", r)

	// Output:
	// v: [4 8 12]
	// r: [4 8 12]
}

func ExampleReduce() {
	r := NewVectorInt([]int{1,2,3})
	s := Reduce(r, func(i, j Value) Value {
		return i.(int) + j.(int)
	})

	fmt.Printf("Sum: %v\n", s)

	// Output:
	// Sum: 6
}

func ExampleFilter() {
	r := NewVectorInt([]int{1,2,3})

	even := func(i interface{}) bool {
		return i.(int) % 2 == 0
	}

	evenNumbers := Filter(v, even)

	fmt.Printf("Even numbers: %v\n", evenNumbers)

	// Output:
	// Even numbers: [2]
}
