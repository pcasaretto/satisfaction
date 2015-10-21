package solver_test

import (
	"fmt"
	"sort"

	. "github.com/pcasaretto/satisfaction"
	"github.com/pcasaretto/satisfaction/problem"
	. "github.com/pcasaretto/satisfaction/solver"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sudoku", func() {
	It("works", func() {
		p := &problem.Sudoku{Rank: 3}
		done := make(chan struct{})
		solver := AlgorithmX{}
		ch := make(chan Solution)
		solver.Solve(p, ch, done)
		Expect(len(<-ch)).To(Equal(81))
		close(done)
	})

	It("works with Givens", func() {
		p := &problem.Sudoku{Rank: 3}
		p.Givens = []problem.SudokuCell{{1, 1, 1}}
		done := make(chan struct{})
		solver := AlgorithmX{}
		ch := make(chan Solution)
		solver.Solve(p, ch, done)
		for _, c := range <-ch {
			p.Solution = append(p.Solution, c.(problem.SudokuCell))
		}
		fmt.Println(p)
		Expect(len(<-ch)).To(Equal(80))
		close(done)
	})

	FIt("works with my diabolical puzzle", func() {
		p := &problem.Sudoku{Rank: 3}
		p.Givens = []problem.SudokuCell{
			{1, 1, 1},
			{1, 7, 5},
			{1, 9, 3},
			{2, 5, 2},
			{2, 7, 1},
			{3, 3, 3},
			{3, 6, 6},
			{3, 9, 8},
			{4, 1, 4},
			{4, 3, 2},
			{4, 4, 8},
			{4, 6, 3},
			{5, 2, 6},
			{5, 9, 2},
			{6, 1, 7},
			{6, 8, 9},
			{6, 9, 5},
			{7, 2, 5},
			{7, 3, 6},
			{8, 3, 4},
			{8, 8, 3},
			{8, 9, 7},
			{9, 1, 2},
			{9, 6, 8},
			{9, 7, 6},
		}
		done := make(chan struct{})
		solver := AlgorithmX{}
		ch := make(chan Solution)
		solver.Solve(p, ch, done)
		var solutions []Solution
		for solution := range ch {
			solutions = append(solutions, solution)
		}
		Expect(len(solutions)).To(Equal(1))
		var cells problem.SudokuCells
		for _, c := range solutions[0] {
			cells = append(cells, c.(problem.SudokuCell))
		}
		sort.Sort(cells)
		expected := problem.SudokuCells{
			{1, 2, 2},
			{1, 3, 9},
			{1, 4, 4},
			{1, 5, 8},
			{1, 6, 7},
			{1, 8, 6},
			{2, 1, 6},
			{2, 2, 4},
			{2, 3, 8},
			{2, 4, 3},
			{2, 6, 5},
			{2, 8, 7},
			{2, 9, 9},
			{3, 1, 5},
			{3, 2, 7},
			{3, 4, 1},
			{3, 5, 9},
			{3, 7, 4},
			{3, 8, 2},
			{4, 2, 9},
			{4, 5, 5},
			{4, 7, 7},
			{4, 8, 1},
			{4, 9, 6},
			{5, 1, 8},
			{5, 3, 5},
			{5, 4, 7},
			{5, 5, 1},
			{5, 6, 9},
			{5, 7, 3},
			{5, 8, 4},
			{6, 2, 3},
			{6, 3, 1},
			{6, 4, 6},
			{6, 5, 4},
			{6, 6, 2},
			{6, 7, 8},
			{7, 1, 3},
			{7, 4, 2},
			{7, 5, 7},
			{7, 6, 4},
			{7, 7, 9},
			{7, 8, 8},
			{7, 9, 1},
			{8, 1, 9},
			{8, 2, 8},
			{8, 4, 5},
			{8, 5, 6},
			{8, 6, 1},
			{8, 7, 2},
			{9, 2, 1},
			{9, 3, 7},
			{9, 4, 9},
			{9, 5, 3},
			{9, 8, 5},
			{9, 9, 4},
		}
		Expect(cells).To(Equal(expected))
	})
})
