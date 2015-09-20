package satisfaction

import "fmt"

var _id uint64 = 0

type CellInterface interface {
	Up() cell
	Down() cell
	Left() cell
	Right() cell
	PushCellDown(toAdd cell)
	PushCellUp(toAdd cell)
	PushCellLeft(toAdd cell)
	PushCellRight(toAdd cell)
	RemoveVertically()
	RestoreVertically()
	RemoveHorizontally()
	RestoreHorizontally()
	Value() interface{}
	setDown(toAdd cell)
	setUp(toAdd cell)
	setLeft(toAdd cell)
	setRight(toAdd cell)
}

type cell struct {
	left  *cell
	right *cell
	up    *cell
	down  *cell
	value interface{}
	size  uint
	id    uint64
}

func NewCell(v interface{}) *cell {
	cell := &cell{value: v}
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

func (c *cell) String() string {
	return fmt.Sprintf("Cell id %d", c.id)
}

func (c *cell) Value() interface{} {
	return c.value
}

func (c *cell) PushCellDown(other *cell) {
	c.down.setUp(other)
	other.setDown(c.down)
	c.down = other
	other.setUp(c)
}

func (c *cell) PushCellUp(other *cell) {
	c.up.setDown(other)
	other.setUp(c.up)
	c.up = other
	other.setDown(c)
}

func (c *cell) PushCellLeft(other *cell) {
	c.left.setRight(other)
	other.setLeft(c.left)
	c.left = other
	other.setRight(c)
}

func (c *cell) PushCellRight(other *cell) {
	c.right.setLeft(other)
	other.setRight(c.right)
	c.right = other
	other.setLeft(c)
}

func (c *cell) setDown(toAdd *cell) {
	c.down = toAdd
}
func (c *cell) setUp(toAdd *cell) {
	c.up = toAdd
}
func (c *cell) setLeft(toAdd *cell) {
	c.left = toAdd
}
func (c *cell) setRight(toAdd *cell) {
	c.right = toAdd
}

func (c *cell) RemoveVertically() {
	c.up.setDown(c.down)
	c.down.setUp(c.up)
}

func (c *cell) RestoreVertically() {
	c.up.setDown(c)
	c.down.setUp(c)
}

func (c *cell) RemoveHorizontally() {
	c.right.setLeft(c.left)
	c.left.setRight(c.right)
}

func (c *cell) RestoreHorizontally() {
	c.right.setLeft(c)
	c.left.setRight(c)
}

func (c *cell) cover() {
	c.RemoveHorizontally()
	for i := c.down; i != c; i = i.down {
		for j := i.right; j != i; j = j.right {
			logger.Println(c, i, j)
			logger.Println(matrix)
			j.RemoveVertically()
		}
	}
}

func (c *cell) uncover() {
	for i := c.up; i != c; i = i.up {
		for j := i.left; j != i; j = j.left {
			logger.Println(c, i, j)
			j.RestoreVertically()
		}
	}
	c.RestoreHorizontally()
}
