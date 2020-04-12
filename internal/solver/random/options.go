package random

import (
	"uci-impl/internal/solver"
)

func availableOptions() []*solver.Option {
	options := make([]*solver.Option, 2, 2)

	HashOption := &solver.Option{
		Name:    "hash",
		Type:    solver.OptionSpinType,
		Default: "1024",
		Min:     "1024",
		Max:     "4096"}

	UCI_EngineAboutOption := &solver.Option{
		Name:    "UCI_EngineAboutOption",
		Type:    solver.OptionStringType,
		Default: "A UCI Chess engine written in Go by mhv2109"}

	options[0] = HashOption
	options[1] = UCI_EngineAboutOption

	return options
}

func newDefaultOptions() solver.Options {
	options := solver.NewOptions()
	for _, option := range availableOptions() {
		options.Set(option.Name, option.Default)
	}
	return options
}
