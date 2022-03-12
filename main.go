package main

import (
	"fmt"

	"github.com/Fawers/eight-queens-go/solver"
)

func main() {
	s := solver.New(8)
	solutions := s.Solve()
	fmt.Printf("Found %d solutions\n", solutions)
}
