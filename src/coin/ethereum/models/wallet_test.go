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

func TestWalletDirectoryEncrypt(t *testing.T) {
	wltDir, err := NewWalletDirectory("testdata")
	require.Nil(t, err)

	err = wltDir.Encrypt("wrong_wallet_id", func(string, core.KeyValueStore) (string, error) {
		return "", nil
	})
	require.EqualError(t, err, errors.ErrWalletNotFound.Error())

	var wlt *KeystoreWallet
	for _, val := range wltDir.wallets {
		wlt = val
	}

	err = wltDir.Encrypt(wlt.dirName, func(string, core.KeyValueStore) (string, error) {
		return "", nil
	})
	require.EqualError(t, err, errors.ErrWalletIsNotDecrypted.Error())

	err = wltDir.Decrypt(wlt.dirName, func(string, core.KeyValueStore) (string, error) {
		return "test", nil
	})
	require.Nil(t, err)
	require.Equal(t, "test", wltDir.walletsPasswords[wlt.dirName])

	err = wltDir.Encrypt(wlt.dirName, func(string, core.KeyValueStore) (string, error) {
		return "", error_pkg.New("Error in password reader")
	})
	require.EqualError(t, err, "Error in password reader")

	err = wltDir.Encrypt(wlt.dirName, func(string, core.KeyValueStore) (string, error) {
		return "test2", nil
	})
	require.Nil(t, err)
	_, ok := wltDir.walletsPasswords[wlt.dirName]
	require.Equal(t, false, ok)

	err = wltDir.Decrypt(wlt.dirName, func(string, core.KeyValueStore) (string, error) {
		return "test2", nil
	})
	require.Nil(t, err)
	require.Equal(t, "test2", wltDir.walletsPasswords[wlt.dirName])

	err = wltDir.Encrypt(wlt.dirName, func(string, core.KeyValueStore) (string, error) {
		return "test", nil
	})
	require.Nil(t, err)
	_, ok = wltDir.walletsPasswords[wlt.dirName]
	require.Equal(t, false, ok)

}

func TestWalletDirectoryIsEncrypted(t *testing.T) {
	wltDir, err := NewWalletDirectory("testdata")
	require.Nil(t, err)

	encrypted, err := wltDir.IsEncrypted("wrong_wallet_name")
	require.EqualError(t, err, errors.ErrNotFound.Error())

	var wlt *KeystoreWallet

	for _, w := range wltDir.wallets {
		wlt = w
	}

	encrypted, err = wltDir.IsEncrypted(wlt.dirName)
	require.Nil(t, err)
	require.Equal(t, true, encrypted)

	err = wltDir.Decrypt(wlt.dirName, func(string, core.KeyValueStore) (string, error) {
		return "test", nil
	})
	require.Nil(t, err)

	encrypted, err = wltDir.IsEncrypted(wlt.dirName)
	require.Nil(t, err)
	require.Equal(t, false, encrypted)

	err = wltDir.Encrypt(wlt.dirName, func(string, core.KeyValueStore) (string, error) {
		return "test", nil
	})
	require.Nil(t, err)
}

func TestWalletDirectoryListWallet(t *testing.T) {
	wltDir, err := NewWalletDirectory("testdata")
	require.Nil(t, err)

	wltIter := wltDir.ListWallets()
	require.NotNil(t, wltIter)

	require.Equal(t, true, wltIter.Next())

	wlt := wltIter.Value()
	require.NotNil(t, wlt)
	require.Equal(t, "test", wlt.GetLabel())
}

func TestWalletDirectoryGetWallet(t *testing.T) {
	wltDir, err := NewWalletDirectory("testdata")
	require.Nil(t, err)

	wlt, err := wltDir.GetWallet("wrong_wallet_id")
	require.EqualError(t, err, errors.ErrNotFound.Error())

	var origWlt *KeystoreWallet
	for _, sWlt := range wltDir.wallets {
		origWlt = sWlt
	}

	wlt, err = wltDir.GetWallet(origWlt.GetId())
	require.Nil(t, err)

	require.Equal(t, origWlt, wlt)
}
