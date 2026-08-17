package regression

import "testing"

func TestScenario06(t *testing.T) {
	in := Scenario06Inputs{A: 1, B: 2, C: 3, D: 4, E: 5, F: 6, G: 7, H: 8, I: 9, J: 10, K: 11, L: 12}
	if got, want := Scenario06(in), in.F+in.G; got != want {
		t.Fatalf("Scenario06 = %d, want %d", got, want)
	}
}
