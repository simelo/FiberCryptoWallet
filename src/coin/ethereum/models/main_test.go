package ethereum

import (
	"os"
	"testing"

	"github.com/fibercrypto/fibercryptowallet/src/core"

	"github.com/fibercrypto/fibercryptowallet/src/util/logging"

	"github.com/stretchr/testify/mock"
)

var global_mock *EthereumApiMock

var logModelTest = logging.MustGetLogger("Ethereum Model Test")

func CleanGlobalMock() {
	global_mock.ExpectedCalls = []*mock.Call{}
}

func TestMain(m *testing.M) {
	if global_mock == nil {
		global_mock = new(EthereumApiMock)
	}

	err := core.GetMultiPool().CreateSection(PoolSection, global_mock)
	if err != nil {
		logModelTest.WithError(err).Warn("Error creating pool section")
		return
	}

	//TODO Register ethereum plugin
	//util.RegisterAltcoin()
	os.Exit(m.Run())

}
