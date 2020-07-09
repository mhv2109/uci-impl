package solver

import (
	"log"
	"sync"

	"github.com/notnil/chess"
)

// SearchParams is a struct that holds values for commands that follow the "Go"
// command in the UCI protocol.
type SearchParams struct {
	// start searching in pondering mode.
	// Do not exit the search in ponder mode, even if it's mate!
	// This means that the last move sent in in the position string is the ponder move.
	// The engine can do what it wants to do, but after a "ponderhit" command
	// it should execute the suggested move to ponder on. This means that the ponder move sent by
	// the GUI can be interpreted as a recommendation about which move to ponder. However, if the
	// engine decides to ponder on a different move, it should not display any mainlines as they are
	// likely to be misinterpreted by the GUI because the GUI expects the engine to ponder on the suggested move.
	Ponder bool
	Wtime  int // white has x msec left on the clock
	Btime  int // black has x msec left on the clock
	Winc   int // white increment per move in mseconds if x > 0
	Binc   int // black increment per move in mseconds if x > 0
	// there are x moves to the next time control,
	// this will only be sent if x > 0,
	// if you don't get this and get the wtime and btime it's sudden death
	Movestogo int
	Depth     int  // search x plies only.
	Nodes     int  // search x nodes only,
	Mate      int  // search for a mate in x moves
	Movetime  int  // search exactly x mseconds
	Infinite  bool // search until the "stop" command. Do not exit the search without being told so in this mode!
}

// NewSearchParams returns a pointer to a SearchParams with default parameters
func NewSearchParams() *SearchParams {
	return &SearchParams{
		false, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, false}
}

// Solver is the interface that must be implemented to be accepted and
// coordinated by the Handler.
type Solver interface {
	GetOption(string) *string      // get Option value, returns nil if not present
	SetOption(string, string)      // set Option value, as a string regardless of interpreted type
	GetOptions() []*Option         // get all Options available for the Solver implementation
	SetPosition(string, ...string) // set game position with FEN string & individual moves in Long-Algebraic format
	SetStartPosition(...string)    // set game position at "start", plus individual moves in Long-Algebraic format
	DoMove(string)                 // do an individual move in Long-Algebraic format
	// Start searching asynchronously, and put results on the returned channel.
	// The search algorithm can place the "best current move" on the channel
	// as they are found.  When StopSearch is called, or the time limit
	// reached, the last result on the channel is interpreted as the best
	// move.  The first element in the result slice is the selected best
	// move, and the following entries are moves the engine plans to ponder
	// on.
	StartSearch(*SearchParams, ...string) chan []string
	StopSearch() // end current running search
	PonderHit()  // signal that opponent made the move the current search is solving for (in pondering mode)
}

// AbstractSolver is a base solver boilerplate to remove some of the repetition.
type AbstractSolver struct {
	Options  Options
	Game     *chess.Game
	Notation chess.Notation

	resultCh      chan []string
	resultChMutex sync.RWMutex
	ponderCh      chan []string
	ponderChMutex sync.RWMutex
}

// NewAbstractSolver returns a pointer to a new initialized AbstractSolver.
func NewAbstractSolver(options Options) *AbstractSolver {
	notation := chess.LongAlgebraicNotation{}
	return &AbstractSolver{
		Options:  options,
		Game:     chess.NewGame(chess.UseNotation(notation)),
		Notation: notation}
}

// GetOption gets Option value, returns nil if not present.
func (solver *AbstractSolver) GetOption(key string) *string {
	return solver.Options.Get(key)
}

// SetOption sets Option value, as a string, regardless of interpreted type.
func (solver *AbstractSolver) SetOption(key, value string) {
	solver.Options.Set(key, value)
}

// SetPosition sets game position with FEN string & individual moves in Long-Algebraic format.
func (solver *AbstractSolver) SetPosition(pos string, moves ...string) {
	fen, err := chess.FEN(pos)
	if err != nil {
		log.Panicln(err)
	}
	solver.Game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))
	solver.doMoves(moves...)
}

