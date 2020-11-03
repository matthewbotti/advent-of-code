package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/matthewbotti/advent-of-code/2019/day-02/opcode"
	"github.com/matthewbotti/advent-of-code/datafile"
)

// stringsToInts converts a string slice to an int slice.
func stringsToInts(s []string) []int {
	var intSlice []int
	for _, value := range s {
		num, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Could not convert string to int")
		}
		intSlice = append(intSlice, num)
	}
	return intSlice
}

func main() {
	s, err := datafile.GetStrings("input.txt")
	if err != nil {
		log.Fatal("Could not extract from file")
	}
	stringSlice := strings.Split(s[0], ",")
	intcode := stringsToInts(stringSlice)
	intcode[1], intcode[2] = 12, 2
	fmt.Printf("2019: Day2 - Part1 = %d\n", opcode.ComputeOpcode(intcode)[0])
}
