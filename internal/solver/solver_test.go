package solver_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mhv2109/uci-impl/internal/solver"
)

var _ = Describe("Option", func() {

	var option Option

	BeforeEach(func() {
		option = Option{}
	})

	It("Spin Option stringifies correctly", func() {
		option.Name = "testoption"
		option.Type = OptionSpinType
		option.Default = "50"
		option.Min = "0"
		option.Max = "100"

		Expect(option.String()).
			To(Equal("option name testoption type spin default 50 min 0 max 100"))
	})

	It("String Option stringifies correctly", func() {
		option.Name = "testoption"
		option.Type = OptionStringType
		option.Default = "testdefault"

		Expect(option.String()).
			To(Equal("option name testoption type string default testdefault"))
	})
})
