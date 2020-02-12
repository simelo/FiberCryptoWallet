package util

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
)

func TestSimpleWalletAddress(t *testing.T) {
	wllt, addr := &mocks.Wallet{}, &mocks.Address{}
	wa := &SimpleWalletAddress{Wallet: wllt, Address: addr}
	addr.On("String").Return("A1234")
	addr.On("IsBip32").Return(true)
	addr.On("GetCryptoAccount").Return(new(mocks.CryptoAccount))
	addr.On("Bytes").Return([]byte{1, 2, 3, 4})
	addr.On("Checksum").Return([]byte{5, 6, 7, 8})
	pk1 := new(mocks.PubKey)
	pk2 := new(mocks.PubKey)
	addr.On("Verify", pk1).Return(true)
	addr.On("Verify", pk2).Return(false)

	require.Equal(t, wllt, wa.GetWallet())
	require.Equal(t, "A1234", wa.String())
	require.Equal(t, true, wa.IsBip32())
	require.Equal(t, addr.GetCryptoAccount(), wa.GetCryptoAccount())
	require.Equal(t, addr.Bytes(), wa.Bytes())
	require.Equal(t, addr.Checksum(), wa.Checksum())
	require.Equal(t, true, wa.Verify(pk1))
	require.Equal(t, false, wa.Verify(pk2))
}

func TestSimpleWalletOutput(t *testing.T) {
	wllt, out := &mocks.Wallet{}, &mocks.TransactionOutput{}
	wllt.On("GetId").Return("ID1")
	wllt.On("IsSpent").Return(true)
	wllt.On("GetAddress").Return(new(mocks.Address))
	wllt.On("GetCoins", "X").Return(2000, nil)
	wllt.On("SupportedAssets").Return([]string{"X", "Y", "Z"})
	wa := &SimpleWalletOutput{Wallet: wllt, UxOut: out}

	require.Equal(t, wllt, wa.GetWallet())
	require.Equal(t, "ID1", wa.GetId())
	require.Equal(t, true, wa.IsSpent())
	require.Equal(t, out.GetAddress(), wa.GetAddress())
	coins, err := wa.GetCoins("X")
	require.NoError(t, err)
	require.Equal(t, 2000, coins)
	require.Equal(t, out.SupportedAssets(), wa.SupportedAssets())
}
