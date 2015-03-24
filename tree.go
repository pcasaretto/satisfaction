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

type Possibility interface{}

type HeaderCell struct {
	*BasicCell
	size int
}

func NewHeaderCell(v interface{}) *HeaderCell {
	return &HeaderCell{NewCell(v), 0}
}

func (cell *HeaderCell) PushCellRight(c Cell) {
	toAdd := c.(*HeaderCell)
	cell.BasicCell.PushCellRight(toAdd)
	cell.size++
}

type ConstraintMatrix struct {
	root *HeaderCell
}

func (c *ConstraintMatrix) FindSolution() []Possibility {

	return []Possibility{}
}

func (c *ConstraintMatrix) String() string {
	var b bytes.Buffer
	var i int
	for _, header := range c.root.CellsRight() {
		for _, cell := range header.CellsDown() {
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
		for _, header := range root.CellsRight() {
			constraint := header.Value().(Constraint)
			if constraint(possibility) {
				newCell := NewCell(possibility)
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
