package satisfaction

type SudokuCell struct {
	row    int
	column int
	value  int
}

type Sudoku struct {
	Rank int
}

func (s *Sudoku) Constraints() []Constraint {
	constraints := []Constraint{}
	size := s.Rank * s.Rank
	var f Constraint

	for n := 1; n <= size; n++ {
		for i := 1; i <= s.Rank; i++ {
			for j := 1; j <= s.Rank; j++ {
				n := n
				i := i
				j := j
				// The number 'n' must appear in grid '(i,j)'
				f = func(c Possibility) bool {
					cell := c.(SudokuCell)
					rightY := (cell.row-1)/s.Rank+1 == i
					rightX := (cell.column-1)/s.Rank+1 == j
					rightValue := cell.value == n
					return rightX && rightY && rightValue
				}
				constraints = append(constraints, f)
			}
		}
	}
	for row := 1; row <= size; row++ {
		for column := 1; column <= size; column++ {

			row := row
			column := column

			// There must be a number in each cell
			f = func(c Possibility) bool {
				cell := c.(SudokuCell)
				return cell.column == column && cell.row == row
			}
			constraints = append(constraints, f)

			// The number 'row' must appear in row 'column'
			f = func(c Possibility) bool {
				cell := c.(SudokuCell)
				return cell.row == column && cell.value == row
			}
			constraints = append(constraints, f)

			// The number 'row' must appear in column 'column'
			f = func(c Possibility) bool {
				cell := c.(SudokuCell)
				return cell.column == column && cell.value == row
			}
			constraints = append(constraints, f)

		}
	}
	return constraints
}

func (s *Sudoku) Possibilities() []Possibility {
	array := []Possibility{}
	size := s.Rank * s.Rank
	for number := 1; number <= size; number++ {
		for row := 1; row <= size; row++ {
			for column := 1; column <= size; column++ {
				array = append(array, SudokuCell{value: number, row: row, column: column})
			}
		}
	}
	return array
}
