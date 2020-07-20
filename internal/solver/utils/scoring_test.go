package utils_test

import (
	"testing"

	"github.com/notnil/chess"
	. "github.com/onsi/gomega"

	. "github.com/mhv2109/uci-impl/internal/solver/utils"
)

func TestTakenPawnScore(t *testing.T) {
	g := NewGomegaWithT(t)

	fen, _ := chess.FEN("rnbqkbnr/ppppppp1/8/6p1/8/8/PPPPPP1P/RNBQKBNR w KQkq - 0 3")
	game := chess.NewGame(fen)
	board := game.Position().Board()

	expected, actual := PawnValue, BlackAdvantage(board)
	g.Expect(actual).
		To(Equal(expected))

	inverse := WhiteAdvantage(board)
	g.Expect(actual).
		To(Equal(-inverse))
}

func TestBeforePawnTakenScore(t *testing.T) {
	g := NewGomegaWithT(t)

	fen, _ := chess.FEN("rnbqkbnr/ppppppp1/7p/6P1/8/8/PPPPPP1P/RNBQKBNR b KQkq - 0 2")
	game := chess.NewGame(fen)
	board := game.Position().Board()

	expected, actual := CentiPawns(0), BlackAdvantage(board)
	g.Expect(actual).
		To(Equal(expected))
}
