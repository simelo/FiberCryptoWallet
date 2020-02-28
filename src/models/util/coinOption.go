package util

import (
	qtCore "github.com/therecipe/qt/core"
	"reflect"
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
	_        func(key string)        `slot:"removeVal"`
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

func (coinOpt *Map) SetValueAsync(key, value string) {
	var exist = false
	// Verify if the KeyList contains the key
	for e := range coinOpt.keyList {
		if coinOpt.keyList[e] == key {
			exist = true
			break
		}
	}

	// If not contains the key, add the key to the KeyList
	if !exist {
		coinOpt.keyList = append(coinOpt.keyList, key)
	}

	coinOpt.keyValue[key] = value
}

func (coinOpt *Map) setValue(key, value string) {
	var exist = false
	// Verify if the KeyList contains the key
	for e := range coinOpt.keyList {
		if coinOpt.keyList[e] == key {
			exist = true
			break
		}
	}

	// If not contains the key, add the key to the KeyList
	if !exist {
		coinOpt.keyList = append(coinOpt.keyList, key)
	}

	coinOpt.keyValue[key] = value
}

func (coinOpt *Map) removeVal(key string) {
	var pos = -1
	// Find the position of the key in the list
	for e := range coinOpt.keyList {
		if coinOpt.keyList[e] == key {
			pos = e
			break
		}
	}

	if pos != -1 {
		// Delete the element in the position
		coinOpt.keyList = append(coinOpt.keyList[:pos], coinOpt.keyList[pos+1:]...)
		delete(coinOpt.keyValue, key)
	}
}

func CompareMaps(a, b *Map) bool {
	return reflect.DeepEqual(a.keyValue, b.keyValue)
}
