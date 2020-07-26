package minimax

import (
	"testing"

	"github.com/notnil/chess"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mhv2109/uci-impl/internal/handler"
	hf "github.com/mhv2109/uci-impl/internal/handler/handlerfakes"
	"github.com/mhv2109/uci-impl/internal/solver"
)

var _ = Describe("MinimaxSolver", func() {
	var (
		emitter       handler.Emitter
		minimaxSolver solver.Solver
	)

	BeforeEach(func() {
		emitter = &hf.FakeEmitter{}
		minimaxSolver = NewMinimaxSolverWithEmitter(emitter)
	})

	It("Returns results", func() {
		sp := solver.NewSearchParams()
		sp.Wtime = 300000
		sp.Btime = 300000

		ch := minimaxSolver.StartSearch(sp)
		Eventually(ch).
			Should(Receive())
	})
})

var _ = Describe("MinimaxAlgo", func() {
	var (
		emitter   handler.Emitter
		called    bool
		submitted [][]string
	)

	BeforeEach(func() {
		called = false
		submitted = make([][]string, 0)
		emitter = &hf.FakeEmitter{}
	})

	submit := func(move []string) bool {
		called = true
		submitted = append(submitted, move)
		return true
	}

	It("Calls submit", func() {
		game := chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))

		algo := newMinimaxAlgo(1, 32, submit, emitter)
		algo.Start(game.Position())

		Expect(called).
			To(BeTrue())
	})

	It("Solver only submits valid moves", func() {
		game := chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))

		expected := len(game.ValidMoves())
		valid := make([]string, 0, expected)
		for _, move := range game.ValidMoves() {
			valid = append(valid, move.String())
		}

		algo := newMinimaxAlgo(1, 32, submit, emitter)
		algo.Start(game.Position())

		actual := len(submitted)
		Expect(actual > expected).
			To(BeFalse(), "Actual %d > Expected %d", actual, expected)
		for _, move := range submitted {
			Expect(valid).
				To(ContainElement(move[0]))
		}
	})

	It("Takes Pawn", func() {
		fen, _ := chess.FEN("rnbqkbnr/ppppppp1/7p/6P1/8/8/PPPPPP1P/RNBQKBNR b KQkq - 0 2")
		game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

		algo := newMinimaxAlgo(3, 128, submit, emitter)
		algo.Start(game.Position())

		best := submitted[len(submitted)-1]
		Expect(best[0]).To(Equal("h6g5"))
	})

})

func BenchmarkTakePawnSelected(b *testing.B) {
	emitter := &hf.FakeEmitter{}
	fen, _ := chess.FEN("rnbqkbnr/ppppppp1/7p/6P1/8/8/PPPPPP1P/RNBQKBNR b KQkq - 0 2")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	submit := func(move []string) bool {
		return true
	}

	for i := 0; i < b.N; i++ {
		minimax := newMinimaxAlgo(3, 32, submit, emitter)
		minimax.Start(game.Position())
	}
}

func Benchmark2(b *testing.B) {
	emitter := &hf.FakeEmitter{}
	fen, _ := chess.FEN("rnb1k2r/pppp1ppp/5n2/8/P7/R1PP4/1P1K2Pq/1NBQ1BR1 w kq - 0 11")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	submit := func(move []string) bool {
		return true
	}

	for i := 0; i < b.N; i++ {
		minimax := newMinimaxAlgo(3, 32, submit, emitter)
		minimax.Start(game.Position())
	}
}
