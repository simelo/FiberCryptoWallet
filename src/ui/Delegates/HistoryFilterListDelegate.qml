import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import WalletsManager 1.0

// Resource imports
// import "qrc:/ui/src/ui/"
import "../" // For quick UI development, switch back to resources when making a release

Item {
    id: root

    property alias tristate: checkDelegate.tristate
    property alias walletText: checkDelegate.text

    clip: true
    implicitWidth: width
    implicitHeight: height
    width: 300
    height: checkDelegate.height +  columnLayout.spacing + listViewFilterAddress.height
    
    ColumnLayout {
        id: columnLayout
        anchors.fill: parent

        CheckDelegate {
            id: checkDelegate

            Layout.fillWidth: true
            tristate: true
            text: name
            LayoutMirroring.enabled: true

            nextCheckState: function() {
                if (checkState === Qt.Checked) {
                    listViewFilterAddress.allChecked(Qt.Unchecked)
                    return Qt.Unchecked
                } else {
                    listViewFilterAddress.allChecked(Qt.Checked)
                    return Qt.Checked
                }
            }

            contentItem: Label {
                leftPadding: checkDelegate.indicator.width + checkDelegate.spacing
                verticalAlignment: Qt.AlignVCenter
                text: checkDelegate.text
                color: checkDelegate.enabled ? checkDelegate.Material.foreground : checkDelegate.Material.hintTextColor
            }
        } // CheckDelegate

        ListView {
            id: listViewFilterAddress
            property int checkedDelegates: 0
            signal allChecked(int isChecked)

            onAllChecked:{
                console.log(this.itemAtIndex(0))
               for (let i=0;i< this.count; i++){
               this.itemAtIndex(i).checkState = isChecked
               console.log( this.itemAtIndex(i).checkState )
               }
            }
            model: addresses

            Layout.fillWidth: true
            implicitHeight: contentHeight
            interactive: false

            delegate: CheckDelegate {
                clip:true
                // BUG: Checking the wallet does not change the check state of addresses
                // Is `checked: marked` ok? Or it should be the opposite?

                LayoutMirroring.enabled: true

                width: parent.width
                text: address

                onCheckedChanged: {
                    ListView.view.checkedDelegates += checked ? 1: -1

                    if (checked == true) {
                        historyManager.addFilter(address)
                    } else {
                        historyManager.removeFilter(address)
                    }
                }
            } // HistoryFilterListAddressDelegate (delegate)
        } // ListView
    } // ColumnLayout
}
