package handler

import (
	"testing"

	"github.com/golang/mock/gomock"

	mh "github.com/mhv2109/uci-impl/internal/handler/mock"
	s "github.com/mhv2109/uci-impl/internal/solver"
	ms "github.com/mhv2109/uci-impl/internal/solver/mock"
)

func NewMocks(t *testing.T) (*ms.MockSolver, *mh.MockEmitter, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	solver := ms.NewMockSolver(ctrl)
	emitter := mh.NewMockEmitter(ctrl)
	return solver, emitter, ctrl
}

func TestWaitGroupIncrementsAndDecrements(t *testing.T) {
	solver, emitter, ctrl := NewMocks(t)
	defer ctrl.Finish()

	h := NewHandlerWithEmitter(solver, emitter)

	// Handler method does nothing, but should increment & decrement WaitGroup
	h.Handle([]string{"ucinewgame"})

	// Expect "readyok".  In the event that the WaitGroup isn't decremented, hangs forever
	emitter.EXPECT().EmitReadyOK()

	h.Handle([]string{"isready"})
}

func TestIsReadyCanBeCalledSuccessively(t *testing.T) {
	solver, emitter, ctrl := NewMocks(t)
	defer ctrl.Finish()

	h := NewHandlerWithEmitter(solver, emitter)

	input := []string{"ucinewgame"}
	for i := 0; i < 100; i++ {
		h.Handle(input)
	}
}

func TestHandleSetOptionEx1(t *testing.T) {
	solver, emitter, ctrl := NewMocks(t)
	defer ctrl.Finish()

	h := NewHandlerWithEmitter(solver, emitter)

	input := []string{"setoption", "name", "Nullmove", "value", "true"}

	solver.
		EXPECT().
		SetOption(gomock.Eq("nullmove"), gomock.Eq("true"))

	h.Handle(input)
}

func TestHandleSetOptionEx2(t *testing.T) {
	solver, emitter, ctrl := NewMocks(t)
	defer ctrl.Finish()

	h := NewHandlerWithEmitter(solver, emitter)

	input := []string{"setoption", "name", "Style", "value", "Risky"}

	solver.
		EXPECT().
		SetOption(gomock.Eq("style"), gomock.Eq("Risky"))

	h.Handle(input)
}

func TestHandleSetOptionEx3(t *testing.T) {
	solver, emitter, ctrl := NewMocks(t)
	defer ctrl.Finish()

	h := NewHandlerWithEmitter(solver, emitter)

	input := []string{"setoption", "name", "Clear", "Hash"}

	solver.
		EXPECT().
		SetOption(gomock.Eq("clear hash"), gomock.Eq(""))

	h.Handle(input)
}

func TestHandleSetOptionEx4(t *testing.T) {
	solver, emitter, ctrl := NewMocks(t)
	defer ctrl.Finish()

	h := NewHandlerWithEmitter(solver, emitter)

	input := []string{"setoption", "name", "NalimovPath", "value",
		"c:\\chess\\tb\\4;c:\\chess\\tb\\5\\n"}

	solver.
		EXPECT().
		SetOption(gomock.Eq("nalimovpath"),
			gomock.Eq("c:\\chess\\tb\\4;c:\\chess\\tb\\5\\n"))

	h.Handle(input)
}

func TestPositionFEN(t *testing.T) {
	solver, emitter, ctrl := NewMocks(t)
	defer ctrl.Finish()

	h := NewHandlerWithEmitter(solver, emitter)

	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	input := []string{"position", "fen", fen, "moves", "e4", "Nf3", "Bb5"}

	solver.
		EXPECT().
		SetPosition(gomock.Eq(fen), gomock.Eq("e4"), gomock.Eq("Nf3"),
			gomock.Eq("Bb5"))

	h.Handle(input)
}

func TestStartPosition(t *testing.T) {
	solver, emitter, ctrl := NewMocks(t)
	defer ctrl.Finish()

	h := NewHandlerWithEmitter(solver, emitter)

	input := []string{"position", "startpos", "moves", "e4", "Nf3", "Bb5"}

	solver.
		EXPECT().
		SetStartPosition(gomock.Eq("e4"), gomock.Eq("Nf3"),
			gomock.Eq("Bb5"))

	h.Handle(input)
}

func TestGoSearchmoves(t *testing.T) {
	solver, emitter, ctrl := NewMocks(t)
	defer ctrl.Finish()

	h := NewHandlerWithEmitter(solver, emitter)

	input := []string{"go", "searchmoves", "e4", "Nf3", "Bb5"}

	solver.
		EXPECT().
		StartSearch(gomock.Any(), gomock.Eq("e4"),
			gomock.Eq("Nf3"), gomock.Eq("Bb5"))

	h.Handle(input)
}

func TestGoSearchParams(t *testing.T) {
	solver, emitter, ctrl := NewMocks(t)
	defer ctrl.Finish()

	h := NewHandlerWithEmitter(solver, emitter)

	input := []string{"go", "ponder", "wtime", "1", "btime", "2", "winc",
		"3", "binc", "4", "movestogo", "5", "depth", "6", "mate", "7",
		"movetime", "8", "infinite"}
	expected := s.NewSearchParams()
	expected.Ponder = true
	expected.Wtime = 1
	expected.Btime = 2
	expected.Winc = 3
	expected.Binc = 4
	expected.Movestogo = 5
	expected.Depth = 6
	expected.Mate = 7
	expected.Movetime = 8
	expected.Infinite = true

	solver.
		EXPECT().
		StartSearch(gomock.Eq(expected))

	h.Handle(input)
}
