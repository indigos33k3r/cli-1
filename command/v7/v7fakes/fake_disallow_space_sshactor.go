// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeDisallowSpaceSSHActor struct {
	UpdateSpaceFeatureStub        func(string, string, bool, string) (v7action.Warnings, error)
	updateSpaceFeatureMutex       sync.RWMutex
	updateSpaceFeatureArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
		arg4 string
	}
	updateSpaceFeatureReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	updateSpaceFeatureReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDisallowSpaceSSHActor) UpdateSpaceFeature(arg1 string, arg2 string, arg3 bool, arg4 string) (v7action.Warnings, error) {
	fake.updateSpaceFeatureMutex.Lock()
	ret, specificReturn := fake.updateSpaceFeatureReturnsOnCall[len(fake.updateSpaceFeatureArgsForCall)]
	fake.updateSpaceFeatureArgsForCall = append(fake.updateSpaceFeatureArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("UpdateSpaceFeature", []interface{}{arg1, arg2, arg3, arg4})
	fake.updateSpaceFeatureMutex.Unlock()
	if fake.UpdateSpaceFeatureStub != nil {
		return fake.UpdateSpaceFeatureStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateSpaceFeatureReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDisallowSpaceSSHActor) UpdateSpaceFeatureCallCount() int {
	fake.updateSpaceFeatureMutex.RLock()
	defer fake.updateSpaceFeatureMutex.RUnlock()
	return len(fake.updateSpaceFeatureArgsForCall)
}

func (fake *FakeDisallowSpaceSSHActor) UpdateSpaceFeatureCalls(stub func(string, string, bool, string) (v7action.Warnings, error)) {
	fake.updateSpaceFeatureMutex.Lock()
	defer fake.updateSpaceFeatureMutex.Unlock()
	fake.UpdateSpaceFeatureStub = stub
}

func (fake *FakeDisallowSpaceSSHActor) UpdateSpaceFeatureArgsForCall(i int) (string, string, bool, string) {
	fake.updateSpaceFeatureMutex.RLock()
	defer fake.updateSpaceFeatureMutex.RUnlock()
	argsForCall := fake.updateSpaceFeatureArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeDisallowSpaceSSHActor) UpdateSpaceFeatureReturns(result1 v7action.Warnings, result2 error) {
	fake.updateSpaceFeatureMutex.Lock()
	defer fake.updateSpaceFeatureMutex.Unlock()
	fake.UpdateSpaceFeatureStub = nil
	fake.updateSpaceFeatureReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDisallowSpaceSSHActor) UpdateSpaceFeatureReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.updateSpaceFeatureMutex.Lock()
	defer fake.updateSpaceFeatureMutex.Unlock()
	fake.UpdateSpaceFeatureStub = nil
	if fake.updateSpaceFeatureReturnsOnCall == nil {
		fake.updateSpaceFeatureReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.updateSpaceFeatureReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDisallowSpaceSSHActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.updateSpaceFeatureMutex.RLock()
	defer fake.updateSpaceFeatureMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDisallowSpaceSSHActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.DisallowSpaceSSHActor = new(FakeDisallowSpaceSSHActor)
