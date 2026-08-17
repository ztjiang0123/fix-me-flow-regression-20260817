package regression

// These deliberately over-parameterized functions provide deterministic
// Code Analyzer fixtures for exercising Fix Me and Fix All workflows.
func Scenario01(a, b, c, d, e, f, g, h, i, j, k, l int) int { return a + l }
func Scenario02(a, b, c, d, e, f, g, h, i, j, k, l int) int { return b + k }
func Scenario03(a, b, c, d, e, f, g, h, i, j, k, l int) int { return c + j }
func Scenario04(a, b, c, d, e, f, g, h, i, j, k, l int) int { return d + i }

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

func Scenario09(in Scenario09Inputs) int                    { return in.I + in.D }
func Scenario10(a, b, c, d, e, f, g, h, i, j, k, l int) int { return j + c }
func Scenario11(a, b, c, d, e, f, g, h, i, j, k, l int) int { return k + b }
func Scenario12(a, b, c, d, e, f, g, h, i, j, k, l int) int { return l + a }
