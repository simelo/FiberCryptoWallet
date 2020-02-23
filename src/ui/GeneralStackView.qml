import QtQuick 2.12
import QtQuick.Controls 2.12

Item {
    id: generalStackView

    property alias depth: stackView.depth
    property alias busy: stackView.busy

    signal backRequested()

    function openOutputsPage() {
        if (stackView.depth > 1) {
            stackView.replace(componentOutputs)
        } else {
            stackView.push(componentOutputs)
        }
    }

    function openPendingTransactionsPage() {
        if (stackView.depth > 1) {
            stackView.replace(componentPendingTransactions)
        } else {
            stackView.push(componentPendingTransactions)
        }
    }

        Component.onCompleted: {
            let walletModel = walletManager.getWallets();
            if (walletModel.length > 0 ) {
                if (stackView.depth > 1) {
                    stackView.replace(componentGeneralSwipeView)
                } else {
                    stackView.push(componentGeneralSwipeView)
                }
            } else{
                if (stackView.depth > 1) {
                    stackView.replace(componentPageCreateLoadWallet)
                } else {
                    stackView.push(componentPageCreateLoadWallet)
                }

            }
        }

    function openBlockchainPage() {
        if (stackView.depth > 1) {
            stackView.replace(componentBlockchain)
        } else {
            stackView.push(componentBlockchain)
        }
    }

    function openNetworkingPage() {
        if (stackView.depth > 1) {
            stackView.replace(componentNetworking)
        } else {
            stackView.push(componentNetworking)
        }
    }

    function openSettingsPage() {
        if (stackView.depth > 1) {
            stackView.replace(componentSettings)
        } else {
            stackView.push(componentSettings)
        }
    }

    function openSettingsAddressBookPage() {
        stackView.push(componentSettingsAddressBook)
    }

    function openAddressBookPage() {
        if (stackView.depth > 1) {
            stackView.replace(componentAddressBook)
        } else {
            stackView.push(componentAddressBook)
        }
    }

    function pop() {
        stackView.pop()
    }

    
    StackView {
        id: stackView
        anchors.fill: parent
    }

    Component {
        id: componentPageCreateLoadWallet

        Item {
            PageCreateLoadWallet {
                id: pageCreateLoadWallet
                anchors.centerIn: parent
                width: 400
                height: 400

                onWalletCreationRequested: {
                    stackView.replace(componentGeneralSwipeView)
                    walletManager.createUnencryptedWallet(pageCreateLoadWallet.seed, pageCreateLoadWallet.name, walletManager.getDefaultWalletType() ,0)
                }

                onWalletLoadingRequested:{
                    stackView.replace(componentGeneralSwipeView)
                    walletManager.createUnencryptedWallet(pageCreateLoadWallet.seed, pageCreateLoadWallet.name, walletManager.getDefaultWalletType(), 10)
                }
            }
        }
    }

    Component {
        id: componentGeneralSwipeView

        GeneralSwipeView {
            id: generalSwipeView
        }
    }

    Component {
        id: componentOutputs

        Outputs {
            id: outputs
        }
    }

    Component {
        id: componentPendingTransactions

        PendingTransactions {
            id: pendingTransactions
        }
    }

    Component {
        id: componentBlockchain

        Blockchain {
            id: blockchain
        }
    }

    Component {
        id: componentNetworking

        Networking {
            id: networking
        }
    }

    Component {
        id: componentAddressBook

        AddressBook {
            id: addressBook

            onCanceled: {
                backRequested()
            }
        }
    }

    Component {
        id: componentSettings

        Settings {
            id: settings
        }
    }

    Component {
        id: componentSettingsAddressBook

        SettingsAddressBook {
            id: settingsAddressBook
        }
    }
}
