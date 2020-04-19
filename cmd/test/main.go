package main

import (
	"fmt"

	"uci-impl/internal/solver"
	"uci-impl/internal/solver/random"
)

func main() {
	sp := solver.NewSearchParams()
	moves := []string{"e2e4", "g1f3"}
	randomSolver := random.NewRandomSolver()

	result := <-randomSolver.StartSearch(sp, moves...)
	fmt.Println(result)
}
