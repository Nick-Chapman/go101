package main

import "fmt"

func runEvalExample() {
	fmt.Println("Eval example...")
	a := &leafsum{1, 2}
	b := &leaf{33}
	c := &leafsum{4, 5}

	d := &sum{b, c}
	e := makeFibTree(8)
	ys := []expr{a, b, c, d, e}

	for i, y := range ys {
		fmt.Printf("expr: (%v) %v --> %d\n", i, y, y.eval())
	}
}

func makeFibTree(n int) expr {
	if n <= 1 {
		return &leaf{n}
	} else {
		return &sum{makeFibTree(n - 1), makeFibTree(n - 2)}
	}
}

type expr interface {
	eval() int
}

type leaf struct {
	value int
}

type leafsum struct {
	left  int
	right int
}

type sum struct {
	left  expr
	right expr
}

func (x *leafsum) eval() int {
	return x.left + x.right
}

func (x *sum) eval() int {
	return x.left.eval() + x.right.eval()
}

func (x *leaf) eval() int {
	return x.value
}
