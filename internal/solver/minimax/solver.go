package minimax

import (
	"log"
	"strconv"
	"time"

	"github.com/mhv2109/uci-impl/internal/handler"
	"github.com/mhv2109/uci-impl/internal/solver"
	"github.com/notnil/chess"
)

type MinimaxSolver struct {
	base *solver.AbstractSolver

	emitter handler.Emitter
	algo    *minimaxAlgo
}

func NewMinimaxSolver() solver.Solver {
	return &MinimaxSolver{
		base:    solver.NewAbstractSolver(newDefaultOptions()),
		emitter: handler.NewEmitter()}
}

func NewMinimaxSolverWithEmitter(emitter handler.Emitter) solver.Solver {
	return &MinimaxSolver{
		base:    solver.NewAbstractSolver(newDefaultOptions()),
		emitter: emitter}
}

func (solver *MinimaxSolver) GetOption(key string) *string {
	return solver.base.GetOption(key)
}

func (solver *MinimaxSolver) SetOption(key, value string) {
	solver.base.SetOption(key, value)
}

func (solver *MinimaxSolver) GetOptions() []*solver.Option {
	return availableOptions()
}

func (solver *MinimaxSolver) getHashSize() int {
	return solver.optionToInt("Hash", 32)
}

func (solver *MinimaxSolver) getDepth() int {
	return solver.optionToInt("Search Depth", 2)
}

func (solver *MinimaxSolver) optionToInt(name string, def int) int {
	opt := solver.GetOption(name)
	if opt == nil {
		return def
	}

	i, e := strconv.Atoi(*opt)
	if e != nil {
		log.Fatalf("Error casting option %s: %s", name, e)
	}
	return i
}

func (solver *MinimaxSolver) SetPosition(pos string, moves ...string) {
	solver.base.SetPosition(pos, moves...)
}

func (solver *MinimaxSolver) SetStartPosition(moves ...string) {
	solver.base.SetStartPosition(moves...)
}

func (solver *MinimaxSolver) DoMove(move string) {
	solver.base.DoMove(move)
}

func (solver *MinimaxSolver) StartSearch(sp *solver.SearchParams, moves ...string) chan []string {
	solver.base.StartMove()

	ret := solver.base.GetResultCh()

	go solver.minimax(sp, moves...)
	if !sp.Infinite && !sp.Ponder {
		var duration int
		if movetime := sp.Movetime; movetime > 0 {
			duration = movetime
		} else {
			if turn := solver.base.Game.Position().Turn(); turn == chess.White && sp.Wtime > 0 {
				duration = sp.Wtime
			} else if turn == chess.Black && sp.Btime > 0 {
				duration = sp.Btime
			}
		}

		if duration > 0 {
			go func() {
				// TODO: represent SearchParams.Movetime as time.Duration
				time.Sleep(time.Duration(duration) * time.Millisecond)
				solver.base.CloseMove()
			}()
		}
	}

	return ret
}

func (solver *MinimaxSolver) minimax(sp *solver.SearchParams, moves ...string) {
	submit := func(move []string) bool {
		if sp.Ponder {
			return solver.base.SubmitPonderCh(move)
		}
		return solver.base.SubmitResultCh(move)
	}

	var depth int
	if depthConfig := solver.getDepth(); depth < 1 || depth > depthConfig {
		depth = depthConfig
	}

	if hashSize := solver.getHashSize(); solver.algo == nil || solver.algo.MaxDepth != depth || solver.algo.HashSize != hashSize {
		solver.algo = newMinimaxAlgo(depth, hashSize, submit, solver.emitter)
	}

	solver.algo.Start(solver.base.Game.Position(), solver.base.GetValidMoves(moves...)...)
	solver.base.CloseMove()
}

func (solver *MinimaxSolver) StopSearch() {
	solver.base.CloseMove()
}

func (solver *MinimaxSolver) PonderHit() {
	solver.base.PonderHit()
}
