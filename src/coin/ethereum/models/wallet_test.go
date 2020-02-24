package ethereum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWalletDirectoryGetStorage(t *testing.T) {
	dir := "wallet-dir"
	test := []struct {
		wlt  *WalletsDirectory
		want *WalletsDirectory
	}{
		{wlt: NewWalletDirectory(dir), want: NewWalletDirectory(dir)},
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
	test := []struct {
		wlt  *WalletsDirectory
		want *WalletsDirectory
	}{
		{wlt: NewWalletDirectory(dir), want: NewWalletDirectory(dir)},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("GetWalletSet_%d", i), func(t *testing.T) {
			storage := tt.wlt.GetWalletSet()
			require.Equal(t, tt.want, storage)
		})
	}
}
