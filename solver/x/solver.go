package x

import "github.com/pcasaretto/satisfaction"

type Solver struct {
}

func (a Solver) Solve(p satisfaction.Problem, out chan<- satisfaction.Solution, done <-chan struct{}) error {
	cMatrix := newConstraintMatrix(p)

	solution := make(satisfaction.Solution, 0, cMatrix.Len())
	possibleSolutions := func(s satisfaction.Possibility) {
		solution = append(solution, s)
	}
	backtrack := func() {
		solution = solution[:len(solution)-1]
	}
	foundSolution := func() {
		b := make(satisfaction.Solution, len(solution))
		copy(b, solution)
		out <- b
	}

	go func() {
		cMatrix.solve(possibleSolutions, backtrack, foundSolution)
		close(out)
	}()
	return nil
}

type constraintMatrix struct {
	root *cell
}

func (c *constraintMatrix) Len() int {
	var i int
	for cell := c.root.left; cell != c.root; cell = cell.right {
		i++
	}
	return i
}

func (c *constraintMatrix) chooseUnsatisfiedConstraint() *cell {
	minCell := c.root.right
	if minCell == c.root {
		return nil
	}

	for i := c.root.right; i != c.root; i = i.right {
		if i.size < minCell.size {
			minCell = i
		}
	}

	return minCell
}

func newConstraintMatrix(problem satisfaction.Problem) *constraintMatrix {
	root := newCell(nil)

	for _, constraint := range problem.Constraints() {
		root.pushCellRight(newCell(constraint))
	}

	for _, possibility := range problem.Possibilities() {
		var lastCell *cell
		for header := root.right; header != root; header = header.right {
			constraint := header.value.(satisfaction.Constraint)
			if constraint(possibility) {
				c := newCell(possibility)
				c.header = header
				if lastCell != nil {
					lastCell.pushCellRight(c)
				}
				lastCell = c
				header.pushCellDown(c)
				header.size++
			}
		}
	}

	return &constraintMatrix{root}
}

var signal = struct{}{}

func (c *constraintMatrix) solve(out func(satisfaction.Possibility), btrack func(), found func()) {
	headerCell := c.chooseUnsatisfiedConstraint()
	if headerCell == nil {
		// Reached end of tree
		found()
		return
	}
	headerCell.cover()
	for cell := headerCell.down; cell != headerCell; cell = cell.down {
		out(cell.value)
		for neighbor := cell.right; neighbor != cell; neighbor = neighbor.right {
			neighbor.header.cover()
		}
		c.solve(out, btrack, found)
		btrack()
		for neighbor := cell.left; neighbor != cell; neighbor = neighbor.left {
			neighbor.header.uncover()
		}
	}
	headerCell.uncover()
}
