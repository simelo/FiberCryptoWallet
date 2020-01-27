import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import ModelUtils 1.0

Item {
    id: root
    implicitHeight: 90

  Component.onCompleted:{
            let keyList=coinOptions.getKeys()
            for (var i=0;i<keyList.length;i++){

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
