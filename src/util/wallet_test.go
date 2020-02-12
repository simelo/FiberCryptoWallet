package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

func TestSimpleWalletAddress(t *testing.T) {
	wllt, addr := new(mocks.Wallet), new(mocks.Address)
	wa := &SimpleWalletAddress{Wallet: wllt, Address: addr}
	addr.On("String").Return("A1234")
	addr.On("IsBip32").Return(true)
	addr.On("GetCryptoAccount").Return(new(mocks.CryptoAccount))
	addr.On("Bytes").Return([]byte{1, 2, 3, 4})
	addr.On("Checksum").Return(core.Checksum([]byte{5, 6, 7, 8}))
	pk1 := new(mocks.PubKey)
	addr.On("Verify", pk1).Return(nil)

	require.Equal(t, wllt, wa.GetWallet())
	require.Equal(t, "A1234", wa.String())
	require.Equal(t, true, wa.IsBip32())
	require.Equal(t, addr.GetCryptoAccount(), wa.GetCryptoAccount())
	require.Equal(t, addr.Bytes(), wa.Bytes())
	require.Equal(t, addr.Checksum(), wa.Checksum())
	require.Nil(t, wa.Verify(pk1))

	addr = new(mocks.Address)
	addr.On("Verify", pk1).Return(errors.New("Error"))
	wa = &SimpleWalletAddress{Wallet: wllt, Address: addr}
	require.NotNil(t, addr.Verify(pk1))
	require.NotNil(t, wa.Verify(pk1))
}

func TestSimpleWalletOutput(t *testing.T) {
	wllt, out := &mocks.Wallet{}, &mocks.TransactionOutput{}
	out.On("GetId").Return("ID1")
	out.On("IsSpent").Return(true)
	out.On("GetAddress").Return(new(mocks.Address))
	out.On("GetCoins", "X").Return(uint64(2000), nil)
	out.On("SupportedAssets").Return([]string{"X", "Y", "Z"})
	wa := &SimpleWalletOutput{Wallet: wllt, UxOut: out}

	require.Equal(t, wllt, wa.GetWallet())
	require.Equal(t, "ID1", wa.GetId())
	require.Equal(t, true, wa.IsSpent())
	require.Equal(t, out.GetAddress(), wa.GetAddress())
	coins, err := wa.GetCoins("X")
	require.NoError(t, err)
	require.Equal(t, uint64(2000), coins)
	require.Equal(t, out.SupportedAssets(), wa.SupportedAssets())
}
