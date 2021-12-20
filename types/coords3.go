package types

type Coord3 struct {
	X, Y, Z int
}

func cosine(angle int) int {
	switch angle {
	case 0:
		return 1
	case 90:
		return 0
	case 180:
		return -1
	case 270:
		return 0
	}
	panic("Invalid angle to cos")
}

func sine(angle int) int {
	switch angle {
	case 0:
		return 0
	case 90:
		return 1
	case 180:
		return 0
	case 270:
		return -1
	}
	panic("Invalid angle to sin")
}

func (c Coord3) RotateX(theta int) Coord3 {
	cos, sin := cosine(theta), sine(theta)
	//	return (x, y*c - z*s, y*s + z*c)
	return Coord3{c.X, c.Y*cos - c.Z*sin, c.Y*sin + c.Z*cos}
}

func (c Coord3) RotateY(theta int) Coord3 {
	cos, sin := cosine(theta), sine(theta)
	// 	return (x*c + z*s, y, -x*s + z*c)
	return Coord3{X: c.X*cos + c.Z*sin, Y: c.Y, Z: -c.X*sin + c.Z*cos}
}

func (c Coord3) RotateZ(theta int) Coord3 {
	cos, sin := cosine(theta), sine(theta)
	//return (x*c - y*s, x*s+y*c, z)
	return Coord3{X: c.X*cos - c.Y*sin, Y: c.X*sin + c.Y*cos, Z: c.Z}
}

func (c Coord3) Rotate(x, y, z int) Coord3 {
	c = c.RotateX(x)
	c = c.RotateY(y)
	return c.RotateZ(z)
}

func (c Coord3) Diff(other Coord3) Vector3 {
	return Vector3{
		X: c.X - other.X,
		Y: c.Y - other.Y,
		Z: c.Z - other.Z,
	}
}

func (c Coord3) Translate(v Vector3) Coord3 {
	return Coord3{
		X: c.X + v.X,
		Y: c.Y + v.Y,
		Z: c.Z + v.Z,
	}
}
