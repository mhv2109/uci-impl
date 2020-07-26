package info

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Info", func() {
	var info *Info

	BeforeEach(func() {
		info = NewInfo()
	})

	testUint := func(name string, value uint, info *Info) {
		Expect(info.String()).
			To(Equal(fmt.Sprintf("info %s %d", name, value)))
	}

	var _ = Describe("uint", func() {
		It("depth", func() {
			info.SetDepth(99)
			testUint("depth", 99, info)
		})

		It("seldepth", func() {
			info.SetSeldepth(99)
			testUint("seldepth", 99, info)
		})

		It("time", func() {
			info.SetTime(99)
			testUint("time", 99, info)
		})

		It("nodes", func() {
			info.SetNodes(99)
			testUint("nodes", 99, info)
		})

		It("currmovenumber", func() {
			info.SetCurrmovenumber(99)
			testUint("currmovenumber", 99, info)
		})

		It("hashfull", func() {
			info.SetHashfull(99)
			testUint("hashfull", 99, info)
		})

		It("nps", func() {
			info.SetNps(99)
			testUint("nps", 99, info)
		})

		It("tbhits", func() {
			info.SetTbhits(99)
			testUint("tbhits", 99, info)
		})

		It("sbhits", func() {
			info.SetSbhits(99)
			testUint("sbhits", 99, info)
		})

		It("cpuload", func() {
			info.SetCpuload(99)
			testUint("cpuload", 99, info)
		})
	})

	It("pv", func() {
		for _, pv := range []string{"e4", "Nf3", "Bb5"} {
			info.AddPv(pv)
		}

		a, e := info.String(), "info pv e4 Nf3 Bb5"
		Expect(a).
			To(Equal(e))
	})

	Describe("score", func() {
		var st ScoreType
		sts := []ScoreType{CP, Mate, Lowerbound, Upperbound}
		i := 0

		BeforeEach(func() {
			st = sts[i]
			info.SetScore(st, i)
		})

		AfterEach(func() {
			i++
		})

		It("Score serializes", func() {
			a, e := info.String(), fmt.Sprintf("info score %s %d", st, i)
			Expect(a).
				To(Equal(e))
		})
	})

	It("currmove", func() {
		info.SetCurrmove("Nf3")

		a, e := info.String(), "info currmove Nf3"
		Expect(a).
			To(Equal(e))
	})

	It("refutation", func() {
		for _, pv := range []string{"e4", "Nf3", "Bb5"} {
			info.AddRefutation(pv)
		}

		a, e := info.String(), "info refutation e4 Nf3 Bb5"
		Expect(a).
			To(Equal(e))
	})

	It("currline", func() {
		info.SetCurrline(1, "e4", "Nf3", "Bb5")

		s, e := info.String(), "info currline 1 e4 Nf3 Bb5"
		Expect(s).
			To(Equal(e))
	})
})
