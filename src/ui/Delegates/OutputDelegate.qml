import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import ModelUtils 1.0

Item {
    id: outputDelegate
    property bool expanded: false
    property bool animateDisplacement: false
    property string outId: ""
    property string address: ""
    property Map traits
    property string walletLbl: ""
    signal showOrHide(bool isActive)
    Behavior on implicitHeight { NumberAnimation { duration: expanded ? 500 : 250 ; easing.type: Easing.OutQuint; onRunningChanged: animateDisplacement = false } }
    Layout.fillWidth: true
    onShowOrHide:{
        implicitHeight = isActive ? Qt.binding(function(){return delegate.height + (expanded ? gridLayoutBasicInfo.height + 10 : 0)}) : 0;
        visible=isActive
    }

    Component.onCompleted:{
        traits.getKeys().forEach(optKey=>{
            Qt.createQmlObject(`import QtQuick 2.12;
                                        import QtQuick.Controls 2.12
                                        import QtQuick.Controls.Material 2.12
                                        import QtQuick.Layouts 1.12;

                                        Label {
                                            text: qsTr("${optKey}: ${traits.getValue(optKey)}")
                                            Layout.leftMargin: 10
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
                text: outId
            }
            Label{
                text: address
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
        columns: 2
        visible: expanded
        columnSpacing: 30
        anchors.top: delegate.bottom
        Layout.fillWidth: true

    } // GridLayout
}
