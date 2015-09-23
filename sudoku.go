package satisfaction

type SudokuCell struct {
	Row    int
	Column int
	Value  int
}

type Sudoku struct {
	Rank   int
	Givens []SudokuCell
}

func (s *Sudoku) gridConstraint(n, i, j int) func(Possibility) bool {
	return func(c Possibility) bool {
		cell := c.(SudokuCell)
		rightY := (cell.Row-1)/s.Rank+1 == i
		rightX := (cell.Column-1)/s.Rank+1 == j
		rightValue := cell.Value == n
		return rightX && rightY && rightValue
	}
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
				constraints = append(constraints, s.gridConstraint(n, i, j))
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
				return cell.Column == column && cell.Row == row
			}
			constraints = append(constraints, f)

			// The number 'row' must appear in row 'column'
			f = func(c Possibility) bool {
				cell := c.(SudokuCell)
				return cell.Row == column && cell.Value == row
			}
			constraints = append(constraints, f)

			// The number 'row' must appear in column 'column'
			f = func(c Possibility) bool {
				cell := c.(SudokuCell)
				return cell.Column == column && cell.Value == row
			}
			constraints = append(constraints, f)

		}
	}
	//eliminate satisfied constraints
	var finalConstraints []Constraint
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

func (s *Sudoku) Possibilities() []Possibility {
	array := []Possibility{}
	size := s.Rank * s.Rank
	for number := 1; number <= size; number++ {
		for row := 1; row <= size; row++ {
			for column := 1; column <= size; column++ {
				array = append(array, SudokuCell{Value: number, Row: row, Column: column})
			}
		}
	}
	//eliminate givens
	var possiblities []Possibility
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
