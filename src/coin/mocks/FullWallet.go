// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import core "github.com/fibercrypto/fibercryptowallet/src/core"
import mock "github.com/stretchr/testify/mock"

// FullWallet is an autogenerated mock type for the FullWallet type
type FullWallet struct {
	mock.Mock
}

// GenAddresses provides a mock function with given fields: addrType, startIndex, count, pwd
func (_m *FullWallet) GenAddresses(addrType core.AddressType, startIndex uint32, count uint32, pwd core.PasswordReader) core.AddressIterator {
	ret := _m.Called(addrType, startIndex, count, pwd)

	var r0 core.AddressIterator
	if rf, ok := ret.Get(0).(func(core.AddressType, uint32, uint32, core.PasswordReader) core.AddressIterator); ok {
		r0 = rf(addrType, startIndex, count, pwd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.AddressIterator)
		}
	}

	return r0
}

// GetCryptoAccount provides a mock function with given fields:
func (_m *FullWallet) GetCryptoAccount() core.CryptoAccount {
	ret := _m.Called()

	var r0 core.CryptoAccount
	if rf, ok := ret.Get(0).(func() core.CryptoAccount); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.CryptoAccount)
		}
	}

	return r0
}

// GetId provides a mock function with given fields:
func (_m *FullWallet) GetId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetLabel provides a mock function with given fields:
func (_m *FullWallet) GetLabel() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetLoadedAddresses provides a mock function with given fields:
func (_m *FullWallet) GetLoadedAddresses() (core.AddressIterator, error) {
	ret := _m.Called()

	var r0 core.AddressIterator
	if rf, ok := ret.Get(0).(func() core.AddressIterator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.AddressIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendFromAddress provides a mock function with given fields: from, to, change, options
func (_m *FullWallet) SendFromAddress(from []core.Address, to []core.TransactionOutput, change core.Address, options core.KeyValueStorage) (core.Transaction, error) {
	ret := _m.Called(from, to, change, options)

	var r0 core.Transaction
	if rf, ok := ret.Get(0).(func([]core.Address, []core.TransactionOutput, core.Address, core.KeyValueStore) core.Transaction); ok {
		r0 = rf(from, to, change, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]core.Address, []core.TransactionOutput, core.Address, core.KeyValueStore) error); ok {
		r1 = rf(from, to, change, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetLabel provides a mock function with given fields: wltName
func (_m *FullWallet) SetLabel(wltName string) {
	_m.Called(wltName)
}

// Sign provides a mock function with given fields: txn, signer, pwd, index
func (_m *FullWallet) Sign(txn core.Transaction, signer core.TxnSigner, pwd core.PasswordReader, index []string) (core.Transaction, error) {
	ret := _m.Called(txn, signer, pwd, index)

	var r0 core.Transaction
	if rf, ok := ret.Get(0).(func(core.Transaction, core.TxnSigner, core.PasswordReader, []string) core.Transaction); ok {
		r0 = rf(txn, signer, pwd, index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.Transaction, core.TxnSigner, core.PasswordReader, []string) error); ok {
		r1 = rf(txn, signer, pwd, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Spend provides a mock function with given fields: unspent, new, change, options
func (_m *FullWallet) Spend(unspent []core.TransactionOutput, new []core.TransactionOutput, change core.Address, options core.KeyValueStorage) (core.Transaction, error) {
	ret := _m.Called(unspent, new, change, options)

	var r0 core.Transaction
	if rf, ok := ret.Get(0).(func([]core.TransactionOutput, []core.TransactionOutput, core.Address, core.KeyValueStore) core.Transaction); ok {
		r0 = rf(unspent, new, change, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]core.TransactionOutput, []core.TransactionOutput, core.Address, core.KeyValueStore) error); ok {
		r1 = rf(unspent, new, change, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transfer provides a mock function with given fields: to, options
func (_m *FullWallet) Transfer(to core.TransactionOutput, options core.KeyValueStorage) (core.Transaction, error) {
	ret := _m.Called(to, options)

	var r0 core.Transaction
	if rf, ok := ret.Get(0).(func(core.TransactionOutput, core.KeyValueStore) core.Transaction); ok {
		r0 = rf(to, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.TransactionOutput, core.KeyValueStore) error); ok {
		r1 = rf(to, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
