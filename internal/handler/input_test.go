package handler_test

import (
	"testing"

	. "github.com/onsi/gomega"

	. "github.com/mhv2109/uci-impl/internal/handler"
	hf "github.com/mhv2109/uci-impl/internal/handler/handlerfakes"
	s "github.com/mhv2109/uci-impl/internal/solver"
	sf "github.com/mhv2109/uci-impl/internal/solver/solverfakes"
)

func TestWaitGroupIncrementsAndDecrements(t *testing.T) {
	g := NewGomegaWithT(t)

	solver := &sf.FakeSolver{}
	emitter := &hf.FakeEmitter{}

	h := NewHandlerWithEmitter(solver, emitter)

	// Handler method does nothing, but should increment & decrement WaitGroup
	h.Handle([]string{"ucinewgame"})

	h.Handle([]string{"isready"})

	// Expect "readyok".  In the event that the WaitGroup isn't decremented, hangs forever
	g.Expect(emitter.EmitReadyOKCallCount()).To(Equal(1))
}

func TestIsReadyCanBeCalledSuccessively(t *testing.T) {
	solver := &sf.FakeSolver{}
	emitter := &hf.FakeEmitter{}

	h := NewHandlerWithEmitter(solver, emitter)

	input := []string{"ucinewgame"}
	for i := 0; i < 100; i++ {
		h.Handle(input)
	}
}

func TestHandleSetOptionEx1(t *testing.T) {
	g := NewGomegaWithT(t)

	solver := &sf.FakeSolver{}
	emitter := &hf.FakeEmitter{}

	h := NewHandlerWithEmitter(solver, emitter)

	input := []string{"setoption", "name", "Nullmove", "value", "true"}

	h.Handle(input)

	key, value := solver.SetOptionArgsForCall(0)
	g.Expect(key).To(Equal("nullmove"))
	g.Expect(value).To(Equal("true"))
}

func TestHandleSetOptionEx2(t *testing.T) {
	g := NewGomegaWithT(t)

	solver := &sf.FakeSolver{}
	emitter := &hf.FakeEmitter{}

	h := NewHandlerWithEmitter(solver, emitter)

	input := []string{"setoption", "name", "Style", "value", "Risky"}

	h.Handle(input)

	key, value := solver.SetOptionArgsForCall(0)
	g.Expect(key).To(Equal("style"))
	g.Expect(value).To(Equal("Risky"))
}

func TestHandleSetOptionEx3(t *testing.T) {
	g := NewGomegaWithT(t)

	solver := &sf.FakeSolver{}
	emitter := &hf.FakeEmitter{}

	h := NewHandlerWithEmitter(solver, emitter)

	input := []string{"setoption", "name", "Clear", "Hash"}

	h.Handle(input)

	key, value := solver.SetOptionArgsForCall(0)
	g.Expect(key).To(Equal("clear hash"))
	g.Expect(value).To(Equal(""))
}

func TestHandleSetOptionEx4(t *testing.T) {
	g := NewGomegaWithT(t)

	solver := &sf.FakeSolver{}
	emitter := &hf.FakeEmitter{}

	h := NewHandlerWithEmitter(solver, emitter)

	input := []string{"setoption", "name", "NalimovPath", "value",
		"c:\\chess\\tb\\4;c:\\chess\\tb\\5\\n"}

	h.Handle(input)

	key, value := solver.SetOptionArgsForCall(0)
	g.Expect(key).To(Equal("nalimovpath"))
	g.Expect(value).To(Equal("c:\\chess\\tb\\4;c:\\chess\\tb\\5\\n"))
}

func TestPositionFEN(t *testing.T) {
	g := NewGomegaWithT(t)

	solver := &sf.FakeSolver{}
	emitter := &hf.FakeEmitter{}

	h := NewHandlerWithEmitter(solver, emitter)

	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	e0, e1, e2 := "e4", "Nf3", "Bb5"
	input := []string{"position", "fen", fen, "moves", e0, e1, e2}

	h.Handle(input)

	f, p := solver.SetPositionArgsForCall(0)
	g.Expect(f).To(Equal(fen))
	g.Expect(p[0]).To(Equal(e0))
	g.Expect(p[1]).To(Equal(e1))
	g.Expect(p[2]).To(Equal(e2))
}

func TestStartPosition(t *testing.T) {
	g := NewGomegaWithT(t)

	solver := &sf.FakeSolver{}
	emitter := &hf.FakeEmitter{}

	h := NewHandlerWithEmitter(solver, emitter)

	pos, e0, e1, e2 := "startpos", "e4", "Nf3", "Bb5"
	input := []string{"position", pos, "moves", e0, e1, e2}

	h.Handle(input)

	p := solver.SetStartPositionArgsForCall(0)
	g.Expect(p[0]).To(Equal(e0))
	g.Expect(p[1]).To(Equal(e1))
	g.Expect(p[2]).To(Equal(e2))
}

func TestGoSearchmoves(t *testing.T) {
	g := NewGomegaWithT(t)

	solver := &sf.FakeSolver{}
	emitter := &hf.FakeEmitter{}

	h := NewHandlerWithEmitter(solver, emitter)

	e0, e1, e2 := "e4", "Nf3", "Bb5"
	input := []string{"go", "searchmoves", e0, e1, e2}

	h.Handle(input)

	_, a := solver.StartSearchArgsForCall(0)
	g.Expect(a[0]).To(Equal(e0))
	g.Expect(a[1]).To(Equal(e1))
	g.Expect(a[2]).To(Equal(e2))
}

func TestGoSearchParams(t *testing.T) {
	g := NewGomegaWithT(t)

	solver := &sf.FakeSolver{}
	emitter := &hf.FakeEmitter{}

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

	h.Handle(input)

	actual, _ := solver.StartSearchArgsForCall(0)
	g.Expect(*actual).
		To(Equal(*expected))
}
