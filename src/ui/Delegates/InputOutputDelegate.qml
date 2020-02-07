import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import ModelUtils 1.0

Item {
    id: root
    implicitHeight: 90
    property bool copyOn: false
    Component.onCompleted:{
            let keyList=coinOptions.getKeys()
            for (let i=0;i<keyList.length;i++){

                Qt.createQmlObject("import QtQuick 2.12;
                                    import QtQuick.Controls 2.12;
                                    import QtQuick.Controls.Material 2.12;
                                    import QtQuick.Layouts 1.12
                                        Label {
                                            text: qsTr(\""+keyList[i]+":\")
                                            font.pointSize: Qt.application.font.pointSize * 0.9
                                            font.bold: true
                                        }
                                    ",gridLayoutInputAndOutputInfo)

                Qt.createQmlObject("import QtQuick 2.12;
                                    import QtQuick.Controls 2.12;
                                    import QtQuick.Controls.Material 2.12;
                                    import QtQuick.Layouts 1.12
                                        Label {
                                            text: \""+coinOptions.getValue(keyList[i])+"\"
                                            font.pointSize: Qt.application.font.pointSize * 0.9
                                        }
                                    ",gridLayoutInputAndOutputInfo)
            }
  }

    ColumnLayout {
        id: columnLayoutRoot
        anchors.fill: parent

        RowLayout {
            id: rowLayoutHeader
            spacing: 10
            Layout.alignment: Qt.AlignTop
            Layout.fillWidth: true

            Label {
                id: labelIndex
                text: index + 1
                font.pointSize: Qt.application.font.pointSize * 0.9
                font.bold: true
            }
            Label {
                id: labelAddress
                text: address
                font.family: "Code New Roman"
                font.pointSize: Qt.application.font.pointSize * 0.9
                font.bold: true
                Layout.fillWidth: true
            }

            ToolButton {
                id: toolButtonCopy
                icon.source: "qrc:/images/resources/images/icons/copy.svg"
                visible: labelAddress.text && copyOn
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
                    if (labelAddress.text) {
                        auxTextInp.text = labelAddress.text
                        auxTextInp.selectAll()
                        auxTextInp.copy()
                        auxTextInp.deselect()
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

        }
        TextInput{
            id:auxTextInp
            visible: false
        }

        GridLayout {
            columns: 2
            Layout.alignment: Qt.AlignTop
            Layout.leftMargin: labelIndex.width + rowLayoutHeader.spacing
            Layout.fillWidth: true
            id :gridLayoutInputAndOutputInfo
        }
    } // ColumnLayout
}
