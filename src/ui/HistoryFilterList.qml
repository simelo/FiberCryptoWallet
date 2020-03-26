import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import Address 1.0

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

ScrollView {
    id: historyFilterDelegate
    signal loadWallets()
    clip: true

    ListView {
        id: listViewFilters
        Component.onCompleted:{
            console.log(this.cacheBuffer)
        }
        signal allChecked(int isChecked, string wltId)
        onAllChecked:{
            for (let i=0;i< this.count; i++){
                if (modelFilters.addresses[i].walletId === wltId){
                    this.itemAtIndex(i).checkState = isChecked
                }
            }
        }
        width: parent.width
        spacing: 10

        model: modelFilters
        section.property: "walletId"
        section.criteria: ViewSection.FullString
        section.delegate: CheckDelegate{
            id: sect
            LayoutMirroring.enabled: true
            text: section
            width: parent.width
            contentItem: Label {
                leftPadding: sect.indicator.width + sect.spacing
                verticalAlignment: Qt.AlignVCenter
                font.bold: true
                font.pointSize: 12
                text: sect.text
                color: sect.enabled ? sect.Material.foreground : sect.Material.hintTextColor
            }
            nextCheckState: function() {
                if (checkState === Qt.Checked) {
                    listViewFilters.allChecked(Qt.Unchecked, section)
                    return Qt.Unchecked
                } else {
                    listViewFilters.allChecked(Qt.Checked, section)
                    return Qt.Checked
                }
            }
        }
        delegate: CheckDelegate {
            id:deleg
            leftPadding: 30
            width: parent.width
            LayoutMirroring.enabled: true
            text: address
            contentItem: Label {
                leftPadding: deleg.indicator.width + deleg.spacing
                verticalAlignment: Qt.AlignVCenter
                text: deleg.text
                color: deleg.enabled ? deleg.Material.foreground : deleg.Material.hintTextColor
            }
            onCheckedChanged: {
                if (checked){
                    historyManager.addFilter(address)
                }else{
                    historyManager.removeFilter(address)
                }
            }
        }

        QAddressList{
            id: modelFilters
            Component.onCompleted:{
                this.loadModel( walletManager.loadAddressForAllWallets())
                listViewFilters.cacheBuffer = this.addresses.length * 45
            }
        }
    }
}