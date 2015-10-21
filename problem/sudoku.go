package problem

import (
	"bytes"
	"fmt"
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
	Rank     int
	Givens   SudokuCells
	Solution SudokuCells
}

func (s *Sudoku) String() string {
	all_cells := append(s.Givens, s.Solution...)
	sort.Sort(all_cells)
	var b bytes.Buffer
	last_index := s.Rank * s.Rank
	for _, cell := range all_cells {
		if cell.Column == last_index {
			b.Write([]byte(fmt.Sprintf(" %d \n", cell.Value)))
		} else {
			b.Write([]byte(fmt.Sprintf(" %d |", cell.Value)))
		}
	}
	return string(b.Bytes())
}

func (s *Sudoku) gridConstraint(n, i, j int) func(satisfaction.Possibility) bool {
	return func(c satisfaction.Possibility) bool {
		cell := c.(SudokuCell)
		rightY := (cell.Row-1)/s.Rank+1 == i
		rightX := (cell.Column-1)/s.Rank+1 == j
		rightValue := cell.Value == n
		return rightX && rightY && rightValue
	}
}

func (s *Sudoku) Constraints() []satisfaction.Constraint {
	constraints := []satisfaction.Constraint{}
	size := s.Rank * s.Rank
	var f satisfaction.Constraint

	for n := 1; n <= size; n++ {
		for i := 1; i <= s.Rank; i++ {
			for j := 1; j <= s.Rank; j++ {
				n := n
				i := i
				j := j
				// The number 'n' must appear in grid '(i,j)'
				constraints = append(constraints, s.gridConstraint(n, i, j))
			}
		}
	}
	for row := 1; row <= size; row++ {
		for column := 1; column <= size; column++ {

			row := row
			column := column

			// There must be a number in each cell
			f = func(c satisfaction.Possibility) bool {
				cell := c.(SudokuCell)
				return cell.Column == column && cell.Row == row
			}
			constraints = append(constraints, f)

			// The number 'row' must appear in row 'column'
			f = func(c satisfaction.Possibility) bool {
				cell := c.(SudokuCell)
				return cell.Row == column && cell.Value == row
			}
			constraints = append(constraints, f)

			// The number 'row' must appear in column 'column'
			f = func(c satisfaction.Possibility) bool {
				cell := c.(SudokuCell)
				return cell.Column == column && cell.Value == row
			}
			constraints = append(constraints, f)

		}
	}
	//eliminate satisfied constraints
	var finalConstraints []satisfaction.Constraint
	for _, constraint := range constraints {
		noneSatisfy := true
		for _, given := range s.Givens {
			if constraint(given) {
				noneSatisfy = false
				break
			}
		}
		if noneSatisfy {
			finalConstraints = append(finalConstraints, constraint)
		}
	}
	return finalConstraints
}

func (s *Sudoku) Possibilities() []satisfaction.Possibility {
	array := []satisfaction.Possibility{}
	size := s.Rank * s.Rank
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
		for _, given := range s.Givens {
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
