package minimax

import (
	"testing"

	"github.com/notnil/chess"
	. "github.com/onsi/gomega"

	hf "github.com/mhv2109/uci-impl/internal/handler/handlerfakes"
	"github.com/mhv2109/uci-impl/internal/solver"
)

func TestMinimaxSolverReturnsResults(t *testing.T) {
	g := NewGomegaWithT(t)

	emitter := &hf.FakeEmitter{}

	sp := solver.NewSearchParams()
	sp.Wtime = 300000
	sp.Btime = 300000
	minimaxSolver := NewMinimaxSolverWithEmitter(emitter)

	ch := minimaxSolver.StartSearch(sp)
	g.Eventually(ch).
		Should(Receive())
}

func TestMinimaxCallsSubmit(t *testing.T) {
	g := NewGomegaWithT(t)

	emitter := &hf.FakeEmitter{}
	game := chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))

	called := false
	submit := func(move []string) bool {
		called = true
		return true
	}

	minimax := newMinimaxAlgo(1, 32, submit, emitter)
	minimax.Start(game.Position())

	g.Expect(called).
		To(BeTrue())

	if !called {
		t.Fail()
	}
}

func TestOnlyValidMovesSubmitted(t *testing.T) {
	g := NewGomegaWithT(t)

	emitter := &hf.FakeEmitter{}
	game := chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))

	expected := len(game.ValidMoves())
	valid := make([]string, 0, expected)
	for _, move := range game.ValidMoves() {
		valid = append(valid, move.String())
	}

	submitted := make([][]string, 0, expected)
	submit := func(move []string) bool {
		submitted = append(submitted, move)
		return true
	}

	minimax := newMinimaxAlgo(1, 32, submit, emitter)
	minimax.Start(game.Position())

	actual := len(submitted)
	g.Expect(actual > expected).
		To(BeFalse(), "Actual %d > Expected %d", actual, expected)
	for _, move := range submitted {
		g.Expect(valid).
			To(ContainElement(move[0]))
	}
}

func TestTakePawnSelected(t *testing.T) {
	g := NewGomegaWithT(t)

	emitter := &hf.FakeEmitter{}
	fen, _ := chess.FEN("rnbqkbnr/ppppppp1/7p/6P1/8/8/PPPPPP1P/RNBQKBNR b KQkq - 0 2")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	submitted := make([][]string, 0, len(game.ValidMoves()))
	submit := func(move []string) bool {
		submitted = append(submitted, move)
		return true
	}

	minimax := newMinimaxAlgo(3, 128, submit, emitter)

	minimax.Start(game.Position())

	best := submitted[len(submitted)-1]
	g.Expect(best[0]).To(Equal("h6g5"))
}

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
