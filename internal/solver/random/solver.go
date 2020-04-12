package random

import (
	"log"
	"math/rand"

	"github.com/notnil/chess"

	"uci-impl/internal/solver"
)

type RandomSolver struct {
	options  solver.Options
	game     *chess.Game
	resultCh chan []string
	ponderCh chan []string
}

func NewRandomSolver() solver.Solver {
	return &RandomSolver{
		options: newDefaultOptions(),
		game:    chess.NewGame()}
}

func (solver *RandomSolver) GetOption(key string) *string {
	return solver.options.Get(key)
}

func (solver *RandomSolver) SetOption(key, value string) {
	solver.options.Set(key, value)
}

func (solver *RandomSolver) GetOptions() []*solver.Option {
	return availableOptions()
}

func (solver *RandomSolver) SetPosition(pos string, moves ...string) {
	fen, err := chess.FEN(pos)
	if err != nil {
		log.Panicln(err)
	}
	solver.game = chess.NewGame(fen)
	solver.doMoves(moves...)
}

func (solver *RandomSolver) doMoves(moves ...string) {
	for _, m := range moves {
		solver.DoMove(m)
	}
}

func (solver *RandomSolver) SetStartPosition(moves ...string) {
	solver.game = chess.NewGame()
	solver.doMoves(moves...)
}

func (solver *RandomSolver) DoMove(move string) {
	solver.game.MoveStr(move)
}

func (solver *RandomSolver) StartSearch(sp *solver.SearchParams, moves ...string) chan []string {
	solver.closeMove()

	solver.resultCh = make(chan []string, 1)
	var validMoves []*chess.Move
	if len(moves) == 0 {
		validMoves = solver.game.ValidMoves()
	} else {
		validMoves = solver.decodeAlgNotations(moves...)
	}

	if sp.Ponder {
		solver.ponderCh = make(chan []string, 1)
		go solver.search(sp, solver.ponderCh, validMoves)
	} else {
		go solver.search(sp, solver.resultCh, validMoves)
	}

	return solver.resultCh
}

func (solver *RandomSolver) search(sp *solver.SearchParams, ch chan []string, moves []*chess.Move) {
	bestMove := moves[rand.Intn(len(moves))]
	if res := []string{bestMove.String()}; sp.Infinite || sp.Ponder {
		ch <- res
		solver.closePonder()
	} else {
		ch <- res
		solver.closeMove()
	}
}

func (solver *RandomSolver) decodeAlgNotation(moveStr string) (*chess.Move, error) {
	position := solver.game.Position()
	notation := chess.AlgebraicNotation{}
	return notation.Decode(position, moveStr)
}

func (solver *RandomSolver) decodeAlgNotations(movesStr ...string) []*chess.Move {
	moves := make([]*chess.Move, len(movesStr))
	for i, move := range movesStr {
		if vm, err := solver.decodeAlgNotation(move); err == nil {
			moves[i] = vm
		}
	}

	return moves
}

func (solver *RandomSolver) StopSearch() {
	solver.closeMove()
}

func (solver *RandomSolver) closeMove() {
	solver.closeSearch()
	solver.closePonder()
}

func (solver *RandomSolver) closeSearch() {
	if solver.resultCh != nil {
		close(solver.resultCh)
		solver.resultCh = nil
	}
}

func (solver *RandomSolver) closePonder() {
	if solver.ponderCh != nil {
		close(solver.ponderCh)
		solver.ponderCh = nil
	}
}

func (solver *RandomSolver) PonderHit() {
	if solver.ponderCh == nil {
		return
	}

	// get best ponder move and return result
	var result []string
	for result = range solver.ponderCh {
	}

	if result != nil {
		solver.resultCh <- result
	}

	solver.closeMove()
}
