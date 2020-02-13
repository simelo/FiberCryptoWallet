import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

ScrollView {
    id: historyFilterDelegate

    ListView {
        id: listViewFilters
        
        width: parent.width
        spacing: 10
        
        model: modelFilters
        delegate: HistoryFilterListDelegate {
            width: parent.width
        }

    }
}