package info

import "fmt"

type ScoreType string

// ScoreType constants identify the scoring method.
const (
	CP         ScoreType = "cp"         // the score from the engine's point of view in centipawns.
	Mate       ScoreType = "mate"       // mate in y moves, not plies. If the engine is getting mated use negative values for y.
	Lowerbound ScoreType = "lowerbound" // the score is just a lower bound.
	Upperbound ScoreType = "upperbound" // the score is just an upper bound.
)

// score struct contains scoring metadata to share with the GUI.
type score struct {
	scoretype ScoreType // what the value signifies
	value     int       // score value
}

func newScore(scoretype ScoreType, value int) *score {
	return &score{scoretype, value}
}

func (s *score) String() string {
	return fmt.Sprintf("score %s %d", s.scoretype, s.value)
}
