package main

import "fmt"

type nums chan int

func runPrimes() {
	fmt.Printf("primes...\n")
	c := make(nums)
	go nats(1000, c)
	printer(c)
	fmt.Printf("\nprimes...done\n")
}

func printer(in nums) {
	for {
		p, more := <-in
		if !more {
			return
		}
		fmt.Printf("%v ", p)
		cNew := make(nums)
		go filter(p, in, cNew)
		in = cNew
	}
}

func nats(max int, out nums) {
	for i := 2; i <= max; i++ {
		out <- i
	}
	close(out)
}

func filter(n int, in nums, out nums) {
	for {
		i, more := <-in
		if !more {
			close(out)
			return
		}
		if i%n != 0 {
			out <- i
		}
	}
}
