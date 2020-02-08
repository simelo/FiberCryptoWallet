package ethereum

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
)

var logWallet = logging.MustGetLogger("Ethereum Wallet")

const (
	WalletTypeStandard = "standard"
)

func NewWalletDirectory(path string) *WalletsDirectory {
	return &WalletsDirectory{
		path: path,
	}
}

// Implements WalletEnv, WalletStorage and WalletSet interfaces
type WalletsDirectory struct {
	path              string
	wallets           map[string]*KeystoreWallet
	mutexForWallets   sync.Mutex
	walletsPasswords  map[string]string
	mutexForPasswords sync.Mutex
}

//WalletEnv methods set
func (walletDir *WalletsDirectory) GetStorage() core.WalletStorage {
	return walletDir
}

func (walletDir *WalletsDirectory) GetWalletSet() core.WalletSet {
	return walletDir
}

//Remove unecessary line
func (walletDir *WalletsDirectory) getWallet(wltId string) (*KeystoreWallet, bool) {
	walletDir.mutexForWallets.Lock()
	defer walletDir.mutexForWallets.Unlock()
	wlt, exist := walletDir[wltId]
	return wlt, exist
}

func (walletDir *WalletsDirectory) setWallet(wltId string, wlt *KeystoreWallet) {
	walletDir.mutexForWallets.Lock()
	defer walletDir.mutexForWallets.Unlock()
	walletDir.wallets[wltId] = wlt
}

//Remove unecessary line
func (walletDir *WalletsDirectory) getPassword(wltId string) (string, bool) {
	walletDir.mutexForPasswords.Lock()
	defer walletDir.mutexForPasswords.Unlock()
	pass, exist := walletDir.walletsPasswords[wltId]
	return pass, exist
}

func (walletDir *WalletsDirectory) setPassword(wltId, passowrd string) {
	walletDir.mutexForWallets.Lock()
	defer walletDir.mutexForWallets.Unlock()
	walletDir.walletsPasswords[wltId] = passowrd
}

//WalletStorage methods set
func (walletDir *WalletsDirectory) Encrypt(walletName string, password core.PasswordReader) error {
	wlt, exist := walletDir.getWallet(walletName)
	if !exist {
		logWallet.WithError(errors.ErrWalletNotFound).Error("Error encrypting wallet")
		return errors.ErrWalletNotFound
	}

	actualPassword, decrypt := walletDir.getPassword(walletName)
	if !decrypt {
		logWallet.WithError(errors.ErrWalletIsNotDecrypted).Error("Error encrypting wallet")
		return errors.ErrWalletIsNotDecrypted
	}
	pwdCtx := util.NewKeyValueMap()
	pwdCtx.SetValue(core.StrTypeName, core.TypeNameWalletStorage)
	pwdCtx.SetValue(core.StrMethodName, "Encrypt")
	pwdCtx.SetValue(core.StrWalletName, walletName)
	pwdCtx.SetValue(core.StrWalletLabel, wlt.GetName())
	newPassword, err := password(fmt.Sprintf("Enter password for %s", wlt.GetName()), pwdCtx)
	if err != nil {
		logWallet.WithError(err).Error("Error encrypting wallet")
		return err
	}

	if actualPassword == newPassword {
		for _, acc := range wlt.Accounts() {
			err := wlt.Lock(acc.Address)
			if err != nil {
				logWallet.WithError(err).Error("Couldn't encrypt wallet properly")
				return err
			}
		}
		delete(walletDir.walletsPasswords, walletName)
		return nil
	}

	err = updateWallet(wlt, actualPassword, newPassword)
	if err != nil {
		logWallet.WithError(err).Error("Couldn't encrypt wallet")
		return err
	}

	delete(walletDir.walletsPasswords, walletName)
	return nil
}

