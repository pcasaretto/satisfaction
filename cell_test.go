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
	Convey("Cell#addCellDown", t, func() {
		cell := NewCell()
		cell2 := NewCell()
		Convey("should set up the connections properly", func() {
			cell.addCellDown(cell2)
			So(cell.down, ShouldEqual, cell2)
			So(cell.up, ShouldEqual, cell2)
			So(cell2.down, ShouldEqual, cell)
			So(cell2.up, ShouldEqual, cell)
		})
	})
}

func TestAddCellUp(t *testing.T) {
	Convey("Cell#addCellUp", t, func() {
		cell := NewCell()
		cell2 := NewCell()
		Convey("should set up the connections properly", func() {
			cell.addCellUp(cell2)
			So(cell.down, ShouldEqual, cell2)
			So(cell.up, ShouldEqual, cell2)
			So(cell2.down, ShouldEqual, cell)
			So(cell2.up, ShouldEqual, cell)
		})
	})
}

func TestAddCellLeft(t *testing.T) {
	Convey("Cell#addCellLeft", t, func() {
		cell := NewCell()
		cell2 := NewCell()
		Convey("should set up the connections properly", func() {
			cell.addCellLeft(cell2)
			So(cell.left, ShouldEqual, cell2)
			So(cell.right, ShouldEqual, cell2)
			So(cell2.left, ShouldEqual, cell)
			So(cell2.right, ShouldEqual, cell)
		})
	})
}

func TestAddCellRight(t *testing.T) {
	Convey("Cell#addCellRight", t, func() {
		cell := NewCell()
		cell2 := NewCell()
		Convey("should set up the connections properly", func() {
			cell.addCellRight(cell2)
			So(cell.left, ShouldEqual, cell2)
			So(cell.right, ShouldEqual, cell2)
			So(cell2.left, ShouldEqual, cell)
			So(cell2.right, ShouldEqual, cell)
		})
	})
}

func TestCellsDown(t *testing.T) {
	Convey("Cell#cellsDown", t, func() {
		cell := NewCell()

		Convey("When there are no cells down", func() {
			Convey("Returns an empty slice", func() {
				So(cell.cellsDown(), ShouldBeEmpty)
			})
		})

		Convey("When there are some cells down", func() {
			Convey("Returns the cells ordered", func() {
				cell2 := NewCell()
				cell3 := NewCell()
				cell.addCellDown(cell3)
				cell.addCellDown(cell2)
				cells := cell.cellsDown()
				So(len(cells), ShouldEqual, 2)
				So(cells[0], ShouldEqual, cell2)
				So(cells[1], ShouldEqual, cell3)
			})
		})
	})
}
