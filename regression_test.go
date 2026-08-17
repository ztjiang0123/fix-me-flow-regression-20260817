package regression

import "testing"

func TestScenario05(t *testing.T) {
	in := Scenario05Inputs{A: 1, B: 2, C: 3, D: 4, E: 5, F: 6, G: 7, H: 8, I: 9, J: 10, K: 11, L: 12}
	if got, want := Scenario05(in), in.E+in.H; got != want {
		t.Fatalf("Scenario05 = %d, want %d", got, want)
	}
}

func TestScenario06(t *testing.T) {
	in := Scenario06Inputs{A: 1, B: 2, C: 3, D: 4, E: 5, F: 6, G: 7, H: 8, I: 9, J: 10, K: 11, L: 12}
	if got, want := Scenario06(in), in.F+in.G; got != want {
		t.Fatalf("Scenario06 = %d, want %d", got, want)
	}
}

func TestScenario07(t *testing.T) {
	in := Scenario07Inputs{A: 1, B: 2, C: 3, D: 4, E: 5, F: 6, G: 7, H: 8, I: 9, J: 10, K: 11, L: 12}
	if got, want := Scenario07(in), in.G+in.F; got != want {
		t.Fatalf("Scenario07 = %d, want %d", got, want)
	}
}

func TestScenario08(t *testing.T) {
	in := Scenario08Inputs{A: 1, B: 2, C: 3, D: 4, E: 5, F: 6, G: 7, H: 8, I: 9, J: 10, K: 11, L: 12}
	if got, want := Scenario08(in), in.H+in.E; got != want {
		t.Fatalf("Scenario08 = %d, want %d", got, want)
	}
}

func TestScenario09(t *testing.T) {
	in := Scenario09Inputs{A: 1, B: 2, C: 3, D: 4, E: 5, F: 6, G: 7, H: 8, I: 9, J: 10, K: 11, L: 12}
	if got, want := Scenario09(in), in.I+in.D; got != want {
		t.Fatalf("Scenario09 = %d, want %d", got, want)
	}
}
