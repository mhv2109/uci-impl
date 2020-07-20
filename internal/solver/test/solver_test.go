package solver

import (
	"testing"

	. "github.com/mhv2109/uci-impl/internal/solver"
	. "github.com/onsi/gomega"
)

func TestSpinOptionToString(t *testing.T) {
	g := NewGomegaWithT(t)

	option := Option{}
	option.Name = "testoption"
	option.Type = OptionSpinType
	option.Default = "50"
	option.Min = "0"
	option.Max = "100"

	g.Expect(option.String()).
		To(Equal("option name testoption type spin default 50 min 0 max 100"))
}

func TestStringOptionToString(t *testing.T) {
	g := NewGomegaWithT(t)

	option := Option{}
	option.Name = "testoption"
	option.Type = OptionStringType
	option.Default = "testdefault"

	g.Expect(option.String()).
		To(Equal("option name testoption type string default testdefault"))
}
