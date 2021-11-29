package utils

import "github.com/biesnecker/godvent/types"

func ManhattanDistance(start types.Coord, end types.Coord) int {
	return IntAbs(start.X-end.X) + IntAbs(start.Y-end.Y)
}

func ManhattanDistance3(start types.Coord3, end types.Coord3) int {
	return IntAbs(start.X-end.X) + IntAbs(start.Y-end.Y) + IntAbs(start.Z-end.Z)
}

func HexDistance(h, g types.HexCoord) int {
	return (IntAbs(h.Q-g.Q) + IntAbs(h.R-g.R) + IntAbs(h.S-g.S)) / 2
}
