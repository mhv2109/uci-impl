package utils

import (
	"github.com/notnil/chess"
)

type CentiPawns int

const nSquares = 64

const (
	PawnValue   CentiPawns = 100
	KnightValue CentiPawns = 300
	BishopValue CentiPawns = 300
	RookValue   CentiPawns = 500
	QueenValue  CentiPawns = 900
	KingValue   CentiPawns = 1000000
)

var MaxScore CentiPawns = 1003900

func BlackAdvantage(board *chess.Board) CentiPawns {
	white, black := getAllPiecesByColor(board)
	return score(black) - score(white)
}

func WhiteAdvantage(board *chess.Board) CentiPawns {
	white, black := getAllPiecesByColor(board)
	return score(white) - score(black)
}

func score(s []chess.Piece) CentiPawns {
	score := CentiPawns(0)
	for _, piece := range s {
		score += scorePiece(piece)
	}
	return score
}

func scorePiece(piece chess.Piece) CentiPawns {
	var ret CentiPawns
	switch piece.Type() {
	case chess.Pawn:
		ret = PawnValue
	case chess.Knight:
		ret = KnightValue
	case chess.Bishop:
		ret = BishopValue
	case chess.Rook:
		ret = RookValue
	case chess.Queen:
		ret = QueenValue
	case chess.King:
		ret = KingValue
	default:
		panic("Invalid type")
	}

	return ret
}

func getAllPiecesByColor(board *chess.Board) (white, black []chess.Piece) {
	white, black = make([]chess.Piece, 0, 16), make([]chess.Piece, 0, 16)
	for sq := chess.Square(0); sq < nSquares; sq++ {
		if piece := board.Piece(sq); piece != chess.NoPiece {
			if color := piece.Color(); color == chess.White {
				white = append(white, piece)
			} else {
				black = append(black, piece)
			}
		}
	}
	return white, black
}
