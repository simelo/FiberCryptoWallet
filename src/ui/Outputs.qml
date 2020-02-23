import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import Outputs 1.0

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

Page {
    id: outputs
    readonly property real listOutputsLeftMargin: 20
    readonly property real listOutputsRightMargin: 50
    readonly property real listOutputsSpacing: 20
    readonly property real internalLabelsWidth: 60
    property alias model: listOutputs
    property var walletList: []

    function loadWallets(pos){
        var wlt = walletList.find(wlt => { return wlt.name === pos.walletLbl })
        if (!wlt){
            wlt = { name: pos.walletLbl , active: false };
            walletList.push(wlt)
        }
        pos.showOrHide(wlt.active)
        pos.visible = wlt.active
    }

    function changeVisibilityByWltName( wltName ){
        let index = walletList.findIndex(wlt => wlt.name===wltName );
        walletList[index].active = !walletList[index].active;

        for( let i=0; i< model.count ; i++ ){
            if ( model.itemAtIndex(i).walletLbl === wltName ){
                  model.itemAtIndex(i).showOrHide(walletList[index].active)
            }
        }
    }

    QOutputs{
        id: outputModel
        Component.onCompleted:{
            walletManager.loadOutputsAsync(this)
        }
    }

    GroupBox{
        anchors.fill: parent
        anchors.margins: 20
        clip: true
        ColumnLayout{
            id: columnLayoutFrame
            anchors.fill: parent
            ColumnLayout {
                id: columnLayoutHeader

                RowLayout {
                    Layout.topMargin: 30

                    Label {
                        text: qsTr("Output ID")
                        font.pointSize: 9
                        Layout.leftMargin: 10 // Empirically obtained
                        Layout.fillWidth: true
                    }
                    Label {
                        text: qsTr("Address")
                        font.pointSize: 9
                        Layout.rightMargin: 98
                    }
                } // RowLayout

                Rectangle {
                    Layout.fillWidth: true
                    height: 1
                    color: "#DDDDDD"
                }
            } // ColumnLayout (header)

            ListView{
                id: listOutputs
                Layout.fillWidth: true
                Layout.fillHeight: true
                clip: true
                model: outputModel

                section.property: "walletOwner"
                section.criteria: ViewSection.FullString
                section.delegate: OutputSectionDelegate{
                    width: parent.width
                }
                delegate: OutputDelegate{
                    width : parent.width
                    outId : outputID
                    address : addressOwner
                    traits : coinOpts
                    walletLbl : walletOwner

                    Component.onCompleted:{
                        loadWallets(this) //TODO
                    }
                }
            }//ListView


        }//columnLayoutFrame
    }//GroupBox


    BusyIndicator {
        id: busyIndicator
        anchors.centerIn: parent
//         Create a `busy` property in the backend and bind it to `running` here:
        running: outputModel.loading
    }
}
