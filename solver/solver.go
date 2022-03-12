package solver

import "fmt"

type Solver struct {
	Dimensions uint8
}

func New(dimensions uint8) Solver {
	var min, max uint8 = 4, 13

	if dimensions < min || dimensions > max {
		panic(fmt.Sprintf("Dimensions must be in [%d..%d]; is: %d", min, max, dimensions))
	}

	return Solver{Dimensions: dimensions}
}

func (s *Solver) Solve() uint {
	var solutions uint = 0
	s.solveForRow(0, make([]Position, s.Dimensions), &solutions)
	return solutions
}

func (s *Solver) solveForRow(row uint8, ps []Position, solutions *uint) {
	if row == s.Dimensions {
		*solutions += 1
		return
	}

	for col := uint8(0); col < s.Dimensions; col++ {
		pos := NewPosition(row, col)
		overlaps := false

		for i := uint8(0); i < row; i++ {
			if overlaps = pos.Overlaps(&ps[i]); overlaps {
				break
			}
		}

		if !overlaps {
			ps[row] = pos
			s.solveForRow(row+1, ps, solutions)
		}
	}
}
