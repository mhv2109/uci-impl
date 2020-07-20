// Code generated by counterfeiter. DO NOT EDIT.
package handlerfakes

import (
	"sync"

	"github.com/mhv2109/uci-impl/internal/handler"
	"github.com/mhv2109/uci-impl/internal/handler/info"
	"github.com/mhv2109/uci-impl/internal/solver"
)

type FakeEmitter struct {
	EmitBestmoveStub        func(...string)
	emitBestmoveMutex       sync.RWMutex
	emitBestmoveArgsForCall []struct {
		arg1 []string
	}
	EmitCopyProtectionCheckingStub        func()
	emitCopyProtectionCheckingMutex       sync.RWMutex
	emitCopyProtectionCheckingArgsForCall []struct {
	}
	EmitCopyProtectionErrorStub        func()
	emitCopyProtectionErrorMutex       sync.RWMutex
	emitCopyProtectionErrorArgsForCall []struct {
	}
	EmitCopyProtectionOkStub        func()
	emitCopyProtectionOkMutex       sync.RWMutex
	emitCopyProtectionOkArgsForCall []struct {
	}
	EmitIDStub        func()
	emitIDMutex       sync.RWMutex
	emitIDArgsForCall []struct {
	}
	EmitInfoStub        func(info.Info)
	emitInfoMutex       sync.RWMutex
	emitInfoArgsForCall []struct {
		arg1 info.Info
	}
	EmitOptionStub        func(solver.Solver)
	emitOptionMutex       sync.RWMutex
	emitOptionArgsForCall []struct {
		arg1 solver.Solver
	}
	EmitReadyOKStub        func()
	emitReadyOKMutex       sync.RWMutex
	emitReadyOKArgsForCall []struct {
	}
	EmitRegistrationCheckingStub        func()
	emitRegistrationCheckingMutex       sync.RWMutex
	emitRegistrationCheckingArgsForCall []struct {
	}
	EmitRegistrationErrorStub        func()
	emitRegistrationErrorMutex       sync.RWMutex
	emitRegistrationErrorArgsForCall []struct {
	}
	EmitRegistrationOkStub        func()
	emitRegistrationOkMutex       sync.RWMutex
	emitRegistrationOkArgsForCall []struct {
	}
	EmitUCIOKStub        func()
	emitUCIOKMutex       sync.RWMutex
	emitUCIOKArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEmitter) EmitBestmove(arg1 ...string) {
	fake.emitBestmoveMutex.Lock()
	fake.emitBestmoveArgsForCall = append(fake.emitBestmoveArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.recordInvocation("EmitBestmove", []interface{}{arg1})
	fake.emitBestmoveMutex.Unlock()
	if fake.EmitBestmoveStub != nil {
		fake.EmitBestmoveStub(arg1...)
	}
}

func (fake *FakeEmitter) EmitBestmoveCallCount() int {
	fake.emitBestmoveMutex.RLock()
	defer fake.emitBestmoveMutex.RUnlock()
	return len(fake.emitBestmoveArgsForCall)
}

func (fake *FakeEmitter) EmitBestmoveCalls(stub func(...string)) {
	fake.emitBestmoveMutex.Lock()
	defer fake.emitBestmoveMutex.Unlock()
	fake.EmitBestmoveStub = stub
}

func (fake *FakeEmitter) EmitBestmoveArgsForCall(i int) []string {
	fake.emitBestmoveMutex.RLock()
	defer fake.emitBestmoveMutex.RUnlock()
	argsForCall := fake.emitBestmoveArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEmitter) EmitCopyProtectionChecking() {
	fake.emitCopyProtectionCheckingMutex.Lock()
	fake.emitCopyProtectionCheckingArgsForCall = append(fake.emitCopyProtectionCheckingArgsForCall, struct {
	}{})
	fake.recordInvocation("EmitCopyProtectionChecking", []interface{}{})
	fake.emitCopyProtectionCheckingMutex.Unlock()
	if fake.EmitCopyProtectionCheckingStub != nil {
		fake.EmitCopyProtectionCheckingStub()
	}
}

func (fake *FakeEmitter) EmitCopyProtectionCheckingCallCount() int {
	fake.emitCopyProtectionCheckingMutex.RLock()
	defer fake.emitCopyProtectionCheckingMutex.RUnlock()
	return len(fake.emitCopyProtectionCheckingArgsForCall)
}

func (fake *FakeEmitter) EmitCopyProtectionCheckingCalls(stub func()) {
	fake.emitCopyProtectionCheckingMutex.Lock()
	defer fake.emitCopyProtectionCheckingMutex.Unlock()
	fake.EmitCopyProtectionCheckingStub = stub
}

func (fake *FakeEmitter) EmitCopyProtectionError() {
	fake.emitCopyProtectionErrorMutex.Lock()
	fake.emitCopyProtectionErrorArgsForCall = append(fake.emitCopyProtectionErrorArgsForCall, struct {
	}{})
	fake.recordInvocation("EmitCopyProtectionError", []interface{}{})
	fake.emitCopyProtectionErrorMutex.Unlock()
	if fake.EmitCopyProtectionErrorStub != nil {
		fake.EmitCopyProtectionErrorStub()
	}
}

func (fake *FakeEmitter) EmitCopyProtectionErrorCallCount() int {
	fake.emitCopyProtectionErrorMutex.RLock()
	defer fake.emitCopyProtectionErrorMutex.RUnlock()
	return len(fake.emitCopyProtectionErrorArgsForCall)
}

func (fake *FakeEmitter) EmitCopyProtectionErrorCalls(stub func()) {
	fake.emitCopyProtectionErrorMutex.Lock()
	defer fake.emitCopyProtectionErrorMutex.Unlock()
	fake.EmitCopyProtectionErrorStub = stub
}

func (fake *FakeEmitter) EmitCopyProtectionOk() {
	fake.emitCopyProtectionOkMutex.Lock()
	fake.emitCopyProtectionOkArgsForCall = append(fake.emitCopyProtectionOkArgsForCall, struct {
	}{})
	fake.recordInvocation("EmitCopyProtectionOk", []interface{}{})
	fake.emitCopyProtectionOkMutex.Unlock()
	if fake.EmitCopyProtectionOkStub != nil {
		fake.EmitCopyProtectionOkStub()
	}
}

func (fake *FakeEmitter) EmitCopyProtectionOkCallCount() int {
	fake.emitCopyProtectionOkMutex.RLock()
	defer fake.emitCopyProtectionOkMutex.RUnlock()
	return len(fake.emitCopyProtectionOkArgsForCall)
}

func (fake *FakeEmitter) EmitCopyProtectionOkCalls(stub func()) {
	fake.emitCopyProtectionOkMutex.Lock()
	defer fake.emitCopyProtectionOkMutex.Unlock()
	fake.EmitCopyProtectionOkStub = stub
}

func (fake *FakeEmitter) EmitID() {
	fake.emitIDMutex.Lock()
	fake.emitIDArgsForCall = append(fake.emitIDArgsForCall, struct {
	}{})
	fake.recordInvocation("EmitID", []interface{}{})
	fake.emitIDMutex.Unlock()
	if fake.EmitIDStub != nil {
		fake.EmitIDStub()
	}
}

func (fake *FakeEmitter) EmitIDCallCount() int {
	fake.emitIDMutex.RLock()
	defer fake.emitIDMutex.RUnlock()
	return len(fake.emitIDArgsForCall)
}

func (fake *FakeEmitter) EmitIDCalls(stub func()) {
	fake.emitIDMutex.Lock()
	defer fake.emitIDMutex.Unlock()
	fake.EmitIDStub = stub
}

func (fake *FakeEmitter) EmitInfo(arg1 info.Info) {
	fake.emitInfoMutex.Lock()
	fake.emitInfoArgsForCall = append(fake.emitInfoArgsForCall, struct {
		arg1 info.Info
	}{arg1})
	fake.recordInvocation("EmitInfo", []interface{}{arg1})
	fake.emitInfoMutex.Unlock()
	if fake.EmitInfoStub != nil {
		fake.EmitInfoStub(arg1)
	}
}

func (fake *FakeEmitter) EmitInfoCallCount() int {
	fake.emitInfoMutex.RLock()
	defer fake.emitInfoMutex.RUnlock()
	return len(fake.emitInfoArgsForCall)
}

func (fake *FakeEmitter) EmitInfoCalls(stub func(info.Info)) {
	fake.emitInfoMutex.Lock()
	defer fake.emitInfoMutex.Unlock()
	fake.EmitInfoStub = stub
}

func (fake *FakeEmitter) EmitInfoArgsForCall(i int) info.Info {
	fake.emitInfoMutex.RLock()
	defer fake.emitInfoMutex.RUnlock()
	argsForCall := fake.emitInfoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEmitter) EmitOption(arg1 solver.Solver) {
	fake.emitOptionMutex.Lock()
	fake.emitOptionArgsForCall = append(fake.emitOptionArgsForCall, struct {
		arg1 solver.Solver
	}{arg1})
	fake.recordInvocation("EmitOption", []interface{}{arg1})
	fake.emitOptionMutex.Unlock()
	if fake.EmitOptionStub != nil {
		fake.EmitOptionStub(arg1)
	}
}

func (fake *FakeEmitter) EmitOptionCallCount() int {
	fake.emitOptionMutex.RLock()
	defer fake.emitOptionMutex.RUnlock()
	return len(fake.emitOptionArgsForCall)
}

func (fake *FakeEmitter) EmitOptionCalls(stub func(solver.Solver)) {
	fake.emitOptionMutex.Lock()
	defer fake.emitOptionMutex.Unlock()
	fake.EmitOptionStub = stub
}

func (fake *FakeEmitter) EmitOptionArgsForCall(i int) solver.Solver {
	fake.emitOptionMutex.RLock()
	defer fake.emitOptionMutex.RUnlock()
	argsForCall := fake.emitOptionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEmitter) EmitReadyOK() {
	fake.emitReadyOKMutex.Lock()
	fake.emitReadyOKArgsForCall = append(fake.emitReadyOKArgsForCall, struct {
	}{})
	fake.recordInvocation("EmitReadyOK", []interface{}{})
	fake.emitReadyOKMutex.Unlock()
	if fake.EmitReadyOKStub != nil {
		fake.EmitReadyOKStub()
	}
}

func (fake *FakeEmitter) EmitReadyOKCallCount() int {
	fake.emitReadyOKMutex.RLock()
	defer fake.emitReadyOKMutex.RUnlock()
	return len(fake.emitReadyOKArgsForCall)
}

func (fake *FakeEmitter) EmitReadyOKCalls(stub func()) {
	fake.emitReadyOKMutex.Lock()
	defer fake.emitReadyOKMutex.Unlock()
	fake.EmitReadyOKStub = stub
}

func (fake *FakeEmitter) EmitRegistrationChecking() {
	fake.emitRegistrationCheckingMutex.Lock()
	fake.emitRegistrationCheckingArgsForCall = append(fake.emitRegistrationCheckingArgsForCall, struct {
	}{})
	fake.recordInvocation("EmitRegistrationChecking", []interface{}{})
	fake.emitRegistrationCheckingMutex.Unlock()
	if fake.EmitRegistrationCheckingStub != nil {
		fake.EmitRegistrationCheckingStub()
	}
}

func (fake *FakeEmitter) EmitRegistrationCheckingCallCount() int {
	fake.emitRegistrationCheckingMutex.RLock()
	defer fake.emitRegistrationCheckingMutex.RUnlock()
	return len(fake.emitRegistrationCheckingArgsForCall)
}

func (fake *FakeEmitter) EmitRegistrationCheckingCalls(stub func()) {
	fake.emitRegistrationCheckingMutex.Lock()
	defer fake.emitRegistrationCheckingMutex.Unlock()
	fake.EmitRegistrationCheckingStub = stub
}

func (fake *FakeEmitter) EmitRegistrationError() {
	fake.emitRegistrationErrorMutex.Lock()
	fake.emitRegistrationErrorArgsForCall = append(fake.emitRegistrationErrorArgsForCall, struct {
	}{})
	fake.recordInvocation("EmitRegistrationError", []interface{}{})
	fake.emitRegistrationErrorMutex.Unlock()
	if fake.EmitRegistrationErrorStub != nil {
		fake.EmitRegistrationErrorStub()
	}
}

func (fake *FakeEmitter) EmitRegistrationErrorCallCount() int {
	fake.emitRegistrationErrorMutex.RLock()
	defer fake.emitRegistrationErrorMutex.RUnlock()
	return len(fake.emitRegistrationErrorArgsForCall)
}

func (fake *FakeEmitter) EmitRegistrationErrorCalls(stub func()) {
	fake.emitRegistrationErrorMutex.Lock()
	defer fake.emitRegistrationErrorMutex.Unlock()
	fake.EmitRegistrationErrorStub = stub
}

func (fake *FakeEmitter) EmitRegistrationOk() {
	fake.emitRegistrationOkMutex.Lock()
	fake.emitRegistrationOkArgsForCall = append(fake.emitRegistrationOkArgsForCall, struct {
	}{})
	fake.recordInvocation("EmitRegistrationOk", []interface{}{})
	fake.emitRegistrationOkMutex.Unlock()
	if fake.EmitRegistrationOkStub != nil {
		fake.EmitRegistrationOkStub()
	}
}

func (fake *FakeEmitter) EmitRegistrationOkCallCount() int {
	fake.emitRegistrationOkMutex.RLock()
	defer fake.emitRegistrationOkMutex.RUnlock()
	return len(fake.emitRegistrationOkArgsForCall)
}

func (fake *FakeEmitter) EmitRegistrationOkCalls(stub func()) {
	fake.emitRegistrationOkMutex.Lock()
	defer fake.emitRegistrationOkMutex.Unlock()
	fake.EmitRegistrationOkStub = stub
}

func (fake *FakeEmitter) EmitUCIOK() {
	fake.emitUCIOKMutex.Lock()
	fake.emitUCIOKArgsForCall = append(fake.emitUCIOKArgsForCall, struct {
	}{})
	fake.recordInvocation("EmitUCIOK", []interface{}{})
	fake.emitUCIOKMutex.Unlock()
	if fake.EmitUCIOKStub != nil {
		fake.EmitUCIOKStub()
	}
}

func (fake *FakeEmitter) EmitUCIOKCallCount() int {
	fake.emitUCIOKMutex.RLock()
	defer fake.emitUCIOKMutex.RUnlock()
	return len(fake.emitUCIOKArgsForCall)
}

func (fake *FakeEmitter) EmitUCIOKCalls(stub func()) {
	fake.emitUCIOKMutex.Lock()
	defer fake.emitUCIOKMutex.Unlock()
	fake.EmitUCIOKStub = stub
}

func (fake *FakeEmitter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.emitBestmoveMutex.RLock()
	defer fake.emitBestmoveMutex.RUnlock()
	fake.emitCopyProtectionCheckingMutex.RLock()
	defer fake.emitCopyProtectionCheckingMutex.RUnlock()
	fake.emitCopyProtectionErrorMutex.RLock()
	defer fake.emitCopyProtectionErrorMutex.RUnlock()
	fake.emitCopyProtectionOkMutex.RLock()
	defer fake.emitCopyProtectionOkMutex.RUnlock()
	fake.emitIDMutex.RLock()
	defer fake.emitIDMutex.RUnlock()
	fake.emitInfoMutex.RLock()
	defer fake.emitInfoMutex.RUnlock()
	fake.emitOptionMutex.RLock()
	defer fake.emitOptionMutex.RUnlock()
	fake.emitReadyOKMutex.RLock()
	defer fake.emitReadyOKMutex.RUnlock()
	fake.emitRegistrationCheckingMutex.RLock()
	defer fake.emitRegistrationCheckingMutex.RUnlock()
	fake.emitRegistrationErrorMutex.RLock()
	defer fake.emitRegistrationErrorMutex.RUnlock()
	fake.emitRegistrationOkMutex.RLock()
	defer fake.emitRegistrationOkMutex.RUnlock()
	fake.emitUCIOKMutex.RLock()
	defer fake.emitUCIOKMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEmitter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ handler.Emitter = new(FakeEmitter)
