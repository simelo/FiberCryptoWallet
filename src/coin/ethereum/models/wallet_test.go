package ethereum

import (
	"fmt"
	"testing"

	error_pkg "errors"

	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/stretchr/testify/require"
)

func TestWalletDirectoryGetStorage(t *testing.T) {
	dir := "wallet-dir"
	wlt, err := NewWalletDirectory(dir)
	require.Nil(t, err)
	want, err := NewWalletDirectory(dir)
	require.Nil(t, err)
	test := []struct {
		wlt  *WalletsDirectory
		want *WalletsDirectory
	}{
		{wlt: wlt, want: want},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("GetStorage_%d", i), func(t *testing.T) {
			storage := tt.wlt.GetStorage()
			require.Equal(t, tt.want, storage)
		})
	}
}

func TestWalletDirectoryGetWalletSet(t *testing.T) {
	dir := "wallet-dir"
	wlt, err := NewWalletDirectory(dir)
	require.Nil(t, err)
	want, err := NewWalletDirectory(dir)
	require.Nil(t, err)
	test := []struct {
		wlt  *WalletsDirectory
		want *WalletsDirectory
	}{
		{wlt: wlt, want: want},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("GetWalletSet_%d", i), func(t *testing.T) {
			storage := tt.wlt.GetWalletSet()
			require.Equal(t, tt.want, storage)
		})
	}
}

func TestWalletDirectoryDecrypt(t *testing.T) {
	wltDir, err := NewWalletDirectory("testdata")
	require.Nil(t, err)

	err = wltDir.Decrypt("wrong_wallet_id", func(string, core.KeyValueStore) (string, error) {
		return "", nil
	})
	require.EqualError(t, err, errors.ErrWalletNotFound.Error())

	wltIter := wltDir.ListWallets()
	var wlt core.Wallet
	for wltIter.Next() {
		wlt = wltIter.Value()
	}

	err = wltDir.Decrypt(wlt.GetId(), func(string, core.KeyValueStore) (string, error) {
		return "", error_pkg.New("Error in password reader")
	})
	require.EqualError(t, err, "Error in password reader")

	pass := wltDir.walletsPasswords[wlt.GetId()]
	require.Equal(t, pass, "")

	err = wltDir.Decrypt(wlt.GetId(), func(string, core.KeyValueStore) (string, error) {
		return "wrong_pass", nil
	})

	require.NotNil(t, err)

	err = wltDir.Decrypt(wlt.GetId(), func(string, core.KeyValueStore) (string, error) {
		return "test", nil
	})
	require.Nil(t, err)

	pass = wltDir.walletsPasswords[wlt.GetId()]
	require.Equal(t, pass, "test")

}
