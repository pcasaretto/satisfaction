package satisfaction

// A Possibility inside the universe of possibilities
type Possibility interface{}

// A Constraint takes a Possibility and returns true
// if it is satisfied
type Constraint func(Possibility) bool

// A Problem consists of a number of Constraints and a universe of
// Possibilities
type Problem interface {
	Constraints() []Constraint
	Possibilities() []Possibility
}

// A Solution is subset of the Problems possibilities
type Solution []Possibility

// A Solver takes a problem, a chan to send Solutions on and a done channel
// that means the caller is no longer interested in any more solutions.
type Solver interface {
	Solve(p Problem, solutions chan<- Solution, done <-chan struct{}) error
}
