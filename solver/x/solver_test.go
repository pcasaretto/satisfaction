package x_test

import (
	"testing"

	"github.com/pcasaretto/satisfaction"
	"github.com/pcasaretto/satisfaction/problem"
	"github.com/pcasaretto/satisfaction/solver/x"

	. "github.com/onsi/gomega"
)

func TestSolver(t *testing.T) {
	var (
		p      satisfaction.Problem
		solver satisfaction.Solver
	)

	solver = x.Solver{}

	p = problem.NewLatinSquare(1)

	expected := problem.LatinSquareCell{1, 1, 1}
	ch := make(chan satisfaction.Solution)
	solver.Solve(p, ch, make(chan struct{}))
	Expect(<-ch).To(ConsistOf(expected))

	p = problem.NewLatinSquare(2)

	solutions := make([]satisfaction.Solution, 0, 2)
	ch = make(chan satisfaction.Solution)
	solver.Solve(p, ch, make(chan struct{}))
	for s := range ch {
		Expect(len(s)).To(Equal(4))
		solutions = append(solutions, s)
	}
	Expect(len(solutions)).To(Equal(2))

	p = problem.NewLatinSquare(3)

	solutions = make([]satisfaction.Solution, 0, 12)
	ch = make(chan satisfaction.Solution)
	solver.Solve(p, ch, make(chan struct{}))
	for s := range ch {
		Expect(len(s)).To(Equal(9))
		solutions = append(solutions, s)
	}
	Expect(len(solutions)).To(Equal(12))
}

func BenchmarkLatinSquare1Solution(b *testing.B) {
	p := problem.NewLatinSquare(5)
	solver := x.Solver{}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		ch := make(chan satisfaction.Solution)
		done := make(chan struct{})
		solver.Solve(p, ch, done)
		<-ch
		close(done)
	}
}

func BenchmarkLatinSquareAllSolutions(b *testing.B) {
	p := problem.NewLatinSquare(5)
	solver := x.Solver{}
	ch := make(chan satisfaction.Solution)
	done := make(chan struct{})
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		solver.Solve(p, ch, make(chan struct{}))
		for range ch {
		}
		close(done)
	}
}
