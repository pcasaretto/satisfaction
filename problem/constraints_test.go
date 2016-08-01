package problem

import "testing"

func TestSudokuGridConstraint(t *testing.T) {
	rank := 2
	sudoku := NewSudoku(rank, nil)
	for i := 1; i <= rank; i++ {
		for j := 1; j <= rank; j++ {
			c := sudoku.gridConstraint(1, i, j)

			xoffset := rank * (i - 1)
			yoffset := rank * (j - 1)
			for k := 1; k <= rank; k++ {
				for l := 1; l <= rank; l++ {
					p := SudokuCell{Value: 1, Row: xoffset + k, Column: yoffset + l}
					if !c(p) {
						t.Errorf("Cell inside %d, %d quadrant failed constraint:\n\t%+v", i, j, p)
					}
				}
			}

			for k := 1; k <= rank; k++ {
				for l := 1; l <= rank; l++ {
					if k == i && l == j {
						continue
					}
					for m := 0; m < rank; m++ {
						p := SudokuCell{Value: 1, Row: k + m + rank*(k-1), Column: l + rank*(l-1)}
						if c(p) {
							t.Errorf("Cell outside %d, %d quadrant satisfied constraint:\n\t%+v", i, j, p)
						}
						p = SudokuCell{Value: 1, Row: k + rank*(k-1), Column: l + m + rank*(l-1)}
						if c(p) {
							t.Errorf("Cell outside %d, %d quadrant satisfied constraint:\n\t%+v", i, j, p)
						}
					}
				}
			}

		}
	}
}
