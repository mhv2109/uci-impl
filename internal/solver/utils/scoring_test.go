package utils_test

import (
	"github.com/notnil/chess"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mhv2109/uci-impl/internal/solver/utils"
)

var _ = Describe("Score", func() {
	It("Take pawn", func() {
		fen, _ := chess.FEN("rnbqkbnr/ppppppp1/8/6p1/8/8/PPPPPP1P/RNBQKBNR w KQkq - 0 3")
		game := chess.NewGame(fen)
		board := game.Position().Board()

		expected, actual := PawnValue, BlackAdvantage(board)
		Expect(actual).
			To(Equal(expected))

		inverse := WhiteAdvantage(board)
		Expect(actual).
			To(Equal(-inverse))
	})

	It("Before take pawn", func() {
		fen, _ := chess.FEN("rnbqkbnr/ppppppp1/7p/6P1/8/8/PPPPPP1P/RNBQKBNR b KQkq - 0 2")
		game := chess.NewGame(fen)
		board := game.Position().Board()

		expected, actual := CentiPawns(0), BlackAdvantage(board)
		Expect(actual).
			To(Equal(expected))
	})
})
