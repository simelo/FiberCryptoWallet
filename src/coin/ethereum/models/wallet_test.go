package ethereum

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/fibercrypto/fibercryptowallet/src/core"

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
		{wlt: wlt, want: wlt},
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

func TestWalletDirectory(t *testing.T) {
	wltDir, err := NewWalletDirectory("testdata")
	require.Nil(t, err)
	wltIter := wltDir.ListWallets()
	var wlt core.Wallet
	for wltIter.Next() {
		wlt = wltIter.Value()
	}
	spew.Dump(wlt)
	fmt.Println(wlt.GetId(), " ", wlt.GetLabel())

}
