// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	graph "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/graph"
	liquidityrebalancer "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/liquidityrebalancer"
	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/models"
)

// Rebalancer is an autogenerated mock type for the Rebalancer type
type Rebalancer struct {
	mock.Mock
}

// ComputeTransfersToBalance provides a mock function with given fields: g, unexecuted
func (_m *Rebalancer) ComputeTransfersToBalance(g graph.Graph, unexecuted []liquidityrebalancer.UnexecutedTransfer) ([]models.ProposedTransfer, error) {
	ret := _m.Called(g, unexecuted)

	if len(ret) == 0 {
		panic("no return value specified for ComputeTransfersToBalance")
	}

	var r0 []models.ProposedTransfer
	var r1 error
	if rf, ok := ret.Get(0).(func(graph.Graph, []liquidityrebalancer.UnexecutedTransfer) ([]models.ProposedTransfer, error)); ok {
		return rf(g, unexecuted)
	}
	if rf, ok := ret.Get(0).(func(graph.Graph, []liquidityrebalancer.UnexecutedTransfer) []models.ProposedTransfer); ok {
		r0 = rf(g, unexecuted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ProposedTransfer)
		}
	}

	if rf, ok := ret.Get(1).(func(graph.Graph, []liquidityrebalancer.UnexecutedTransfer) error); ok {
		r1 = rf(g, unexecuted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRebalancer creates a new instance of Rebalancer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRebalancer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Rebalancer {
	mock := &Rebalancer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
