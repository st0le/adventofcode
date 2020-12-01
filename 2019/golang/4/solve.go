package main

import (
	"fmt"
	"strconv"
)

// Input 240298-784956

func isValid(n int) bool {
	repeats := 0
	num := strconv.Itoa(n)
	for i := 0; i < len(num); i++ {
		if i > 0 {
			if num[i-1] <= num[i] {
				if num[i-1] == num[i] {
					repeats++
				}
			} else {
				return false
			}

		}
	}
	return repeats > 0
}

func isValid2(n int) bool {
	repeats := make([]int, 10)
	num := strconv.Itoa(n)
	for i := 0; i < len(num); i++ {
		if i > 0 {
			if num[i-1] <= num[i] {
				if num[i-1] == num[i] {
					repeats[num[i]-'0']++
				}
			} else {
				return false
			}

		}
	}

	found := false
	for _, v := range repeats {
		if v == 1 {
			found = true
		}
	}
	return found
}

func main() {
	low, high, count := 240298, 784956, 0

	// Solve Part 1
	for i := low; i <= high; i++ {
		if isValid(i) {
			// fmt.Println(i)
			count++
		}
	}
	fmt.Println(count)

	// Solve Part 2
	count = 0
	for i := low; i <= high; i++ {
		if isValid2(i) {
			fmt.Println(i)
			count++
		}
	}
	fmt.Println(count)
}
