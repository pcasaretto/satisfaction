package dancinglinks

import (
	"bytes"
	"fmt"
)

type ExactCoverProblem interface {
	Constraints() []Constraint
	Possibilities() []Possibility
}

type Constraint func(Possibility) bool

type Thing struct {
	header *HeaderCell
	value  Possibility
}

type Possibility interface{}

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

type ConstraintMatrix struct {
	root *HeaderCell
}

func (c *ConstraintMatrix) Len() int {
	var i int
	for cell := c.root.Left(); cell != c.root; cell = cell.Right() {
		i++
	}
	return i
}

func (c *ConstraintMatrix) chooseUnsatisfiedConstraint() *HeaderCell {
	cell := c.root.Right().(*HeaderCell)
	if cell != c.root {
		return cell
	}
	return nil
}

func (c *ConstraintMatrix) solve(out chan Possibility, btrack chan bool, found chan bool) {
	headerCell := c.chooseUnsatisfiedConstraint()
	if headerCell == nil {
		// Reached end of tree, backtrack and find next solution
		found <- true
		return
	}
	headerCell.cover()
	for cell := headerCell.Down(); cell != headerCell; cell = cell.Down() {
		thing := cell.Value().(Thing)
		out <- thing.value
		for neighbor := cell.Right(); neighbor != cell; neighbor = neighbor.Right() {
			thing = neighbor.Value().(Thing)
			thing.header.cover()
		}
		c.solve(out, btrack, found)
		btrack <- true
		for neighbor := cell.Right(); neighbor != cell; neighbor = neighbor.Right() {
			thing = neighbor.Value().(Thing)
			thing.header.uncover()
		}
		headerCell.uncover()
	}
}

func (c *ConstraintMatrix) FindSolution() []Possibility {

	solution_channel := make(chan Possibility)
	backtrack := make(chan bool)
	found_solution := make(chan bool)

	solutions := make([]Possibility, 0, c.Len())
	go c.solve(solution_channel, backtrack, found_solution)
loop:
	for {
		select {
		case s := <-solution_channel:
			solutions = append(solutions, s)
		case <-backtrack:
			solutions = solutions[:len(solutions)-1]
		case <-found_solution:
			break loop
		}
	}
	return solutions
}

func (c *ConstraintMatrix) String() string {
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

func NewConstraintMatrix(problem ExactCoverProblem) *ConstraintMatrix {
	root := NewHeaderCell(nil)

	for _, constraint := range problem.Constraints() {
		root.PushCellRight(NewHeaderCell(constraint))
	}

	for _, possibility := range problem.Possibilities() {
		var lastCell Cell
		for header := root.Right(); header != root; header = header.Right() {
			constraint := header.Value().(Constraint)
			if constraint(possibility) {
				thing := Thing{
					header: header.(*HeaderCell),
					value:  possibility,
				}
				newCell := NewCell(thing)
				if lastCell != nil {
					lastCell.PushCellRight(newCell)
				}
				lastCell = newCell
				header.PushCellDown(newCell)
			}
		}
	}

	return &ConstraintMatrix{root}
}
