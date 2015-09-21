package satisfaction

import (
	"bytes"
	"fmt"
)

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

func (c *constraintMatrix) String() string {
	var b bytes.Buffer
	var k int
	outerRound := false
	for i := c.root.right; ; i = i.down {
		if i == c.root.right {
			if outerRound {
				break
			} else {
				outerRound = true
			}
		}
		round := false
		for j := i; ; j = j.right {
			if j == i {
				if round {
					break
				} else {
					round = true
				}
			}
			b.WriteString(fmt.Sprint(j))
			b.WriteString(" | ")
			k++
		}
		b.WriteString(" ∑ ")
	}
	b.WriteString(fmt.Sprintf("\n Size: %d", k))
	return b.String()
}

func newConstraintMatrix(problem Problem) *constraintMatrix {
	root := NewCell(nil)

	for _, constraint := range problem.Constraints() {
		root.PushCellRight(NewCell(constraint))
	}

	for _, possibility := range problem.Possibilities() {
		var lastCell *cell
		for header := root.right; header != root; header = header.right {
			constraint := header.value.(Constraint)
			if constraint(possibility) {
				newCell := NewCell(possibility)
				newCell.header = header
				if lastCell != nil {
					lastCell.PushCellRight(newCell)
				}
				lastCell = newCell
				header.PushCellDown(newCell)
				header.size++
			}
		}
	}

	return &constraintMatrix{root}
}

func (c *constraintMatrix) solve(out chan Possibility, btrack chan struct{}, found chan struct{}) {
	headerCell := c.chooseUnsatisfiedConstraint()
	if headerCell == nil {
		// Reached end of tree
		found <- struct{}{}
		return
	}
	headerCell.cover()
	for cell := headerCell.down; cell != headerCell; cell = cell.down {
		out <- cell.value
		for neighbor := cell.right; neighbor != cell; neighbor = neighbor.right {
			neighbor.header.cover()
		}
		c.solve(out, btrack, found)
		btrack <- struct{}{}
		for neighbor := cell.left; neighbor != cell; neighbor = neighbor.left {
			neighbor.header.uncover()
		}
	}
	headerCell.uncover()
}

type AlgorithmX struct {
}

func (a AlgorithmX) Solve(p Problem, out chan<- Solution, done <-chan struct{}) error {
	cMatrix := newConstraintMatrix(p)

	possibleSolutions := make(chan Possibility)
	backtrack := make(chan struct{})
	foundSolution := make(chan struct{})
	allFound := make(chan struct{})

	solution := make(Solution, 0, cMatrix.Len())
	go func() {
		cMatrix.solve(possibleSolutions, backtrack, foundSolution)
		close(allFound)
	}()
	go func() {
	loop:
		for {
			select {
			case s := <-possibleSolutions:
				solution = append(solution, s)
			case <-backtrack:
				solution = solution[:len(solution)-1]
			case <-foundSolution:
				out <- solution
			case <-allFound:
				close(out)
				break loop
			case <-done:
				break loop
			}
		}
	}()

	return nil
}
