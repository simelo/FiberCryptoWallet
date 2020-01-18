import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// Resource imports
// import "qrc:/ui/src/ui/"
import "../" // For quick UI development, switch back to resources when making a release

Dialog {
    id: dialogTransactionsDetails

    property alias date: transactionDetails.date
    property alias status: transactionDetails.status
    property alias type: transactionDetails.type
    property alias amount: transactionDetails.amount
    property alias hoursReceived: transactionDetails.hoursReceived
    property alias hoursBurned: transactionDetails.hoursBurned
    property alias transactionID: transactionDetails.transactionID
    property alias blockHeight: transactionDetails.blockHeight

    property alias coinOpts: transactionDetails.modelCoinOpts
    property alias modelInputs: transactionDetails.modelInputs
    property alias modelOutputs: transactionDetails.modelOutputs

    property alias expanded: transactionDetails.expanded

    title: qsTr("Transaction details")
    standardButtons: Dialog.Ok

    onAboutToShow:{
        for(var i = transactionDetails.specificDetails.children.length; i > 0 ; i--) {
            transactionDetails.specificDetails.children[i-1].destroy()
        }
        let keyList=coinOpts.getKeys()
            for (var i=0;i<keyList.length;i++){
                if (coinOpts.getValue(keyList[i])=="0"){
                    continue
                }

                Qt.createQmlObject("import QtQuick 2.12;
                                    import QtQuick.Controls 2.12;
                                    import QtQuick.Controls.Material 2.12
                                    import QtQuick.Layouts 1.12
                                    Label {
                                        text: qsTr(\""+keyList[i]+":\")
                                        font.pointSize: Qt.application.font.pointSize * 0.9
                                        font.bold: true
                                    }",transactionDetails.specificDetails)

                Qt.createQmlObject("import QtQuick 2.12;
                                    import QtQuick.Controls 2.12;
                                    import QtQuick.Controls.Material 2.12
                                    import QtQuick.Layouts 1.12
                                    Label {
                                        text: \""+coinOpts.getValue(keyList[i])+"\"
                                        font.pointSize: Qt.application.font.pointSize * 0.9
                                        Layout.fillWidth: true
                                    }",transactionDetails.specificDetails)
            }
    }

    TransactionDetails {
        id: transactionDetails
        anchors.fill: parent
        clip: true
    }
}
