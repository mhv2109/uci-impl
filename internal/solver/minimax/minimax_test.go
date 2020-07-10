package minimax

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mhv2109/uci-impl/internal/handler/mock"
	"github.com/mhv2109/uci-impl/internal/solver"
	"github.com/notnil/chess"
)

func TestMinimaxSolverReturnsResults(t *testing.T) {
	ctrl := gomock.NewController(t)
	emitter := mock.NewMockEmitter(ctrl)
	defer ctrl.Finish()

	emitter.
		EXPECT().
		EmitInfo(gomock.Any()).
		AnyTimes()

	sp := solver.NewSearchParams()
	sp.Wtime = 300000
	sp.Btime = 300000
	minimaxSolver := NewMinimaxSolverWithEmitter(emitter)

	resultCh := minimaxSolver.StartSearch(sp)
	time.Sleep(1 * time.Second)
	if len(resultCh) <= 0 {
		t.Fail()
	}
}

func TestMinimaxCallsSubmit(t *testing.T) {
	ctrl := gomock.NewController(t)
	emitter := mock.NewMockEmitter(ctrl)
	defer ctrl.Finish()
	game := chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))

	emitter.
		EXPECT().
		EmitInfo(gomock.Any()).
		AnyTimes()

	called := false
	submit := func(move []string) bool {
		called = true
		return true
	}

	minimax := newMinimaxAlgo(1, 32, submit, emitter)
	minimax.Start(game.Position())

	if !called {
		t.Fail()
	}
}

func TestOnlyValidMovesSubmitted(t *testing.T) {
	ctrl := gomock.NewController(t)
	emitter := mock.NewMockEmitter(ctrl)
	defer ctrl.Finish()

	game := chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))

	emitter.
		EXPECT().
		EmitInfo(gomock.Any()).
		AnyTimes()

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

	if actual := len(submitted); actual > expected {
		t.Errorf("Actual %d > Expected %d", expected, actual)
	}
	for _, move := range submitted {
		m := move[0]
		found := false
		for _, vm := range valid {
			if m == vm {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Invalid move: %s", m)
		}
	}
}

func TestTakePawnSelected(t *testing.T) {
	ctrl := gomock.NewController(t)
	emitter := mock.NewMockEmitter(ctrl)
	defer ctrl.Finish()

	fen, _ := chess.FEN("rnbqkbnr/ppppppp1/7p/6P1/8/8/PPPPPP1P/RNBQKBNR b KQkq - 0 2")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	emitter.
		EXPECT().
		EmitInfo(gomock.Any()).
		AnyTimes()

	submitted := make([][]string, 0, len(game.ValidMoves()))
	submit := func(move []string) bool {
		submitted = append(submitted, move)
		return true
	}

	minimax := newMinimaxAlgo(3, 128, submit, emitter)

	minimax.Start(game.Position())

	best := submitted[len(submitted)-1]
	if actual, expected := best[0], "h6g5"; expected != actual {
		t.Errorf("Expected %s, actual %s", expected, actual)
	}
}

func BenchmarkTakePawnSelected(b *testing.B) {
	ctrl := gomock.NewController(b)
	emitter := mock.NewMockEmitter(ctrl)
	defer ctrl.Finish()

	fen, _ := chess.FEN("rnbqkbnr/ppppppp1/7p/6P1/8/8/PPPPPP1P/RNBQKBNR b KQkq - 0 2")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	emitter.
		EXPECT().
		EmitInfo(gomock.Any()).
		AnyTimes()

	submit := func(move []string) bool {
		return true
	}

	for i := 0; i < b.N; i++ {
		minimax := newMinimaxAlgo(3, 32, submit, emitter)
		minimax.Start(game.Position())
	}
}

func Benchmark2(b *testing.B) {
	ctrl := gomock.NewController(b)
	emitter := mock.NewMockEmitter(ctrl)
	defer ctrl.Finish()

	fen, _ := chess.FEN("rnb1k2r/pppp1ppp/5n2/8/P7/R1PP4/1P1K2Pq/1NBQ1BR1 w kq - 0 11")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	emitter.
		EXPECT().
		EmitInfo(gomock.Any()).
		AnyTimes()

	submit := func(move []string) bool {
		return true
	}

	for i := 0; i < b.N; i++ {
		minimax := newMinimaxAlgo(3, 32, submit, emitter)
		minimax.Start(game.Position())
	}
}
