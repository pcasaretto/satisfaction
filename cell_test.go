package dancinglinks_test

import (
	. "github.com/pcasaretto/dancinglinks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)


var _ = Describe("Cell", func() {
	format.UseStringerRepresentation = true
	Describe(".NewCell", func() {
		It("creates a cell pointing to itself in the four directions", func() {
			cell := NewCell(nil)
			Expect(cell.Up()).To(BeEquivalentTo(cell))
			Expect(cell.Down()).To(BeEquivalentTo(cell))
			Expect(cell.Left()).To(BeEquivalentTo(cell))
			Expect(cell.Right()).To(BeEquivalentTo(cell))
		})

		It("sets the given value to Value", func() {
			i := 1
			cell := NewCell(i)
			Expect(cell.Value()).To(BeEquivalentTo(i))
		})
	})

	Describe("#PushCellDown", func() {

		var (
			cell  *BasicCell
			cell2 *BasicCell
		)

		BeforeEach(func() {
			cell = NewCell(nil)
			cell2 = NewCell(nil)
			cell.PushCellDown(cell2)
		})

		It("should set up the connections properly", func() {
			Expect(cell.Up()).To(BeEquivalentTo(cell2))
			Expect(cell.Down()).To(BeEquivalentTo(cell2))
			Expect(cell2.Up()).To(BeEquivalentTo(cell))
			Expect(cell2.Down()).To(BeEquivalentTo(cell))
		})
	})

	Describe("#PushCellUp", func() {

		var (
			cell  *BasicCell
			cell2 *BasicCell
		)

		BeforeEach(func() {
			cell = NewCell(nil)
			cell2 = NewCell(nil)
			cell.PushCellUp(cell2)
		})

		It("should set up the connections properly", func() {
			Expect(cell.Up()).To(BeEquivalentTo(cell2))
			Expect(cell.Down()).To(BeEquivalentTo(cell2))
			Expect(cell2.Up()).To(BeEquivalentTo(cell))
			Expect(cell2.Down()).To(BeEquivalentTo(cell))
		})
	})

	Describe("#PushCellLeft", func() {

		var (
			cell  *BasicCell
			cell2 *BasicCell
		)

		BeforeEach(func() {
			cell = NewCell(nil)
			cell2 = NewCell(nil)
			cell.PushCellLeft(cell2)
		})

		It("should set up the connections properly", func() {
			Expect(cell.Left()).To(BeEquivalentTo(cell2))
			Expect(cell.Right()).To(BeEquivalentTo(cell2))
			Expect(cell2.Left()).To(BeEquivalentTo(cell))
			Expect(cell2.Right()).To(BeEquivalentTo(cell))
		})
	})

	Describe("#PushCellRight", func() {

		var (
			cell  *BasicCell
			cell2 *BasicCell
		)

		BeforeEach(func() {
			cell = NewCell(nil)
			cell2 = NewCell(nil)
			cell.PushCellRight(cell2)
		})

		It("should set up the connections properly", func() {
			Expect(cell.Left()).To(BeEquivalentTo(cell2))
			Expect(cell.Right()).To(BeEquivalentTo(cell2))
			Expect(cell2.Left()).To(BeEquivalentTo(cell))
			Expect(cell2.Right()).To(BeEquivalentTo(cell))
		})
	})

	Describe("removing and restoring vertically", func() {
		var (
			cell, cellUp, cellDown *BasicCell
		)

		BeforeEach(func() {
			cell = NewCell(nil)
			cellUp = NewCell(nil)
			cellDown = NewCell(nil)
			cell.PushCellDown(cellDown)
			cell.PushCellUp(cellUp)
		})

		It("#RemoveVertically removes the cell from the vertical line", func() {
			cell.RemoveVertically()
			Expect(cellUp.Down()).To(BeEquivalentTo(cellDown))
			Expect(cellDown.Up()).To(BeEquivalentTo(cellUp))
		})

		It("Cell#RestoreVertically restores the cell to the vertical line", func() {
			cell.RemoveVertically()
			cell.RestoreVertically()
			Expect(cellUp.Down()).To(BeEquivalentTo(cell))
			Expect(cellDown.Up()).To(BeEquivalentTo(cell))
			Expect(cell.Up()).To(BeEquivalentTo(cellUp))
			Expect(cell.Down()).To(BeEquivalentTo(cellDown))
		})
	})

	Describe("removing and restoring vertically", func() {
		var (
			cell, cellLeft, cellRight *BasicCell
		)

		BeforeEach(func() {
			cell = NewCell(nil)
			cellLeft = NewCell(nil)
			cellRight = NewCell(nil)
			cell.PushCellLeft(cellLeft)
			cell.PushCellRight(cellRight)
		})

		It("#RemoveHorizontally removes the cell from the horizontal line", func() {
			cell.RemoveHorizontally()
			Expect(cellLeft.Left()).To(BeEquivalentTo(cellRight))
			Expect(cellRight.Right()).To(BeEquivalentTo(cellLeft))
		})

		It("Cell#RestoreHorizontally restores the cell to the horizontal line", func() {
			cell.RemoveHorizontally()
			cell.RestoreHorizontally()
			Expect(cellRight.Left()).To(BeEquivalentTo(cell))
			Expect(cellLeft.Right()).To(BeEquivalentTo(cell))
			Expect(cell.Left()).To(BeEquivalentTo(cellLeft))
			Expect(cell.Right()).To(BeEquivalentTo(cellRight))
		})
	})
})
