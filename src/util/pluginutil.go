package util

import (
	"errors"
	"fmt"
	params2 "github.com/fibercrypto/fibercryptowallet/src/params"
	"math"
	"strings"

	"github.com/fibercrypto/fibercryptowallet/src/core"
	local "github.com/fibercrypto/fibercryptowallet/src/main"
)

func AltcoinCaption(ticker string) string {
	if info, isRegistered := local.LoadAltcoinManager().DescribeAltcoin(ticker); isRegistered {
		return info.Name
	}
	return ticker + " <Unregistered>"
}

func AltcoinQuotient(ticker string) (uint64, error) {
	if info, isRegistered := local.LoadAltcoinManager().DescribeAltcoin(ticker); isRegistered {
		return uint64(math.Pow(float64(10), float64(info.Accuracy))), nil
	}
	return uint64(0), errors.New(ticker + " <Unregistered>")
}

func RegisterAltcoin(p core.AltcoinPlugin) {
	local.LoadAltcoinManager().RegisterPlugin(p)
}

// LookupSignerByUID search for signer matching given ID
func LookupSignerByUID(wlt core.Wallet, id core.UID) core.TxnSigner {
	wltSigner, isSigner := wlt.(core.TxnSigner)
	// Reference to self
	if id == core.UID("") {
		if isSigner {
			return wltSigner
		}
		return nil
	}
	// Wallet matches ID
	if isSigner && wltSigner.GetSignerUID() == id {
		return wltSigner
	}
	// Lookup global signers
	return local.LoadAltcoinManager().LookupSignService(id)
}

// AddressFromString returns a core.Address if match with string address.
// If the coinTicket parameter not match with any address type returns 'coinTicket not match' error.
func AddressFromString(addrs, coinTicket string) (core.Address, error) {
	altPlugin, ok := local.LoadAltcoinManager().LookupAltcoinPlugin(coinTicket)
	if !ok {
		return nil, errors.New("coinTicket not match")
	}
	return altPlugin.AddressFromString(addrs)
}

// LoadBlockchainStatus load the blockchainStatus for each ticket
func LoadBlockchainStatus(ticket string) (core.BlockchainStatus, error) {
	altcoinPlugin, ok := local.LoadAltcoinManager().LookupAltcoinPlugin(ticket)
	if !ok {
		return nil, fmt.Errorf("ticket %s don't match", ticket)
	}

	return altcoinPlugin.LoadBlockchainStatus(params2.DataRefreshTimeout), nil
}

// LoadAllExistentCurrencies load all existent currencies
func LoadAllExistentCurrencies() []string {
	var currenciesList = make([]string, 0)
	for _, plugin := range local.LoadAltcoinManager().ListRegisteredPlugins() {
		currenciesList = append(currenciesList, plugin.ListSupportedAltcoins()[0].Name)
	}

	return currenciesList
}

func GetTicketOfCurrncies(currency string) (string, error) {
	for _, plugin := range local.LoadAltcoinManager().ListRegisteredPlugins() {
		for _, altcoin := range plugin.ListSupportedAltcoins() {
			if strings.ToLower(currency) == strings.ToLower(altcoin.Name) {
				return altcoin.Ticker, nil
			}
		}

	}
	return "", fmt.Errorf("currency %s don't match", currency)
}
