package handler

import (
	"fmt"
	"strings"

	"uci-impl/internal/handler/info"
	"uci-impl/internal/solver"
)

type Emitter interface {
	EmitID()
	EmitUCIOK()
	EmitReadyOK()
	EmitBestmove(moves ...string)
	EmitCopyProtectionChecking()
	EmitCopyProtectionOk()
	EmitCopyProtectionError()
	EmitRegistrationChecking()
	EmitRegistrationOk()
	EmitRegistrationError()
	EmitInfo(i info.Info)
	EmitOption(s solver.Solver)
}

type EmitterImpl struct{}

func NewEmitter() Emitter {
	return &EmitterImpl{}
}

func (e *EmitterImpl) EmitID() {
	fmt.Println("id name mhv2109-engine")
	fmt.Println("id author mhv2109")
}

func (e *EmitterImpl) EmitUCIOK() {
	fmt.Println("uciok")
}

func (e *EmitterImpl) EmitReadyOK() {
	fmt.Println("readyok")
}

func (e *EmitterImpl) EmitBestmove(moves ...string) {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("bestmove %s", moves[0]))

	if len(moves) == 2 {
		builder.WriteString(fmt.Sprintf(" ponder %s", moves[1]))
	}

	fmt.Println(builder.String())
}

func (e *EmitterImpl) EmitCopyProtectionChecking() {
	fmt.Println("copyprotection checking")
}

func (e *EmitterImpl) EmitCopyProtectionOk() {
	fmt.Println("copyprotection ok")
}

func (e *EmitterImpl) EmitCopyProtectionError() {
	fmt.Println("copyprotection error")
}

func (e *EmitterImpl) EmitRegistrationChecking() {
	fmt.Println("registration checking")
}

func (e *EmitterImpl) EmitRegistrationOk() {
	fmt.Println("registration ok")
}

func (e *EmitterImpl) EmitRegistrationError() {
	fmt.Println("registration error")
}

func (e *EmitterImpl) EmitInfo(i info.Info) {
	iStr := i.String()

	// don't print empty info
	if iStr != "info" {
		fmt.Println(iStr)
	}
}

func (e *EmitterImpl) EmitOption(s solver.Solver) {
	for _, o := range s.GetOptions() {
		fmt.Println(o)
	}
}
