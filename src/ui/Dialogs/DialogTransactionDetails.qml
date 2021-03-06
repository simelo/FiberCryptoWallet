import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
import "qrc:/ui/src/ui/"

Dialog {
    id: root

    property alias date: transactionDetails.date
    property alias status: transactionDetails.status
    property alias type: transactionDetails.type
    property alias amount: transactionDetails.amount
    property alias hoursReceived: transactionDetails.hoursReceived
    property alias hoursBurned: transactionDetails.hoursBurned
    property alias transactionID: transactionDetails.transactionID

    title: qsTr("Transaction details")

    ColumnLayout {
        anchors.fill: parent
        spacing: 20

        TransactionDetails {
            id: transactionDetails
            implicitWidth: 500
            Layout.fillWidth: true
        }

        Rectangle {
            visible: transactionDetails.expanded
            height: 1
            color: Material.color(Material.Grey)
            Layout.fillWidth: true
        }
    }
}
