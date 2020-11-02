package fuel

import "testing"

type fuelArgs struct {
	mass  float64
	want1 float64
	want2 float64
}

var testCases = []fuelArgs{
	{mass: 12, want1: 2, want2: 2},
	{mass: 14, want1: 2, want2: 2},
	{mass: 1969, want1: 654, want2: 966},
	{mass: 100756, want1: 33583, want2: 50346},
}

func TestFuel(t *testing.T) {
	for _, tc := range testCases {
		got := FuelRequired(tc.mass)
		if tc.want1 != got {
			t.Errorf("fuelRequired(%f): want %f, got %f", tc.mass, tc.want1, got)
		}
	}
}

func TestRecursiveFuel(t *testing.T) {
	for _, tc := range testCases {
		got := RecursiveFuel(tc.mass)
		if tc.want2 != got {
			t.Errorf("RecursiveFuel(%f): want %f, got %f", tc.mass, tc.want2, got)
		}
	}
}
