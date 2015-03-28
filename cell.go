package dancinglinks

import "fmt"

var _id int = 0

type Cell interface {
	Up() Cell
	Down() Cell
	Left() Cell
	Right() Cell
	PushCellDown(toAdd Cell)
	PushCellUp(toAdd Cell)
	PushCellLeft(toAdd Cell)
	PushCellRight(toAdd Cell)
	RemoveVertically()
	RestoreVertically()
	RemoveHorizontally()
	RestoreHorizontally()
	Value() interface{}
	setDown(toAdd Cell)
	setUp(toAdd Cell)
	setLeft(toAdd Cell)
	setRight(toAdd Cell)
}

type BasicCell struct {
	left  Cell
	right Cell
	up    Cell
	down  Cell
	value interface{}
	id    int
}

func NewCell(v interface{}) *BasicCell {
	cell := &BasicCell{value: v}
	cell.down = cell
	cell.up = cell
	cell.left = cell
	cell.right = cell
	_id++
	cell.id = _id
	return cell
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

func (cell *BasicCell) String() string {
	return fmt.Sprintf("Cell id %d", cell.id)
}

func (cell *BasicCell) Value() interface{} {
	return cell.value
}

func (cell *BasicCell) Up() Cell {
	return cell.up
}

func (cell *BasicCell) Down() Cell {
	return cell.down
}

func (cell *BasicCell) Left() Cell {
	return cell.left
}

func (cell *BasicCell) Right() Cell {
	return cell.right
}

func (cell *BasicCell) PushCellDown(c Cell) {
	cell.down.setUp(c)
	c.setDown(cell.down)
	cell.down = c
	c.setUp(cell)
}

func (cell *BasicCell) PushCellUp(c Cell) {
	cell.up.setDown(c)
	c.setUp(cell.up)
	cell.up = c
	c.setDown(cell)
}

func (cell *BasicCell) PushCellLeft(c Cell) {
	cell.left.setRight(c)
	c.setLeft(cell.left)
	cell.left = c
	c.setRight(cell)
}

func (cell *BasicCell) PushCellRight(c Cell) {
	cell.right.setLeft(c)
	c.setRight(cell.right)
	cell.right = c
	c.setLeft(cell)
}

func (c *BasicCell) setDown(toAdd Cell) {
	c.down = toAdd
}
func (c *BasicCell) setUp(toAdd Cell) {
	c.up = toAdd
}
func (c *BasicCell) setLeft(toAdd Cell) {
	c.left = toAdd
}
func (c *BasicCell) setRight(toAdd Cell) {
	c.right = toAdd
}

func (cell *BasicCell) RemoveVertically() {
	cell.up.setDown(cell.down)
	cell.down.setUp(cell.up)
}

func (cell *BasicCell) RestoreVertically() {
	cell.up.setDown(cell)
	cell.down.setUp(cell)
}

func (cell *BasicCell) RemoveHorizontally() {
	cell.right.setLeft(cell.left)
	cell.left.setRight(cell.right)
}

func (cell *BasicCell) RestoreHorizontally() {
	cell.right.setLeft(cell)
	cell.left.setRight(cell)
}
