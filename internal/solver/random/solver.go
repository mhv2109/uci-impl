package random

import (
	"math/rand"

	"github.com/notnil/chess"

	"github.com/mhv2109/uci-impl/internal/solver"
)

type RandomSolver struct {
	base *solver.AbstractSolver
}

func NewRandomSolver() solver.Solver {
	return &RandomSolver{
		base: solver.NewAbstractSolver(newDefaultOptions())}
}

func (solver *RandomSolver) GetOption(key string) *string {
	return solver.base.GetOption(key)
}

func (solver *RandomSolver) SetOption(key, value string) {
	solver.base.SetOption(key, value)
}

func (solver *RandomSolver) GetOptions() []*solver.Option {
	return availableOptions()
}

func (solver *RandomSolver) SetPosition(pos string, moves ...string) {
	solver.base.SetPosition(pos, moves...)
}

func (solver *RandomSolver) SetStartPosition(moves ...string) {
	solver.base.SetStartPosition(moves...)
}

func (solver *RandomSolver) DoMove(move string) {
	solver.base.DoMove(move)
}

func (solver *RandomSolver) StartSearch(sp *solver.SearchParams, moves ...string) chan []string {
	solver.base.StartMove()

	ret := solver.base.GetResultCh()

	go func() {
		validMoves := solver.base.GetValidMoves(moves...)
		if move := getMove(validMoves); sp.Ponder {
			solver.base.SubmitPonderCh(move)
		} else {
			solver.base.SubmitResultCh(move)
			if !sp.Infinite {
				solver.base.CloseMove()
			}
		}
	}()

	return ret
}

func getMove(moves []*chess.Move) []string {
	return []string{moves[rand.Intn(len(moves))].String()}
}

func (solver *RandomSolver) StopSearch() {
	solver.base.CloseMove()
}

func (solver *RandomSolver) PonderHit() {
	solver.base.PonderHit()
}
