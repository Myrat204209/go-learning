package main

import (
	"fmt"
)

func printPrimes(max int) {
	fmt.Println("2")
	for i := 3; i <= max; i += 2 {
		fl := false
		for j := 3; j*j <= i; j += 2 {
			if i%j == 0 {
				fl = true
				continue
			}
		}
		if fl == false {
			fmt.Println(i)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
