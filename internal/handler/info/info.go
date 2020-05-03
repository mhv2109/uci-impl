package info

import (
	"fmt"
	"strings"
)

// Info is a struct that holds data to be sent to the GUI.
type Info struct {
	depth *uint // search depth in plies
	// selective search depth in plies,
	// if the engine sends seldepth there must also be a "depth" present in the same string.
	seldepth       *uint
	time           *uint    // the time searched in ms, this should be sent together with the pv.
	nodes          *uint    // x nodes searched, the engine should send this info regularly
	currmovenumber *uint    // currently searching move number x, for the first move x should be 1 not 0.
	hashfull       *uint    // the hash is x permill full, the engine should send this info regularly
	nps            *uint    // x nodes per second searched, the engine should send this info regularl
	tbhits         *uint    // x positions where found in the endgame table bases
	sbhits         *uint    // x positions where found in the shredder endgame databases
	cpuload        *uint    // the cpu usage of the engine is x permill.
	pv             []string // the best line found
	score          *score
	currmove       *string // currently searching this move
	// move <move1> is refuted by the line <move2> ... <movei>, i can be any number >= 1.
	// Example: after move d1h5 is searched, the engine can send
	// "info refutation d1h5 g6h5"
	// if g6h5 is the best answer after d1h5 or if g6h5 refutes the move d1h5.
	// if there is no refutation for d1h5 found, the engine should just send
	// "info refutation d1h5"
	//   The engine should only send this if the option "UCI_ShowRefutations" is set to true.
	refutation []string
	currline   *currline // this is the current line the engine is calculating.
}

func NewInfo() *Info {
	return &Info{}
}

func (i *Info) SetDepth(depth uint) {
	i.depth = &depth
}

func (i *Info) SetSeldepth(seldepth uint) {
	i.seldepth = &seldepth
}

func (i *Info) SetTime(time uint) {
	i.time = &time
}

func (i *Info) SetNodes(nodes uint) {
	i.nodes = &nodes
}

func (i *Info) SetCurrmovenumber(currmovenumber uint) {
	i.currmovenumber = &currmovenumber
}

func (i *Info) SetHashfull(hashfull uint) {
	i.hashfull = &hashfull
}

func (i *Info) SetNps(nps uint) {
	i.nps = &nps
}

func (i *Info) SetTbhits(tbhits uint) {
	i.tbhits = &tbhits
}

func (i *Info) SetSbhits(sbhits uint) {
	i.sbhits = &sbhits
}

func (i *Info) SetCpuload(cpuload uint) {
	i.cpuload = &cpuload
}

func (i *Info) SetPv(pv []string) {
	i.pv = pv
}

func (i *Info) AddPv(p string) {
	i.pv = append(i.pv, p)
}

func (i *Info) SetScore(scoretype ScoreType, value int) {
	i.score = newScore(scoretype, value)
}

func (i *Info) SetCurrmove(currmove string) {
	i.currmove = &currmove
}

func (i *Info) SetRefutation(refutation []string) {
	i.refutation = refutation
}

func (i *Info) AddRefutation(r string) {
	i.refutation = append(i.refutation, r)
}

func (i *Info) SetCurrline(cpunr uint, moves ...string) {
	i.currline = newCurrline(cpunr, moves...)
}

func (i *Info) String() string {
	var builder strings.Builder
	builder.WriteString("info")

	if i.depth != nil {
		builder.WriteString(fmt.Sprintf(" depth %d", *i.depth))
	}
	if i.seldepth != nil {
		builder.WriteString(fmt.Sprintf(" seldepth %d", *i.seldepth))
	}
	if i.time != nil {
		builder.WriteString(fmt.Sprintf(" time %d", *i.time))
	}
	if i.nodes != nil {
		builder.WriteString(fmt.Sprintf(" nodes %d", *i.nodes))
	}
	if i.currmovenumber != nil {
		builder.WriteString(fmt.Sprintf(" currmovenumber %d",
			*i.currmovenumber))
	}
	if i.hashfull != nil {
		builder.WriteString(fmt.Sprintf(" hashfull %d", *i.hashfull))
	}
	if i.nps != nil {
		builder.WriteString(fmt.Sprintf(" nps %d", *i.nps))
	}
	if i.tbhits != nil {
		builder.WriteString(fmt.Sprintf(" tbhits %d", *i.tbhits))
	}
	if i.sbhits != nil {
		builder.WriteString(fmt.Sprintf(" sbhits %d", *i.sbhits))
	}
	if i.cpuload != nil {
		builder.WriteString(fmt.Sprintf(" cpuload %d", *i.cpuload))
	}

	if len(i.pv) > 0 {
		builder.WriteString(" pv")
		for _, p := range i.pv {
			builder.WriteString(fmt.Sprintf(" %s", p))
		}
	}

	if i.score != nil {
		builder.WriteString(fmt.Sprintf(" %s", i.score))
	}

	if i.currmove != nil {
		builder.WriteString(fmt.Sprintf(" currmove %s", *i.currmove))
	}

	if len(i.refutation) > 0 {
		builder.WriteString(" refutation")
		for _, r := range i.refutation {
			builder.WriteString(fmt.Sprintf(" %s", r))
		}
	}

	if i.currline != nil {
		builder.WriteString(fmt.Sprintf(" %s", i.currline))
	}

	return builder.String()
}
