import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12

Drawer{
id: root
    property alias walletName: walletDetails.name
    property alias walletCoin: walletDetails.coin
    property alias walletCoinOpts: walletDetails.coinOpts
    property alias walletAddressList: walletDetails.addressList
    property alias expanded: walletDetails.expanded
    property alias walletStatus: walletDetails.isEncrypted
    property alias walletFileName: walletDetails.fileName
    property alias walletIndex: walletDetails.index
    property var bind
    signal changedContent()

    onChangedContent:{
        if (walletName){
        generateDetails()
        }
    }

    onAboutToShow:{
        if (walletName){
            generateDetails()
        }
    }

    WalletDetails {
        id: walletDetails
        anchors.fill: parent
        clip: true
    }

    function generateDetails() {
        // Create a list of child of walletDetails.specificDetails object.
	    let childList = Array.from(walletDetails.specificDetails.children)
	    // Remove all existents childs of walletDetails.specificDetails object.
	    childList.forEach(item=>{item.destroy()})

        // Wallet Label. Can be edited.
		Qt.createQmlObject(`import QtQuick 2.12;
                                import QtQuick.Controls 2.12;
                                import QtQuick.Controls.Material 2.12
                                import QtQuick.Layouts 1.12

                                Label {
                                    text: qsTr("Name:")
                                    font.pointSize: Qt.application.font.pointSize * 0.9
                                    font.bold: true
                                }`,walletDetails.specificDetails)

		Qt.createQmlObject(`import QtQuick 2.12;
                                import QtQuick.Controls 2.12;
                                import QtQuick.Controls.Material 2.12
                                import QtQuick.Layouts 1.12
                                import QtGraphicalEffects 1.0

                                RowLayout{
                                Label {
                                    id: labelName
                                    text: walletName
                                    font.pointSize: Qt.application.font.pointSize * 0.9
                                }
                                    Rectangle {
                                        property bool hovered: false
                                        color: hovered ? Material.accent : "transparent"
                                        height: labelName.height
                                        width: height
                                        radius: width / 2
                                        Image {
                                            id: imageEdit
                                            anchors.centerIn: parent
                                            source:"qrc:/images/resources/images/icons/edit.svg"
                                            sourceSize:Qt.size(parent.width * 0.9, parent.height * 0.9)
                                        }
                                        ColorOverlay {
                                            anchors.fill: imageEdit
                                            source: imageEdit
                                            color: Material.background
                                            visible: parent.hovered
                                        }
                                        MouseArea{
                                            id:mouseArea
                                            hoverEnabled : true
                                            anchors.fill : parent
                                            cursorShape: Qt.PointingHandCursor
                                            preventStealing: true
                                            onEntered: parent.hovered = true
                                            onExited: parent.hovered = false
                                        onClicked:{
                                            walletDetails.editWallet.originalWalletName = walletName
                                            walletDetails.editWallet.name = walletName
                                            walletDetails.editWallet.open()
                                        }
                                        }
                                    }
                                }`,walletDetails.specificDetails)
        // Wallet currency.
		Qt.createQmlObject(`import QtQuick 2.12;
                                import QtQuick.Controls 2.12;
                                import QtQuick.Controls.Material 2.12
                                import QtQuick.Layouts 1.12
                                Label {
                                    text: qsTr("Coin:")
                                    font.pointSize: Qt.application.font.pointSize * 0.9
                                    font.bold: true
                                }`,walletDetails.specificDetails)

		Qt.createQmlObject(`import QtQuick 2.12;
                                       import QtQuick.Controls 2.12;
                                       import QtQuick.Controls.Material 2.12
                                       import QtQuick.Layouts 1.12
                                       import QtGraphicalEffects 1.0

                                       RowLayout{
                                       Label {
                                           id: labelCoin
                                           text: walletCoin
                                           font.pointSize: Qt.application.font.pointSize * 0.9
                                       }
                                       Rectangle {
                                           height:labelCoin.height
                                           width:height
                                           radius: width / 2
                                           Image {
                                               id: imageCoin
                                               anchors.centerIn: parent
                                               source:"qrc:/images/resources/images/icons/"+ walletCoin +"-Icon.svg\"
                                               sourceSize:Qt.size(parent.width, parent.height)
                                           }
                                       }
                                    }`,walletDetails.specificDetails)

        // Iterate for each key in the key list of coinOptions and generate the details of the currency.
        walletCoinOpts.getKeys().forEach(key=>{
            Qt.createQmlObject(`import QtQuick 2.12;
                                    import QtQuick.Controls 2.12;
                                    import QtQuick.Controls.Material 2.12
                                    import QtQuick.Layouts 1.12
                                    Label {
                                        text: qsTr("${key}:")
                                        font.pointSize: Qt.application.font.pointSize * 0.9
                                        font.bold: true
                                    }`, walletDetails.specificDetails)

            Qt.createQmlObject(`import QtQuick 2.12;
                                    import QtQuick.Controls 2.12;
                                    import QtQuick.Controls.Material 2.12
                                    import QtQuick.Layouts 1.12
                                    Label {
                                        text: " ${walletCoinOpts.getValue( key )}"
                                        font.pointSize: Qt.application.font.pointSize * 0.9
                                        Layout.fillWidth: true
                                    }`,walletDetails.specificDetails)
        })
    }
}
