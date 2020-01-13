package util

import (
	qtCore "github.com/therecipe/qt/core"
	"sort"
)

func init() {
	Map_QmlRegisterType2("ModelUtils", 1, 0, "Map")
}

type Map struct {
	qtCore.QObject
	keyValue map[string]string

	_ func()                  `constructor:"init"`
	_ func(key string) string `slot:"getValue"`
	_ func(key, value string) `slot:"setValue"`
	_ func() []string         `slot:"getKeys"`
}

func (coinOpt *Map) init() {
	coinOpt.keyValue = make(map[string]string)
	coinOpt.ConnectGetValue(coinOpt.getValue)
	coinOpt.ConnectSetValue(coinOpt.setValue)
	coinOpt.ConnectGetKeys(coinOpt.getKeys)
}

func (coinOpt *Map) getValue(key string) string {
	return coinOpt.keyValue[key]
}

func (coinOpt *Map) getKeys() []string {
	if len(coinOpt.keyValue) == 0 {
		return []string{}
	}
	var keysList = make([]string, 0)
	for e := range coinOpt.keyValue {
		keysList = append(keysList, e)
	}
	sort.Strings(keysList)
	return keysList
}

func (coinOpt *Map) setValue(key, value string) {
	coinOpt.keyValue[key] = value
}
