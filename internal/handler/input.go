package handler

import (
	"os"
	"strconv"
	"strings"

	"uci-impl/internal/config"
	"uci-impl/internal/solver"
)

// UCIInputHandler is responsible for processing commands from the GUI, and
// coordinating those commands to the Solver implementation, the global progam
// state, and ultimately back to the GUI.
type UCIInputHandler struct {
	solver  solver.Solver
	emitter Emitter
}

// NewHandler returns an instance of UCIInputHandler, given a Solver
// implementation.
func NewHandler(s solver.Solver) *UCIInputHandler {
	return NewHandlerWithEmitter(s, NewEmitter())
}

// NewHandlerWithEmitter returns an instance of UCIInputHandler, given a
// Solver and an Emitter implementation.
func NewHandlerWithEmitter(s solver.Solver, e Emitter) *UCIInputHandler {
	return &UCIInputHandler{
		solver:  s,
		emitter: e}
}

// Handle processes commands from the GUI to the engine.
func (handler *UCIInputHandler) Handle(input []string) {
	// return on empty input
	if len(input) < 1 {
		return
	}

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

}

// uci
// Tell engine to use the uci (universal chess interface),
// this will be sent once as a first command after program boot
// to tell the engine to switch to uci mode.
// After receiving the uci command the engine must identify itself with the "id" command
// and send the "option" commands to tell the GUI which engine settings the engine supports if any.
// After that the engine should send "uciok" to acknowledge the uci mode.
// If no uciok is sent within a certain time period, the engine task will be killed by the GUI.
func (handler *UCIInputHandler) handleUci(input []string) {
	handler.emitter.EmitID()
	handler.emitter.EmitOption(handler.solver)
	handler.emitter.EmitUCIOK()
	handler.checkCopyprotection()
}

func (handler *UCIInputHandler) checkCopyprotection() {
	// no copyprotection implemented
	handler.emitter.EmitCopyProtectionChecking()
	handler.emitter.EmitCopyProtectionOk()
}

// debug [ on | off ]
// Switch the debug mode of the engine on and off.
// In debug mode the engine should send additional infos to the GUI, e.g. with the "info string" command,
// to help debugging, e.g. the commands that the engine has received etc.
// This mode should be switched off by default and this command can be sent
// any time, also when the engine is thinking.
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

// isready
// This is used to synchronize the engine with the GUI. When the GUI has sent a command or
// multiple commands that can thandleInputake some time to complete,
// this command can be used to wait for the engine to be ready again or
// to ping the engine to find out if it is still alive.
// E.g. this should be sent after setting the path to the tablebases as this can take some time.
// This command is also required once before the engine is asked to do any search
// to wait for the engine to finish initializing.
// This command must always be answered with "readyok" and can be sent also when the engine is calculating
// in which case the engine should also immediately answer with "readyok" without stopping the search.
func (handler *UCIInputHandler) handleIsReady(input []string) {
	handler.emitter.EmitReadyOK()
}

// setoption name <id> [value <x>]
// This is sent to the engine when the user wants to change the internal parameters
// of the engine. For the "button" type no value is needed.
// One string will be sent for each parameter and this will only be sent when the engine is waiting.
// The name and value of the option in <id> should not be case sensitive and can inlude spaces.
// The substrings "value" and "name" should be avoided in <id> and <x> to allow unambiguous parsing,
// for example do not use <name> = "draw value".
// Here are some strings for the example below:
// 	   "setoption name Nullmove value true\n"
//         "setoption name Selectivity value 3\n"
// 	   "setoption name Style value Risky\n"
// 	   "setoption name Clear Hash\n"
// 	   "setoption name NalimovPath value c:\chess\tb\4;c:\chess\tb\5\n"
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

