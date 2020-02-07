import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

ColumnLayout{
    property bool hide : true
    RowLayout{
        Label {
            Layout.leftMargin: 5
            id: label
            text: section
            font.bold: true

            Layout.fillWidth: true
        }

        ToolButton{
              text:""
            icon.source: "qrc:/images/resources/images/icons/up.svg"
            icon.color: Material.accent
            Behavior on rotation { NumberAnimation { duration: 500 } }

            onClicked:{
                hide=!hide
                this.rotation = hide ? 0 : 180
                changeVisibilityByWltName(section)
//                console.log(listOutputs.itemAtIndex(1).outputID)
//                console.log(listOutputs.indexAt(1/)
            }
        }
    }
    RowLayout{
        Rectangle {
            id: rect
            Layout.fillWidth: true
            height: 1
            color: "#DDDDDD"
        }
    }
}
