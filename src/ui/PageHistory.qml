import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import HistoryModels 1.0
import Transaction 1.0

// Resource imports
// import "qrc:/ui/src/ui/Dialogs"
// import "qrc:/ui/src/ui/Delegates"
import "Dialogs/" // For quick UI development, switch back to resources when making a release
import "Delegates/" // For quick UI development, switch back to resources when making a release

Page {
    id: root

    GroupBox {
        anchors.fill: parent
        anchors.margins: 20
        clip: true

        label: RowLayout {
            SwitchDelegate {
                id: switchFilters
                text: qsTr("Filters")
                onClicked:{
                    if (!checked) {
//                        historyManager.hasChanged(false)
			        }else {
//                        historyManager.hasChanged(true)
                    }
                }
            }

            Button {
                id: buttonFilters
                flat: true
                enabled: switchFilters.checked
                highlighted: true
                text: qsTr("Select filters")

                onClicked: {
                    toolTipFilters.open()
                }
            }
        } // RowLayout (GroupBox label)

        ScrollView {
            anchors.fill: parent
            clip: true
            ListView {
                id: listTransactions

                model: modelTransactions
                delegate: HistoryListDelegate {
                    onClicked: {
                        dialogTransactionDetails.open()
                        listTransactions.currentIndex = index
                    }
                }
            }
        }
    } // GroupBox


    Dialog {
        id: toolTipFilters

        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 440 ? 440 - 40 : applicationWindow.width - 40
        height: Math.min(applicationWindow.height - 40, filter.contentHeight + header.height + footer.height + topPadding + bottomPadding)

        modal: true
        standardButtons: Dialog.Close
        closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutside
        title: qsTr("Available filters")

        onClosed: {
//           historyManager.hasChanged(true)
        }


        onOpened:{
            filter.loadWallets()
        }

        HistoryFilterList {
            id: filter
            anchors.fill: parent
        }
    } // Dialog

    DialogTransactionDetails {
        id: dialogTransactionDetails

        readonly property real maxHeight: expanded ? 650 : 450

        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 640 ? 640 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > maxHeight ? maxHeight - 40 : applicationWindow.height - 40
        Behavior on height { NumberAnimation { duration: 1000; easing.type: Easing.OutQuint } }

        modal: true
        focus: true

        date : listTransactions.currentItem !== null ? listTransactions.currentItem.modelDate : ""
        status : listTransactions.currentItem !== null ? listTransactions.currentItem.modelStatus : 0
        type : listTransactions.currentItem !== null ? listTransactions.currentItem.modelType : 0
        modelAmount  : listTransactions.currentItem !== null ? listTransactions.currentItem.modelAmount : ""
        transactionID : listTransactions.currentItem !== null ? listTransactions.currentItem.modelTransactionID : ""
        modelInputs : listTransactions.currentItem !== null ? listTransactions.currentItem.modelInputs : null
        modelOutputs : listTransactions.currentItem !== null ? listTransactions.currentItem.modelOutputs : null
        blockHeight : listTransactions.currentItem !== null ? listTransactions.currentItem.modelBlockHeight : "0"
        coinOpts : listTransactions.currentItem !== null ? listTransactions.currentItem.coinOpts : null

    }

    QTransactionList {
        id: modelTransactions
        Component.onCompleted:{
//            historyManager.loadTransactionAsync(this)
        }
    }

    HistoryManager {
        id: historyManager
    }
}
