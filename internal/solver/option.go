package solver

import (
	"fmt"
	"strings"
	"uci-impl/internal/config"
)

type Options config.Configuration
type OptionType string

const (
	OptionSpinType   OptionType = "spin"
	OptionStringType OptionType = "string"
)

type Option struct {
	Name    string
	Type    OptionType
	Default string
	Min     string
	Max     string
}

func NewOptions() Options {
	return config.NewConfiguration()
}

func (o *Option) String() string {
	var s strings.Builder
	s.WriteString("option")

	if o.Name != "" {
		concatOptionString(&s, "name", o.Name)
	}
	if o.Type != "" {
		concatOptionString(&s, "type", string(o.Type))
	}
	if o.Default != "" {
		concatOptionString(&s, "default", o.Default)
	}
	if o.Min != "" {
		concatOptionString(&s, "min", o.Min)
	}
	if o.Max != "" {
		concatOptionString(&s, "max", o.Max)
	}

	return s.String()
}

func concatOptionString(s *strings.Builder, name, val string) {
	s.WriteString(fmt.Sprintf(" %s %s", name, val))
}
