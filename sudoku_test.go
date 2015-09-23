package satisfaction_test

import (
	. "github.com/pcasaretto/satisfaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("Sudoku", func() {
	It("works", func() {
		problem := &Sudoku{Rank: 3}
		done := make(chan struct{})
		solver := AlgorithmX{}
		ch := make(chan Solution)
		solver.Solve(problem, ch, done)
		Expect(len(<-ch)).To(Equal(81))
		close(done)
	})

	It("works with Givens", func() {
		problem := &Sudoku{Rank: 3}
		problem.Givens = []SudokuCell{{1, 1, 1}}
		done := make(chan struct{})
		solver := AlgorithmX{}
		ch := make(chan Solution)
		solver.Solve(problem, ch, done)
		Expect(len(<-ch)).To(Equal(80))
		close(done)
	})
})
