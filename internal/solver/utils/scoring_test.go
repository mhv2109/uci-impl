package utils

import (
	"testing"

	"github.com/notnil/chess"
)

func TestTakenPawnScore(t *testing.T) {
	fen, _ := chess.FEN("rnbqkbnr/ppppppp1/8/6p1/8/8/PPPPPP1P/RNBQKBNR w KQkq - 0 3")
	game := chess.NewGame(fen)
	board := game.Position().Board()

	if expected, actual := PawnValue, BlackAdvantage(board); expected != actual {
		t.Errorf("Expected %d, actual %d", expected, actual)
	} else if inverse := WhiteAdvantage(board); actual != -inverse {
		t.Errorf("Expected %d, actual %d", -actual, inverse)
	}
}

func TestBeforePawnTakenScore(t *testing.T) {
	fen, _ := chess.FEN("rnbqkbnr/ppppppp1/7p/6P1/8/8/PPPPPP1P/RNBQKBNR b KQkq - 0 2")
	game := chess.NewGame(fen)
	board := game.Position().Board()

	if expected, actual := CentiPawns(0), BlackAdvantage(board); expected != actual {
		t.Errorf("Expected %d, actual %d", expected, actual)
	}
}
