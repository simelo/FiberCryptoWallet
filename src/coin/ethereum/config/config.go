package config

import (
	"encoding/json"
	"strings"

	local "github.com/fibercrypto/fibercryptowallet/src/main"
	"github.com/fibercrypto/fibercryptowallet/src/util"
)

const (
	SettingPathToNode         = "node"
	SettingPathToWalletSource = "walletSource"
	SectionName               = "ethereum"
	LocalWallet               = "local"
)

var (
	sectionManager *local.SectionManager
)

func RegisterConfig() error {
	cm := local.GetConfigManager()
	node := map[string]string{"address": "https://mainnet.infura.io/v3/18c1bbac31de4d56a4295d4cdb6709ad"}
	nodeBytes, err := json.Marshal(node)
	if err != nil {
		return err
	}

	nodeOpt := local.NewOption(SettingPathToNode, []string{}, false, string(nodeBytes))

	walletsDefaultDirectory := util.GetMultiPlatformUserDirectory([]string{".ethereum", "wallets"})
	wltSrc := &walletSource{
		id:     "1",
		Tp:     string(LocalWallet),
		Source: walletsDefaultDirectory,
	}
	wltSrcBytes, err := json.Marshal(wltSrc)
	if err != nil {
		return err
	}

	wltOpt := local.NewOption(string(wltSrc.id), []string{SettingPathToWalletSource}, false, string(wltSrcBytes))

	sectionManager = cm.RegisterSection(SectionName, []*local.Option{nodeOpt, wltOpt})

	return nil

}

func GetOption(path string) (string, error) {
	stringList := strings.Split(path, "/")
	return sectionManager.GetValue(stringList[len(stringList)-1], stringList[:len(stringList)-1])
}

func GetWalletSources() ([]*walletSource, error) {
	wltsString, err := getValues(SettingPathToWalletSource)
	if err != nil {
		return nil, err
	}
	wltSrcs := make([]*walletSource, len(wltsString))
	for i, wlt := range wltsString {
		wltSrcs[i] = new(walletSource)
		err = json.Unmarshal([]byte(wlt), wltSrcs[i])
		if err != nil {
			return nil, err
		}
	}
	return wltSrcs, nil
}

type walletSource struct {
	id     string
	Tp     string `json:"SourceType"`
	Source string `json:"Source"`
}

func getValues(prefix string) ([]string, error) {
	return sectionManager.GetValues(strings.Split(prefix, "/"))
}
