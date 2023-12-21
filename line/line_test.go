package main

import (
	"math"
	"testing"
)

func TestApproximation(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []point
		want  float64
	}{
		{"1", []point{{100, 200}, {400, 300}, {600, 500}, {200, 500}}, 363},
		{"2", []point{{700, 100}, {3, 700}, {231, 552}}, 550},
		{"3", []point{{1, 2}, {3, 3}, {4, 5}, {6, 7}}, 4},
		{"4", []point{{0, 0}, {800, 600}}, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := approximate(tc.input)
			got := float64(f(0))
			if math.Abs(got-tc.want) > 1 {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

// func main() {
// 	testing.Main(
// 		/* matchString */ func(a, b string) (bool, error) { return a == b, nil },
// 		/* tests */ []testing.InternalTest{
// 			{Name: "Test Approximation", F: TestApproximation},
// 		},
// 		/* benchmarks */ nil /* examples */, nil)
// }
