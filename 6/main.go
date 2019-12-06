package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type body struct {
	parent string
}

var bodies map[string]body

func main() {
	bodies = make(map[string]body)
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ")")
		bodies[row[1]] = body{parent: row[0]}
	}
	totalOrbits := 0
	for k := range bodies {
		totalOrbits = totalOrbits + countOrbits(bodies[k])
	}
	fmt.Println("Part A:", totalOrbits)
	youPath := listPath(bodies["YOU"])
	sanPath := listPath(bodies["SAN"])

	fmt.Println("Part B:", findShortestPath(youPath, sanPath))

}

func findShortestPath(a, b []string) int {
	// Start by finding the first common term
	for i, aVal := range a {
		for k, bVal := range b {
			if aVal == bVal {
				return i + k
			}
		}
	}
	return 1

}

func listPath(b body) []string {
	// fmt.Println(b.parent)
	path := []string{b.parent}
	if b.parent == "COM" {
		return []string{}
	}
	path = append(path, listPath(bodies[b.parent])...)
	return path

}

func countOrbits(b body) int {
	num := 1
	if b.parent == "COM" {
		return 1
	}
	num = num + countOrbits(bodies[b.parent])
	return num

}
