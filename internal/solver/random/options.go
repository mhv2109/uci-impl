package random

import (
	"uci-impl/internal/solver"
)

func availableOptions() []*solver.Option {
	options := make([]*solver.Option, 1, 1)

	UCI_EngineAboutOption := &solver.Option{
		Name:    "UCI_EngineAboutOption",
		Type:    solver.OptionStringType,
		Default: "A UCI Chess engine, written in Go by mhv2109, that chooses a valid move at random"}

	options[0] = UCI_EngineAboutOption

	return options
}

func newDefaultOptions() solver.Options {
	options := solver.NewOptions()
	for _, option := range availableOptions() {
		options.Set(option.Name, option.Default)
	}
	return options
}
