import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import QtGraphicalEffects 1.12
import WalletsManager 1.0
import HistoryModels 1.0
// Resource imports
// import "qrc:/ui/src/ui/Dialogs"
import "../Dialogs/" // For quick UI development, switch back to resources when making a release

Item {
    id: root

    readonly property real delegateHeight: 30
    property bool expanded: expand
    // The following property is used to avoid a binding conflict with the `height` property.
    // Also avoids a bug with the animation when collapsing a wallet
    readonly property real finalViewHeight: expanded ? delegateHeight*(addressList.count) + 50 : 0


    width: walletList.width
    height: itemDelegateMainButton.height + (expanded ? finalViewHeight : 0)

    Behavior on height { NumberAnimation { duration: 250; easing.type: Easing.OutQuint } }

    ColumnLayout {
        id: delegateColumnLayout
        anchors.fill: parent

        ItemDelegate {
            id: itemDelegateMainButton
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop
            font.bold: expanded

            RowLayout {
                id: delegateRowLayout
                anchors.fill: parent
                anchors.leftMargin: listWalletLeftMargin
                anchors.rightMargin: listWalletRightMargin
                spacing: listWalletSpacing

                Image {
                    id: status
                    source: statusIcon
                    sourceSize: "24x24"
                }

                Label {
                    id: labelWalletName
                    text: name // a role of the model
                    Layout.fillWidth: true
                }

                Item {
                    id: coinImg

                    width: coinIcon.width
                    height: coinIcon.height

                    Image {
                        id: coinIcon
                        source: "qrc:/images/resources/images/icons/"+currency+"-Icon.svg"
                        sourceSize: "24x24"
                    }
                }

                Item {
                    id: itemImageLockIcon

                    width: lockIcon.width
                    height: lockIcon.height

                    Image {
                        id: lockIcon
                        source: "qrc:/images/resources/images/icons/lock" + (encryptionEnabled ? "On" : "Off") + ".svg"
                        sourceSize: "24x24"
                    }

                    ColorOverlay {
                        anchors.fill: lockIcon
                        source: lockIcon
                        color: Material.theme === Material.Dark ? Material.foreground : "undefined"
                    }
                }

                Label {
                    id: labelBalance
                    text:  !coinOpts ? "N/A" : coinOpts.getValue(coinOpts.getKeys()[0])==""? "N/A" : coinOpts.getValue(coinOpts.getKeys()[0]) // a role of the model
                    color: Material.accent
                    horizontalAlignment: Text.AlignRight
                    Layout.preferredWidth: internalLabelsWidth
                    BusyIndicator {
                        anchors.verticalCenter: parent.verticalCenter
                        anchors.right: parent.right
//                        running: sky === qsTr("N/A") ? true : false

                        implicitWidth: implicitHeight
                        implicitHeight: parent.height + 10
                    }
                }
            } // RowLayout

            onClicked: {
                drawerWalletDetail.walletName = Qt.binding(function(){return name})
                drawerWalletDetail.walletCoin = Qt.binding(function(){return currency})
                drawerWalletDetail.walletCoinOpts = Qt.binding(function(){return coinOpts})
                drawerWalletDetail.walletAddressList = Qt.binding(function(){return addresses})
                drawerWalletDetail.walletStatus = Qt.binding(function(){return encryptionEnabled})
                drawerWalletDetail.walletFileName = Qt.binding(function(){return fileName})
                drawerWalletDetail.walletIndex = index
                drawerWalletDetail.visible = true
            }
        } // ItemDelegate
    } //ColumnLayout
    // Roles: address, addressSky, addressCoinHours
    // Use listModel.append( { "address": value, "addressSky": value, "addressCoinHours": value } )
    // Or implement the model in the backend (a more recommendable approach)
//    QAddressList {
//
//        id: listAddresses
//        property Timer timer: Timer {
//            id: addressModelTimer
//            interval: 7000
//            repeat: trueGetLoadedAddresses
//            running: true
//            onTriggered: {
//                walletManager.updateModel(fileName,listAddresses);
//            }
//        }
//    }
}
