package info

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func TestInfoToStringDepth(t *testing.T) {
	info := NewInfo()
	info.SetDepth(99)
	testUint(t, "depth", 99, info)
}

func TestInfoToStringSeldepth(t *testing.T) {
	info := NewInfo()
	info.SetSeldepth(99)
	testUint(t, "seldepth", 99, info)
}

func TestInfoToStringTime(t *testing.T) {
	info := NewInfo()
	info.SetTime(99)
	testUint(t, "time", 99, info)
}

func TestInfoToStringNodes(t *testing.T) {
	info := NewInfo()
	info.SetNodes(99)
	testUint(t, "nodes", 99, info)
}

func TestInfoToStringCurrmovenumber(t *testing.T) {
	info := NewInfo()
	info.SetCurrmovenumber(99)
	testUint(t, "currmovenumber", 99, info)
}

func TestInfoToStringHashfull(t *testing.T) {
	info := NewInfo()
	info.SetHashfull(99)
	testUint(t, "hashfull", 99, info)
}

func TestInfoToStringNps(t *testing.T) {
	info := NewInfo()
	info.SetNps(99)
	testUint(t, "nps", 99, info)
}

func TestInfoToStringTbhits(t *testing.T) {
	info := NewInfo()
	info.SetTbhits(99)
	testUint(t, "tbhits", 99, info)
}

func TestInfoToStringSbhits(t *testing.T) {
	info := NewInfo()
	info.SetSbhits(99)
	testUint(t, "sbhits", 99, info)
}

func TestInfoToStringCpuload(t *testing.T) {
	info := NewInfo()
	info.SetCpuload(99)
	testUint(t, "cpuload", 99, info)
}

func TestInfoToStringPv(t *testing.T) {
	g := NewGomegaWithT(t)

	info := NewInfo()
	for _, pv := range []string{"e4", "Nf3", "Bb5"} {
		info.AddPv(pv)
	}

	a, e := info.String(), "info pv e4 Nf3 Bb5"
	g.Expect(a).
		To(Equal(e))
}

func TestInfoToStringScore(t *testing.T) {
	g := NewGomegaWithT(t)

	for i, st := range []ScoreType{CP, Mate, Lowerbound, Upperbound} {
		info := NewInfo()
		info.SetScore(st, i)

		a, e := info.String(), fmt.Sprintf("info score %s %d", st, i)
		g.Expect(a).
			To(Equal(e))
	}
}

func TestInfoToStringCurrmove(t *testing.T) {
	g := NewGomegaWithT(t)

	info := NewInfo()
	info.SetCurrmove("Nf3")

	a, e := info.String(), "info currmove Nf3"
	g.Expect(a).
		To(Equal(e))
}

func TestInfoToStringRefutations(t *testing.T) {
	g := NewGomegaWithT(t)

	info := NewInfo()
	for _, pv := range []string{"e4", "Nf3", "Bb5"} {
		info.AddRefutation(pv)
	}

	a, e := info.String(), "info refutation e4 Nf3 Bb5"
	g.Expect(a).
		To(Equal(e))
}

func TestInfoToStringCurrline(t *testing.T) {
	g := NewGomegaWithT(t)

	info := NewInfo()
	info.SetCurrline(1, "e4", "Nf3", "Bb5")

	s, e := info.String(), "info currline 1 e4 Nf3 Bb5"
	g.Expect(s).
		To(Equal(e))
}

func testUint(t *testing.T, name string, value uint, info *Info) {
	NewGomegaWithT(t).Expect(info.String()).
		To(Equal(fmt.Sprintf("info %s %d", name, value)))
}
