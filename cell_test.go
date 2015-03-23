package dancinglinks_test

import (
	. "github.com/pcasaretto/dancinglinks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cell", func() {
	Describe(".NewCell", func() {
		It("creates a cell pointing to itself in the four directions", func() {
			cell := NewCell(nil)
			Expect(cell.Up()).To(Equal(cell))
			Expect(cell.Down()).To(Equal(cell))
			Expect(cell.Left()).To(Equal(cell))
			Expect(cell.Right()).To(Equal(cell))
		})

		It("sets the given value to Value", func() {
			i := 1
			cell := NewCell(i)
			Expect(cell.Value).To(Equal(i))
		})
	})

	Describe("#PushCellDown", func() {

		var (
			cell  *Cell
			cell2 *Cell
		)

		BeforeEach(func() {
			cell = NewCell(nil)
			cell2 = NewCell(nil)
			cell.PushCellDown(cell2)
		})

		It("should set up the connections properly", func() {
			Expect(cell.Up()).To(Equal(cell2))
			Expect(cell.Down()).To(Equal(cell2))
			Expect(cell2.Up()).To(Equal(cell))
			Expect(cell2.Down()).To(Equal(cell))
		})
	})

	Describe("#PushCellUp", func() {

		var (
			cell  *Cell
			cell2 *Cell
		)

		BeforeEach(func() {
			cell = NewCell(nil)
			cell2 = NewCell(nil)
			cell.PushCellUp(cell2)
		})

		It("should set up the connections properly", func() {
			Expect(cell.Up()).To(Equal(cell2))
			Expect(cell.Down()).To(Equal(cell2))
			Expect(cell2.Up()).To(Equal(cell))
			Expect(cell2.Down()).To(Equal(cell))
		})
	})

	Describe("#PushCellLeft", func() {

		var (
			cell  *Cell
			cell2 *Cell
		)

		BeforeEach(func() {
			cell = NewCell(nil)
			cell2 = NewCell(nil)
			cell.PushCellLeft(cell2)
		})

		It("should set up the connections properly", func() {
			Expect(cell.Left()).To(Equal(cell2))
			Expect(cell.Right()).To(Equal(cell2))
			Expect(cell2.Left()).To(Equal(cell))
			Expect(cell2.Right()).To(Equal(cell))
		})
	})

	Describe("#PushCellRight", func() {

		var (
			cell  *Cell
			cell2 *Cell
		)

		BeforeEach(func() {
			cell = NewCell(nil)
			cell2 = NewCell(nil)
			cell.PushCellRight(cell2)
		})

		It("should set up the connections properly", func() {
			Expect(cell.Left()).To(Equal(cell2))
			Expect(cell.Right()).To(Equal(cell2))
			Expect(cell2.Left()).To(Equal(cell))
			Expect(cell2.Right()).To(Equal(cell))
		})
	})

	Describe("removing and restoring vertically", func() {
		var (
			cell, cellUp, cellDown *Cell
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
			Expect(cellUp.Down()).To(Equal(cellDown))
			Expect(cellDown.Up()).To(Equal(cellUp))
		})

		It("Cell#RestoreVertically restores the cell to the vertical line", func() {
			cell.RemoveVertically()
			cell.RestoreVertically()
			Expect(cellUp.Down()).To(Equal(cell))
			Expect(cellDown.Up()).To(Equal(cell))
			Expect(cell.Up()).To(Equal(cellUp))
			Expect(cell.Down()).To(Equal(cellDown))
		})
	})

	Describe("removing and restoring vertically", func() {
		var (
			cell, cellLeft, cellRight *Cell
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
			Expect(cellLeft.Left()).To(Equal(cellRight))
			Expect(cellRight.Right()).To(Equal(cellLeft))
		})

		It("Cell#RestoreHorizontally restores the cell to the horizontal line", func() {
			cell.RemoveHorizontally()
			cell.RestoreHorizontally()
			Expect(cellRight.Left()).To(Equal(cell))
			Expect(cellLeft.Right()).To(Equal(cell))
			Expect(cell.Left()).To(Equal(cellLeft))
			Expect(cell.Right()).To(Equal(cellRight))
		})
	})
})
