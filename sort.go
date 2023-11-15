package main

import "fmt"
import "slices"
import "math/rand"

//type nums chan int // in other file

func runSort() {
	fmt.Printf("sort...\n")
	n := 70     // number of elements to sort
	max := 1000 // max size of element
	c1 := make(chan int)
	c2 := make(chan int)
	go randomDriver(n, max, c1)
	go splittingSorter("x", c1, c2)
	mes := s_printer(n, c2)
	fmt.Printf("%s\n", mes)
}

func splittingSorter(name string, in <-chan int, out chan<- int) {
	// Assumes in has at least one element!
	i, more := <-in
	j, more := <-in
	if !more {
		//fmt.Printf("(%s:LEAF)", name)
		out <- i
		close(out)
		return
	}
	nameL := name + "L"
	nameR := name + "R"
	left := make(chan int)
	right := make(chan int)
	leftSorted := make(chan int)
	rightSorted := make(chan int)
	go merger(leftSorted, rightSorted, out)
	go splittingSorter(nameL, left, leftSorted)
	go splittingSorter(nameR, right, rightSorted)
	left <- i
	right <- j
	go splitter(in, left, right)
}

func splitter(in <-chan int, left, right chan<- int) {
	for {
		i, more := <-in
		if !more {
			close(left)
			close(right)
			return
		}
		left <- i
		left, right = right, left // FLIP
	}
}

func merger(left, right <-chan int, out chan<- int) {
	// Assumes both left & right have at least one element!
	var l, r int
	l = <-left // get zeros if/when the channel is empty!
	r = <-right
	var more bool
	for {
		if l <= r { // The *only* use of a comparison op in the sort algorithm
			out <- l
			l, more = <-left
			if !more {
				out <- r
				passThrough(right, out)
				break
			}
		} else {
			out <- r
			r, more = <-right
			if !more {
				out <- l
				passThrough(left, out)
				break
			}
		}
	}
	close(out)
}

func passThrough(in <-chan int, out chan<- int) {
	for {
		i, more := <-in
		if !more {
			break
		}
		out <- i
	}
}

// This is not used...
func monoSorter(name string, in <-chan int, out chan<- int) {
	var buf []int
	for {
		i, more := <-in
		if !more {
			fmt.Printf("(%s=%d)", name, len(buf))
			slices.Sort(buf)
			for _, j := range buf {
				out <- j
			}
			close(out)
			return
		}
		buf = append(buf, i)
	}
}

func randomDriver(n int, max int, out chan<- int) {
	for i := 0; i < n; i++ {
		r := rand.Intn(max)
		out <- r
	}
	close(out)
}

func s_printer(nExpected int, in <-chan int) string { //meh. need better way to do namespacs
	n := 0
	last := 0
	for {
		i, more := <-in
		if !more {
			if n != nExpected {
				return fmt.Sprintf("**LOST %d", nExpected-n)
			} else {
				return "ok"
			}
		}
		ok := last <= i
		n++
		fmt.Printf("%d ", i)
		if !ok {
			return "**UNSORTED"
		}
		last = i
	}
}
