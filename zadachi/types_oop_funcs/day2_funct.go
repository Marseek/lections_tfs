package main

import (
	"fmt"
	"sort"
)

func main5() {
	type T struct {
		a int
		b float64
		c string
	}
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
}

func main2() {
	f1, f2 := New(), New()
	f1() // 1
	f2() // 1 (другая n)
	f1() // 2
	f2() // 2
}

func swTest() {
	var t interface{}
	t = 5

	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}

func New() func() {
	n := 0
	return func() {
		n++
		fmt.Println(n)
	}
}

func lambdas() {

	people := []string{"Alice", "Bob", "Dave"}
	sort.Slice(people, func(i, j int) bool {
		return len(people[i]) < len(people[j])
	})
	fmt.Println(people)
}
