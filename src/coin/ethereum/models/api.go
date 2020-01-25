package ethereum

import (
	"github.com/stretchr/testify/mock"
)

type EthereumApiMock struct {
	mock.Mock
}

func (m *EthereumApiMock) Create() (interface{}, error) {
	return m, nil
}
