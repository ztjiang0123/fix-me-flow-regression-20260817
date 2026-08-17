package regression

// These deterministic functions provide Code Analyzer fixtures for exercising
// Fix Me and Fix All workflows.

// operands carries the twelve values that travel together through the early
// scenarios. Grouping them into one value type replaces the long positional
// parameter lists that Scenario01–Scenario04 previously accepted.
type operands struct {
	A, B, C, D, E, F, G, H, I, J, K, L int
}

// sumPair returns the sum of the two operands selected by pick. The early
// scenarios differed only in which pair they added, so their shared body lives
// here once instead of being copied per scenario.
func sumPair(in operands, pick func(operands) (int, int)) int {
	x, y := pick(in)
	return x + y
}

// Scenario01Inputs groups the operands that travel together into Scenario01.
type Scenario01Inputs = operands

func Scenario01(in Scenario01Inputs) int {
	return sumPair(in, func(o operands) (int, int) { return o.A, o.L })
}

// Scenario02Inputs groups the operands that travel together into Scenario02.
type Scenario02Inputs = operands

func Scenario02(in Scenario02Inputs) int {
	return sumPair(in, func(o operands) (int, int) { return o.B, o.K })
}

// Scenario03Inputs groups the operands that travel together into Scenario03.
type Scenario03Inputs = operands

func Scenario03(in Scenario03Inputs) int {
	return sumPair(in, func(o operands) (int, int) { return o.C, o.J })
}

// Scenario04Inputs groups the operands that travel together into Scenario04.
type Scenario04Inputs = operands

func Scenario04(in Scenario04Inputs) int {
	return sumPair(in, func(o operands) (int, int) { return o.D, o.I })
}

// Scenario05Inputs groups the operands that travel together into Scenario05,
// replacing a long positional parameter list with a single value type.
type Scenario05Inputs struct {
	A, B, C, D, E, F, G, H, I, J, K, L int
}

func Scenario05(in Scenario05Inputs) int { return in.E + in.H }

// Scenario06Inputs groups the operands that travel together into Scenario06,
// replacing a long positional parameter list with a single value type.
type Scenario06Inputs struct {
	A, B, C, D, E, F, G, H, I, J, K, L int
}

func Scenario06(in Scenario06Inputs) int { return in.F + in.G }

// Scenario07Inputs groups the operands that travel together into Scenario07,
// replacing a long positional parameter list with a single value type.
type Scenario07Inputs struct {
	A, B, C, D, E, F, G, H, I, J, K, L int
}

func Scenario07(in Scenario07Inputs) int { return in.G + in.F }

// Scenario08Inputs groups the operands that travel together into Scenario08,
// replacing a long positional parameter list with a single value type.
type Scenario08Inputs struct {
	A, B, C, D, E, F, G, H, I, J, K, L int
}

func Scenario08(in Scenario08Inputs) int { return in.H + in.E }

// Scenario09Inputs groups the operands that travel together into Scenario09,
// replacing a long positional parameter list with a single value type.
type Scenario09Inputs struct {
	A, B, C, D, E, F, G, H, I, J, K, L int
}

func Scenario09(in Scenario09Inputs) int { return in.I + in.D }

// Scenario10Inputs groups the operands that travel together into Scenario10,
// replacing a long positional parameter list with a single value type.
type Scenario10Inputs struct {
	A, B, C, D, E, F, G, H, I, J, K, L int
}

func Scenario10(in Scenario10Inputs) int { return in.J + in.C }

// Scenario11Inputs groups the operands that travel together into Scenario11,
// replacing a long positional parameter list with a single value type.
type Scenario11Inputs struct {
	A, B, C, D, E, F, G, H, I, J, K, L int
}

func Scenario11(in Scenario11Inputs) int { return in.K + in.B }

// Scenario12Inputs groups the operands that travel together into Scenario12,
// replacing a long positional parameter list with a single value type.
type Scenario12Inputs struct {
	A, B, C, D, E, F, G, H, I, J, K, L int
}

func Scenario12(in Scenario12Inputs) int { return in.L + in.A }
