package main

import (
	"fmt"
	"strconv"
)

func main() {
	lower := 256310
	upper := 732736

	numOfSolutions := 0
	for i := lower; i < upper; i++ {
		s := strconv.Itoa(i)
		isConsecutive := false
		isIncreasing := true
		for j := 0; j < len(s)-1; j++ {
			cur, _ := strconv.Atoi(string(s[j]))
			next, _ := strconv.Atoi(string(s[j+1]))

			if cur == next {
				isConsecutive = true
			}
			if cur > next {
				isIncreasing = false
			}
		}

		if isConsecutive && isIncreasing {
			numOfSolutions += 1
		}
	}
	fmt.Printf("Number of Solutions: %v", numOfSolutions)
}
