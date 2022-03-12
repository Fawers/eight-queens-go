package solver

type Position struct {
	Row, Col uint8
}

func NewPosition(row, col uint8) Position {
	return Position{Row: row, Col: col}
}

func (p *Position) Overlaps(p2 *Position) bool {
	if p.Row == p2.Row || p.Col == p2.Col {
		return true
	}

	return abs(p2.Row-p.Row) == abs(p2.Col-p.Col)
}

func abs(u uint8) uint8 {
	if u&0x80 != 0 {
		return ^u + 1
	}

	return u
}
