package opcode

import (
	"reflect"
	"testing"
)

type opcode struct {
	intcode []int
	want    []int
}

var testCases = []opcode{
	{intcode: []int{1, 0, 0, 0, 99}, want: []int{2, 0, 0, 0, 99}},
	{intcode: []int{2, 3, 0, 3, 99}, want: []int{2, 3, 0, 6, 99}},
	{intcode: []int{2, 4, 4, 5, 99, 0}, want: []int{2, 4, 4, 5, 99, 9801}},
	{intcode: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, want: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
}

func TestOpcode(t *testing.T) {
	for _, tc := range testCases {
		got := ComputeOpcode(tc.intcode)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("ComputeOpcode(%v): want %v, got %v\n", tc.intcode, tc.want, got)
		}
	}
}
