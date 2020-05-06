package random

import (
	"log"
	"math/rand"
	"sync"

	"github.com/notnil/chess"

	"github.com/mhv2109/uci-impl/internal/solver"
)

type RandomSolver struct {
	options       solver.Options
	game          *chess.Game
	resultCh      chan []string
	resultChMutex sync.RWMutex
	ponderCh      chan []string
	ponderChMutex sync.RWMutex
}

func NewRandomSolver() solver.Solver {
	return &RandomSolver{
		options: newDefaultOptions(),
		game:    chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))}
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
	solver.game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))
	solver.doMoves(moves...)
}

func (solver *RandomSolver) doMoves(moves ...string) {
	for _, m := range moves {
		solver.DoMove(m)
	}
}

func (solver *RandomSolver) SetStartPosition(moves ...string) {
	solver.game = chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))
	solver.doMoves(moves...)
}

func (solver *RandomSolver) DoMove(move string) {
	if err := solver.game.MoveStr(move); err != nil {
		log.Panicln("Invalid move: " + move)
	}
}

func (solver *RandomSolver) StartSearch(sp *solver.SearchParams, moves ...string) chan []string {
	solver.closeMove()
	solver.startMove()

	ret := solver.resultCh

	go func() {
		validMoves := solver.getValidMoves(moves...)
		if move := getMove(validMoves); sp.Ponder {
			solver.submitPonderCh(move)
		} else {
			solver.submitResultCh(move)
			if !sp.Infinite {
				solver.closeMove()
			}
		}
	}()

	return ret
}

func (solver *RandomSolver) getValidMoves(moves ...string) []*chess.Move {
	if len(moves) == 0 {
		return solver.game.ValidMoves()
	} else {
		return solver.decodeAlgNotations(moves...)
	}
}

func (solver *RandomSolver) decodeAlgNotation(moveStr string) (*chess.Move, error) {
	position := solver.game.Position()
	notation := chess.LongAlgebraicNotation{}
	return notation.Decode(position, moveStr)
}

func (solver *RandomSolver) decodeAlgNotations(movesStr ...string) []*chess.Move {
	moves := make([]*chess.Move, 0, len(movesStr))
	for _, move := range movesStr {
		if vm, err := solver.decodeAlgNotation(move); err == nil {
			moves = append(moves, vm)
		}
	}

	return moves
}

func getMove(moves []*chess.Move) []string {
	return []string{moves[rand.Intn(len(moves))].String()}
}

func (solver *RandomSolver) StopSearch() {
	solver.closeMove()
}

func (solver *RandomSolver) PonderHit() {
	if solver.ponderCh == nil || solver.resultCh == nil {
		return
	}

	var result []string
	if len(solver.ponderCh) == 0 {
		result = <-solver.ponderCh
		solver.closePonderCh()
	} else {
		ponderCh := solver.ponderCh
		solver.closePonderCh()
		for result = range ponderCh {
		}
	}

	// get best ponder move and return result
	if result != nil {
		solver.resultCh <- result
	}

	solver.closeResultCh()
}

func (solver *RandomSolver) startMove() {
	solver.setupResultCh()
	solver.setupPonderCh()
}

func (solver *RandomSolver) closeMove() {
	solver.closeResultCh()
	solver.closePonderCh()
}

func (solver *RandomSolver) setupResultCh() {
	solver.resultChMutex.Lock()
	defer solver.resultChMutex.Unlock()

	if solver.resultCh == nil {
		solver.resultCh = make(chan []string, 1)
	}
}

func (solver *RandomSolver) submitResultCh(move []string) bool {
	solver.resultChMutex.RLock()
	defer solver.resultChMutex.RUnlock()

	if solver.resultCh != nil {
		solver.resultCh <- move
		return true
	} else {
		return false
	}
}

func (solver *RandomSolver) closeResultCh() {
	solver.resultChMutex.Lock()
	defer solver.resultChMutex.Unlock()

	if solver.resultCh != nil {
		close(solver.resultCh)
		solver.resultCh = nil
	}
}

func (solver *RandomSolver) setupPonderCh() {
	solver.ponderChMutex.Lock()
	defer solver.ponderChMutex.Unlock()

	if solver.ponderCh == nil {
		solver.ponderCh = make(chan []string, 1)
	}
}

func (solver *RandomSolver) submitPonderCh(move []string) bool {
	solver.ponderChMutex.RLock()
	defer solver.ponderChMutex.RUnlock()

	if solver.ponderCh != nil {
		solver.ponderCh <- move
		return true
	} else {
		return false
	}
}

func (solver *RandomSolver) closePonderCh() {
	solver.ponderChMutex.Lock()
	defer solver.ponderChMutex.Unlock()

	if solver.ponderCh != nil {
		close(solver.ponderCh)
		solver.ponderCh = nil
	}
}
