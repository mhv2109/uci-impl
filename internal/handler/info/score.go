package info

import "fmt"

type ScoreType string

const (
	CP         ScoreType = "cp"
	Mate       ScoreType = "mate"
	Lowerbound ScoreType = "lowerbound"
	Upperbound ScoreType = "upperbound"
)

type score struct {
	scoretype ScoreType
	value     int
}

func newScore(scoretype ScoreType, value int) *score {
	return &score{scoretype, value}
}

func (s *score) String() string {
	return fmt.Sprintf("score %s %d", s.scoretype, s.value)
}
