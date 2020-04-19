package random

import (
	"testing"

	"uci-impl/internal/solver"
)

func TestOnlySelectedMovesReturned(t *testing.T) {
	sp := solver.NewSearchParams()
	moves := []string{"e2e4", "g1f3"}
	randomSolver := NewRandomSolver()

	for i := 0; i < 10; i++ {
		result := <-randomSolver.StartSearch(sp, moves...)
		if len(result) == 0 {
			t.Fail()
		} else if !contains(moves, result[0]) {
			t.Fail()
		}
	}
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func TestPonderHit(t *testing.T) {
	sp := &solver.SearchParams{Ponder: true}
	moves := []string{"e2e4", "g1f3"}
	randomSolver := NewRandomSolver()

	for i := 0; i < 10; i++ {
		ch := randomSolver.StartSearch(sp, moves...)
		if len(ch) != 0 {
			t.Fail()
		}

		randomSolver.PonderHit()
		if len(ch) != 1 {
			t.Fail()
		}

		result := <-ch
		if !contains(moves, result[0]) {
			t.Fail()
		}
	}
}

func TestPonderMiss(t *testing.T) {
	sp := solver.NewSearchParams()
	sp.Ponder = true

	randomSolver := NewRandomSolver()

	ch := randomSolver.StartSearch(sp)
	if len(ch) != 0 {
		t.Fail()
	}

	randomSolver.StopSearch()
	if len(ch) != 0 {
		t.Fail()
	}
}
