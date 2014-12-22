package dancinglinks

var _id int = 0

type Cell struct {
	left  *Cell
	right *Cell
	up    *Cell
	down  *Cell
	id    int
}

func NewCell() *Cell {
	cell := &Cell{}
	cell.down = cell
	cell.up = cell
	cell.left = cell
	cell.right = cell
	_id = _id + 1
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

func (cell *Cell) cellsGivenDirection(dir direction) []*Cell {
  var cells []*Cell

  var next func(c *Cell) *Cell

  switch dir {
  case up:
    next = func(c *Cell) *Cell { return c.up }
  case down:
    next = func(c *Cell) *Cell { return c.down }
  case left:
    next = func(c *Cell) *Cell { return c.left }
  case right:
    next = func(c *Cell) *Cell { return c.right }
  }

	for c := next(cell); c != cell; c = next(c) {
		cells = append(cells, c)
	}
	return cells
}

func (cell *Cell) cellsDown() []*Cell {
  return cell.cellsGivenDirection(down)
}

func (cell *Cell) cellsUp() []*Cell {
  return cell.cellsGivenDirection(up)
}

func (cell *Cell) cellsLeft() []*Cell {
  return cell.cellsGivenDirection(left)
}

func (cell *Cell) cellsRight() []*Cell {
  return cell.cellsGivenDirection(right)
}

func (cell *Cell) addCellDown(toAdd *Cell) {
	cell.down.up = toAdd
	toAdd.down = cell.down
	cell.down = toAdd
	toAdd.up = cell
}

func (cell *Cell) addCellUp(toAdd *Cell) {
	cell.up.down = toAdd
	toAdd.up = cell.up
	cell.up = toAdd
	toAdd.down = cell
}

func (cell *Cell) addCellLeft(toAdd *Cell) {
	cell.left.right = toAdd
	toAdd.left = cell.left
	cell.left = toAdd
	toAdd.right = cell
}

func (cell *Cell) addCellRight(toAdd *Cell) {
	cell.right.left = toAdd
	toAdd.right = cell.right
	cell.right = toAdd
	toAdd.left = cell
}

func (cell *Cell) implodeVertically() {
  cell.up.down = cell.down
  cell.down.up = cell.up
}

func (cell *Cell) restoreVertically() {
  cell.up.down = cell
  cell.down.up = cell
}

func (cell *Cell) implodeHorizontally() {
  cell.right.left = cell.left
  cell.left.right = cell.right
}

func (cell *Cell) restoreHorizontally() {
  cell.right.left = cell
  cell.left.right = cell
}
