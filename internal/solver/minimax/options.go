package minimax

import (
	"github.com/mhv2109/uci-impl/internal/solver"
)

func availableOptions() []*solver.Option {
	options := make([]*solver.Option, 3, 3)

	UCI_EngineAboutOption := &solver.Option{
		Name:    "UCI_EngineAboutOption",
		Type:    solver.OptionStringType,
		Default: "A UCI Chess engine, written in Go by mhv2109, uses a Minimax algorithm with Alpha-Beta pruning"}

	HashOption := &solver.Option{
		Name:    "Hash",
		Type:    solver.OptionSpinType,
		Default: "32",
		Min:     "1",
		Max:     "4096"}

	DepthOption := &solver.Option{
		Name:    "Search Depth",
		Type:    solver.OptionSpinType,
		Default: "2",
		Min:     "1",
		Max:     "4"}

	options[0] = UCI_EngineAboutOption
	options[1] = HashOption
	options[2] = DepthOption

	return options
}

func newDefaultOptions() solver.Options {
	options := solver.NewOptions()
	for _, option := range availableOptions() {
		options.Set(option.Name, option.Default)
	}
	return options
}
