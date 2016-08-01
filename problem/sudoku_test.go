package problem_test

import (
	"testing"

	"github.com/pcasaretto/satisfaction"
	"github.com/pcasaretto/satisfaction/problem"
)

func TestSudokuConstraints(t *testing.T) {
	s := problem.NewSudoku(2, nil)

	constraints := s.Constraints()
	if actual, expected := len(constraints), 64; actual != expected {
		t.Errorf("\n\texpected: %d\n\tgot: %d", expected, actual)
	}

	s = problem.NewSudoku(3, nil)

	constraints = s.Constraints()
	if actual, expected := len(constraints), 324; actual != expected {
		t.Errorf("\n\texpected: %d\n\tgot: %d", expected, actual)
	}
}

func TestSudokuValid(t *testing.T) {
	var solution satisfaction.Solution

	p := problem.NewSudoku(2, nil)

	solution = satisfaction.Solution{
		problem.SudokuCell{1, 1, 1},
	}
	if actual, expected := p.Valid(solution), false; actual != expected {
		t.Errorf("\n\texpected: %t\n\tgot: %t", expected, actual)
	}

	solution = satisfaction.Solution{
		problem.SudokuCell{1, 1, 1},
		problem.SudokuCell{1, 1, 1},
		problem.SudokuCell{1, 1, 1},
		problem.SudokuCell{1, 1, 1},
	}
	if actual, expected := p.Valid(solution), false; actual != expected {
		t.Errorf("\n\texpected: %t\n\tgot: %t", expected, actual)
	}

	solution = satisfaction.Solution{
		problem.SudokuCell{1, 1, 1},
		problem.SudokuCell{1, 2, 2},
		problem.SudokuCell{1, 3, 3},
		problem.SudokuCell{1, 4, 4},
		problem.SudokuCell{2, 1, 1},
		problem.SudokuCell{2, 2, 2},
		problem.SudokuCell{2, 3, 3},
		problem.SudokuCell{2, 4, 4},
		problem.SudokuCell{3, 1, 1},
		problem.SudokuCell{3, 2, 2},
		problem.SudokuCell{3, 3, 3},
		problem.SudokuCell{3, 4, 4},
		problem.SudokuCell{4, 1, 1},
		problem.SudokuCell{4, 2, 2},
		problem.SudokuCell{4, 3, 3},
		problem.SudokuCell{4, 4, 4},
	}
	if actual, expected := p.Valid(solution), false; actual != expected {
		t.Errorf("\n\texpected: %t\n\tgot: %t", expected, actual)
	}

	solution = satisfaction.Solution{
		problem.SudokuCell{2, 4, 4},
		problem.SudokuCell{4, 3, 4},
		problem.SudokuCell{3, 1, 4},
		problem.SudokuCell{1, 2, 4},
		problem.SudokuCell{4, 4, 2},
		problem.SudokuCell{3, 2, 2},
		problem.SudokuCell{4, 2, 1},
		problem.SudokuCell{4, 1, 3},
		problem.SudokuCell{2, 2, 3},
		problem.SudokuCell{3, 4, 3},
		problem.SudokuCell{1, 3, 3},
		problem.SudokuCell{3, 3, 1},
		problem.SudokuCell{2, 3, 2},
		problem.SudokuCell{1, 1, 2},
		problem.SudokuCell{2, 1, 1},
		problem.SudokuCell{1, 4, 1},
	}
	if actual, expected := p.Valid(solution), true; actual != expected {
		t.Errorf("\n\texpected: %t\n\tgot: %t", expected, actual)
	}
	solution = satisfaction.Solution{
		problem.SudokuCell{2, 4, 4},
		problem.SudokuCell{4, 3, 4},
		problem.SudokuCell{3, 1, 4},
		problem.SudokuCell{1, 2, 4},
		problem.SudokuCell{4, 4, 1},
		problem.SudokuCell{3, 2, 1},
		problem.SudokuCell{4, 2, 2},
		problem.SudokuCell{4, 1, 3},
		problem.SudokuCell{2, 2, 3},
		problem.SudokuCell{1, 4, 3},
		problem.SudokuCell{3, 4, 2},
		problem.SudokuCell{3, 3, 3},
		problem.SudokuCell{1, 3, 2},
		problem.SudokuCell{2, 3, 1},
		problem.SudokuCell{2, 1, 2},
		problem.SudokuCell{1, 1, 1},
	}
	if actual, expected := p.Valid(solution), true; actual != expected {
		t.Errorf("\n\texpected: %t\n\tgot: %t", expected, actual)
	}

	solution = satisfaction.Solution{
		problem.SudokuCell{4, 4, 4},
		problem.SudokuCell{2, 3, 4},
		problem.SudokuCell{3, 1, 4},
		problem.SudokuCell{1, 2, 4},
		problem.SudokuCell{4, 3, 2},
		problem.SudokuCell{3, 2, 2},
		problem.SudokuCell{4, 2, 3},
		problem.SudokuCell{4, 1, 1},
		problem.SudokuCell{2, 2, 1},
		problem.SudokuCell{3, 4, 1},
		problem.SudokuCell{3, 3, 3},
		problem.SudokuCell{1, 3, 1},
		problem.SudokuCell{2, 4, 3},
		problem.SudokuCell{1, 1, 3},
		problem.SudokuCell{1, 4, 2},
		problem.SudokuCell{2, 1, 2},
	}
	if actual, expected := p.Valid(solution), true; actual != expected {
		t.Errorf("\n\texpected: %t\n\tgot: %t", expected, actual)
	}

	solution = satisfaction.Solution{
		problem.SudokuCell{1, 1, 3},
		problem.SudokuCell{1, 2, 1},
		problem.SudokuCell{1, 3, 2},
		problem.SudokuCell{1, 4, 4},
		problem.SudokuCell{2, 1, 4},
		problem.SudokuCell{2, 2, 2},
		problem.SudokuCell{2, 3, 3},
		problem.SudokuCell{2, 4, 1},
		problem.SudokuCell{3, 1, 1},
		problem.SudokuCell{3, 2, 3},
		problem.SudokuCell{3, 3, 4},
		problem.SudokuCell{3, 4, 2},
		problem.SudokuCell{4, 1, 2},
		problem.SudokuCell{4, 2, 4},
		problem.SudokuCell{4, 3, 1},
		problem.SudokuCell{4, 4, 3},
	}
	if actual, expected := p.Valid(solution), true; actual != expected {
		t.Errorf("\n\texpected: %t\n\tgot: %t", expected, actual)
	}

}
