import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import ModelUtils 1.0

ItemDelegate {
    id: root

    property string modelTransactionID
    property date modelTimestamp
    property Map coinOpts

    RowLayout {
        id: rowLayoutRoot
        anchors.fill: parent
        anchors.leftMargin: 10
        anchors.rightMargin: 10

        spacing: 20

        Label {
            Material.foreground: Material.Grey
            text: modelTransactionID // model's roles
            font.family: "Code New Roman"
            font.pointSize: Qt.application.font.pointSize * 0.9
            wrapMode: Label.WrapAnywhere
            Layout.fillWidth: true
            Layout.minimumWidth: implicitWidth/2 // As the font is monospaced, this should work fine
        }

        Label {
            text: modelTimestamp // model's role
            Layout.preferredWidth: 150
        }

    } // RowLayout (root)
}
