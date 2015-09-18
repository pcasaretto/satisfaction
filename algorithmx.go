package satisfaction

import (
	"bytes"
	"fmt"
)

type container struct {
	header *HeaderCell
	value  Possibility
}

type HeaderCell struct {
	*BasicCell
	size int
}

func NewHeaderCell(v interface{}) *HeaderCell {
	header := &HeaderCell{}
	cell := NewCell(v)
	cell.up = header
	cell.down = header
	cell.left = header
	cell.right = header
	header.BasicCell = cell
	return header
}

func (cell *HeaderCell) cover() {
	cell.RemoveHorizontally()
	for c1 := cell.Down(); c1 != cell; c1 = c1.Down() {
		for c2 := c1.Right(); c2 != c1; c2 = c2.Right() {
			c2.RemoveVertically()
		}
	}
}

func (cell *HeaderCell) uncover() {
	for c1 := cell.Up(); c1 != cell; c1 = c1.Up() {
		for c2 := c1.Left(); c2 != c1; c2 = c2.Left() {
			c2.RestoreVertically()
		}
	}
	cell.RestoreHorizontally()
}

func (cell *HeaderCell) PushCellRight(c Cell) {
	cell.BasicCell.PushCellRight(c)
	cell.size++
}

type constraintMatrix struct {
	root *HeaderCell
}

func (c *constraintMatrix) Len() int {
	var i int
	for cell := c.root.Left(); cell != c.root; cell = cell.Right() {
		i++
	}
	return i
}

func (c *constraintMatrix) chooseUnsatisfiedConstraint() *HeaderCell {
	cell := c.root.Right().(*HeaderCell)
	if cell != c.root {
		return cell
	}
	return nil
}

func (c *constraintMatrix) solve(out chan Possibility, btrack chan bool, found chan bool) {
	headerCell := c.chooseUnsatisfiedConstraint()
	if headerCell == nil {
		// Reached end of tree, backtrack and find next solution
		found <- true
		return
	}
	headerCell.cover()
	for cell := headerCell.Down(); cell != headerCell; cell = cell.Down() {
		t := cell.Value().(container)
		out <- t.value
		for neighbor := cell.Right(); neighbor != cell; neighbor = neighbor.Right() {
			t = neighbor.Value().(container)
			t.header.cover()
		}
		c.solve(out, btrack, found)
		btrack <- true
		for neighbor := cell.Right(); neighbor != cell; neighbor = neighbor.Right() {
			t = neighbor.Value().(container)
			t.header.uncover()
		}
		headerCell.uncover()
	}
}

func (c *constraintMatrix) String() string {
	var b bytes.Buffer
	var i int
	for header := c.root.Right(); header != c.root; header = header.Right() {
		for cell := header.Down(); cell != header; cell = cell.Down() {
			b.WriteString(cell.(*BasicCell).String())
			b.WriteString(" | ")
			i++
		}
	}
	b.WriteString(fmt.Sprintf("\n Size: %d", i))
	return b.String()
}

func newConstraintMatrix(problem Problem) *constraintMatrix {
	root := NewHeaderCell(nil)

	for _, constraint := range problem.Constraints() {
		root.PushCellRight(NewHeaderCell(constraint))
	}

	for _, possibility := range problem.Possibilities() {
		var lastCell Cell
		for header := root.Right(); header != root; header = header.Right() {
			constraint := header.Value().(Constraint)
			if constraint(possibility) {
				container := container{
					header: header.(*HeaderCell),
					value:  possibility,
				}
				newCell := NewCell(container)
				if lastCell != nil {
					lastCell.PushCellRight(newCell)
				}
				lastCell = newCell
				header.PushCellDown(newCell)
			}
		}
	}

	return &constraintMatrix{root}
}

type AlgorithmX struct {
}

func (a AlgorithmX) Solve(p Problem, out chan<- Solution, done <-chan struct{}) error {
	cMatrix := newConstraintMatrix(p)

	possibleSolutions := make(chan Possibility)
	backtrack := make(chan bool)
	foundSolution := make(chan bool)

	solution := make(Solution, 0, cMatrix.Len())
	go cMatrix.solve(possibleSolutions, backtrack, foundSolution)
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
			case <-done:
				break loop
			}
		}
	}()

	return nil
}
