package x

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestNewCell(t *testing.T) {
	RegisterTestingT(t)
	cell := newCell(nil)
	Expect(cell.up).To(Equal(cell))
	Expect(cell.down).To(Equal(cell))
	Expect(cell.left).To(Equal(cell))
	Expect(cell.right).To(Equal(cell))

	i := 1
	cell = newCell(i)
	Expect(cell.value).To(BeEquivalentTo(i))
}

func TestPushCellDown(t *testing.T) {
	RegisterTestingT(t)
	var c, c2 *cell

	c = newCell(nil)
	c2 = newCell(nil)
	c.pushCellDown(c2)

	Expect(c.up).To(Equal(c2))
	Expect(c.down).To(Equal(c2))
	Expect(c2.up).To(Equal(c))
	Expect(c2.down).To(Equal(c))
}

func TestPushCellUp(t *testing.T) {
	RegisterTestingT(t)
	var c, c2 *cell

	c = newCell(nil)
	c2 = newCell(nil)
	c.pushCellUp(c2)

	Expect(c.up).To(Equal(c2))
	Expect(c.down).To(Equal(c2))
	Expect(c2.up).To(Equal(c))
	Expect(c2.down).To(Equal(c))
}

func TestPushCellLeft(t *testing.T) {
	RegisterTestingT(t)
	var c, c2 *cell

	c = newCell(nil)
	c2 = newCell(nil)
	c.pushCellLeft(c2)

	Expect(c.left).To(Equal(c2))
	Expect(c.right).To(Equal(c2))
	Expect(c2.left).To(Equal(c))
	Expect(c2.right).To(Equal(c))
}

func TestPushCellRight(t *testing.T) {
	RegisterTestingT(t)
	var c, c2 *cell

	c = newCell(nil)
	c2 = newCell(nil)
	c.pushCellRight(c2)

	Expect(c.left).To(Equal(c2))
	Expect(c.right).To(Equal(c2))
	Expect(c2.left).To(Equal(c))
	Expect(c2.right).To(Equal(c))
}

func RemoveAndRestoreVertically(t *testing.T) {
	RegisterTestingT(t)
	var (
		c, cellUp, cellDown *cell
	)

	c = newCell(nil)
	cellUp = newCell(nil)
	cellDown = newCell(nil)
	c.pushCellDown(cellDown)
	c.pushCellUp(cellUp)

	c.removeVertically()
	Expect(cellUp.down).To(Equal(cellDown))
	Expect(cellDown.up).To(Equal(cellUp))

	c.removeVertically()
	c.restoreVertically()
	Expect(cellUp.down).To(Equal(c))
	Expect(cellDown.up).To(Equal(c))
	Expect(c.up).To(Equal(cellUp))
	Expect(c.down).To(Equal(cellDown))
}

func RemoveAndRestoreHorizontally(t *testing.T) {
	RegisterTestingT(t)
	var (
		c, cellLeft, cellRight *cell
	)

	c = newCell(nil)
	cellLeft = newCell(nil)
	cellRight = newCell(nil)
	c.pushCellLeft(cellLeft)
	c.pushCellRight(cellRight)

	c.removeHorizontally()
	Expect(cellLeft.left).To(Equal(cellRight))
	Expect(cellRight.right).To(Equal(cellLeft))

	c.removeHorizontally()
	c.restoreHorizontally()
	Expect(cellRight.left).To(Equal(c))
	Expect(cellLeft.right).To(Equal(c))
	Expect(c.left).To(Equal(cellLeft))
	Expect(c.right).To(Equal(cellRight))
}
