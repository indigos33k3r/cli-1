// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeSpaceQuotasActor struct {
	GetSpaceQuotasByOrgGUIDStub        func(string) ([]v7action.SpaceQuota, v7action.Warnings, error)
	getSpaceQuotasByOrgGUIDMutex       sync.RWMutex
	getSpaceQuotasByOrgGUIDArgsForCall []struct {
		arg1 string
	}
	getSpaceQuotasByOrgGUIDReturns struct {
		result1 []v7action.SpaceQuota
		result2 v7action.Warnings
		result3 error
	}
	getSpaceQuotasByOrgGUIDReturnsOnCall map[int]struct {
		result1 []v7action.SpaceQuota
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpaceQuotasActor) GetSpaceQuotasByOrgGUID(arg1 string) ([]v7action.SpaceQuota, v7action.Warnings, error) {
	fake.getSpaceQuotasByOrgGUIDMutex.Lock()
	ret, specificReturn := fake.getSpaceQuotasByOrgGUIDReturnsOnCall[len(fake.getSpaceQuotasByOrgGUIDArgsForCall)]
	fake.getSpaceQuotasByOrgGUIDArgsForCall = append(fake.getSpaceQuotasByOrgGUIDArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetSpaceQuotasByOrgGUID", []interface{}{arg1})
	fake.getSpaceQuotasByOrgGUIDMutex.Unlock()
	if fake.GetSpaceQuotasByOrgGUIDStub != nil {
		return fake.GetSpaceQuotasByOrgGUIDStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getSpaceQuotasByOrgGUIDReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSpaceQuotasActor) GetSpaceQuotasByOrgGUIDCallCount() int {
	fake.getSpaceQuotasByOrgGUIDMutex.RLock()
	defer fake.getSpaceQuotasByOrgGUIDMutex.RUnlock()
	return len(fake.getSpaceQuotasByOrgGUIDArgsForCall)
}

func (fake *FakeSpaceQuotasActor) GetSpaceQuotasByOrgGUIDCalls(stub func(string) ([]v7action.SpaceQuota, v7action.Warnings, error)) {
	fake.getSpaceQuotasByOrgGUIDMutex.Lock()
	defer fake.getSpaceQuotasByOrgGUIDMutex.Unlock()
	fake.GetSpaceQuotasByOrgGUIDStub = stub
}

func (fake *FakeSpaceQuotasActor) GetSpaceQuotasByOrgGUIDArgsForCall(i int) string {
	fake.getSpaceQuotasByOrgGUIDMutex.RLock()
	defer fake.getSpaceQuotasByOrgGUIDMutex.RUnlock()
	argsForCall := fake.getSpaceQuotasByOrgGUIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSpaceQuotasActor) GetSpaceQuotasByOrgGUIDReturns(result1 []v7action.SpaceQuota, result2 v7action.Warnings, result3 error) {
	fake.getSpaceQuotasByOrgGUIDMutex.Lock()
	defer fake.getSpaceQuotasByOrgGUIDMutex.Unlock()
	fake.GetSpaceQuotasByOrgGUIDStub = nil
	fake.getSpaceQuotasByOrgGUIDReturns = struct {
		result1 []v7action.SpaceQuota
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceQuotasActor) GetSpaceQuotasByOrgGUIDReturnsOnCall(i int, result1 []v7action.SpaceQuota, result2 v7action.Warnings, result3 error) {
	fake.getSpaceQuotasByOrgGUIDMutex.Lock()
	defer fake.getSpaceQuotasByOrgGUIDMutex.Unlock()
	fake.GetSpaceQuotasByOrgGUIDStub = nil
	if fake.getSpaceQuotasByOrgGUIDReturnsOnCall == nil {
		fake.getSpaceQuotasByOrgGUIDReturnsOnCall = make(map[int]struct {
			result1 []v7action.SpaceQuota
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getSpaceQuotasByOrgGUIDReturnsOnCall[i] = struct {
		result1 []v7action.SpaceQuota
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceQuotasActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSpaceQuotasByOrgGUIDMutex.RLock()
	defer fake.getSpaceQuotasByOrgGUIDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSpaceQuotasActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.SpaceQuotasActor = new(FakeSpaceQuotasActor)
