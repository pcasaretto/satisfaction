package problem

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"

	"github.com/pcasaretto/satisfaction"
)

type SudokuCell struct {
	Row    int
	Column int
	Value  int
}

type SudokuCells []SudokuCell

func (a SudokuCells) Len() int      { return len(a) }
func (a SudokuCells) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SudokuCells) Less(i, j int) bool {
	return a[i].Row < a[j].Row || (a[i].Row == a[j].Row && a[i].Column < a[j].Column)
}

type Sudoku struct {
	rank          int
	givens        SudokuCells
	constraints   []satisfaction.Constraint
	possibilities []satisfaction.Possibility
}

func NewSudoku(rank int, givens SudokuCells) *Sudoku {
	if givens == nil {
		givens = make(SudokuCells, 0)
	}
	return &Sudoku{rank: rank, givens: givens}
}

func (s *Sudoku) CSVDump(solution SudokuCells, w io.Writer) {
	cw := csv.NewWriter(w)
	complete := append(s.givens, solution...)
	sort.Sort(complete)

	t := func(in SudokuCells) []string {
		s := make([]string, len(in))
		for i, c := range in {
			s[i] = fmt.Sprintf("%d", c.Value)
		}
		return s
	}

	size := s.rank * s.rank
	for i := 0; i < size; i++ {
		cw.Write(t(complete[i*size : (i+1)*size]))
	}
	cw.Flush()
}

func (s *Sudoku) Valid(solution satisfaction.Solution) bool {
	if len(solution)+len(s.givens) != s.rank*s.rank*s.rank*s.rank {
		return false
	}
	constraints := s.Constraints()

	var k int
	for _, c := range constraints {
		for _, p := range solution {
			if c(p) {
				k++
				break
			}
		}
	}
	return k == len(constraints)
}

func (s *Sudoku) gridConstraint(n, i, j int) func(satisfaction.Possibility) bool {
	return func(c satisfaction.Possibility) bool {
		cell := c.(SudokuCell)
		rightY := (cell.Row-1)/s.rank+1 == i
		rightX := (cell.Column-1)/s.rank+1 == j
		rightValue := cell.Value == n
		return rightX && rightY && rightValue
	}
}

func (s *Sudoku) cellConstraint(column, row int) satisfaction.Constraint {
	return satisfaction.Constraint(func(c satisfaction.Possibility) bool {
		cell := c.(SudokuCell)
		return cell.Column == column && cell.Row == row
	})
}

func (s *Sudoku) rowConstraint(index, value int) satisfaction.Constraint {
	return satisfaction.Constraint(func(c satisfaction.Possibility) bool {
		cell := c.(SudokuCell)
		return cell.Row == index && cell.Value == value
	})
}

func (s *Sudoku) columnConstraint(index, value int) satisfaction.Constraint {
	return satisfaction.Constraint(func(c satisfaction.Possibility) bool {
		cell := c.(SudokuCell)
		return cell.Column == index && cell.Value == value
	})
}

func (s *Sudoku) Constraints() []satisfaction.Constraint {
	if s.constraints != nil {
		return s.constraints
	}
	size := s.rank * s.rank

	for n := 1; n <= size; n++ {
		for i := 1; i <= s.rank; i++ {
			for j := 1; j <= s.rank; j++ {
				n := n
				i := i
				j := j
				// The number 'n' must appear in grid '(i,j)'
				s.constraints = append(s.constraints, s.gridConstraint(n, i, j))
			}
		}
	}

	generators := []func(int, int) satisfaction.Constraint{
		s.cellConstraint,   // There must be a number in each cell
		s.rowConstraint,    // The number 'row' must appear in row 'column'
		s.columnConstraint, // The number 'row' must appear in column 'column'
	}
	for row := 1; row <= size; row++ {
		for column := 1; column <= size; column++ {

			row := row
			column := column

			for _, g := range generators {
				s.constraints = append(s.constraints, g(column, row))
			}

		}
	}
	//eliminate satisfied constraints
	var finalConstraints []satisfaction.Constraint
	for _, constraint := range s.constraints {
		noneSatisfy := true
		for _, given := range s.givens {
			if constraint(given) {
				noneSatisfy = false
				break
			}
		}
		if noneSatisfy {
			finalConstraints = append(finalConstraints, constraint)
		}
	}
	s.constraints = finalConstraints
	return s.constraints
}

func (s *Sudoku) Possibilities() []satisfaction.Possibility {
	array := []satisfaction.Possibility{}
	size := s.rank * s.rank
	for number := 1; number <= size; number++ {
		for row := 1; row <= size; row++ {
			for column := 1; column <= size; column++ {
				array = append(array, SudokuCell{Value: number, Row: row, Column: column})
			}
		}
	}
	//eliminate givens
	var possiblities []satisfaction.Possibility
	for _, cell := range array {
		cell := cell.(SudokuCell)
		isGiven := false
		for _, given := range s.givens {
			if cell.Row == given.Row && cell.Column == given.Column {
				isGiven = true
				break
			}
		}
		if !isGiven {
			possiblities = append(possiblities, cell)
		}
	}
	return possiblities
}
