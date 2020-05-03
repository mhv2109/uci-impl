package solver

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
