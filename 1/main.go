package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var totalFuelPart1 int
	var totalFuelPart2 int
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		oneFuel, _ := strconv.Atoi(scanner.Text())
		twoFuel, _ := strconv.Atoi(scanner.Text())
		totalFuelPart1 = totalFuelPart1 + getFuel(oneFuel)
		totalFuelPart2 = totalFuelPart2 + getFuelRecursive(twoFuel)
	}
	fmt.Println(totalFuelPart1)
	fmt.Println(totalFuelPart2)

}

func getFuel(mass int) int {
	return (mass / 3) - 2
}

func getFuelRecursive(mass int) int {
	totalFuel := 0
	oneFuel := mass
	for oneFuel > 1 {
		oneFuel = getFuel(oneFuel)
		if oneFuel > 0 {
			totalFuel = totalFuel + oneFuel
		}
	}
	return totalFuel

}
