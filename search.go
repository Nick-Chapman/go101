package main

import "fmt"

func runSearch() {
	fmt.Printf("search\n")
	A := []int{2, 3, 4}
	X := 5
	fmt.Printf("search: %v %v...\n", A, X)
	B := search(A, X)
	fmt.Printf("search: res= %v", B)
}

// Challenge set by Mike Taylor
// https://reprog.wordpress.com/2010/04/19/are-you-one-of-the-10-percent

func search(A []int, X int) bool {

	var loop func(i int, j int) bool

	loop = func(i int, j int) bool {
		if i > j {
			return false
		} else if i == j {
			if A[i] == X {
				return true
			} else {
				return false
			}
		} else {
			c := i + ((j - i) / 2)
			e := A[c]
			if e == X {
				return true
			} else if e < X {
				return loop(c+1, j)
			} else {
				return loop(i, c-1)
			}
		}
	}
	return loop(0, len(A)-1)
}
