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
		isConsecutive := make([]bool, 0)
		isIncreasing := true
		for j := 0; j < len(s)-1; j++ {
			cur, _ := strconv.Atoi(string(s[j]))
			next, _ := strconv.Atoi(string(s[j+1]))

			isConsecutive = append(isConsecutive, cur == next)
			if cur > next {
				isIncreasing = false
			}
		}

		if (testEq(isConsecutive) || testEq(isConsecutive)) && isIncreasing {
			numOfSolutions += 1
			fmt.Printf("%v\n", i)
		}
	}
	fmt.Printf("Number of Solutions: %v", numOfSolutions)
}

// 112233
// TFTFT

// 122334
// FTFTF
// 588899
// FTTFT

func testEq(a []bool) bool {

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if a[i] == true && a[i+1] == false {
				return true
			}
		} else if i == len(a)-1 {
			if a[i] == true && a[i-1] == false {
				return true
			}
		} else {
			if a[i] == true && a[i-1] == false && a[i+1] == false {
				return true
			}
		}
	}

	return false
}