func (solver *AbstractSolver) doMoves(moves ...string) {
	for _, m := range moves {
		solver.DoMove(m)
	}
}

// DoMove applies an individual move to the Game in Long-Algebraic format.
func (solver *AbstractSolver) DoMove(move string) {
	if err := solver.Game.MoveStr(move); err != nil {
		log.Panicln("Invalid move: " + move)
	}
}

// SetStartPosition sets game position at "start", plus applies individual moves in Long-Algebraic format using DoMove.
func (solver *AbstractSolver) SetStartPosition(moves ...string) {
	solver.Game = chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))
	solver.doMoves(moves...)
}

// GetValidMoves returns all valid moves for the current Game state.
func (solver *AbstractSolver) GetValidMoves(moves ...string) []*chess.Move {
	if len(moves) == 0 {
		return solver.Game.ValidMoves()
	}
	return solver.decodeAlgNotations(moves...)
}

func (solver *AbstractSolver) decodeAlgNotations(movesStr ...string) []*chess.Move {
	moves := make([]*chess.Move, 0, len(movesStr))
	for _, move := range movesStr {
		if vm, err := solver.decodeAlgNotation(move); err == nil {
			moves = append(moves, vm)
		}
	}
	return moves
}

func (solver *AbstractSolver) decodeAlgNotation(moveStr string) (*chess.Move, error) {
	position := solver.Game.Position()
	return solver.Notation.Decode(position, moveStr)
}

// StartMove prepares result and ponder channels for communicating search results.
func (solver *AbstractSolver) StartMove() {
	solver.CloseMove()
	solver.setupResultCh()
	solver.setupPonderCh()
}

func (solver *AbstractSolver) setupResultCh() {
	solver.resultChMutex.Lock()
	defer solver.resultChMutex.Unlock()

	if solver.resultCh == nil {
		solver.resultCh = make(chan []string, len(solver.Game.ValidMoves()))
	}
}

func (solver *AbstractSolver) setupPonderCh() {
	solver.ponderChMutex.Lock()
	defer solver.ponderChMutex.Unlock()

	if solver.ponderCh == nil {
		solver.ponderCh = make(chan []string, len(solver.Game.ValidMoves()))
	}
}

// CloseMove tears down the result and ponder channels.
func (solver *AbstractSolver) CloseMove() {
	solver.closeResultCh()
	solver.closePonderCh()
}

func (solver *AbstractSolver) closeResultCh() {
	solver.resultChMutex.Lock()
	defer solver.resultChMutex.Unlock()

	if solver.resultCh != nil {
		close(solver.resultCh)
		solver.resultCh = nil
	}
}

func (solver *AbstractSolver) closePonderCh() {
	solver.ponderChMutex.Lock()
	defer solver.ponderChMutex.Unlock()

	if solver.ponderCh != nil {
		close(solver.ponderCh)
		solver.ponderCh = nil
	}
}

// SubmitResultCh submits search result to result channel.  Returns true if move was successfully subitted, false otherwise.
func (solver *AbstractSolver) SubmitResultCh(move []string) bool {
	solver.resultChMutex.RLock()
	defer solver.resultChMutex.RUnlock()

	if solver.resultCh != nil {
		solver.resultCh <- move
		return true
	}
	return false
}

// SubmitPonderCh submits search result to ponder channel.  Returns true if move was successfully subitted, false otherwise.
func (solver *AbstractSolver) SubmitPonderCh(move []string) bool {
	solver.ponderChMutex.RLock()
	defer solver.ponderChMutex.RUnlock()

	if solver.ponderCh != nil {
		solver.ponderCh <- move
		return true
	}
	return false
}

// PonderHit signals that opponent made the move the current search is solving for (in pondering mode).
func (solver *AbstractSolver) PonderHit() {
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

// GetResultCh returns the final result channel to share with UCIHandler.
func (solver *AbstractSolver) GetResultCh() chan []string {
	return solver.resultCh
}
