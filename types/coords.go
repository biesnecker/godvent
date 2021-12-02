package types

type Coord struct {
	X, Y int
}

func (c Coord) GetSurroundingCoords() []Coord {
	res := make([]Coord, 0, 8)
	for x := c.X - 1; x < c.X+2; x++ {
		for y := c.Y - 1; y < c.Y+2; y++ {
			if x == c.X && y == c.Y {
				continue
			}
			res = append(res, Coord{X: x, Y: y})
		}
	}
	return res
}

func (c Coord) GetAdjacentCoords() []Coord {
	return []Coord{c.Up(), c.Down(), c.Left(), c.Right()}
}

func (c Coord) Equals(x, y int) bool {
	return c.X == x && c.Y == y
}

func (c Coord) Up() Coord {
	return c.UpBy(1)
}

func (c Coord) UpBy(i int) Coord {
	return Coord{X: c.X, Y: c.Y + i}
}

func (c Coord) Down() Coord {
	return c.DownBy(1)
}

func (c Coord) DownBy(i int) Coord {
	return Coord{X: c.X, Y: c.Y - i}
}

func (c Coord) Right() Coord {
	return c.RightBy(1)
}

func (c Coord) RightBy(i int) Coord {
	return Coord{X: c.X + i, Y: c.Y}
}

func (c Coord) Left() Coord {
	return c.LeftBy(1)
}

func (c Coord) LeftBy(i int) Coord {
	return Coord{X: c.X - i, Y: c.Y}
}

func (c Coord) IsNegative() bool {
	return c.X < 0 || c.Y < 0
}

func (c Coord) IsInBounds(minX, minY, maxX, maxY int) bool {
	return c.X >= minX && c.X <= maxX && c.Y >= minY && c.Y <= maxY
}
