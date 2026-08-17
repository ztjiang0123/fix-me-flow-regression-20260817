package regression

import "testing"

// sampleInputs assigns each field a distinct value so a scenario picking the
// wrong pair would produce a distinguishable sum.
func sampleInputs() ScenarioInputs {
	return ScenarioInputs{
		A: 1, B: 2, C: 3, D: 4, E: 5, F: 6,
		G: 7, H: 8, I: 9, J: 10, K: 11, L: 12,
	}
}

func TestScenariosPreserveOriginalPairSums(t *testing.T) {
	in := sampleInputs()

	cases := []struct {
		name string
		got  int
		want int
	}{
		{"Scenario01", Scenario01(in), in.A + in.L},
		{"Scenario02", Scenario02(in), in.B + in.K},
		{"Scenario03", Scenario03(in), in.C + in.J},
		{"Scenario04", Scenario04(in), in.D + in.I},
		{"Scenario05", Scenario05(in), in.E + in.H},
		{"Scenario06", Scenario06(in), in.F + in.G},
		{"Scenario07", Scenario07(in), in.G + in.F},
		{"Scenario08", Scenario08(in), in.H + in.E},
		{"Scenario09", Scenario09(in), in.I + in.D},
		{"Scenario10", Scenario10(in), in.J + in.C},
		{"Scenario11", Scenario11(in), in.K + in.B},
		{"Scenario12", Scenario12(in), in.L + in.A},
	}

	for _, tc := range cases {
		if tc.got != tc.want {
			t.Errorf("%s = %d, want %d", tc.name, tc.got, tc.want)
		}
	}
}
