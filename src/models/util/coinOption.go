package util

import (
	qtCore "github.com/therecipe/qt/core"
)

func init() {
	Map_QmlRegisterType2("ModelUtils", 1, 0, "Map")
}

type Map struct {
	qtCore.QObject
	keyValue map[string]string
	keyList  []string
	_        func()                  `constructor:"init"`
	_        func(key string) string `slot:"getValue"`
	_        func(key, value string) `slot:"setValue"`
	_        func() []string         `slot:"getKeys"`
}

func (coinOpt *Map) init() {
	coinOpt.keyValue = make(map[string]string)
	coinOpt.keyList = make([]string, 0)
	coinOpt.ConnectGetValue(coinOpt.getValue)
	coinOpt.ConnectSetValue(coinOpt.setValue)
	coinOpt.ConnectGetKeys(coinOpt.getKeys)
}

func (coinOpt *Map) getValue(key string) string {
	return coinOpt.keyValue[key]
}

func (coinOpt *Map) getKeys() []string {
	return coinOpt.keyList
}

func (coinOpt *Map) setValue(key, value string) {
	coinOpt.keyList = append(coinOpt.keyList, key)
	coinOpt.keyValue[key] = value
}
