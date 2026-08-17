package regression

// These scenarios provide deterministic Code Analyzer fixtures for exercising
// Fix Me and Fix All workflows. They previously duplicated the same body across
// every function and accepted a long positional parameter list; both are now
// resolved by grouping the inputs into a single value type and delegating the
// shared logic to one helper.

// ScenarioInputs groups the twelve values that every scenario operates on.
// Passing this value type replaces the long positional parameter list that
// each scenario used to declare.
type ScenarioInputs struct {
	A, B, C, D, E, F, G, H, I, J, K, L int
}

// values exposes the inputs in positional order so the shared helper can index
// into them without repeating the field list at each call site.
func (in ScenarioInputs) values() [12]int {
	return [12]int{in.A, in.B, in.C, in.D, in.E, in.F, in.G, in.H, in.I, in.J, in.K, in.L}
}

// pairSum is the single implementation shared by every scenario: it returns the
// sum of the two inputs at the given positions. Extracting it removes the
// near-duplicate bodies the scenarios used to carry.
func pairSum(in ScenarioInputs, first, second int) int {
	v := in.values()
	return v[first] + v[second]
}

func Scenario01(in ScenarioInputs) int { return pairSum(in, 0, 11) }
func Scenario02(in ScenarioInputs) int { return pairSum(in, 1, 10) }
func Scenario03(in ScenarioInputs) int { return pairSum(in, 2, 9) }
func Scenario04(in ScenarioInputs) int { return pairSum(in, 3, 8) }
func Scenario05(in ScenarioInputs) int { return pairSum(in, 4, 7) }
func Scenario06(in ScenarioInputs) int { return pairSum(in, 5, 6) }
func Scenario07(in ScenarioInputs) int { return pairSum(in, 6, 5) }
func Scenario08(in ScenarioInputs) int { return pairSum(in, 7, 4) }
func Scenario09(in ScenarioInputs) int { return pairSum(in, 8, 3) }
func Scenario10(in ScenarioInputs) int { return pairSum(in, 9, 2) }
func Scenario11(in ScenarioInputs) int { return pairSum(in, 10, 1) }
func Scenario12(in ScenarioInputs) int { return pairSum(in, 11, 0) }
