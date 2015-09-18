package satisfaction

type Possibility interface{}

type Constraint func(Possibility) bool

type Problem interface {
	Constraints() []Constraint
	Possibilities() []Possibility
}

type Solution []Possibility

type SolveOptions struct {
	Limit int
}

type Solver interface {
	Solve(Problem, SolveOptions) ([]Solution, error)
}
