package random

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/mhv2109/uci-impl/internal/solver"
)

func TestOnlySelectedMovesReturned(t *testing.T) {
	g := NewGomegaWithT(t)

	sp := solver.NewSearchParams()
	moves := []string{"e2e4", "g1f3"}
	randomSolver := NewRandomSolver()

	for i := 0; i < 10; i++ {
		result := <-randomSolver.StartSearch(sp, moves...)

		g.Expect(result).
			ToNot(HaveLen(0))
		g.Expect(moves).
			To(ContainElement(result[0]))
	}
}

func TestPonderHit(t *testing.T) {
	g := NewGomegaWithT(t)

	sp := &solver.SearchParams{Ponder: true}
	moves := []string{"e2e4", "g1f3"}
	randomSolver := NewRandomSolver()

	for i := 0; i < 10; i++ {
		ch := randomSolver.StartSearch(sp, moves...)
		g.Expect(ch).
			To(HaveLen(0))

		randomSolver.PonderHit()
		g.Expect(ch).
			To(HaveLen(1))

		result := <-ch
		g.Expect(moves).
			To(ContainElement(result[0]))
	}
}

func TestPonderMiss(t *testing.T) {
	g := NewGomegaWithT(t)

	sp := solver.NewSearchParams()
	sp.Ponder = true

	randomSolver := NewRandomSolver()

	ch := randomSolver.StartSearch(sp)
	g.Expect(ch).
		To(HaveLen(0))

	randomSolver.StopSearch()
	g.Expect(ch).
		To(HaveLen(0))
}
