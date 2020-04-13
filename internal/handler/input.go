package handler

import (
	"os"
	"strconv"
	"strings"
	"sync"

	"uci-impl/internal/config"
	"uci-impl/internal/solver"
)

type UCIInputHandler struct {
	name, code       string
	isReadyWaitGroup *sync.WaitGroup
	solver           solver.Solver
	emitter          Emitter
}

func NewHandler(s solver.Solver) *UCIInputHandler {
	return NewHandlerWithEmitter(s, NewEmitter())
}

func NewHandlerWithEmitter(s solver.Solver, e Emitter) *UCIInputHandler {
	return &UCIInputHandler{
		isReadyWaitGroup: &sync.WaitGroup{},
		solver:           s,
		emitter:          e}
}

func (handler *UCIInputHandler) Handle(input []string) {
	// return on empty input
	if len(input) < 1 {
		return
	}

	// register goroutine
	handler.isReadyWaitGroup.Add(1)

	switch input[0] {
	case "uci":
		handler.handleUci(input)
	case "debug":
		handler.handleDebug(input)
	case "isready":
		handler.handleIsReady(input)
	case "setoption":
		handler.handleSetOption(input)
	case "register":
		handler.handleRegister(input)
	case "ucinewgame":
		handler.handleUcinewgame(input)
	case "position":
		handler.handlePosition(input)
	case "go":
		handler.handleGo(input)
	case "stop":
		handler.handleStop(input)
	case "ponderhit":
		handler.handlePonderHit(input)
	case "quit":
		handler.handleQuit(input)
	default:
		// invalid input, do nothing and return (TODO: setup logger)
	}

	handler.isReadyWaitGroup.Done()
}

/*
Handler (input) methods
*/

func (handler *UCIInputHandler) handleUci(input []string) {
	handler.emitter.EmitID()
	handler.emitter.EmitOption(handler.solver)
	handler.emitter.EmitUCIOK()
}

func (handler *UCIInputHandler) handleDebug(input []string) {
	if len(input) != 2 {
		// invalid input, do nothing and return (TODO: setup logger)
		return
	}

	arg := input[1]
	if (arg != config.DebugOff) && (arg != config.DebugOn) {
		// invalid argument, do nothing and return (TODO: setup logger)
		return
	}

	config.Config.Set(config.Debug, arg)
}

func (handler *UCIInputHandler) handleIsReady(input []string) {
	// Don't wait for this goroutine
	handler.isReadyWaitGroup.Add(-1)
	// Wait for other threads to finish
	handler.isReadyWaitGroup.Wait()

	handler.emitter.EmitReadyOK()

	// Add counter to wg to handle IsReadyWaitGroup.Done() in calling function
	handler.isReadyWaitGroup.Add(1)
}

func (handler *UCIInputHandler) handleSetOption(input []string) {
	if len(input) < 3 || input[1] != "name" {
		// invalid input, do nothing and return (TODO: setup logger)
		return
	}

	vi := -1
	for i, arg := range input {
		if arg == "value" {
			vi = i
			break
		}
	}

	var valueSlice []string
	var nameSlice []string
	if vi == -1 {
		nameSlice = input[2:]
	} else {
		nameSlice = input[2:vi]
		valueSlice = input[vi+1:]
	}
	name := strings.ToLower(strings.Join(nameSlice, " "))

	if valueSlice == nil {
		handler.solver.SetOption(name, "")
	} else {
		value := strings.Join(valueSlice, " ")
		handler.solver.SetOption(name, value)
	}

}

func (handler *UCIInputHandler) handleRegister(input []string) {
	if len(input) < 2 {
		// invalid input, do nothing and return (TODO: setup logger)
		return
	}

	if input[1] == "later" {
		// don't register engine now
		return
	}

	ni, nj := -1, -1
	for i, arg := range input {
		if arg == "name" {
			ni = i
			break
		}
	}
	if ni != -1 {
		for j, arg := range input[ni+1:] {
			if arg == "code" {
				break
			}
			nj = j
		}
	}

	ci, cj := -1, -1
	for i, arg := range input {
		if arg == "code" {
			ci = i
			break
		}
	}
	if ci != -1 {
		for j, arg := range input[ci+1:] {
			if arg == "name" {
				break
			}
			cj = j
		}
	}

	// set name
	if ni != -1 {
		handler.name = strings.Join(input[ni:nj+1], " ")
	}

	// set code
	if ci != -1 {
		handler.code = strings.Join(input[ci:cj+1], " ")
	}
}

func (handler *UCIInputHandler) handleUcinewgame(input []string) {
	// we don't currently take any specific action upon a new game
}

func (handler *UCIInputHandler) handlePosition(input []string) {
	li := len(input)
	if li < 2 {
		// invalid input, do nothing and return (TODO: setup logger)
		return
	}

	// get moves
	var moves []string
	for i, v := range input {
		if v == "moves" {
			moves = make([]string, len(input)-i-1)
			for j, move := range input[i+1:] {
				moves[j] = move
			}
			break
		}
	}

	// set positions
	switch input[1] {
	case "startpos":
		handler.solver.SetStartPosition(moves...)
	case "fen":
		if li < 3 {
			// invalid input, do nothing and return (TODO: setup logger)
			return
		}
		handler.solver.SetPosition(input[2], moves...)
	default:
		// invalid input, do nothing and return (TODO: setup logger)
		return
	}
}

func (handler *UCIInputHandler) handleGo(input []string) {
	sp := solver.NewSearchParams()

	var searchmoves []string
	for i, v := range input {
		switch v {
		case "searchmoves":
			searchmoves = make([]string, len(input)-i-1)
			for j, searchmove := range input[i+1:] {
				searchmoves[j] = searchmove
			}
			break
		case "ponder":
			sp.Ponder = true
		case "wtime":
			if wtime, err := strconv.Atoi(input[i+1]); err == nil {
				sp.Wtime = wtime
			}
		case "btime":
			if btime, err := strconv.Atoi(input[i+1]); err == nil {
				sp.Btime = btime
			}
		case "winc":
			if winc, err := strconv.Atoi(input[i+1]); err == nil {
				sp.Winc = winc
			}
		case "binc":
			if binc, err := strconv.Atoi(input[i+1]); err == nil {
				sp.Binc = binc
			}
		case "movestogo":
			if movestogo, err := strconv.Atoi(input[i+1]); err == nil {
				sp.Movestogo = movestogo
			}
		case "depth":
			if depth, err := strconv.Atoi(input[i+1]); err == nil {
				sp.Depth = depth
			}
		case "mate":
			if mate, err := strconv.Atoi(input[i+1]); err == nil {
				sp.Mate = mate
			}
		case "movetime":
			if movetime, err := strconv.Atoi(input[i+1]); err == nil {
				sp.Movetime = movetime
			}
		case "infinite":
			sp.Infinite = true
		}
	}

	// Start solver and return move
	/* TODO: should the solver only implement StartSearch/StopSearch and
	move Pondering logic to input/output handler? */
	ch := handler.solver.StartSearch(sp, searchmoves...)
	go func() {
		var result []string
		for result = range ch {
		}

		// in event of "pondermiss", result will be nil
		if result != nil {
			handler.emitter.EmitBestmove(result...)
		}
	}()
}

func (handler *UCIInputHandler) handleStop(input []string) {
	handler.solver.StopSearch()
}

func (handler *UCIInputHandler) handlePonderHit(input []string) {
	handler.solver.PonderHit()
}

func (handler *UCIInputHandler) handleQuit(input []string) {
	os.Exit(0)
}
