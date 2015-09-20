package satisfaction

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type container struct {
	header *cell
	value  Possibility
}

var logger = log.New(os.Stdout, "logger: ", log.Lshortfile)
var matrix *constraintMatrix

type constraintMatrix struct {
	root *cell
}

func (c *constraintMatrix) Len() int {
	var i int
	for cell := c.root.left; cell != c.root; cell = cell.right {
		i++
	}
	return i
}

func (c *constraintMatrix) chooseUnsatisfiedConstraint() *cell {
	cell := c.root.right
	if cell != c.root {
		return cell
	}
	return nil
}

func (c *constraintMatrix) String() string {
	var b bytes.Buffer
	var k int
	outerRound := false
	for i := c.root.right; ; i = i.down {
		if i == c.root.right {
			if outerRound {
				break
			} else {
				outerRound = true
			}
		}
		round := false
		for j := i; ; j = j.right {
			if j == i {
				if round {
					break
				} else {
					round = true
				}
			}
			b.WriteString(fmt.Sprint(j))
			b.WriteString(" | ")
			k++
		}
		b.WriteString("\n")
	}
	b.WriteString(fmt.Sprintf("\n Size: %d", k))
	return b.String()
}

func newConstraintMatrix(problem Problem) *constraintMatrix {
	root := NewCell(nil)

	for _, constraint := range problem.Constraints() {
		root.PushCellRight(NewCell(constraint))
	}

	for _, possibility := range problem.Possibilities() {
		var lastCell *cell
		for header := root.right; header != root; header = header.right {
			constraint := header.value.(Constraint)
			if constraint(possibility) {
				container := container{
					header: header,
					value:  possibility,
				}
				newCell := NewCell(container)
				if lastCell != nil {
					lastCell.PushCellRight(newCell)
				}
				lastCell = newCell
				header.PushCellDown(newCell)
			}
		}
	}

	return &constraintMatrix{root}
}

func (c *constraintMatrix) solve(out chan Possibility, btrack chan struct{}, found chan struct{}) {
	fmt.Println("recursing")
	headerCell := c.chooseUnsatisfiedConstraint()
	if headerCell == nil {
		// Reached end of tree
		found <- struct{}{}
		return
	}
	headerCell.cover()
	for cell := headerCell.down; cell != headerCell; cell = cell.down {
		fmt.Println("125")
		t := cell.value.(container)
		out <- t.value
		for neighbor := cell.right; neighbor != cell; neighbor = neighbor.right {
			fmt.Println("129")
			t = neighbor.value.(container)
			t.header.cover()
		}
		c.solve(out, btrack, found)
		btrack <- struct{}{}
		for neighbor := cell.left; neighbor != cell; neighbor = neighbor.left {
			fmt.Println("135")
			t = neighbor.value.(container)
			t.header.uncover()
		}
		headerCell.uncover()
	}
	fmt.Println("141")
}

type AlgorithmX struct {
}

func (a AlgorithmX) Solve(p Problem, out chan<- Solution, done <-chan struct{}) error {
	cMatrix := newConstraintMatrix(p)
	matrix = cMatrix

	possibleSolutions := make(chan Possibility)
	backtrack := make(chan struct{})
	foundSolution := make(chan struct{})
	allFound := make(chan struct{})

	solution := make(Solution, 0, cMatrix.Len())
	go func() {
		cMatrix.solve(possibleSolutions, backtrack, foundSolution)
		close(allFound)
	}()
	go func() {
	loop:
		for {
			select {
			case s := <-possibleSolutions:
				fmt.Println(s)
				solution = append(solution, s)
			case <-backtrack:
				fmt.Println("backtracking")
				solution = solution[:len(solution)-1]
			case <-foundSolution:
				fmt.Println("found a solution")
				out <- solution
			case <-allFound:
				fmt.Println("found all solutions")
				break loop
			case <-done:
				break loop
			}
		}
	}()

	return nil
}
