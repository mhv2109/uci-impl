package solver

import (
	"fmt"
	"strings"

	"github.com/mhv2109/uci-impl/internal/config"
)

// Options is an alias for Configuration.
type Options config.Configuration

// OptionType is an alias for string, supported OptionTypes are below.
type OptionType string

// Supported OptionsTypes
const (
	// a checkbox that can either be true or false
	OptionCheckType OptionType = "check"

	// a spin wheel that can be an integer in a certain range
	OptionSpinType OptionType = "spin"

	// a combo box that can have different predefined strings as a value
	OptionComboType OptionType = "combo"

	// a button that can be pressed to send a command to the engine
	OptionButtonType OptionType = "button"

	// a text field that has a string as a value, an empty string has the
	// value "<empty>"
	OptionStringType OptionType = "string"
)

// Option struct contains all the metadata for a particular option accepted by
// the Engine.
type Option struct {
	Name    string
	Type    OptionType
	Default string
	Min     string
	Max     string
	Vars    []string
}

// NewOptions returns a new instance of Options.
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
	for _, v := range o.Vars {
		concatOptionString(&s, "var", v)
	}

	return s.String()
}

func concatOptionString(s *strings.Builder, name, val string) {
	s.WriteString(fmt.Sprintf(" %s %s", name, val))
}
