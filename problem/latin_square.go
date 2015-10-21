package problem

import "github.com/pcasaretto/satisfaction"

type LatinSquare struct {
	size int
}

type LatinSquareCell struct {
	Row    int
	Column int
	Value  int
}

func NewLatinSquare(size int) *LatinSquare {
	return &LatinSquare{size}
}

func (l *LatinSquare) Constraints() []satisfaction.Constraint {
	constraints := []satisfaction.Constraint{}
	var f satisfaction.Constraint
	for row := 1; row <= l.size; row++ {
		for column := 1; column <= l.size; column++ {

			// loop variable is reused for each iteration
			row := row
			column := column

			// There must be a number in each cell
			f = func(c satisfaction.Possibility) bool {
				cell := c.(LatinSquareCell)
				return cell.Column == column && cell.Row == row
			}
			constraints = append(constraints, f)

			// The number 'row' must appear in row 'column'
			f = func(c satisfaction.Possibility) bool {
				cell := c.(LatinSquareCell)
				return cell.Row == column && cell.Value == row
			}
			constraints = append(constraints, f)

			// The number 'row' must appear in column 'column'
			f = func(c satisfaction.Possibility) bool {
				cell := c.(LatinSquareCell)
				return cell.Column == column && cell.Value == row
			}
			constraints = append(constraints, f)
		}
	}
	return constraints
}

func (l *LatinSquare) Possibilities() []satisfaction.Possibility {
	array := []satisfaction.Possibility{}
	for number := 1; number <= l.size; number++ {
		for row := 1; row <= l.size; row++ {
			for column := 1; column <= l.size; column++ {
				array = append(array, LatinSquareCell{Value: number, Row: row, Column: column})
			}
		}
	}
	return array
}
