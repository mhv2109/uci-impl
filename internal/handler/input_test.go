package handler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mhv2109/uci-impl/internal/handler"
	hf "github.com/mhv2109/uci-impl/internal/handler/handlerfakes"
	s "github.com/mhv2109/uci-impl/internal/solver"
	sf "github.com/mhv2109/uci-impl/internal/solver/solverfakes"
)

var _ = Describe("Handler", func() {
	var (
		solver  *sf.FakeSolver
		emitter *hf.FakeEmitter
		handler *UCIInputHandler
	)

	BeforeEach(func() {
		solver = &sf.FakeSolver{}
		emitter = &hf.FakeEmitter{}
		handler = NewHandlerWithEmitter(solver, emitter)
	})

	It("Waitgroup increments and decrements", func() {
		// Handler method does nothing, but should increment & decrement WaitGroup
		handler.Handle([]string{"ucinewgame"})

		handler.Handle([]string{"isready"})

		// Expect "readyok".  In the event that the WaitGroup isn't decremented, hangs forever
		Expect(emitter.EmitReadyOKCallCount()).To(Equal(1))
	})

	It("ucinewgame", func() {
		input := []string{"ucinewgame"}
		for i := 0; i < 100; i++ {
			handler.Handle(input)
		}
	})

	var _ = Describe("setoption", func() {
		It("Set Nullmove option", func() {
			input := []string{"setoption", "name", "Nullmove", "value", "true"}

			handler.Handle(input)

			key, value := solver.SetOptionArgsForCall(0)
			Expect(key).To(Equal("nullmove"))
			Expect(value).To(Equal("true"))
		})

		It("Set Style option", func() {
			input := []string{"setoption", "name", "Style", "value", "Risky"}

			handler.Handle(input)

			key, value := solver.SetOptionArgsForCall(0)
			Expect(key).To(Equal("style"))
			Expect(value).To(Equal("Risky"))
		})

		It("Set Clear option", func() {
			input := []string{"setoption", "name", "Clear", "Hash"}

			handler.Handle(input)

			key, value := solver.SetOptionArgsForCall(0)
			Expect(key).To(Equal("clear hash"))
			Expect(value).To(Equal(""))
		})

		It("Set NalimovPath option", func() {
			input := []string{"setoption", "name", "NalimovPath", "value",
				"c:\\chess\\tb\\4;c:\\chess\\tb\\5\\n"}

			handler.Handle(input)

			key, value := solver.SetOptionArgsForCall(0)
			Expect(key).To(Equal("nalimovpath"))
			Expect(value).To(Equal("c:\\chess\\tb\\4;c:\\chess\\tb\\5\\n"))
		})
	})

	var _ = Describe("position", func() {
		It("Set position", func() {
			fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
			e0, e1, e2 := "e4", "Nf3", "Bb5"
			input := []string{"position", "fen", fen, "moves", e0, e1, e2}

			handler.Handle(input)

			f, p := solver.SetPositionArgsForCall(0)
			Expect(f).To(Equal(fen))
			Expect(p[0]).To(Equal(e0))
			Expect(p[1]).To(Equal(e1))
			Expect(p[2]).To(Equal(e2))
		})

		It("Set start position", func() {
			pos, e0, e1, e2 := "startpos", "e4", "Nf3", "Bb5"
			input := []string{"position", pos, "moves", e0, e1, e2}

			handler.Handle(input)

			p := solver.SetStartPositionArgsForCall(0)
			Expect(p[0]).To(Equal(e0))
			Expect(p[1]).To(Equal(e1))
			Expect(p[2]).To(Equal(e2))
		})
	})

	var _ = Describe("search", func() {
		It("search moves", func() {
			e0, e1, e2 := "e4", "Nf3", "Bb5"
			input := []string{"go", "searchmoves", e0, e1, e2}

			handler.Handle(input)

			_, a := solver.StartSearchArgsForCall(0)
			Expect(a[0]).To(Equal(e0))
			Expect(a[1]).To(Equal(e1))
			Expect(a[2]).To(Equal(e2))
		})

		It("search params", func() {
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

			handler.Handle(input)

			actual, _ := solver.StartSearchArgsForCall(0)
			Expect(*actual).
				To(Equal(*expected))
		})
	})
})
