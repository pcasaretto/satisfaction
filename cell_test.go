package satisfaction

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

var _ = Describe("Cell", func() {
	format.UseStringerRepresentation = true
	Describe(".NewCell", func() {
		It("creates a cell pointing to itself in the four directions", func() {
			cell := NewCell(nil)
			Expect(cell.up).To(Equal(cell))
			Expect(cell.down).To(Equal(cell))
			Expect(cell.left).To(Equal(cell))
			Expect(cell.right).To(Equal(cell))
		})

		It("sets the given value to Value", func() {
			i := 1
			cell := NewCell(i)
			Expect(cell.value).To(BeEquivalentTo(i))
		})
	})

	Describe("#PushCellDown", func() {

		var (
			c, c2 *cell
		)

		BeforeEach(func() {
			c = NewCell(nil)
			c2 = NewCell(nil)
			c.PushCellDown(c2)
		})

		It("should set up the connections properly", func() {
			Expect(c.up).To(Equal(c2))
			Expect(c.down).To(Equal(c2))
			Expect(c2.up).To(Equal(c))
			Expect(c2.down).To(Equal(c))
		})
	})

	Describe("#PushCellUp", func() {

		var (
			c, c2 *cell
		)

		BeforeEach(func() {
			c = NewCell(nil)
			c2 = NewCell(nil)
			c.PushCellUp(c2)
		})

		It("should set up the connections properly", func() {
			Expect(c.up).To(Equal(c2))
			Expect(c.down).To(Equal(c2))
			Expect(c2.up).To(Equal(c))
			Expect(c2.down).To(Equal(c))
		})
	})

	Describe("#PushCellLeft", func() {

		var (
			c, c2 *cell
		)

		BeforeEach(func() {
			c = NewCell(nil)
			c2 = NewCell(nil)
			c.PushCellLeft(c2)
		})

		It("should set up the connections properly", func() {
			Expect(c.left).To(Equal(c2))
			Expect(c.right).To(Equal(c2))
			Expect(c2.left).To(Equal(c))
			Expect(c2.right).To(Equal(c))
		})
	})

	Describe("#PushCellRight", func() {

		var (
			c, c2 *cell
		)

		BeforeEach(func() {
			c = NewCell(nil)
			c2 = NewCell(nil)
			c.PushCellRight(c2)
		})

		It("should set up the connections properly", func() {
			Expect(c.left).To(Equal(c2))
			Expect(c.right).To(Equal(c2))
			Expect(c2.left).To(Equal(c))
			Expect(c2.right).To(Equal(c))
		})
	})

	Describe("removing and restoring vertically", func() {
		var (
			c, cellUp, cellDown *cell
		)

		BeforeEach(func() {
			c = NewCell(nil)
			cellUp = NewCell(nil)
			cellDown = NewCell(nil)
			c.PushCellDown(cellDown)
			c.PushCellUp(cellUp)
		})

		It("#RemoveVertically removes the cell from the vertical line", func() {
			c.RemoveVertically()
			Expect(cellUp.down).To(Equal(cellDown))
			Expect(cellDown.up).To(Equal(cellUp))
		})

		It("Cell#RestoreVertically restores the cell to the vertical line", func() {
			c.RemoveVertically()
			c.RestoreVertically()
			Expect(cellUp.down).To(Equal(c))
			Expect(cellDown.up).To(Equal(c))
			Expect(c.up).To(Equal(cellUp))
			Expect(c.down).To(Equal(cellDown))
		})
	})

	Describe("removing and restoring horizontally", func() {
		var (
			c, cellLeft, cellRight *cell
		)

		BeforeEach(func() {
			c = NewCell(nil)
			cellLeft = NewCell(nil)
			cellRight = NewCell(nil)
			c.PushCellLeft(cellLeft)
			c.PushCellRight(cellRight)
		})

		It("#RemoveHorizontally removes the cell from the horizontal line", func() {
			c.RemoveHorizontally()
			Expect(cellLeft.left).To(Equal(cellRight))
			Expect(cellRight.right).To(Equal(cellLeft))
		})

		It("Cell#RestoreHorizontally restores the cell to the horizontal line", func() {
			c.RemoveHorizontally()
			c.RestoreHorizontally()
			Expect(cellRight.left).To(Equal(c))
			Expect(cellLeft.right).To(Equal(c))
			Expect(c.left).To(Equal(cellLeft))
			Expect(c.right).To(Equal(cellRight))
		})
	})
})
