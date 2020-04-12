package info

import (
	"strings"
	"fmt"
)

type currline struct {
	cpunr uint
	moves []string
}

func newCurrline(cpunr uint, moves ...string) *currline {
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