func (walletDir *WalletsDirectory) Decrypt(walletName string, password core.PasswordReader) error {
	wlt, exist := walletDir.getWallet(walletName)
	if !exist {
		logWallet.WithError(errors.ErrWalletNotFound).Error("Couldn't decrypt wallet")
		return errors.ErrWalletNotFound
	}

	pwdCtx := util.NewKeyValueMap()
	pwdCtx.SetValue(core.StrTypeName, core.TypeNameWalletStorage)
	pwdCtx.SetValue(core.StrMethodName, "Decrypt")
	pwdCtx.SetValue(core.StrWalletName, walletName)
	pwdCtx.SetValue(core.StrWalletLabel, wlt.GetName())
	pwd, err := password(fmt.Sprintf("Enter password for %s", wlt.GetName()), pwdCtx)
	if err != nil {
		logWallet.WithError(err).Error("Error decrypting wallet")
		return err
	}
	for _, acc := range wlt.Accounts() {
		err = wlt.Unlock(acc, pwd)
		if err != nil {
			logWallet.WithError(err).Error("Error decrypting wallet")
			return err
		}
	}

	walletDir.setPassword(walletName, pwd)
	return nil
}

func (walletDir *WalletsDirectory) IsEncrypted(walletName string) (bool, error) {
	_, exist := walletDir.getWallet(walletName)
	if !exist {
		logWallet.WithError(errors.ErrNotFound).Error("Couldn't check if wallet is encrypted")
		return false, errors.ErrNotFound
	}
	_, decrypt := walletDir.getPassword(walletName)

	return decrypt, nil
}

func updateWallet(wlt *KeystoreWallet, password, newPassword string) error {
	tempDir, err := ioutil.TempDir("", "ethWlt")
	if err != nil {
		return err
	}

	tempWallet := keystore.NewKeyStore(tempDir, keystore.StandardScryptN, keystore.StandardScryptP)
	defer os.RemoveAll(tempDir)

	for _, acc := range wlt.Accounts() {
		accBytes, err := wlt.Export(acc, password, newPassword)
		if err != nil {
			return err
		}
		_, err = tempWallet.ImportPreSaleKey(accBytes, newPassword)
		if err != nil {
			return err
		}
	}
	newAccs, err := ioutil.ReadDir(tempDir)
	if err != nil {
		return err
	}

	for _, accFile := range newAccs {
		data, err := ioutil.ReadFile(accFile.Name())
		if err != nil {
			return err
		}
		name := filepath.Dir(accFile.Name())
		err = ioutil.WriteFile(filepath.Join(wlt.dirName, name), data, 0644)
		if err != nil {
			return err
		}
	}
	allFiles, err := ioutil.ReadDir(wlt.dirName)
	if err != nil {
		return err
	}
	for _, file := range allFiles {
		founded := false
		for _, f := range newAccs {
			if file.Name() == f.Name() {
				founded = true
				break
			}
		}
		if !founded {
			os.Remove(file.Name())
		}
	}
	return nil

}

//WalletSet methods set
func (walletDir *WalletsDirectory) ListWallets() core.WalletIterator {
	wallets := make([]core.Wallet, 0)
	walletDir.mutexForWallets.Lock()
	for _, wlt := range walletDir.wallets {
		wallets = append(wallets, wlt)
	}
	walletDir.mutexForWallets.Unlock()
	return NewKeyStoreWalletIterator(wallets)
}

func (walletDir *WalletsDirectory) GetWallet(id string) (core.Wallet, error) {
	wlt, exist := walletDir.getWallet(id)
	if !exist {
		logWallet.WithError(errors.ErrNotFound).Error("Error getting wallet")
		return nil, errors.ErrNotFound
	}
	return wlt, nil
}

func (walletDir *WalletsDirectory) SupportedWalletType() []string {
	return []string{WalletTypeStandard}
}

