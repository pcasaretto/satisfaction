package satisfaction_test

import (
	. "github.com/pcasaretto/satisfaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type LatinSquare struct {
	size int
}

type LatinSquareCell struct {
	row    int
	column int
	value  int
}

func NewLatinSquare(size int) *LatinSquare {
	return &LatinSquare{size}
}

func (l *LatinSquare) Constraints() []Constraint {
	constraints := []Constraint{}
	var f Constraint
	for row := 1; row <= l.size; row++ {
		for column := 1; column <= l.size; column++ {

			// loop variable is reused for each iteration
			row := row
			column := column

			// There must be a number in each cell
			f = func(c Possibility) bool {
				cell := c.(LatinSquareCell)
				return cell.column == column && cell.row == row
			}
			constraints = append(constraints, f)

			// The number 'row' must appear in row 'column'
			f = func(c Possibility) bool {
				cell := c.(LatinSquareCell)
				return cell.row == column && cell.value == row
			}
			constraints = append(constraints, f)

			// The number 'row' must appear in column 'column'
			f = func(c Possibility) bool {
				cell := c.(LatinSquareCell)
				return cell.column == column && cell.value == row
			}
			constraints = append(constraints, f)
		}
	}
	return constraints
}

func (l *LatinSquare) Possibilities() []Possibility {
	array := []Possibility{}
	for number := 1; number <= l.size; number++ {
		for row := 1; row <= l.size; row++ {
			for column := 1; column <= l.size; column++ {
				array = append(array, LatinSquareCell{value: number, row: row, column: column})
			}
		}
	}
	return array
}

var _ = Describe("AlgorithmX", func() {
	Describe("#Solve", func() {

		var (
			problem Problem
			solver  Solver
		)

		Context("for a size 1 Latin Square", func() {

			BeforeEach(func() {
				problem = NewLatinSquare(1)
				solver = AlgorithmX{}
			})

			It("finds the right solutions", func() {
				expected := LatinSquareCell{1, 1, 1}
				ch := make(chan Solution)
				solver.Solve(problem, ch, make(chan struct{}))
				Expect(<-ch).To(ConsistOf(expected))
			})
		})
	})

})