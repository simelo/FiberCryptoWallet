import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import Transaction 1.0
import ModelUtils 1.0
import Address 1.0

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

// Backend imports

Item {
    id: root

    property date date: "2000-01-01 00:00"
    property int type: TransactionDetails.Type.Send
    property int status: TransactionDetails.Status.Preview
    property var statusString: [ qsTr("Confirmed"), qsTr("Pending"), qsTr("Preview") ]
    property string transactionID
    property string blockHeight
    property string amount
    property QAddressList modelInputs
    property QAddressList modelOutputs
    property Map modelCoinOpts
    property alias specificDetails : gridLayoutSpecificInfo
    readonly property bool expanded: buttonMoreDetails.checked

    enum Status {
        Confirmed,
        Pending,
        Preview
    }

    enum Type {
        Send,
        Receive,
        Internal,
        Generic
    }

    readonly property real basicHeight: 80 + rowLayoutBasicDetails.height + rowLayoutSpecificDetails.height
    implicitHeight: Math.min(basicHeight + (expanded ? Math.max(listViewInputs.height, listViewOutputs.height) : 0), maxHeight)
    Behavior on implicitHeight { NumberAnimation { duration: 1000; easing.type: Easing.OutQuint } }

    implicitWidth: 600
    clip: true

    function getIconPath(){
        if (type == TransactionDetails.Type.Receive || type == TransactionDetails.Type.Send){
        return "qrc:/images/resources/images/icons/send-" + (type == TransactionDetails.Type.Send ? "blue" : "amber") + ".svg"
        }
        if( !coinOpts){
        return "qrc:/images/resources/images/icons/error.svg"
        }
        return "qrc:/images/resources/images/icons/"+coinOpts.getKeys()[0].split(' ')[0].toLowerCase()+"-Icon.svg"
    }

    function setMainBalance(){
        let text = "";
        if (type === TransactionDetails.Type.Recive || type === TransactionDetails.Type.Send){
        text+= type === TransactionDetails.Type.Recive ? "Recive " : "Send "
        }

        if(!coinOpts){
             return text
        }

        return text + amount + " " + coinOpts.getKeys()[0]
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

                Label {
                    text: qsTr("Basic Transaction Details")
                    font.bold: true
                    Layout.alignment: Qt.AlignTop
                    Layout.fillWidth: true
                }

                GridLayout {
                    id: gridLayoutBasicInfo
                    Material.foreground: Material.Grey
                    columns: 2
                    columnSpacing: 10

                    Layout.alignment: Qt.AlignTop
                    Layout.fillWidth: true

                    Label {
                        text: qsTr("Date:")
                        font.pointSize: Qt.application.font.pointSize * 0.9
                        font.bold: true
                    }
                    Label {
                        text: Qt.formatDateTime(root.date, Qt.DefaultLocaleShortDate)
                        font.pointSize: Qt.application.font.pointSize * 0.9
                    }

                    Label {
                        text: qsTr("Height:")
                        font.pointSize: Qt.application.font.pointSize * 0.9
                        font.bold: true
                    }
                    Label {
                        text: blockHeight
                        font.pointSize: Qt.application.font.pointSize * 0.9
                    }

                    Label {
                        text: qsTr("Status:")
                        font.pointSize: Qt.application.font.pointSize * 0.9
                        font.bold: true
                    }
                    Label {
                        text: statusString[root.status]
                        font.pointSize: Qt.application.font.pointSize * 0.9
                    }

                    Label {
                        text: qsTr("Tx ID:")
                        font.pointSize: Qt.application.font.pointSize * 0.9
                        font.bold: true
                    }
                    Label {
                        text: root.transactionID
                        font.pointSize: Qt.application.font.pointSize * 0.9
                        Layout.fillWidth: true
                    }
                } // GridLayout
            }

            ColumnLayout {
            id: rightUpPanel
                Layout.alignment: Qt.AlignTop
                Layout.topMargin: 0
                Layout.rightMargin: 20
                Image {
                    source: getIconPath()
                    sourceSize: "64x64"
                    fillMode: Image.PreserveAspectFit
                    mirror: type === TransactionDetails.Type.Receive
                    Layout.fillWidth: true
                }

                Label {
                    id : labelMainBalance
                    text: setMainBalance()
                    font.bold: true
                    font.pointSize: Qt.application.font.pointSize * 1.15
                    horizontalAlignment: Label.AlignHCenter
                    Layout.fillWidth: true
                }
            }
        } // RowLayout

        RowLayout{
            id: rowLayoutSpecificDetails
            Layout.alignment: Qt.AlignTop
            Layout.fillWidth: true
            Layout.fillHeight: true
            ColumnLayout {
                id:columnLayoutSpecificDetails
                Layout.alignment: Qt.AlignTop
                Layout.fillWidth: true
                Label {
                    text: qsTr("Specific Transaction Details")
                    font.bold: true
                    Layout.alignment: Qt.AlignTop
                    Layout.fillWidth: true
                }
                GridLayout {
                    id: gridLayoutSpecificInfo
                    Material.foreground: Material.Grey
                    columns: 2
                    columnSpacing: 10
                    Layout.alignment: Qt.AlignTop
                    Layout.fillWidth: true
                }
            }
        }

        RowLayout {
            id: rowLayoutDetailsSeparator

            Layout.alignment: Qt.AlignTop
            Layout.fillWidth: true

            Button {
                id: buttonMoreDetails
                implicitWidth: 200
                flat: true
                checkable: true
                text: checked ? qsTr("Less details") : qsTr("More details")
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

                Label {
                    text: qsTr("Inputs")
                    font.bold: true
                    font.italic: true
                    Layout.fillWidth: true
                }

                ScrollView {
                    Layout.alignment: Qt.AlignTop
                    Layout.fillWidth: true
                    Layout.fillHeight: true

                    ListView {
                        id: listViewInputs

                        Material.foreground: Material.Grey
                        model: modelInputs
                        clip: true
                        delegate: InputOutputDelegate {
                            Component.onCompleted : {
                                coinFtr = Qt.binding(function(){return coinOptions})
                            }
                            width: parent.width
                        }
                    }
                } // ScrollView
            } // ColumnLayout (inputs)

            ColumnLayout {
                id: columnLayoutOutputs
                Layout.fillWidth: true
                Layout.fillHeight: true

                Label {
                    text: qsTr("Outputs")
                    font.bold: true
                    font.italic: true
                    Layout.fillWidth: true
                }

                ScrollView {
                    Layout.alignment: Qt.AlignTop
                    Layout.fillWidth: true
                    Layout.fillHeight: true

                    ListView {
                        id: listViewOutputs

                        Material.foreground: Material.Grey
                        model: modelOutputs
                        clip: true
                        delegate: InputOutputDelegate {
                            Component.onCompleted : {
                                coinFtr = Qt.binding(function(){return coinOptions})
                            }
                            width: parent.width
                        }
                    }
                } // ScrollView
            } // ColumnLayout (outputs)
        } // RowLayout (details)
    } // ColumnLayout (root)

    // Roles: address, addressSky, addressCoinHours
    // Use listModel.append( { "address": value, "addressSky": value, "addressCoinHours": value } )
    // Or implement the model in the backend (a more recommendable approach)
    // ListModel {
    //     id: listModelInputs
    //     ListElement { address: "qrxw7364w8xerusftaxkw87ues"; addressSky: 30; addressCoinHours: 1049 }
    //     ListElement { address: "8745yuetsrk8tcsku4ryj48ije"; addressSky: 12; addressCoinHours: 16011 }
    // }
    // ListModel {
    //     id: listModelOutputs
    //     ListElement { address: "734irweaweygtawieta783cwyc"; addressSky: 38; addressCoinHours: 5048 }
    //     ListElement { address: "ekq03i3qerwhjqoqh9823yurig"; addressSky: 61; addressCoinHours: 9456 }
    //     ListElement { address: "1kjher73yiner7wn32nkuwe94v"; addressSky: 1; addressCoinHours: 24 }
    //     ListElement { address: "oopfwwklfd34iuhjwe83w3h28r"; addressSky: 111; addressCoinHours: 13548 }
    // }
    // QAddressList{
    //     id: listModelInputs
    // }
    // QAddressList{
    //     id: listModelOutputs
    // }
}
