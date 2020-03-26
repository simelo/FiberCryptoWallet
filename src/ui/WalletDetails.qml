import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import Address 1.0
import ModelUtils 1.0

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

// Backend imports

import "./Dialogs"
Item {
    id: root

    property string name
    property string coin
    property string fileName
    property Map coinOpts
    property bool isEncrypted
    property int index
    property alias specificDetails : gridLayoutBasicInfo
    property alias editWallet : dialogEditWallet
    property bool emptyAddressVisible: true
    readonly property bool expanded: buttonMoreDetails.checked


    readonly property real basicHeight: 80 + rowLayoutBasicDetails.height
    Behavior on implicitHeight { NumberAnimation { duration: 1000; easing.type: Easing.OutQuint } }

    function updateWalletEncryption(){
        if (isEncrypted){
            dialogGetPassword.open()
        }else{
            dialogSetPassword.open()
        }
    }

    implicitWidth: 600
    clip: true

    QAddressList{
        id: addressList
        Component.onCompleted:{
            addressList.start() //start Async
        }
    }

    onFileNameChanged:{
        walletManager.loadAddressModelByWallet(fileName, addressList)
    // if wallet is different change the model
    // else continue updating
    }

    ColumnLayout {
        id: columnLayoutRoot
        anchors.fill: parent
        spacing: 20

        RowLayout {
            id: rowLayoutBasicDetails
            Layout.alignment: Qt.AlignTop
            Layout.fillWidth: true

            ColumnLayout {
            id: columnLayoutBasicDetails
                Layout.alignment: Qt.AlignTop
                Layout.fillWidth: true
                Layout.topMargin: 5

                GridLayout {
                    id: gridLayoutBasicInfo
                    Material.foreground: Material.Grey
                    columns: 2
                    columnSpacing: 10
                    Layout.alignment: Qt.AlignTop
                    Layout.fillWidth: true
                } // GridLayout
                RowLayout {
                    Layout.alignment: Qt.AlignTop
                    Layout.fillWidth: true
                    Label {
                        text: qsTr("Encryption")
                        font.bold: true
                        color: Material.hintTextColor
                    }
                    Switch {
                        id: encryptSwitch
                        checked: isEncrypted
                        onToggled: {
                        updateWalletEncryption();
                        }
                    }

                    Image {
                        id: lockIcon
                        source: "qrc:/images/resources/images/icons/lock" + (isEncrypted ? "On" : "Off") + ".svg"
                        sourceSize: "24x24"
                    }
                } // RowLayout
            }
        } // RowLayout

        RowLayout {
            id: rowLayoutDetailsSeparator

            Layout.alignment: Qt.AlignTop
            Layout.fillWidth: true

            Button {
                id: buttonMoreDetails
                implicitWidth: 200
                flat: true
                checkable: true
                text: checked ? qsTr("Hide Addresses") : qsTr("Show Addresses")
            }

            Rectangle {
                height: 1
                color: Material.color(Material.Grey)
                Layout.alignment: Qt.AlignVCenter
                Layout.fillWidth: true
            }
        }

        RowLayout {
            id: rowLayoutMoreDetails

            Layout.alignment: Qt.AlignTop
            Layout.fillWidth: true
            Layout.fillHeight: true

            opacity: expanded ? 1 : 0
            Behavior on opacity { NumberAnimation { duration: 1000; easing.type: Easing.OutQuint } }

            ColumnLayout {
                id: columnLayoutInputs
                Layout.fillWidth: true
                Layout.fillHeight: true
                RowLayout{
                    Layout.fillWidth: true
                     ToolButton {
                        id: buttonToggleVisibility
                        text: qsTr("Show empty")
                        checkable: true
                        checked: emptyAddressVisible
                        icon.source: "qrc:/images/resources/images/icons/visible" + (checked ? "On" : "Off") + ".svg"
                        Material.accent: Material.Indigo
                        Material.foreground: Material.Grey
                        Layout.fillWidth: true

                        onCheckedChanged: {
                            emptyAddressVisible = checked
                        }
                    }
                    ToolButton {
                        id: buttonAddAddress
                        text: qsTr("Add address")
                        icon.source: "qrc:/images/resources/images/icons/add.svg"
                        Material.accent: Material.Teal
                        Material.foreground: Material.accent
                        Layout.fillWidth: true

                        onClicked: {
                         dialogAddAddresses.open()
                        }
                    }
                }
                ScrollView {
                    Layout.alignment: Qt.AlignTop
                    Layout.fillWidth: true
                    Layout.fillHeight: true

                    ListView {
                        id: listViewInputs

                        Material.foreground: Material.Grey
                        model: addressList
                        clip: true
                        delegate: InputOutputDelegate {
                            copyOn: true
                            strAddr:  address
                            Component.onCompleted : {
                                coinFtr = Qt.binding(function(){return coinOptions})
                                implicitHeight = Qt.binding(function(){return !(emptyAddressVisible && isEmpty ) ? 90 : 0 })
                                }
                            width: parent.width
                            clip: true
                        }
                    }
                } // ScrollView
            } // ColumnLayout (inputs)
        } // RowLayout (details)
    } // ColumnLayout (root)

 DialogAddAddresses {
        id: dialogAddAddresses
        anchors.centerIn: Overlay.overlay
        modal: true
        focus: true
        onAccepted: {
            if (isEncrypted){
                dialogGetPasswordForAddAddresses.title = "Enter Password"
                dialogGetPasswordForAddAddresses.height = dialogGetPassword.height - 20
                dialogGetPasswordForAddAddresses.nAddress = spinValue
                dialogGetPasswordForAddAddresses.open()

            } else{
                walletManager.newWalletAddress(fileName, spinValue, "")
                walletManager.loadAddressModelByWallet(fileName, addressList)

            }


        }
        onRejected: {
            console.log("Adding rejected")
        }
    } // DialogAddAddresses
    DialogGetPassword {
            id: dialogGetPasswordForAddAddresses
            anchors.centerIn: Overlay.overlay
            property int nAddress
            width: applicationWindow.width > 540 ? 540 - 120 : applicationWindow.width - 40
            height: applicationWindow.height > 570 ? 570 - 180 : applicationWindow.height - 40

            focus: true
            modal: true

            onAccepted: {
                walletManager.newWalletAddress(fileName, nAddress, dialogGetPasswordForAddAddresses.password)
                walletManager.loadAddressModelByWallet(fileName, addressList)
            }
        }

        DialogSetPassword {
            id: dialogSetPassword

            anchors.centerIn: Overlay.overlay
            width: applicationWindow.width > 540 ? 540 - 40 : applicationWindow.width - 40
            height: applicationWindow.height > implicitHeight + 40 ? implicitHeight : applicationWindow.height - 40

            headerMessage: qsTr("We suggest that you encrypt each one of your wallets with a password. If you forget your password, you can reset it with your seed. Make sure you have your seed saved somewhere safe before encrypting your wallet.")
            headerMessageColor: Material.color(Material.Red)
            focus: true
            modal: true

            onAccepted: {
                let encrypt = walletManager.encryptWallet(fileName, password)
                isEncrypted = encrypt
            }
            onRejected: {
                encryptSwitch.checked = isEncrypted
            }
        } // DialogSetPassword
        DialogGetPassword {
            id: dialogGetPassword

            anchors.centerIn: Overlay.overlay
            width: applicationWindow.width > 440 ? 440 - 40 : applicationWindow.width - 40
            height: applicationWindow.height > 320 ? 320 - 40 : applicationWindow.height - 40

            headerMessage: qsTr("<b>Warning:</b> for security reasons, it is not recommended to keep the wallets unencrypted. Caution is advised.")
            headerMessageColor: Material.color(Material.Red)
            focus: true
            modal: true
            onAccepted: {
                let encrypt = walletManager.decryptWallet(fileName, password)
                isEncrypted = encrypt
            }
            onRejected: {
             encryptSwitch.checked = isEncrypted
            }
        }
        DialogEditWallet {
            id: dialogEditWallet
            anchors.centerIn: Overlay.overlay

            focus: true
            modal: true

            onAccepted: {
                 walletManager.editWalletLbl(fileName, name)
            }
        } // DialogEditWallet
}
