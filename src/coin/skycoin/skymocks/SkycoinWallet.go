// Code generated by mockery v1.0.0. DO NOT EDIT.

package skymocks

import mock "github.com/stretchr/testify/mock"

// SkycoinWallet is an autogenerated mock type for the SkycoinWallet type
type SkycoinWallet struct {
	mock.Mock
}

// GetSkycoinWalletType provides a mock function with given fields:
func (_m *SkycoinWallet) GetSkycoinWalletType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
