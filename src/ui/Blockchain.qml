import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import BlockchainModels 1.0

Page {
    id: blockchain
    property int selectedIndex : 0

    property string numberOfBlocks: model.numberOfBlocks
    property date timestampLastBlock: model.timestampLastBlock
    property string hashLastBlock: model.hashLastBlock

    function updateModel(){
        for(let i = gridLayoutDetails.children.length; i > 0 ; i--) {
            gridLayoutDetails.children[i-1].destroy()
        }
        model.update(utils.getTicketOfCurrency(currencyList.model[selectedIndex]))
            model.coinOpts.getKeys().forEach(key=>{
                Qt.createQmlObject(`import QtQuick 2.12;
                    import QtQuick.Controls 2.12;
                    import QtQuick.Controls.Material 2.12
                    import QtQuick.Layouts 1.12
                    ColumnLayout {
                        Label {
                            text: qsTr("${key}")
                            font.bold: true
                        }
                        Label {
                            text: "${model.coinOpts.getValue(key)}"
                        }
                    }`,gridLayoutDetails)
            })
    }

    Drawer {
        id: drawer
        interactive: true
        width: parent.width > 300 ? 200 : parent.width* 0.5
        edge: Qt.LeftEdge
        visible:true
        position: 1
        topMargin: parent.height - blockchain.height
        bottomPadding: 20
        height: blockchain.height

        ColumnLayout{
            anchors.fill: parent
            clip:true
            RowLayout{
                Layout.alignment: Qt.AlignTop
                Layout.fillWidth: true
//                Layout.bottomMargin: -3
//                Layout.topMargin: -5
//                Layout.leftMargin: -3
//                Layout.rightMargin: -3
                Rectangle{
//                    border { width: 3; color: applicationWindow.Material.background }
                    Layout.fillWidth: true
                    height: lblTitle.height * 1.2
                    color: applicationWindow.accentColor


                    Label{
                        color: applicationWindow.Material.background
                        topPadding: 20
                        id: lblTitle
                        leftPadding: 10
                        Layout.fillWidth: true
                        font.pointSize: Qt.application.font.pointSize * 1.1
                        text: qsTr("Supported currencies:")
                        font.bold: true
//                        wrapMode: Label.WrapAnywhere
                    }
                }
            }

            Rectangle{
                height: 1
                color: Material.color(Material.Grey)
                Layout.alignment: Qt.AlignVCenter
                Layout.topMargin: -6
                Layout.fillWidth: true
            }

            ScrollView{
                Layout.fillWidth: true
                Layout.fillHeight: true
                spacing:10
                ListView{
                    id: currencyList
                    clip: true
                    model:  utils.getSupportedCurrencies()
                    Component.onCompleted:{
                        currencyList.itemAtIndex(selectedIndex).selected = true
                    }

                    delegate:ItemDelegate{
                        property alias selected: imgSelected.visible
                        width: parent.width
                        RowLayout{
                            anchors.fill: parent
                            anchors.leftMargin: 10
                            anchors.rightMargin: 15

                            Image {
                                height: lblCurrency.height
                                width: height
                                Layout.alignment: Qt.AlignLeft
                                visible: false
                                id: imgSelected
                                source:"qrc:/images/resources/images/icons/check-simple.svg"
                                sourceSize:Qt.size(parent.height * 0.7 , parent.height * 0.7 )
                            }


                            Label{
                                id: lblCurrency
                                text: modelData
                            }
                            Image {
                                height: lblCurrency.height
                                width: height
                                Layout.alignment: Qt.AlignRight
                                source:"qrc:/images/resources/images/icons/"+modelData.toLowerCase()+"-Icon.svg"
//                                source:"qrc:/images/resources/images/icons/"+"skycoin"+"-Icon.svg"
                                sourceSize:Qt.size(parent.height * 0.7 , parent.height * 0.7 )
                            }
                            }
                            onClicked:{
                                if(selectedIndex != index){
                                    currencyList.itemAtIndex(selectedIndex).selected = false
                                    selectedIndex = index
                                    selected=true
                                }
                            }
                        }
                    }
                }
            }
        }
    BlockchainStatusModel{
        id: model
        Component.onCompleted:{
            updateModel()
        }
    }

    ColumnLayout {
        id: columnLayoutRoot
        anchors.top: parent.top
        anchors.left: parent.left
        anchors.right: parent.right
        anchors.margins: 20
        spacing: 20

        GroupBox {
            id: groupBoxBlockDetails
            label:RowLayout{
                Label{
                    text: currencyList.model[selectedIndex]
                    font.pointSize: Qt.application.font.pointSize * 1.2
                    color: applicationWindow.Material.accent
                    font.underline: true
                    MouseArea{
                        id:mouseArea
                        hoverEnabled : true
                        anchors.fill : parent
                        cursorShape: Qt.PointingHandCursor
                        preventStealing: true

                        onEntered: {
                            parent.font.bold=true
                            parent.color = Qt.darker(applicationWindow.Material.accent)
                        }

                        onExited: {
                            parent.font.bold=false
                            parent.color = applicationWindow.Material.accent
                        }

                        onClicked:{
                            drawer.visible = true
                        }
                    }
                }
                Label{
                    text: qsTr(" block details")
                }
            }
            clip: true
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop | Qt.AlignHCenter

            ColumnLayout {
                id: columnLayoutBlockDetails
                anchors.fill: parent
                spacing: 20

                ColumnLayout {
                    Label {
                        text: qsTr("Number of blocks")
                        font.bold: true
                    }
                    Label {
                        id: labelNumberOfBlocks
                        text: numberOfBlocks
                    }
                }

                RowLayout {
                    spacing: 20

                    ColumnLayout {
                        Label {
                            text: qsTr("Timestamp of last block")
                            font.bold: true
                        }
                        Label {
                            id: labelTimestampLastBlock
                            text: Qt.formatDateTime(timestampLastBlock, Qt.DefaultLocaleShortDate)
                        }
                    }

                    ColumnLayout {
                        Layout.fillWidth: true
                        Label {
                            text: qsTr("Hash of last block")
                            font.bold: true

                            ToolButton {
                                id: toolButtonCopy
                                anchors.left: parent.right
                                anchors.verticalCenter: parent.verticalCenter
                                icon.source: "qrc:/images/resources/images/icons/copy.svg"
                                visible: textInputHashLastBlock.text
                                ToolTip.text: qsTr("Copy to clipboard")
                                ToolTip.visible: hovered // TODO: pressed when mobile?
                                ToolTip.delay: Qt.styleHints.mousePressAndHoldInterval

                                Image {
                                    id: imageCopied
                                    anchors.centerIn: parent
                                    source: "qrc:/images/resources/images/icons/check-simple.svg"
                                    fillMode: Image.PreserveAspectFit
                                    sourceSize: Qt.size(toolButtonCopy.icon.width*1.5, toolButtonCopy.icon.height*1.5)
                                    z: 1

                                    opacity: 0.0
                                }

                                onClicked: {
                                    if (textInputHashLastBlock.text) {
                                        textInputHashLastBlock.selectAll()
                                        textInputHashLastBlock.copy()
                                        textInputHashLastBlock.deselect()
                                        if (copyAnimation.running) {
                                            copyAnimation.restart()
                                        } else {
                                            copyAnimation.start()
                                        }
                                    }
                                }

                                SequentialAnimation {
                                    id: copyAnimation
                                    NumberAnimation { target: imageCopied; property: "opacity"; to: 1.0; easing.type: Easing.OutCubic }
                                    PauseAnimation { duration: 1000 }
                                    NumberAnimation { target: imageCopied; property: "opacity"; to: 0.0; easing.type: Easing.OutCubic }
                                }
                            } // ToolButton
                        } // Label
                        TextInput {
                            id: textInputHashLastBlock
                            Layout.fillWidth: true
                            text: hashLastBlock
                            readOnly: true
                            font.family: "Code New Roman"
                            wrapMode: Label.WrapAnywhere
                        }
                    } // ColumnLayout
                } // RowLayout
            } // ColumnLayout (block details)
        } // GroupBox (block details)

        GroupBox {
            id: groupBoxSkyDetails
            title: qsTr(utils.getTicketOfCurrency(currencyList.model[selectedIndex])+" details")
            clip: true
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop | Qt.AlignHCenter

            GridLayout {
                id: gridLayoutDetails
                rows: 2
                columns: 2
            } // GridLayout
        } // GroupBox (sky details)
    } // ColumnLayout (root)

    BusyIndicator {
        id: busyIndicator

        anchors.centerIn: parent
        // Create a `busy` property in the backend and bind it to `running` here:
        running: model.loading
		property Timer timer: Timer{
			id: blockchainTimer
			repeat: true
			running: true
			interval: 3000
			onTriggered: {
                updateModel()
			}

		} 
    }
}
