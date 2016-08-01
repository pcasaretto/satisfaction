package x

import "fmt"

var _id uint64 = 0

type cell struct {
	left   *cell
	right  *cell
	up     *cell
	down   *cell
	header *cell
	value  interface{}
	size   uint
	id     uint64
}

func newCell(v interface{}) *cell {
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

func (c *cell) pushCellDown(other *cell) {
	c.down.setUp(other)
	other.setDown(c.down)
	c.down = other
	other.setUp(c)
}

func (c *cell) pushCellUp(other *cell) {
	c.up.setDown(other)
	other.setUp(c.up)
	c.up = other
	other.setDown(c)
}

func (c *cell) pushCellLeft(other *cell) {
	c.left.setRight(other)
	other.setLeft(c.left)
	c.left = other
	other.setRight(c)
}

func (c *cell) pushCellRight(other *cell) {
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

func (c *cell) removeVertically() {
	c.up.setDown(c.down)
	c.down.setUp(c.up)
}

func (c *cell) restoreVertically() {
	c.up.setDown(c)
	c.down.setUp(c)
}

func (c *cell) removeHorizontally() {
	c.right.setLeft(c.left)
	c.left.setRight(c.right)
}

func (c *cell) restoreHorizontally() {
	c.right.setLeft(c)
	c.left.setRight(c)
}

func (c *cell) cover() {
	c.removeHorizontally()
	for i := c.down; i != c; i = i.down {
		for j := i.right; j != i; j = j.right {
			j.removeVertically()
			j.header.size--
		}
	}
}

func (c *cell) uncover() {
	for i := c.up; i != c; i = i.up {
		for j := i.left; j != i; j = j.left {
			j.restoreVertically()
			j.header.size++
		}
	}
	c.restoreHorizontally()
}
