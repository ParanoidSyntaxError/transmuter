// Code generated by mockery v2.38.0. DO NOT EDIT.

package mock_node_interface

import (
	big "math/big"

	arb_node_interface "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/arb_node_interface"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// NodeInterfaceInterface is an autogenerated mock type for the NodeInterfaceInterface type
type NodeInterfaceInterface struct {
	mock.Mock
}

// Address provides a mock function with given fields:
func (_m *NodeInterfaceInterface) Address() common.Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Address")
	}

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// BlockL1Num provides a mock function with given fields: opts, l2BlockNum
func (_m *NodeInterfaceInterface) BlockL1Num(opts *bind.CallOpts, l2BlockNum uint64) (uint64, error) {
	ret := _m.Called(opts, l2BlockNum)

	if len(ret) == 0 {
		panic("no return value specified for BlockL1Num")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, uint64) (uint64, error)); ok {
		return rf(opts, l2BlockNum)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, uint64) uint64); ok {
		r0 = rf(opts, l2BlockNum)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, uint64) error); ok {
		r1 = rf(opts, l2BlockNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConstructOutboxProof provides a mock function with given fields: opts, size, leaf
func (_m *NodeInterfaceInterface) ConstructOutboxProof(opts *bind.CallOpts, size uint64, leaf uint64) (arb_node_interface.ConstructOutboxProof, error) {
	ret := _m.Called(opts, size, leaf)

	if len(ret) == 0 {
		panic("no return value specified for ConstructOutboxProof")
	}

	var r0 arb_node_interface.ConstructOutboxProof
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, uint64, uint64) (arb_node_interface.ConstructOutboxProof, error)); ok {
		return rf(opts, size, leaf)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, uint64, uint64) arb_node_interface.ConstructOutboxProof); ok {
		r0 = rf(opts, size, leaf)
	} else {
		r0 = ret.Get(0).(arb_node_interface.ConstructOutboxProof)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, uint64, uint64) error); ok {
		r1 = rf(opts, size, leaf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EstimateRetryableTicket provides a mock function with given fields: opts, sender, deposit, to, l2CallValue, excessFeeRefundAddress, callValueRefundAddress, data
func (_m *NodeInterfaceInterface) EstimateRetryableTicket(opts *bind.TransactOpts, sender common.Address, deposit *big.Int, to common.Address, l2CallValue *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, data []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, sender, deposit, to, l2CallValue, excessFeeRefundAddress, callValueRefundAddress, data)

	if len(ret) == 0 {
		panic("no return value specified for EstimateRetryableTicket")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address, *big.Int, common.Address, *big.Int, common.Address, common.Address, []byte) (*types.Transaction, error)); ok {
		return rf(opts, sender, deposit, to, l2CallValue, excessFeeRefundAddress, callValueRefundAddress, data)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address, *big.Int, common.Address, *big.Int, common.Address, common.Address, []byte) *types.Transaction); ok {
		r0 = rf(opts, sender, deposit, to, l2CallValue, excessFeeRefundAddress, callValueRefundAddress, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address, *big.Int, common.Address, *big.Int, common.Address, common.Address, []byte) error); ok {
		r1 = rf(opts, sender, deposit, to, l2CallValue, excessFeeRefundAddress, callValueRefundAddress, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBatchContainingBlock provides a mock function with given fields: opts, blockNum
func (_m *NodeInterfaceInterface) FindBatchContainingBlock(opts *bind.CallOpts, blockNum uint64) (uint64, error) {
	ret := _m.Called(opts, blockNum)

	if len(ret) == 0 {
		panic("no return value specified for FindBatchContainingBlock")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, uint64) (uint64, error)); ok {
		return rf(opts, blockNum)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, uint64) uint64); ok {
		r0 = rf(opts, blockNum)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, uint64) error); ok {
		r1 = rf(opts, blockNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GasEstimateComponents provides a mock function with given fields: opts, to, contractCreation, data
func (_m *NodeInterfaceInterface) GasEstimateComponents(opts *bind.TransactOpts, to common.Address, contractCreation bool, data []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, to, contractCreation, data)

	if len(ret) == 0 {
		panic("no return value specified for GasEstimateComponents")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address, bool, []byte) (*types.Transaction, error)); ok {
		return rf(opts, to, contractCreation, data)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address, bool, []byte) *types.Transaction); ok {
		r0 = rf(opts, to, contractCreation, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address, bool, []byte) error); ok {
		r1 = rf(opts, to, contractCreation, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GasEstimateL1Component provides a mock function with given fields: opts, to, contractCreation, data
func (_m *NodeInterfaceInterface) GasEstimateL1Component(opts *bind.TransactOpts, to common.Address, contractCreation bool, data []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, to, contractCreation, data)

	if len(ret) == 0 {
		panic("no return value specified for GasEstimateL1Component")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address, bool, []byte) (*types.Transaction, error)); ok {
		return rf(opts, to, contractCreation, data)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address, bool, []byte) *types.Transaction); ok {
		r0 = rf(opts, to, contractCreation, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address, bool, []byte) error); ok {
		r1 = rf(opts, to, contractCreation, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL1Confirmations provides a mock function with given fields: opts, blockHash
func (_m *NodeInterfaceInterface) GetL1Confirmations(opts *bind.CallOpts, blockHash [32]byte) (uint64, error) {
	ret := _m.Called(opts, blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetL1Confirmations")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, [32]byte) (uint64, error)); ok {
		return rf(opts, blockHash)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, [32]byte) uint64); ok {
		r0 = rf(opts, blockHash)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, [32]byte) error); ok {
		r1 = rf(opts, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2BlockRangeForL1 provides a mock function with given fields: opts, blockNum
func (_m *NodeInterfaceInterface) L2BlockRangeForL1(opts *bind.CallOpts, blockNum uint64) (arb_node_interface.L2BlockRangeForL1, error) {
	ret := _m.Called(opts, blockNum)

	if len(ret) == 0 {
		panic("no return value specified for L2BlockRangeForL1")
	}

	var r0 arb_node_interface.L2BlockRangeForL1
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, uint64) (arb_node_interface.L2BlockRangeForL1, error)); ok {
		return rf(opts, blockNum)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, uint64) arb_node_interface.L2BlockRangeForL1); ok {
		r0 = rf(opts, blockNum)
	} else {
		r0 = ret.Get(0).(arb_node_interface.L2BlockRangeForL1)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, uint64) error); ok {
		r1 = rf(opts, blockNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LegacyLookupMessageBatchProof provides a mock function with given fields: opts, batchNum, index
func (_m *NodeInterfaceInterface) LegacyLookupMessageBatchProof(opts *bind.CallOpts, batchNum *big.Int, index uint64) (arb_node_interface.LegacyLookupMessageBatchProof, error) {
	ret := _m.Called(opts, batchNum, index)

	if len(ret) == 0 {
		panic("no return value specified for LegacyLookupMessageBatchProof")
	}

	var r0 arb_node_interface.LegacyLookupMessageBatchProof
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int, uint64) (arb_node_interface.LegacyLookupMessageBatchProof, error)); ok {
		return rf(opts, batchNum, index)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int, uint64) arb_node_interface.LegacyLookupMessageBatchProof); ok {
		r0 = rf(opts, batchNum, index)
	} else {
		r0 = ret.Get(0).(arb_node_interface.LegacyLookupMessageBatchProof)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int, uint64) error); ok {
		r1 = rf(opts, batchNum, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NitroGenesisBlock provides a mock function with given fields: opts
func (_m *NodeInterfaceInterface) NitroGenesisBlock(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for NitroGenesisBlock")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (*big.Int, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewNodeInterfaceInterface creates a new instance of NodeInterfaceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNodeInterfaceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *NodeInterfaceInterface {
	mock := &NodeInterfaceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
