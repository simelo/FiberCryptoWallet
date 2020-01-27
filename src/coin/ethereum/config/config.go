package config

import (
	"encoding/json"
	"strings"

	local "github.com/fibercrypto/fibercryptowallet/src/main"
)

const (
	SettingPathToNode = "node"
	SectionName       = "ethereum"
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

	sectionManager = cm.RegisterSection(SectionName, []*local.Option{nodeOpt})

	return nil

}

func GetOption(path string) (string, error) {
	stringList := strings.Split(path, "/")
	return sectionManager.GetValue(stringList[len(stringList)-1], stringList[:len(stringList)-1])
}
