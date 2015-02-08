package dancinglinks

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewCell(t *testing.T) {
	Convey("A new cell should point to itself in the four directions", t, func() {
		cell := NewCell()
		So(cell.up, ShouldEqual, cell)
		So(cell.down, ShouldEqual, cell)
		So(cell.left, ShouldEqual, cell)
		So(cell.right, ShouldEqual, cell)
	})
}

func TestAddCellDown(t *testing.T) {
	Convey("Cell#AddCellDown", t, func() {
		cell := NewCell()
		cell2 := NewCell()
		Convey("should set up the connections properly", func() {
			cell.AddCellDown(cell2)
			So(cell.down, ShouldEqual, cell2)
			So(cell.up, ShouldEqual, cell2)
			So(cell2.down, ShouldEqual, cell)
			So(cell2.up, ShouldEqual, cell)
		})
	})
}

func TestAddCellUp(t *testing.T) {
	Convey("Cell#AddCellUp", t, func() {
		cell := NewCell()
		cell2 := NewCell()
		Convey("should set up the connections properly", func() {
			cell.AddCellUp(cell2)
			So(cell.down, ShouldEqual, cell2)
			So(cell.up, ShouldEqual, cell2)
			So(cell2.down, ShouldEqual, cell)
			So(cell2.up, ShouldEqual, cell)
		})
	})
}

func TestAddCellLeft(t *testing.T) {

	Convey("Cell#AddCellLeft", t, func() {

		cell := NewCell()
		cell2 := NewCell()

		Convey("should set up the connections properly", func() {
			cell.AddCellLeft(cell2)
			So(cell.left, ShouldEqual, cell2)
			So(cell.right, ShouldEqual, cell2)
			So(cell2.left, ShouldEqual, cell)
			So(cell2.right, ShouldEqual, cell)
		})

	})

}

func TestAddCellRight(t *testing.T) {
	Convey("Cell#AddCellRight", t, func() {
		cell := NewCell()
		cell2 := NewCell()
		Convey("should set up the connections properly", func() {
			cell.AddCellRight(cell2)
			So(cell.left, ShouldEqual, cell2)
			So(cell.right, ShouldEqual, cell2)
			So(cell2.left, ShouldEqual, cell)
			So(cell2.right, ShouldEqual, cell)
		})
	})
}

func TestCellsDown(t *testing.T) {
	Convey("Cell#CellsDown", t, func() {
		cell := NewCell()

		Convey("When there are no cells down", func() {
			Convey("Returns an empty slice", func() {
				So(cell.CellsDown(), ShouldBeEmpty)
			})
		})

		Convey("When there are some cells down", func() {
			Convey("Returns the cells ordered", func() {
				cell2 := NewCell()
				cell3 := NewCell()
				cell.AddCellDown(cell3)
				cell.AddCellDown(cell2)
				cells := cell.CellsDown()
				So(len(cells), ShouldEqual, 2)
				So(cells[0], ShouldEqual, cell2)
				So(cells[1], ShouldEqual, cell3)
			})
		})
	})
}

func TestRemoveAndRestoreVertically(t *testing.T) {
	Convey("", t, func() {
		cell := NewCell()
		cellUp := NewCell()
		cellDown := NewCell()
		cell.AddCellDown(cellDown)
		cell.AddCellUp(cellUp)
		cell.RemoveVertically()
		Convey("Cell#RemoveVertically removes the cell from the horizontal line", func() {
			So(cellUp.down, ShouldEqual, cellDown)
			So(cellDown.up, ShouldEqual, cellUp)
		})
		cell.RestoreVertically()
		Convey("Cell#RestoreVertically restores the cell to the horizontal line", func() {
			So(cellUp.down, ShouldEqual, cell)
			So(cell.up, ShouldEqual, cellUp)
			So(cellDown.up, ShouldEqual, cell)
			So(cell.down, ShouldEqual, cellDown)
		})
	})
}

func TestRemoveAndRestoreHorizontally(t *testing.T) {
	Convey("", t, func() {
		cell := NewCell()
		cellLeft := NewCell()
		cellRight := NewCell()
		cell.AddCellRight(cellRight)
		cell.AddCellLeft(cellLeft)
		cell.RemoveHorizontally()
		Convey("Cell#RemoveHorizontally removes the cell from the horizontal line", func() {
			So(cellLeft.right, ShouldEqual, cellRight)
			So(cellRight.left, ShouldEqual, cellLeft)
		})
		cell.RestoreHorizontally()
		Convey("Cell#RestoreHorizontally restores the cell to the horizontal line", func() {
			So(cellLeft.right, ShouldEqual, cell)
			So(cell.left, ShouldEqual, cellLeft)
			So(cellRight.left, ShouldEqual, cell)
			So(cell.right, ShouldEqual, cellRight)
		})
	})
}