// register
// This is the command to try to register an engine or to tell the engine that registration
// will be done later. This command should always be sent if the engine	has sent "registration error"
// at program startup.
// The following tokens are allowed:
// * later
//    the user doesn't want to register the engine now.
// * name <x>
//    the engine should be registered with the name <x>
// * code <y>
//    the engine should be registered with the code <y>
// Example:
//    "register later"
//    "register name Stefan MK code 4359874324"
func (handler *UCIInputHandler) handleRegister(input []string) {
	if len(input) < 2 || input[1] == "later" {
		return
	}
	// don't implement any form of registration checking
	handler.emitter.EmitRegistrationChecking()
	handler.emitter.EmitRegistrationOk()
}

// ucinewgame
// This is sent to the engine when the next search (started with "position" and "go") will be from
// a different game. This can be a new game the engine should play or a new game it should analyse but
// also the next position from a testsuite with positions only.
// If the GUI hasn't sent a "ucinewgame" before the first "position" command, the engine shouldn't
// expect any further ucinewgame commands as the GUI is probably not supporting the ucinewgame command.
// So the engine should not rely on this command even though all new GUIs should support it.
// As the engine's reaction to "ucinewgame" can take some time the GUI should always send "isready"
// after "ucinewgame" to wait for the engine to finish its operation.
func (handler *UCIInputHandler) handleUcinewgame(input []string) {
	// we don't currently take any specific action upon a new game
}

// position [fen <fenstring> | startpos ]  moves <move1> .... <movei>
// Set up the position described in fenstring on the internal board and
// play the moves on the internal chess board.
// if the game was played  from the start position the string "startpos" will be sent
// Note: no "new" command is needed. However, if this position is from a different game than
// the last position sent to the engine, the GUI should have sent a "ucinewgame" inbetween.
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

// go
// Start calculating on the current position set up with the "position" command.
// There are a number of commands that can follow this command, all will be sent in the same string.
// If one command is not sent its value should be interpreted as it would not influence the search.
// * searchmoves <move1> .... <movei>
// 	restrict search to this moves only
// 	Example: After "position startpos" and "go infinite searchmoves e2e4 d2d4"
// 	the engine should only search the two moves e2e4 and d2d4 in the initial position.
// * ponder
// 	start searching in pondering mode.
// 	Do not exit the search in ponder mode, even if it's mate!
// 	This means that the last move sent in in the position string is the ponder move.
// 	The engine can do what it wants to do, but after a "ponderhit" command
// 	it should execute the suggested move to ponder on. This means that the ponder move sent by
// 	the GUI can be interpreted as a recommendation about which move to ponder. However, if the
// 	engine decides to ponder on a different move, it should not display any mainlines as they are
// 	likely to be misinterpreted by the GUI because the GUI expects the engine to ponder
//      on the suggested move.
// * wtime <x>
// 	white has x msec left on the clock
// * btime <x>
// 	black has x msec left on the clock
// * winc <x>
// 	white increment per move in mseconds if x > 0
// * binc <x>
// 	black increment per move in mseconds if x > 0
// * movestogo <x>
//      there are x moves to the next time control,
// 	this will only be sent if x > 0,
// 	if you don't get this and get the wtime and btime it's sudden death
// * depth <x>
// 	search x plies only.
// * nodes <x>
//      search x nodes only,
// * mate <x>
// 	search for a mate in x moves
// * movetime <x>
// 	search exactly x mseconds
// * infinite
// 	search until the "stop" command. Do not exit the search without being told so in this mode!
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

// stop
// Stop calculating as soon as possible,
// don't forget the "bestmove" and possibly the "ponder" token when finishing the search.
func (handler *UCIInputHandler) handleStop(input []string) {
	handler.solver.StopSearch()
}

// ponderhit
// The user has played the expected move. This will be sent if the engine was told to ponder on the same move
// the user has played. The engine should continue searching but switch from pondering to normal search.
func (handler *UCIInputHandler) handlePonderHit(input []string) {
	handler.solver.PonderHit()
}

// quit
// Quit the program as soon as possible.
func (handler *UCIInputHandler) handleQuit(input []string) {
	os.Exit(0)
}
