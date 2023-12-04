package main

import "fmt"
import "cmp"

func runPlayWithSets() {
	fmt.Printf("play with sets...\n")
	s := empty[int]()
	see := func(tag string) {
		fmt.Printf("\n%v...\n", tag)
		fmt.Printf("set: %v\n", s)
		fmt.Printf("size? %v\n", size(s))
		fmt.Printf("A15? %v\n", member(15,s))
	}
	see("empty")
	s = makeIntRangeSet(10, 20)
	see("10..20")
	s = remove(15, s)
	see("removed-15")
}

func makeIntRangeSet(a, b int) set[int] {
	if b < a {
		return empty[int]()
	} else {
		return insert(b, makeIntRangeSet(a, b-1))
	}
}

func member[T elem](x T, set set[T]) bool {
	return set.member(x)
}
func empty[T elem]() set[T] {
	return emptySet[T]{}
}
func size[T elem](set set[T]) int {
	return set.size()
}
func insert[T elem](x T, set set[T]) set[T] {
	return set.insert(x)
}
func remove[T elem](x T, set set[T]) set[T] {
	return set.remove(x)
}
type elem = cmp.Ordered

type set[T elem] interface {
	null() bool
	size() int
	member(T) bool
	insert(T) set[T]
	remove(T) set[T]
}

type emptySet[T elem] struct{}
type binNode[T elem] struct {
	left  set[T]
	elem  T
	right set[T]
}

func (emptySet[T]) null() bool {
	return true
}
func (binNode[T]) null() bool {
	return false
}

func (emptySet[T]) size() int {
	return 0
}
func (b binNode[T]) size() int {
	return 1 + b.left.size() + b.right.size()
}

func (emptySet[T]) member(T) bool {
	return false
}
func (b binNode[T]) member(x T) bool {
	if x == b.elem {
		return true
	} else if x < b.elem {
		return b.left.member(x)
	} else {
		return b.right.member(x)
	}
}

func (emptySet[T]) insert(x T) set[T] {
	return makeBinNode(empty[T](), x, empty[T]())
}
func (b binNode[T]) insert(x T) set[T] {
	if x == b.elem {
		return b
	} else if x < b.elem {
		return makeBinNode(b.left.insert(x), b.elem, b.right)
	} else {
		return makeBinNode(b.left, b.elem, b.right.insert(x))
	}
}

func (b emptySet[T]) remove(T) set[T] {
	return b
}
func (b binNode[T]) remove(x T) set[T] {
	if x == b.elem {
		return abut(b.left, b.right)
	} else if x < b.elem {
		return makeBinNode(b.left.remove(x), b.elem, b.right)
	} else {
		return makeBinNode(b.left, b.elem, b.right.remove(x))
	}
}

func abut[T elem](left, right set[T]) set[T] {
	//panic("abut")
	return left // TODO: wrong!
}

func makeBinNode[T elem](left set[T], x T, right set[T]) set[T] {
	return binNode[T]{left,x,right}
}
