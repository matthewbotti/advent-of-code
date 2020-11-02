// Package santafuel contains functions for solving Santa's fuel requirements
// per https://adventofcode.com/2019/day/1 exercise.
package fuel

import "math"

// FuelRequired determines fuel needed from a given mass.
func FuelRequired(mass float64) float64 {
	return math.Floor(mass/3) - 2
}

// RecursiveFuel determines fuel needed recursivley given given a mass.
func RecursiveFuel(mass float64) (sum float64) {
	fuelNeeded := FuelRequired(mass)
	for sum = fuelNeeded; FuelRequired(fuelNeeded) > 0; sum += fuelNeeded {
		fuelNeeded = FuelRequired(fuelNeeded)
	}
	return
}
