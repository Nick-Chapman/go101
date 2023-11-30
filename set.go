package main

import "fmt"

func runPlayWithSets() {
	fmt.Printf("play with sets...\n")
	s := empty[int]()
	see := func(tag string) {
		fmt.Printf("\n%v...\n", tag)
		fmt.Printf("set: %s\n", s)
		fmt.Printf("null? %v\n", s.null())
		fmt.Printf("4? %v\n", s.member(4))
		fmt.Printf("5? %v\n", s.member(5))
	}
	see("empty")
	s = insert(5, s)
	see("added-5")
	s = insert(4, s)
	see("added-4")
	s = delete(5, s)
	see("deleted-5")
}

// API...

/*
type elem = comparable

type set[T elem] interface {
	member(T) bool
	null() bool
}

func empty[T elem]() set[T] {
	panic("empty")
}

func insert[T elem](x T, set set[T]) set[T] {
	panic("insert")
}

func delete[T elem](x T, set set[T]) set[T] {
	panic("delete")
}
*/


// simplistic implementation...

type elem = comparable

type set[T elem] interface {
	member(T) bool
	null() bool
	delete(T) set[T]
}

func empty[T elem]() set[T] {
	return &emptySet[T]{}
}

func insert[T elem](x T, set set[T]) set[T] {
	return &consSet[T]{x,set}
}

func delete[T elem](x T, set set[T]) set[T] {
	return set.delete(x)
}

type emptySet[T elem] struct{}

func (*emptySet[T]) null() bool {
	return true
}
func (*emptySet[T]) member(T) bool {
	return false
}
func (me *emptySet[T]) delete(T) set[T] {
	return me
}

type consSet[T elem] struct{
	x T;
	xs set[T];
}

func (*consSet[T]) null() bool {
	return false
}
func (me *consSet[T]) member(x T) bool {
	return me.x == x || me.xs.member(x)
}
func (me *consSet[T]) delete(x T) set[T] {
	if me.x == x {
		return me.xs
	} else {
		return insert(me.x, delete (x,me.xs))
	}
}