func (walletDir *WalletsDirectory) CreateWallet(name, seed, walletType string, isEncrypted bool, password core.PasswordReader, scanAddressesN int) (core.Wallet, error) {
	if !util.IsValidWalletType(walletType, walletDir) {
		logWallet.WithError(errors.ErrInvalidWalletType).Error("Error creating wallet")
		return nil, errors.ErrInvalidWalletType
	}

	wltId := walletDir.generateUniqueId(name)

	pwd := ""
	if isEncrypted {
		pwdCtx := util.NewKeyValueMap()
		pwdCtx.SetValue(core.StrTypeName, core.TypeNameWalletStorage)
		pwdCtx.SetValue(core.StrMethodName, "CreateWallet")
		pwdCtx.SetValue(core.StrWalletName, walletName)
		pwdCtx.SetValue(core.StrWalletLabel, wlt.GetName())
		pwd, err := password(fmt.Sprintf("Enter password for %s", wlt.GetName()), pwdCtx)
		if err != nil {
			logWallet.WithError(err).Error("Error creating wallet")
			return nil, err
		}
	}

	wltDir := filepath.Join(walletDir.path)
	// Review file mode
	err = os.Mkdir(wltDir, 0755)
	if err != nil {
		logWallet.WithError(err).Error("Error creating wallet")
		return nil, err
	}

	nWlt := NewKeystoreWallet(wltDir, name)
	walletDir.setWallet(wltId, nWlt)

	_, err := nWlt.NewAccount(pwd)
	if err != nil {
		logWallet.WithError(err).Error("Error creating wallet")
		os.RemoveAll(wltDir)
		delete(walletDir.wallets, wltId)
		return nil, err
	}

	if !isEncrypted {
		walletDir.setPassword(wltId, "")
	}

	return nWlt, nil
}

func (walletDir *WalletsDirectory) DefaultWalletType() string {
	return WalletTypeStandard
}

func (walletDir *WalletsDirectory) generateUniqueId(name string) string {
	idBase := fmt.Sprintf("%s_%d", name, time.Now().UTC().UnixNano())
	id := idBase
	cont = 0
	for {
		validId := true
		walletDir.mutexForWallets.Lock()
		for wltId, _ := range walletDir.wallets {
			if wltId == id {
				validId = false
				break
			}
		}
		walletDir.mutexForWallets.Unlock()
		if !validId {
			cont++
			id := fmt.Sprintf("%s_%d", idBase, cont)
			continue
		}
		break
	}
	return id
}

func NewKeystoreWallet(dir, name string) *KeystoreWallet {
	ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	return &KeystoreWallet{
		KeyStore: ks,
		name:     name,
		dirName:  dir,
	}
}

type KeystoreWallet struct {
	*keystore.KeyStore
	name    string
	dirName string
}

//Wallet methods set
func (kw *KeystoreWallet) GetId() string {
	return kw.dirName
}

func (kw *KeystoreWallet) GetLabel() string {
	return kw.name
}

func (kw *KeystoreWallet) SetLabel(name string) {
	return
}

func (kw *KeystoreWallet) SendFromAddress(from []core.Address, to []core.TransactionOutput, change core.Address, options core.KeyValueStore) (core.Transaction, error) {
	if len(from) != 1 {
		logWallet.WithError(errors.ErrInvalidFromAddresses).Error("Couldn't send transaction")
	}

	types.NewTransa

}

func (kw *KeystoreWallet) Spend(unspent, new []core.TransactionOutput, change core.Address, options core.KeyValueStore) (core.Transaction, error) {
	return nil, errors.ErrNotImplemented
}

func (kw *KeystoreWallet) Transfer(to core.TransactionOutput, options core.KeyValueStore) (core.Transaction, error) {
	return nil, errors.ErrNotImplemented
}
func NewKeyStoreWalletIterator(wallets []core.Wallet) *KeystoreWalletIterator {
	return &KeystoreWalletIterator{
		wallets: wallets,
		current: -1,
	}
}

type KeystoreWalletIterator struct {
	current int
	wallets []core.Wallet
}

func (it *KeystoreWalletIterator) Value() core.Wallet {
	return it.wallets[it.current]
}

func (it *KeystoreWalletIterator) Next() bool {
	if it.HasNext() {
		it.current++
		return true
	}
	return false
}

func (it *KeystoreWalletIterator) HasNext() bool {
	return !((it.current + 1) >= len(it.wallets))
}
