package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x     int
	y     int
	steps int
}

type wire struct {
	points []point
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func main() {
	var wireA wire
	var wireB wire
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	instructions := strings.Split(scanner.Text(), ",")
	wireA = drawOnGrid(instructions, wireA)
	scanner.Scan()
	instructions = strings.Split(scanner.Text(), ",")
	wireB = drawOnGrid(instructions, wireB)
	compareWires(wireA, wireB)

}

func drawOnGrid(instructions []string, wire wire) wire {
	var xPos, yPos int
	steps := 0
	for _, instr := range instructions {
		direction := string(instr[0])
		length, _ := strconv.Atoi(instr[1:])
		switch direction {
		case "L":
			for i := 0; i < length; i++ {
				wire.points = append(wire.points, point{xPos - i, yPos, steps})
				steps++
			}
			xPos = xPos - length
		case "R":
			for i := 0; i < length; i++ {
				wire.points = append(wire.points, point{xPos + i, yPos, steps})
				steps++
			}
			xPos = xPos + length
		case "U":
			for i := 0; i < length; i++ {
				wire.points = append(wire.points, point{xPos, yPos + i, steps})
				steps++
			}
			yPos = yPos + length
		case "D":
			for i := 0; i < length; i++ {
				wire.points = append(wire.points, point{xPos, yPos - i, steps})
				steps++
			}
			yPos = yPos - length

		}
	}
	return wire
}

func compareWires(w1, w2 wire) {
	lowestStep := 100000000000
	lowestManhattan := 100000000
	for _, w1Point := range w1.points {
		for _, w2Point := range w2.points {
			if w1Point.x == w2Point.x && w1Point.y == w2Point.y {
				if w1Point.steps+w2Point.steps < lowestStep && w1Point.steps != 0 {
					lowestStep = w1Point.steps + w2Point.steps
				}
				if abs(w1Point.x)+abs(w1Point.y) < lowestManhattan && w1Point.steps != 0 {
					lowestManhattan = abs(w1Point.x) + abs(w1Point.y)
				}
			}

		}
	}
	fmt.Println("Part A:", lowestManhattan)
	fmt.Println("Part B:", lowestStep)
}
