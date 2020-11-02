// Solutions to Advent of Code 2019 - Day 1.
package main

import (
	"fmt"
	"log"

	"github.com/matthewbotti/advent-of-code/2019/day-01/fuel"
	"github.com/matthewbotti/advent-of-code/datafile"
)

func main() {
	modules, err := datafile.GetFloats("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Part 1 Solution
	var sum1 int
	for _, module := range modules {
		sum1 += int(fuel.FuelRequired(module))
	}
	fmt.Printf("2019: Day1 - Part1 = %d\n", sum1)

	// Part 2 Solution
	var sum2 int
	for _, module := range modules {
		sum2 += int(fuel.RecursiveFuel(module))
	}
	fmt.Printf("2019: Day1 - Part2 = %d\n", sum2)
}
