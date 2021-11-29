package types

type HexCoord struct {
	Q, R, S int
}

func (h HexCoord) North() HexCoord {
	return HexCoord{h.Q, h.R - 1, h.S + 1}
}

func (h HexCoord) Northeast() HexCoord {
	return HexCoord{h.Q + 1, h.R - 1, h.S}
}

func (h HexCoord) Southeast() HexCoord {
	return HexCoord{h.Q + 1, h.R, h.S - 1}
}

func (h HexCoord) South() HexCoord {
	return HexCoord{h.Q, h.R + 1, h.S - 1}
}

func (h HexCoord) Southwest() HexCoord {
	return HexCoord{h.Q - 1, h.R + 1, h.S}
}

func (h HexCoord) Northwest() HexCoord {
	return HexCoord{h.Q - 1, h.R, h.S + 1}
}
