package solver

import "fmt"

type Solver struct {
	Dimensions uint8
	solutions  uint
}

func New(dimensions uint8) Solver {
	if dimensions <= 3 || dimensions > 13 {
		panic(fmt.Sprintf("Dimensions must be in [3..10[; Is: %d", dimensions))
	}

	return Solver{Dimensions: dimensions, solutions: 0}
}

func (s *Solver) Solve() uint {
	s.solutions = 0
	s.solveForRow(0, make([]Position, s.Dimensions))
	return s.solutions
}

func (s *Solver) solveForRow(row uint8, ps []Position) {
	if row == s.Dimensions {
		s.solutions += 1
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
			s.solveForRow(row+1, ps)
		}
	}
}
