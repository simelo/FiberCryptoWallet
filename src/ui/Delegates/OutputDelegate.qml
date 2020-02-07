import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import OutputsModels 1.0


Item {
    id: outputDelegate
    property bool expanded: false
    property bool animateDisplacement: false

    implicitHeight: delegate.height + (expanded ? gridLayoutBasicInfo.height + 10 : 0)
    Behavior on implicitHeight { NumberAnimation { duration: expanded ? 500 : 250 ; easing.type: Easing.OutQuint; onRunningChanged: animateDisplacement = false } }
    Layout.fillWidth: true

    Component.onCompleted:{
        coinOpts.getKeys().forEach(optKey=>{
            Qt.createQmlObject(`import QtQuick 2.12;
                                        import QtQuick.Controls 2.12;
                                        import QtQuick.Controls.Material 2.12
                                        import QtQuick.Layouts 1.12

                                        Label {
                                            text: qsTr("${optKey}")
                                            font.pointSize: Qt.application.font.pointSize * 0.9
                                            font.bold: true
                                        }`, gridLayoutBasicInfo)

            Qt.createQmlObject(`import QtQuick 2.12;
                                        import QtQuick.Controls 2.12;
                                        import QtQuick.Controls.Material 2.12
                                        import QtQuick.Layouts 1.12

                                        Label {
                                            text: qsTr("${coinOpts.getValue(optKey)}")
                                            font.pointSize: Qt.application.font.pointSize * 0.9
                                            font.bold: true
                                        }`, gridLayoutBasicInfo)

        })

    }

    ItemDelegate{
        id: delegate
        width: parent.width
        RowLayout{
            anchors.fill: parent
            anchors.leftMargin: 10
            anchors.rightMargin: 10
            spacing: 20
            Layout.fillWidth: true
            Label{
                Material.foreground: Material.Grey
                font.family: "Code New Roman"
                font.pointSize: Qt.application.font.pointSize * 0.9
                wrapMode: Label.WrapAnywhere
                Layout.fillWidth: true
                Layout.minimumWidth: implicitWidth/2 // As the font is monospaced, this should work fine
                text: outputID
            }
            Label{
                text: addressOwner
                Layout.fillWidth: true
                Layout.preferredWidth: 150
                Layout.alignment: Qt.AlignRight
                wrapMode: Label.WrapAnywhere
            }
        }//delegateRowLayout
        onClicked:{
        expanded= !expanded

        }
    }//delegate

    GridLayout {
        id: gridLayoutBasicInfo
        Material.foreground: Material.Grey
        columns: 4
        visible: expanded
        columnSpacing: 10
        anchors.top: delegate.bottom
        Layout.fillWidth: true

    } // GridLayout
}
