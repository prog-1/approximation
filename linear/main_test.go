package main

import (
	"testing"
)

func TestApproximation(t *testing.T) {
	for _, tc := range []struct {
		name         string
		input        []Point
		wantA, wantB float64
	}{
		{"1", []Point{{100, 200}, {400, 300}}, 0.3333333333333333, 166.66666666666669},
	} {
		t.Run(tc.name, func(t *testing.T) {
			gotA, gotB := approximation(tc.input)
			if gotA != tc.wantA || gotB != tc.wantB {
				t.Errorf("got = %v, %v, want = %v, %v", gotA, gotB, tc.wantA, tc.wantB)
			}
		})
	}
}
