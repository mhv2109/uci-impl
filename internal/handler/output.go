package handler

import (
	"fmt"
	"strings"

	"uci-impl/internal/handler/info"
	"uci-impl/internal/solver"
)

func emitID() {
	fmt.Println("id name mhv2109-engine")
	fmt.Println("id author mhv2109")
}

func emitUCIOK() {
	fmt.Println("uciok")
}

func emitReadyOK() {
	fmt.Println("readyok")
}

func emitBestmove(moves ...string) {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("bestmove %s", moves[0]))

	if len(moves) == 2 {
		builder.WriteString(fmt.Sprintf(" ponder %s", moves[1]))
	}

	fmt.Println(builder.String())
}

func emitCopyprotection() {
	panic("Not Implemented!")
}

func emitRegistration() {
	panic("Not Implemented!")
}

func emitInfo(i info.Info) {
	iStr := i.String()

	// don't print empty info
	if iStr != "info" {
		fmt.Println(iStr)
	}
}

func emitOption(s solver.Solver) {
	for _, o := range s.GetOptions() {
		fmt.Println(o)
	}
}
