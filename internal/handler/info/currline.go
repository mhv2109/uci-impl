package info

import (
	"fmt"
	"strings"
)

// currline is the current line the engine is calculating.
type currline struct {
	cpunr int // the number of the cpu if the engine is running on more than one cpu
	moves []string
}

func newCurrline(cpunr int, moves ...string) *currline {
	return &currline{cpunr, moves}
}

func (c *currline) String() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("currline %d", c.cpunr))
	for _, move := range c.moves {
		builder.WriteString(fmt.Sprintf(" %s", move))
	}

	return builder.String()
}
