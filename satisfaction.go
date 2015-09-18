package satisfaction

type Possibility interface{}

type Constraint func(Possibility) bool

type Problem interface {
	Constraints() []Constraint
	Possibilities() []Possibility
}

type Solution []Possibility

type Solver interface {
	Solve(p Problem, solutions chan<- Solution, done <-chan struct{}) error
}
