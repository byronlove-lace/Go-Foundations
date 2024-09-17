package main

import "fmt"

var x any = 5

func SumIntOrFloat[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, num := range m {
		sum += num
	}
	return sum
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	return r
}

// Linked List using Generics
type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}

	return elems
}

func main() {
	var m = map[int]string{1: "hi", 2: "howareye", 3: "hullo"}

	fmt.Println("keys:", MapKeys(m))

	// Don't need to specifiy types
	// _ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(29)
	lst.Push(31)

	fmt.Println(lst.GetAll())
}
