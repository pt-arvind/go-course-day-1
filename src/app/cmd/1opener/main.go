package main

import "log"

/*
If we list all the natural numbers below 30 that are multiples of 7 or 9, we get 7, 9, 14, 18, 21, 27, and 28.
The sum of these multiples is 124.
Find the sum of all the multiples of 7 or 9 below BUT NOT INCLUDING 30,000.
*/
func main() {
	sum := 0
	for i := 7; i < 30000; i++ {
		if i%7 == 0 {
			sum += i
			continue
		}
		if i%9 == 0 {
			sum += i
		}
	}

	log.Printf("sum: %+v\n", sum)
}

//ans: 107132146
