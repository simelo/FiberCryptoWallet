package util

import (
	"github.com/skycoin/skycoin/src/api"
)

//Return address of daemon node----TODO
func nodeAddress() string {
	return "https://node.skycoin.net"
}

//Return username of daemon node----TODO
func nodeUsername() string {
	return "Kid"
}

//Return password of daemon node-----TODO
func nodePassword() string {
	return "P@ssw0rd!"
}

// NewClient example
func NewClient() *api.Client {
	c := api.NewClient(nodeAddress())
	//c.SetAuth(nodeUsername(), nodePassword())
	return c
}
