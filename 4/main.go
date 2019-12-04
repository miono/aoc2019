package main

import (
	"fmt"
	"strconv"
)

func main() {
	low := 372304
	high := 847060
	var workingPWsA []int
	var workingPWsB []int

	for i := low; i < high; i++ {
		if doubleDigits(i) && increasing(i) {
			workingPWsA = append(workingPWsA, i)
		}
		if doubleDigits2(i) && increasing(i) {
			workingPWsB = append(workingPWsB, i)
		}

	}
	fmt.Println("Part A:", len(workingPWsA))
	fmt.Println("Part B:", len(workingPWsB))

}

func doubleDigits(pw int) bool {
	pwString := strconv.Itoa(pw)
	for i := 0; i < len(pwString)-1; i++ {
		if pwString[i] == pwString[i+1] {
			return true
		}
	}
	return false
}

func doubleDigits2(pw int) bool {
	pwString := strconv.Itoa(pw)
	freq := make(map[int]int)
	for _, j := range pwString {
		jInt, _ := strconv.Atoi(string(j))
		freq[jInt]++

	}
	for _, v := range freq {
		if v == 2 {
			return true
		}

	}
	return false
}

func increasing(pw int) bool {
	pwString := strconv.Itoa(pw)
	for i := 0; i < len(pwString)-1; i++ {
		iInt, _ := strconv.Atoi(string(pwString[i]))
		jInt, _ := strconv.Atoi(string(pwString[i+1]))
		if iInt > jInt {
			return false
		}
	}
	return true
}
