package random_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mhv2109/uci-impl/internal/solver"
	"github.com/mhv2109/uci-impl/internal/solver/random"
)

var _ = Describe("RandomSolver", func() {
	var (
		sp           *solver.SearchParams
		randomSolver solver.Solver
	)

	BeforeEach(func() {
		sp = solver.NewSearchParams()
		randomSolver = random.NewRandomSolver()
	})

	It("Only returs selected moves", func() {
		moves := []string{"e2e4", "g1f3"}

		for i := 0; i < 10; i++ {
			result := <-randomSolver.StartSearch(sp, moves...)

			Expect(result).
				ToNot(HaveLen(0))
			Expect(moves).
				To(ContainElement(result[0]))
		}
	})

	It("Returns move when ponder hit", func() {
		sp.Ponder = true
		moves := []string{"e2e4", "g1f3"}

		for i := 0; i < 10; i++ {
			ch := randomSolver.StartSearch(sp, moves...)
			Expect(ch).
				To(HaveLen(0))

			randomSolver.PonderHit()
			Expect(ch).
				To(HaveLen(1))

			result := <-ch
			Expect(moves).
				To(ContainElement(result[0]))
		}
	})

	It("Doesn't return move when ponder miss", func() {
		sp.Ponder = true

		ch := randomSolver.StartSearch(sp)
		Expect(ch).
			To(HaveLen(0))

		randomSolver.StopSearch()
		Expect(ch).
			To(HaveLen(0))
	})
})
