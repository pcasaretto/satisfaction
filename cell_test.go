package dancinglinks

import (
  "testing"
  . "github.com/smartystreets/goconvey/convey"
)

func TestNewCell(t *testing.T) {
  Convey("A new cell should point to itself in the four directions", t, func(){
    cell := NewCell()
    So(cell.up, ShouldEqual, cell)
    So(cell.down, ShouldEqual, cell)
    So(cell.left, ShouldEqual, cell)
    So(cell.right, ShouldEqual, cell)
  })
}
