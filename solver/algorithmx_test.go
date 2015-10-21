package solver_test

import (
	"github.com/pcasaretto/satisfaction"
	"github.com/pcasaretto/satisfaction/problem"
	. "github.com/pcasaretto/satisfaction/solver"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AlgorithmX", func() {
	Describe("#Solve", func() {

		var (
			p      satisfaction.Problem
			solver satisfaction.Solver
		)

		BeforeEach(func() {
			solver = AlgorithmX{}
		})

		Context("for a size 1 Latin Square", func() {

			BeforeEach(func() {
				p = problem.NewLatinSquare(1)
			})

			It("finds the right solutions", func() {
				expected := problem.LatinSquareCell{1, 1, 1}
				ch := make(chan satisfaction.Solution)
				solver.Solve(p, ch, make(chan struct{}))
				Expect(<-ch).To(ConsistOf(expected))
			})
		})

		Context("for a size 2 Latin Square", func() {

			BeforeEach(func() {
				p = problem.NewLatinSquare(2)
			})

			It("finds all solutions", func() {
				solutions := make([]satisfaction.Solution, 0, 2)
				ch := make(chan satisfaction.Solution)
				solver.Solve(p, ch, make(chan struct{}))
				for s := range ch {
					Expect(len(s)).To(Equal(4))
					solutions = append(solutions, s)
				}
				Expect(len(solutions)).To(Equal(2))
			})
		})

		Context("for a size 3 Latin Square", func() {

			BeforeEach(func() {
				p = problem.NewLatinSquare(3)
			})

			It("finds all solutions", func() {
				solutions := make([]satisfaction.Solution, 0, 12)
				ch := make(chan satisfaction.Solution)
				solver.Solve(p, ch, make(chan struct{}))
				for s := range ch {
					Expect(len(s)).To(Equal(9))
					solutions = append(solutions, s)
				}
				Expect(len(solutions)).To(Equal(12))
			})
		})

		Measure("it should do something hard efficiently", func(b Benchmarker) {
			ch := make(chan satisfaction.Solution)
			p = problem.NewLatinSquare(5)

			b.Time("first solution", func() {
				done := make(chan struct{})
				solver.Solve(p, ch, done)
				<-ch
				close(done)
			})

			solutions := 0
			b.Time("all solutions", func() {
				solver.Solve(p, ch, make(chan struct{}))
				for range ch {
					solutions++
				}
				Expect(solutions).To(Equal(161280))
			})
		}, 10)
	})

})
