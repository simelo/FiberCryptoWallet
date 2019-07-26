

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractListModel>
#include <QByteArray>
#include <QChildEvent>
#include <QDateTime>
#include <QEvent>
#include <QHash>
#include <QMap>
#include <QMetaMethod>
#include <QMimeData>
#include <QModelIndex>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QPersistentModelIndex>
#include <QSize>
#include <QString>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


typedef QMap<qint32, QByteArray> type378cdd;
typedef QMap<qint32, QByteArray> type378cdd;
class WalletModel1c4313: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<Wallet1c4313*> wallets READ wallets WRITE setWallets NOTIFY walletsChanged)
public:
	WalletModel1c4313(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");WalletModel1c4313_WalletModel1c4313_QRegisterMetaType();WalletModel1c4313_WalletModel1c4313_QRegisterMetaTypes();callbackWalletModel1c4313_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackWalletModel1c4313_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackWalletModel1c4313_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackWalletModel1c4313_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<Wallet1c4313*> wallets() { return ({ QList<Wallet1c4313*>* tmpP = static_cast<QList<Wallet1c4313*>*>(callbackWalletModel1c4313_Wallets(this)); QList<Wallet1c4313*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setWallets(QList<Wallet1c4313*> wallets) { callbackWalletModel1c4313_SetWallets(this, ({ QList<Wallet1c4313*>* tmpValue = new QList<Wallet1c4313*>(wallets); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_WalletsChanged(QList<Wallet1c4313*> wallets) { callbackWalletModel1c4313_WalletsChanged(this, ({ QList<Wallet1c4313*>* tmpValue = new QList<Wallet1c4313*>(wallets); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~WalletModel1c4313() { callbackWalletModel1c4313_DestroyWalletModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackWalletModel1c4313_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackWalletModel1c4313_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackWalletModel1c4313_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackWalletModel1c4313_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackWalletModel1c4313_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackWalletModel1c4313_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackWalletModel1c4313_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackWalletModel1c4313_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackWalletModel1c4313_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackWalletModel1c4313_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel1c4313_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackWalletModel1c4313_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackWalletModel1c4313_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel1c4313_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackWalletModel1c4313_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackWalletModel1c4313_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackWalletModel1c4313_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackWalletModel1c4313_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackWalletModel1c4313_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackWalletModel1c4313_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackWalletModel1c4313_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackWalletModel1c4313_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackWalletModel1c4313_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackWalletModel1c4313_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackWalletModel1c4313_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackWalletModel1c4313_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackWalletModel1c4313_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackWalletModel1c4313_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackWalletModel1c4313_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackWalletModel1c4313_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackWalletModel1c4313_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackWalletModel1c4313_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackWalletModel1c4313_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackWalletModel1c4313_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackWalletModel1c4313_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackWalletModel1c4313_ResetInternalData(this); };
	void revert() { callbackWalletModel1c4313_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackWalletModel1c4313_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackWalletModel1c4313_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackWalletModel1c4313_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackWalletModel1c4313_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel1c4313_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackWalletModel1c4313_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackWalletModel1c4313_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel1c4313_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackWalletModel1c4313_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackWalletModel1c4313_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackWalletModel1c4313_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackWalletModel1c4313_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackWalletModel1c4313_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackWalletModel1c4313_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackWalletModel1c4313_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackWalletModel1c4313_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackWalletModel1c4313_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackWalletModel1c4313_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackWalletModel1c4313_CustomEvent(this, event); };
	void deleteLater() { callbackWalletModel1c4313_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackWalletModel1c4313_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackWalletModel1c4313_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackWalletModel1c4313_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackWalletModel1c4313_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackWalletModel1c4313_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackWalletModel1c4313_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<Wallet1c4313*> walletsDefault() { return _wallets; };
	void setWalletsDefault(QList<Wallet1c4313*> p) { if (p != _wallets) { _wallets = p; walletsChanged(_wallets); } };
signals:
	void rolesChanged(type378cdd roles);
	void walletsChanged(QList<Wallet1c4313*> wallets);
public slots:
	void addWallet(Wallet1c4313* v0) { callbackWalletModel1c4313_AddWallet(this, v0); };
	void editWallet(qint32 row, QString name, bool encryptionEnabled, qint32 sky, qint32 coinHours) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackWalletModel1c4313_EditWallet(this, row, namePacked, encryptionEnabled, sky, coinHours); };
	void removeWallet(qint32 row) { callbackWalletModel1c4313_RemoveWallet(this, row); };
	void loadModel() { callbackWalletModel1c4313_LoadModel(this); };
private:
	type378cdd _roles;
	QList<Wallet1c4313*> _wallets;
};

Q_DECLARE_METATYPE(WalletModel1c4313*)


void WalletModel1c4313_WalletModel1c4313_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
	qRegisterMetaType<QList<QObject*>>("QList<Wallet1c4313*>");
}

class AddressDetails1c4313: public QObject
{
Q_OBJECT
Q_PROPERTY(QString address READ address WRITE setAddress NOTIFY addressChanged)
Q_PROPERTY(float addressSky READ addressSky WRITE setAddressSky NOTIFY addressSkyChanged)
Q_PROPERTY(qint32 addressCoinHours READ addressCoinHours WRITE setAddressCoinHours NOTIFY addressCoinHoursChanged)
public:
	AddressDetails1c4313(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");AddressDetails1c4313_AddressDetails1c4313_QRegisterMetaType();AddressDetails1c4313_AddressDetails1c4313_QRegisterMetaTypes();callbackAddressDetails1c4313_Constructor(this);};
	QString address() { return ({ Moc_PackedString tempVal = callbackAddressDetails1c4313_Address(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setAddress(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackAddressDetails1c4313_SetAddress(this, addressPacked); };
	void Signal_AddressChanged(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackAddressDetails1c4313_AddressChanged(this, addressPacked); };
	float addressSky() { return callbackAddressDetails1c4313_AddressSky(this); };
	void setAddressSky(float addressSky) { callbackAddressDetails1c4313_SetAddressSky(this, addressSky); };
	void Signal_AddressSkyChanged(float addressSky) { callbackAddressDetails1c4313_AddressSkyChanged(this, addressSky); };
	qint32 addressCoinHours() { return callbackAddressDetails1c4313_AddressCoinHours(this); };
	void setAddressCoinHours(qint32 addressCoinHours) { callbackAddressDetails1c4313_SetAddressCoinHours(this, addressCoinHours); };
	void Signal_AddressCoinHoursChanged(qint32 addressCoinHours) { callbackAddressDetails1c4313_AddressCoinHoursChanged(this, addressCoinHours); };
	 ~AddressDetails1c4313() { callbackAddressDetails1c4313_DestroyAddressDetails(this); };
	void childEvent(QChildEvent * event) { callbackAddressDetails1c4313_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackAddressDetails1c4313_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackAddressDetails1c4313_CustomEvent(this, event); };
	void deleteLater() { callbackAddressDetails1c4313_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackAddressDetails1c4313_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackAddressDetails1c4313_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackAddressDetails1c4313_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackAddressDetails1c4313_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackAddressDetails1c4313_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackAddressDetails1c4313_TimerEvent(this, event); };
	QString addressDefault() { return _address; };
	void setAddressDefault(QString p) { if (p != _address) { _address = p; addressChanged(_address); } };
	float addressSkyDefault() { return _addressSky; };
	void setAddressSkyDefault(float p) { if (p != _addressSky) { _addressSky = p; addressSkyChanged(_addressSky); } };
	qint32 addressCoinHoursDefault() { return _addressCoinHours; };
	void setAddressCoinHoursDefault(qint32 p) { if (p != _addressCoinHours) { _addressCoinHours = p; addressCoinHoursChanged(_addressCoinHours); } };
signals:
	void addressChanged(QString address);
	void addressSkyChanged(float addressSky);
	void addressCoinHoursChanged(qint32 addressCoinHours);
public slots:
private:
	QString _address;
	float _addressSky;
	qint32 _addressCoinHours;
};

Q_DECLARE_METATYPE(AddressDetails1c4313*)


void AddressDetails1c4313_AddressDetails1c4313_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class AddressList1c4313: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<AddressDetails1c4313*> addresses READ addresses WRITE setAddresses NOTIFY addressesChanged)
public:
	AddressList1c4313(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");AddressList1c4313_AddressList1c4313_QRegisterMetaType();AddressList1c4313_AddressList1c4313_QRegisterMetaTypes();callbackAddressList1c4313_Constructor(this);};
	void Signal_AddAddress(AddressDetails1c4313* transaction) { callbackAddressList1c4313_AddAddress(this, transaction); };
	void Signal_RemoveAddress(qint32 index) { callbackAddressList1c4313_RemoveAddress(this, index); };
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackAddressList1c4313_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackAddressList1c4313_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackAddressList1c4313_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<AddressDetails1c4313*> addresses() { return ({ QList<AddressDetails1c4313*>* tmpP = static_cast<QList<AddressDetails1c4313*>*>(callbackAddressList1c4313_Addresses(this)); QList<AddressDetails1c4313*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setAddresses(QList<AddressDetails1c4313*> addresses) { callbackAddressList1c4313_SetAddresses(this, ({ QList<AddressDetails1c4313*>* tmpValue = new QList<AddressDetails1c4313*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_AddressesChanged(QList<AddressDetails1c4313*> addresses) { callbackAddressList1c4313_AddressesChanged(this, ({ QList<AddressDetails1c4313*>* tmpValue = new QList<AddressDetails1c4313*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~AddressList1c4313() { callbackAddressList1c4313_DestroyAddressList(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackAddressList1c4313_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackAddressList1c4313_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackAddressList1c4313_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackAddressList1c4313_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressList1c4313_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackAddressList1c4313_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackAddressList1c4313_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackAddressList1c4313_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackAddressList1c4313_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackAddressList1c4313_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressList1c4313_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackAddressList1c4313_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackAddressList1c4313_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressList1c4313_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackAddressList1c4313_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackAddressList1c4313_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackAddressList1c4313_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackAddressList1c4313_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackAddressList1c4313_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackAddressList1c4313_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackAddressList1c4313_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackAddressList1c4313_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackAddressList1c4313_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressList1c4313_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressList1c4313_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackAddressList1c4313_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackAddressList1c4313_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackAddressList1c4313_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackAddressList1c4313_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackAddressList1c4313_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressList1c4313_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressList1c4313_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressList1c4313_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackAddressList1c4313_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackAddressList1c4313_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackAddressList1c4313_ResetInternalData(this); };
	void revert() { callbackAddressList1c4313_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackAddressList1c4313_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackAddressList1c4313_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackAddressList1c4313_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackAddressList1c4313_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressList1c4313_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackAddressList1c4313_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackAddressList1c4313_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressList1c4313_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackAddressList1c4313_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackAddressList1c4313_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackAddressList1c4313_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackAddressList1c4313_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackAddressList1c4313_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackAddressList1c4313_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackAddressList1c4313_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackAddressList1c4313_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackAddressList1c4313_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackAddressList1c4313_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackAddressList1c4313_CustomEvent(this, event); };
	void deleteLater() { callbackAddressList1c4313_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackAddressList1c4313_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackAddressList1c4313_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackAddressList1c4313_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackAddressList1c4313_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackAddressList1c4313_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackAddressList1c4313_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<AddressDetails1c4313*> addressesDefault() { return _addresses; };
	void setAddressesDefault(QList<AddressDetails1c4313*> p) { if (p != _addresses) { _addresses = p; addressesChanged(_addresses); } };
signals:
	void addAddress(AddressDetails1c4313* transaction);
	void removeAddress(qint32 index);
	void rolesChanged(type378cdd roles);
	void addressesChanged(QList<AddressDetails1c4313*> addresses);
public slots:
private:
	type378cdd _roles;
	QList<AddressDetails1c4313*> _addresses;
};

Q_DECLARE_METATYPE(AddressList1c4313*)


void AddressList1c4313_AddressList1c4313_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
	qRegisterMetaType<QList<QObject*>>("QList<AddressDetails1c4313*>");
}

class TransactionDetails1c4313: public QObject
{
Q_OBJECT
Q_PROPERTY(QDateTime date READ date WRITE setDate NOTIFY dateChanged)
Q_PROPERTY(qint32 status READ status WRITE setStatus NOTIFY statusChanged)
Q_PROPERTY(qint32 type READ type WRITE setType NOTIFY typeChanged)
Q_PROPERTY(qint32 amount READ amount WRITE setAmount NOTIFY amountChanged)
Q_PROPERTY(qint32 hoursReceived READ hoursReceived WRITE setHoursReceived NOTIFY hoursReceivedChanged)
Q_PROPERTY(qint32 hoursBurned READ hoursBurned WRITE setHoursBurned NOTIFY hoursBurnedChanged)
Q_PROPERTY(QString transactionID READ transactionID WRITE setTransactionID NOTIFY transactionIDChanged)
Q_PROPERTY(QString sentAddress READ sentAddress WRITE setSentAddress NOTIFY sentAddressChanged)
Q_PROPERTY(QString receivedAddress READ receivedAddress WRITE setReceivedAddress NOTIFY receivedAddressChanged)
Q_PROPERTY(AddressList1c4313* inputs READ inputs WRITE setInputs NOTIFY inputsChanged)
Q_PROPERTY(AddressList1c4313* outputs READ outputs WRITE setOutputs NOTIFY outputsChanged)
public:
	TransactionDetails1c4313(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");TransactionDetails1c4313_TransactionDetails1c4313_QRegisterMetaType();TransactionDetails1c4313_TransactionDetails1c4313_QRegisterMetaTypes();callbackTransactionDetails1c4313_Constructor(this);};
	QDateTime date() { return *static_cast<QDateTime*>(callbackTransactionDetails1c4313_Date(this)); };
	void setDate(QDateTime date) { callbackTransactionDetails1c4313_SetDate(this, new QDateTime(date)); };
	void Signal_DateChanged(QDateTime date) { callbackTransactionDetails1c4313_DateChanged(this, new QDateTime(date)); };
	qint32 status() { return callbackTransactionDetails1c4313_Status(this); };
	void setStatus(qint32 status) { callbackTransactionDetails1c4313_SetStatus(this, status); };
	void Signal_StatusChanged(qint32 status) { callbackTransactionDetails1c4313_StatusChanged(this, status); };
	qint32 type() { return callbackTransactionDetails1c4313_Type(this); };
	void setType(qint32 ty) { callbackTransactionDetails1c4313_SetType(this, ty); };
	void Signal_TypeChanged(qint32 ty) { callbackTransactionDetails1c4313_TypeChanged(this, ty); };
	qint32 amount() { return callbackTransactionDetails1c4313_Amount(this); };
	void setAmount(qint32 amount) { callbackTransactionDetails1c4313_SetAmount(this, amount); };
	void Signal_AmountChanged(qint32 amount) { callbackTransactionDetails1c4313_AmountChanged(this, amount); };
	qint32 hoursReceived() { return callbackTransactionDetails1c4313_HoursReceived(this); };
	void setHoursReceived(qint32 hoursReceived) { callbackTransactionDetails1c4313_SetHoursReceived(this, hoursReceived); };
	void Signal_HoursReceivedChanged(qint32 hoursReceived) { callbackTransactionDetails1c4313_HoursReceivedChanged(this, hoursReceived); };
	qint32 hoursBurned() { return callbackTransactionDetails1c4313_HoursBurned(this); };
	void setHoursBurned(qint32 hoursBurned) { callbackTransactionDetails1c4313_SetHoursBurned(this, hoursBurned); };
	void Signal_HoursBurnedChanged(qint32 hoursBurned) { callbackTransactionDetails1c4313_HoursBurnedChanged(this, hoursBurned); };
	QString transactionID() { return ({ Moc_PackedString tempVal = callbackTransactionDetails1c4313_TransactionID(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTransactionID(QString transactionID) { QByteArray tec2ac1 = transactionID.toUtf8(); Moc_PackedString transactionIDPacked = { const_cast<char*>(tec2ac1.prepend("WHITESPACE").constData()+10), tec2ac1.size()-10 };callbackTransactionDetails1c4313_SetTransactionID(this, transactionIDPacked); };
	void Signal_TransactionIDChanged(QString transactionID) { QByteArray tec2ac1 = transactionID.toUtf8(); Moc_PackedString transactionIDPacked = { const_cast<char*>(tec2ac1.prepend("WHITESPACE").constData()+10), tec2ac1.size()-10 };callbackTransactionDetails1c4313_TransactionIDChanged(this, transactionIDPacked); };
	QString sentAddress() { return ({ Moc_PackedString tempVal = callbackTransactionDetails1c4313_SentAddress(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setSentAddress(QString sentAddress) { QByteArray ted7223 = sentAddress.toUtf8(); Moc_PackedString sentAddressPacked = { const_cast<char*>(ted7223.prepend("WHITESPACE").constData()+10), ted7223.size()-10 };callbackTransactionDetails1c4313_SetSentAddress(this, sentAddressPacked); };
	void Signal_SentAddressChanged(QString sentAddress) { QByteArray ted7223 = sentAddress.toUtf8(); Moc_PackedString sentAddressPacked = { const_cast<char*>(ted7223.prepend("WHITESPACE").constData()+10), ted7223.size()-10 };callbackTransactionDetails1c4313_SentAddressChanged(this, sentAddressPacked); };
	QString receivedAddress() { return ({ Moc_PackedString tempVal = callbackTransactionDetails1c4313_ReceivedAddress(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setReceivedAddress(QString receivedAddress) { QByteArray t55bc8c = receivedAddress.toUtf8(); Moc_PackedString receivedAddressPacked = { const_cast<char*>(t55bc8c.prepend("WHITESPACE").constData()+10), t55bc8c.size()-10 };callbackTransactionDetails1c4313_SetReceivedAddress(this, receivedAddressPacked); };
	void Signal_ReceivedAddressChanged(QString receivedAddress) { QByteArray t55bc8c = receivedAddress.toUtf8(); Moc_PackedString receivedAddressPacked = { const_cast<char*>(t55bc8c.prepend("WHITESPACE").constData()+10), t55bc8c.size()-10 };callbackTransactionDetails1c4313_ReceivedAddressChanged(this, receivedAddressPacked); };
	AddressList1c4313* inputs() { return static_cast<AddressList1c4313*>(callbackTransactionDetails1c4313_Inputs(this)); };
	void setInputs(AddressList1c4313* inputs) { callbackTransactionDetails1c4313_SetInputs(this, inputs); };
	void Signal_InputsChanged(AddressList1c4313* inputs) { callbackTransactionDetails1c4313_InputsChanged(this, inputs); };
	AddressList1c4313* outputs() { return static_cast<AddressList1c4313*>(callbackTransactionDetails1c4313_Outputs(this)); };
	void setOutputs(AddressList1c4313* outputs) { callbackTransactionDetails1c4313_SetOutputs(this, outputs); };
	void Signal_OutputsChanged(AddressList1c4313* outputs) { callbackTransactionDetails1c4313_OutputsChanged(this, outputs); };
	 ~TransactionDetails1c4313() { callbackTransactionDetails1c4313_DestroyTransactionDetails(this); };
	void childEvent(QChildEvent * event) { callbackTransactionDetails1c4313_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackTransactionDetails1c4313_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackTransactionDetails1c4313_CustomEvent(this, event); };
	void deleteLater() { callbackTransactionDetails1c4313_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackTransactionDetails1c4313_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackTransactionDetails1c4313_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackTransactionDetails1c4313_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackTransactionDetails1c4313_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackTransactionDetails1c4313_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackTransactionDetails1c4313_TimerEvent(this, event); };
	QDateTime dateDefault() { return _date; };
	void setDateDefault(QDateTime p) { if (p != _date) { _date = p; dateChanged(_date); } };
	qint32 statusDefault() { return _status; };
	void setStatusDefault(qint32 p) { if (p != _status) { _status = p; statusChanged(_status); } };
	qint32 typeDefault() { return _type; };
	void setTypeDefault(qint32 p) { if (p != _type) { _type = p; typeChanged(_type); } };
	qint32 amountDefault() { return _amount; };
	void setAmountDefault(qint32 p) { if (p != _amount) { _amount = p; amountChanged(_amount); } };
	qint32 hoursReceivedDefault() { return _hoursReceived; };
	void setHoursReceivedDefault(qint32 p) { if (p != _hoursReceived) { _hoursReceived = p; hoursReceivedChanged(_hoursReceived); } };
	qint32 hoursBurnedDefault() { return _hoursBurned; };
	void setHoursBurnedDefault(qint32 p) { if (p != _hoursBurned) { _hoursBurned = p; hoursBurnedChanged(_hoursBurned); } };
	QString transactionIDDefault() { return _transactionID; };
	void setTransactionIDDefault(QString p) { if (p != _transactionID) { _transactionID = p; transactionIDChanged(_transactionID); } };
	QString sentAddressDefault() { return _sentAddress; };
	void setSentAddressDefault(QString p) { if (p != _sentAddress) { _sentAddress = p; sentAddressChanged(_sentAddress); } };
	QString receivedAddressDefault() { return _receivedAddress; };
	void setReceivedAddressDefault(QString p) { if (p != _receivedAddress) { _receivedAddress = p; receivedAddressChanged(_receivedAddress); } };
	AddressList1c4313* inputsDefault() { return _inputs; };
	void setInputsDefault(AddressList1c4313* p) { if (p != _inputs) { _inputs = p; inputsChanged(_inputs); } };
	AddressList1c4313* outputsDefault() { return _outputs; };
	void setOutputsDefault(AddressList1c4313* p) { if (p != _outputs) { _outputs = p; outputsChanged(_outputs); } };
signals:
	void dateChanged(QDateTime date);
	void statusChanged(qint32 status);
	void typeChanged(qint32 ty);
	void amountChanged(qint32 amount);
	void hoursReceivedChanged(qint32 hoursReceived);
	void hoursBurnedChanged(qint32 hoursBurned);
	void transactionIDChanged(QString transactionID);
	void sentAddressChanged(QString sentAddress);
	void receivedAddressChanged(QString receivedAddress);
	void inputsChanged(AddressList1c4313* inputs);
	void outputsChanged(AddressList1c4313* outputs);
public slots:
private:
	QDateTime _date;
	qint32 _status;
	qint32 _type;
	qint32 _amount;
	qint32 _hoursReceived;
	qint32 _hoursBurned;
	QString _transactionID;
	QString _sentAddress;
	QString _receivedAddress;
	AddressList1c4313* _inputs;
	AddressList1c4313* _outputs;
};

Q_DECLARE_METATYPE(TransactionDetails1c4313*)


void TransactionDetails1c4313_TransactionDetails1c4313_QRegisterMetaTypes() {
	qRegisterMetaType<QDateTime>();
	qRegisterMetaType<QString>();
}

class Wallet1c4313: public QObject
{
Q_OBJECT
Q_PROPERTY(QString name READ name WRITE setName NOTIFY nameChanged)
Q_PROPERTY(qint32 encryptionEnabled READ encryptionEnabled WRITE setEncryptionEnabled NOTIFY encryptionEnabledChanged)
Q_PROPERTY(qint32 sky READ sky WRITE setSky NOTIFY skyChanged)
Q_PROPERTY(qint32 coinHours READ coinHours WRITE setCoinHours NOTIFY coinHoursChanged)
Q_PROPERTY(QString fileName READ fileName WRITE setFileName NOTIFY fileNameChanged)
public:
	Wallet1c4313(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");Wallet1c4313_Wallet1c4313_QRegisterMetaType();Wallet1c4313_Wallet1c4313_QRegisterMetaTypes();callbackWallet1c4313_Constructor(this);};
	QString name() { return ({ Moc_PackedString tempVal = callbackWallet1c4313_Name(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setName(QString name) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackWallet1c4313_SetName(this, namePacked); };
	void Signal_NameChanged(QString name) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackWallet1c4313_NameChanged(this, namePacked); };
	qint32 encryptionEnabled() { return callbackWallet1c4313_EncryptionEnabled(this); };
	void setEncryptionEnabled(qint32 encryptionEnabled) { callbackWallet1c4313_SetEncryptionEnabled(this, encryptionEnabled); };
	void Signal_EncryptionEnabledChanged(qint32 encryptionEnabled) { callbackWallet1c4313_EncryptionEnabledChanged(this, encryptionEnabled); };
	qint32 sky() { return callbackWallet1c4313_Sky(this); };
	void setSky(qint32 sky) { callbackWallet1c4313_SetSky(this, sky); };
	void Signal_SkyChanged(qint32 sky) { callbackWallet1c4313_SkyChanged(this, sky); };
	qint32 coinHours() { return callbackWallet1c4313_CoinHours(this); };
	void setCoinHours(qint32 coinHours) { callbackWallet1c4313_SetCoinHours(this, coinHours); };
	void Signal_CoinHoursChanged(qint32 coinHours) { callbackWallet1c4313_CoinHoursChanged(this, coinHours); };
	QString fileName() { return ({ Moc_PackedString tempVal = callbackWallet1c4313_FileName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setFileName(QString fileName) { QByteArray td83e09 = fileName.toUtf8(); Moc_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackWallet1c4313_SetFileName(this, fileNamePacked); };
	void Signal_FileNameChanged(QString fileName) { QByteArray td83e09 = fileName.toUtf8(); Moc_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackWallet1c4313_FileNameChanged(this, fileNamePacked); };
	 ~Wallet1c4313() { callbackWallet1c4313_DestroyWallet(this); };
	void childEvent(QChildEvent * event) { callbackWallet1c4313_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackWallet1c4313_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackWallet1c4313_CustomEvent(this, event); };
	void deleteLater() { callbackWallet1c4313_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackWallet1c4313_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackWallet1c4313_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackWallet1c4313_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackWallet1c4313_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackWallet1c4313_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackWallet1c4313_TimerEvent(this, event); };
	QString nameDefault() { return _name; };
	void setNameDefault(QString p) { if (p != _name) { _name = p; nameChanged(_name); } };
	qint32 encryptionEnabledDefault() { return _encryptionEnabled; };
	void setEncryptionEnabledDefault(qint32 p) { if (p != _encryptionEnabled) { _encryptionEnabled = p; encryptionEnabledChanged(_encryptionEnabled); } };
	qint32 skyDefault() { return _sky; };
	void setSkyDefault(qint32 p) { if (p != _sky) { _sky = p; skyChanged(_sky); } };
	qint32 coinHoursDefault() { return _coinHours; };
	void setCoinHoursDefault(qint32 p) { if (p != _coinHours) { _coinHours = p; coinHoursChanged(_coinHours); } };
	QString fileNameDefault() { return _fileName; };
	void setFileNameDefault(QString p) { if (p != _fileName) { _fileName = p; fileNameChanged(_fileName); } };
signals:
	void nameChanged(QString name);
	void encryptionEnabledChanged(qint32 encryptionEnabled);
	void skyChanged(qint32 sky);
	void coinHoursChanged(qint32 coinHours);
	void fileNameChanged(QString fileName);
public slots:
private:
	QString _name;
	qint32 _encryptionEnabled;
	qint32 _sky;
	qint32 _coinHours;
	QString _fileName;
};

Q_DECLARE_METATYPE(Wallet1c4313*)


void Wallet1c4313_Wallet1c4313_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

struct Moc_PackedString AddressDetails1c4313_Address(void* ptr)
{
	return ({ QByteArray t686357 = static_cast<AddressDetails1c4313*>(ptr)->address().toUtf8(); Moc_PackedString { const_cast<char*>(t686357.prepend("WHITESPACE").constData()+10), t686357.size()-10 }; });
}

struct Moc_PackedString AddressDetails1c4313_AddressDefault(void* ptr)
{
	return ({ QByteArray tbaed36 = static_cast<AddressDetails1c4313*>(ptr)->addressDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tbaed36.prepend("WHITESPACE").constData()+10), tbaed36.size()-10 }; });
}

void AddressDetails1c4313_SetAddress(void* ptr, struct Moc_PackedString address)
{
	static_cast<AddressDetails1c4313*>(ptr)->setAddress(QString::fromUtf8(address.data, address.len));
}

void AddressDetails1c4313_SetAddressDefault(void* ptr, struct Moc_PackedString address)
{
	static_cast<AddressDetails1c4313*>(ptr)->setAddressDefault(QString::fromUtf8(address.data, address.len));
}

void AddressDetails1c4313_ConnectAddressChanged(void* ptr)
{
	QObject::connect(static_cast<AddressDetails1c4313*>(ptr), static_cast<void (AddressDetails1c4313::*)(QString)>(&AddressDetails1c4313::addressChanged), static_cast<AddressDetails1c4313*>(ptr), static_cast<void (AddressDetails1c4313::*)(QString)>(&AddressDetails1c4313::Signal_AddressChanged));
}

void AddressDetails1c4313_DisconnectAddressChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressDetails1c4313*>(ptr), static_cast<void (AddressDetails1c4313::*)(QString)>(&AddressDetails1c4313::addressChanged), static_cast<AddressDetails1c4313*>(ptr), static_cast<void (AddressDetails1c4313::*)(QString)>(&AddressDetails1c4313::Signal_AddressChanged));
}

void AddressDetails1c4313_AddressChanged(void* ptr, struct Moc_PackedString address)
{
	static_cast<AddressDetails1c4313*>(ptr)->addressChanged(QString::fromUtf8(address.data, address.len));
}

float AddressDetails1c4313_AddressSky(void* ptr)
{
	return static_cast<AddressDetails1c4313*>(ptr)->addressSky();
}

float AddressDetails1c4313_AddressSkyDefault(void* ptr)
{
	return static_cast<AddressDetails1c4313*>(ptr)->addressSkyDefault();
}

void AddressDetails1c4313_SetAddressSky(void* ptr, float addressSky)
{
	static_cast<AddressDetails1c4313*>(ptr)->setAddressSky(addressSky);
}

void AddressDetails1c4313_SetAddressSkyDefault(void* ptr, float addressSky)
{
	static_cast<AddressDetails1c4313*>(ptr)->setAddressSkyDefault(addressSky);
}

void AddressDetails1c4313_ConnectAddressSkyChanged(void* ptr)
{
	QObject::connect(static_cast<AddressDetails1c4313*>(ptr), static_cast<void (AddressDetails1c4313::*)(float)>(&AddressDetails1c4313::addressSkyChanged), static_cast<AddressDetails1c4313*>(ptr), static_cast<void (AddressDetails1c4313::*)(float)>(&AddressDetails1c4313::Signal_AddressSkyChanged));
}

void AddressDetails1c4313_DisconnectAddressSkyChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressDetails1c4313*>(ptr), static_cast<void (AddressDetails1c4313::*)(float)>(&AddressDetails1c4313::addressSkyChanged), static_cast<AddressDetails1c4313*>(ptr), static_cast<void (AddressDetails1c4313::*)(float)>(&AddressDetails1c4313::Signal_AddressSkyChanged));
}

void AddressDetails1c4313_AddressSkyChanged(void* ptr, float addressSky)
{
	static_cast<AddressDetails1c4313*>(ptr)->addressSkyChanged(addressSky);
}

int AddressDetails1c4313_AddressCoinHours(void* ptr)
{
	return static_cast<AddressDetails1c4313*>(ptr)->addressCoinHours();
}

int AddressDetails1c4313_AddressCoinHoursDefault(void* ptr)
{
	return static_cast<AddressDetails1c4313*>(ptr)->addressCoinHoursDefault();
}

void AddressDetails1c4313_SetAddressCoinHours(void* ptr, int addressCoinHours)
{
	static_cast<AddressDetails1c4313*>(ptr)->setAddressCoinHours(addressCoinHours);
}

void AddressDetails1c4313_SetAddressCoinHoursDefault(void* ptr, int addressCoinHours)
{
	static_cast<AddressDetails1c4313*>(ptr)->setAddressCoinHoursDefault(addressCoinHours);
}

void AddressDetails1c4313_ConnectAddressCoinHoursChanged(void* ptr)
{
	QObject::connect(static_cast<AddressDetails1c4313*>(ptr), static_cast<void (AddressDetails1c4313::*)(qint32)>(&AddressDetails1c4313::addressCoinHoursChanged), static_cast<AddressDetails1c4313*>(ptr), static_cast<void (AddressDetails1c4313::*)(qint32)>(&AddressDetails1c4313::Signal_AddressCoinHoursChanged));
}

void AddressDetails1c4313_DisconnectAddressCoinHoursChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressDetails1c4313*>(ptr), static_cast<void (AddressDetails1c4313::*)(qint32)>(&AddressDetails1c4313::addressCoinHoursChanged), static_cast<AddressDetails1c4313*>(ptr), static_cast<void (AddressDetails1c4313::*)(qint32)>(&AddressDetails1c4313::Signal_AddressCoinHoursChanged));
}

void AddressDetails1c4313_AddressCoinHoursChanged(void* ptr, int addressCoinHours)
{
	static_cast<AddressDetails1c4313*>(ptr)->addressCoinHoursChanged(addressCoinHours);
}

int AddressDetails1c4313_AddressDetails1c4313_QRegisterMetaType()
{
	return qRegisterMetaType<AddressDetails1c4313*>();
}

int AddressDetails1c4313_AddressDetails1c4313_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<AddressDetails1c4313*>(const_cast<const char*>(typeName));
}

int AddressDetails1c4313_AddressDetails1c4313_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressDetails1c4313>();
#else
	return 0;
#endif
}

int AddressDetails1c4313_AddressDetails1c4313_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressDetails1c4313>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* AddressDetails1c4313___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressDetails1c4313___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressDetails1c4313___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* AddressDetails1c4313___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressDetails1c4313___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* AddressDetails1c4313___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* AddressDetails1c4313___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressDetails1c4313___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressDetails1c4313___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressDetails1c4313___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressDetails1c4313___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressDetails1c4313___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressDetails1c4313___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressDetails1c4313___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressDetails1c4313___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressDetails1c4313_NewAddressDetails(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new AddressDetails1c4313(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new AddressDetails1c4313(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new AddressDetails1c4313(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new AddressDetails1c4313(static_cast<QWindow*>(parent));
	} else {
		return new AddressDetails1c4313(static_cast<QObject*>(parent));
	}
}

void AddressDetails1c4313_DestroyAddressDetails(void* ptr)
{
	static_cast<AddressDetails1c4313*>(ptr)->~AddressDetails1c4313();
}

void AddressDetails1c4313_DestroyAddressDetailsDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void AddressDetails1c4313_ChildEventDefault(void* ptr, void* event)
{
	static_cast<AddressDetails1c4313*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void AddressDetails1c4313_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressDetails1c4313*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void AddressDetails1c4313_CustomEventDefault(void* ptr, void* event)
{
	static_cast<AddressDetails1c4313*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void AddressDetails1c4313_DeleteLaterDefault(void* ptr)
{
	static_cast<AddressDetails1c4313*>(ptr)->QObject::deleteLater();
}

void AddressDetails1c4313_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressDetails1c4313*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char AddressDetails1c4313_EventDefault(void* ptr, void* e)
{
	return static_cast<AddressDetails1c4313*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char AddressDetails1c4313_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<AddressDetails1c4313*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void AddressDetails1c4313_TimerEventDefault(void* ptr, void* event)
{
	static_cast<AddressDetails1c4313*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void AddressList1c4313_ConnectAddAddress(void* ptr)
{
	QObject::connect(static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(AddressDetails1c4313*)>(&AddressList1c4313::addAddress), static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(AddressDetails1c4313*)>(&AddressList1c4313::Signal_AddAddress));
}

void AddressList1c4313_DisconnectAddAddress(void* ptr)
{
	QObject::disconnect(static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(AddressDetails1c4313*)>(&AddressList1c4313::addAddress), static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(AddressDetails1c4313*)>(&AddressList1c4313::Signal_AddAddress));
}

void AddressList1c4313_AddAddress(void* ptr, void* transaction)
{
	static_cast<AddressList1c4313*>(ptr)->addAddress(static_cast<AddressDetails1c4313*>(transaction));
}

void AddressList1c4313_ConnectRemoveAddress(void* ptr)
{
	QObject::connect(static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(qint32)>(&AddressList1c4313::removeAddress), static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(qint32)>(&AddressList1c4313::Signal_RemoveAddress));
}

void AddressList1c4313_DisconnectRemoveAddress(void* ptr)
{
	QObject::disconnect(static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(qint32)>(&AddressList1c4313::removeAddress), static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(qint32)>(&AddressList1c4313::Signal_RemoveAddress));
}

void AddressList1c4313_RemoveAddress(void* ptr, int index)
{
	static_cast<AddressList1c4313*>(ptr)->removeAddress(index);
}

struct Moc_PackedList AddressList1c4313_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressList1c4313*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressList1c4313_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressList1c4313*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressList1c4313_SetRoles(void* ptr, void* roles)
{
	static_cast<AddressList1c4313*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressList1c4313_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<AddressList1c4313*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressList1c4313_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(QMap<qint32, QByteArray>)>(&AddressList1c4313::rolesChanged), static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(QMap<qint32, QByteArray>)>(&AddressList1c4313::Signal_RolesChanged));
}

void AddressList1c4313_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(QMap<qint32, QByteArray>)>(&AddressList1c4313::rolesChanged), static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(QMap<qint32, QByteArray>)>(&AddressList1c4313::Signal_RolesChanged));
}

void AddressList1c4313_RolesChanged(void* ptr, void* roles)
{
	static_cast<AddressList1c4313*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList AddressList1c4313_Addresses(void* ptr)
{
	return ({ QList<AddressDetails1c4313*>* tmpValue = new QList<AddressDetails1c4313*>(static_cast<AddressList1c4313*>(ptr)->addresses()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressList1c4313_AddressesDefault(void* ptr)
{
	return ({ QList<AddressDetails1c4313*>* tmpValue = new QList<AddressDetails1c4313*>(static_cast<AddressList1c4313*>(ptr)->addressesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressList1c4313_SetAddresses(void* ptr, void* addresses)
{
	static_cast<AddressList1c4313*>(ptr)->setAddresses(({ QList<AddressDetails1c4313*>* tmpP = static_cast<QList<AddressDetails1c4313*>*>(addresses); QList<AddressDetails1c4313*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressList1c4313_SetAddressesDefault(void* ptr, void* addresses)
{
	static_cast<AddressList1c4313*>(ptr)->setAddressesDefault(({ QList<AddressDetails1c4313*>* tmpP = static_cast<QList<AddressDetails1c4313*>*>(addresses); QList<AddressDetails1c4313*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressList1c4313_ConnectAddressesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(QList<AddressDetails1c4313*>)>(&AddressList1c4313::addressesChanged), static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(QList<AddressDetails1c4313*>)>(&AddressList1c4313::Signal_AddressesChanged));
}

void AddressList1c4313_DisconnectAddressesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(QList<AddressDetails1c4313*>)>(&AddressList1c4313::addressesChanged), static_cast<AddressList1c4313*>(ptr), static_cast<void (AddressList1c4313::*)(QList<AddressDetails1c4313*>)>(&AddressList1c4313::Signal_AddressesChanged));
}

void AddressList1c4313_AddressesChanged(void* ptr, void* addresses)
{
	static_cast<AddressList1c4313*>(ptr)->addressesChanged(({ QList<AddressDetails1c4313*>* tmpP = static_cast<QList<AddressDetails1c4313*>*>(addresses); QList<AddressDetails1c4313*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int AddressList1c4313_AddressList1c4313_QRegisterMetaType()
{
	return qRegisterMetaType<AddressList1c4313*>();
}

int AddressList1c4313_AddressList1c4313_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<AddressList1c4313*>(const_cast<const char*>(typeName));
}

int AddressList1c4313_AddressList1c4313_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressList1c4313>();
#else
	return 0;
#endif
}

int AddressList1c4313_AddressList1c4313_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressList1c4313>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int AddressList1c4313_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList1c4313_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressList1c4313_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressList1c4313_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList1c4313_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressList1c4313_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressList1c4313_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList1c4313_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressList1c4313_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressList1c4313___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList1c4313___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressList1c4313___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressList1c4313___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList1c4313___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressList1c4313___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int AddressList1c4313___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void AddressList1c4313___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* AddressList1c4313___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* AddressList1c4313___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressList1c4313___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressList1c4313___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressList1c4313___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressList1c4313___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList1c4313___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressList1c4313___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressList1c4313___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList1c4313___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressList1c4313___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressList1c4313___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList1c4313___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressList1c4313___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressList1c4313___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList1c4313___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressList1c4313___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressList1c4313___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList1c4313___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressList1c4313___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressList1c4313___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void AddressList1c4313___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressList1c4313___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList AddressList1c4313___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressList1c4313___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressList1c4313___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressList1c4313___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressList1c4313___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressList1c4313_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList1c4313_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressList1c4313_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressList1c4313_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList1c4313_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressList1c4313_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressList1c4313___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList1c4313___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressList1c4313___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* AddressList1c4313___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList1c4313___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* AddressList1c4313___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* AddressList1c4313___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList1c4313___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressList1c4313___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressList1c4313___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList1c4313___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressList1c4313___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressList1c4313___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList1c4313___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressList1c4313___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressList1c4313___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressList1c4313___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressList1c4313___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressList1c4313___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressList1c4313___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressList1c4313___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressList1c4313___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressList1c4313___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressList1c4313___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressList1c4313___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressList1c4313___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressList1c4313___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressList1c4313___addresses_atList(void* ptr, int i)
{
	return ({AddressDetails1c4313* tmp = static_cast<QList<AddressDetails1c4313*>*>(ptr)->at(i); if (i == static_cast<QList<AddressDetails1c4313*>*>(ptr)->size()-1) { static_cast<QList<AddressDetails1c4313*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList1c4313___addresses_setList(void* ptr, void* i)
{
	static_cast<QList<AddressDetails1c4313*>*>(ptr)->append(static_cast<AddressDetails1c4313*>(i));
}

void* AddressList1c4313___addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<AddressDetails1c4313*>();
}

void* AddressList1c4313___setAddresses_addresses_atList(void* ptr, int i)
{
	return ({AddressDetails1c4313* tmp = static_cast<QList<AddressDetails1c4313*>*>(ptr)->at(i); if (i == static_cast<QList<AddressDetails1c4313*>*>(ptr)->size()-1) { static_cast<QList<AddressDetails1c4313*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList1c4313___setAddresses_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<AddressDetails1c4313*>*>(ptr)->append(static_cast<AddressDetails1c4313*>(i));
}

void* AddressList1c4313___setAddresses_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<AddressDetails1c4313*>();
}

void* AddressList1c4313___addressesChanged_addresses_atList(void* ptr, int i)
{
	return ({AddressDetails1c4313* tmp = static_cast<QList<AddressDetails1c4313*>*>(ptr)->at(i); if (i == static_cast<QList<AddressDetails1c4313*>*>(ptr)->size()-1) { static_cast<QList<AddressDetails1c4313*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList1c4313___addressesChanged_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<AddressDetails1c4313*>*>(ptr)->append(static_cast<AddressDetails1c4313*>(i));
}

void* AddressList1c4313___addressesChanged_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<AddressDetails1c4313*>();
}

int AddressList1c4313_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList1c4313_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressList1c4313_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressList1c4313_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList1c4313_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressList1c4313_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressList1c4313_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList1c4313_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressList1c4313_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* AddressList1c4313_NewAddressList(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new AddressList1c4313(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new AddressList1c4313(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new AddressList1c4313(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new AddressList1c4313(static_cast<QWindow*>(parent));
	} else {
		return new AddressList1c4313(static_cast<QObject*>(parent));
	}
}

void AddressList1c4313_DestroyAddressList(void* ptr)
{
	static_cast<AddressList1c4313*>(ptr)->~AddressList1c4313();
}

void AddressList1c4313_DestroyAddressListDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char AddressList1c4313_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long AddressList1c4313_FlagsDefault(void* ptr, void* index)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* AddressList1c4313_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* AddressList1c4313_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* AddressList1c4313_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char AddressList1c4313_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char AddressList1c4313_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int AddressList1c4313_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* AddressList1c4313_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void AddressList1c4313_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char AddressList1c4313_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* AddressList1c4313_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char AddressList1c4313_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressList1c4313_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList AddressList1c4313_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressList1c4313_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressList1c4313_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString AddressList1c4313_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t29e12e = static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t29e12e.prepend("WHITESPACE").constData()+10), t29e12e.size()-10 }; });
}

char AddressList1c4313_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char AddressList1c4313_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* AddressList1c4313_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char AddressList1c4313_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressList1c4313_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void AddressList1c4313_ResetInternalDataDefault(void* ptr)
{
	static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::resetInternalData();
}

void AddressList1c4313_RevertDefault(void* ptr)
{
	static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList AddressList1c4313_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressList1c4313_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char AddressList1c4313_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char AddressList1c4313_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char AddressList1c4313_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void AddressList1c4313_SortDefault(void* ptr, int column, long long order)
{
	static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* AddressList1c4313_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char AddressList1c4313_SubmitDefault(void* ptr)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::submit();
}

long long AddressList1c4313_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long AddressList1c4313_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::supportedDropActions();
}

void AddressList1c4313_ChildEventDefault(void* ptr, void* event)
{
	static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void AddressList1c4313_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void AddressList1c4313_CustomEventDefault(void* ptr, void* event)
{
	static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void AddressList1c4313_DeleteLaterDefault(void* ptr)
{
	static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::deleteLater();
}

void AddressList1c4313_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char AddressList1c4313_EventDefault(void* ptr, void* e)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char AddressList1c4313_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void AddressList1c4313_TimerEventDefault(void* ptr, void* event)
{
	static_cast<AddressList1c4313*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void* TransactionDetails1c4313_Date(void* ptr)
{
	return new QDateTime(static_cast<TransactionDetails1c4313*>(ptr)->date());
}

void* TransactionDetails1c4313_DateDefault(void* ptr)
{
	return new QDateTime(static_cast<TransactionDetails1c4313*>(ptr)->dateDefault());
}

void TransactionDetails1c4313_SetDate(void* ptr, void* date)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setDate(*static_cast<QDateTime*>(date));
}

void TransactionDetails1c4313_SetDateDefault(void* ptr, void* date)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setDateDefault(*static_cast<QDateTime*>(date));
}

void TransactionDetails1c4313_ConnectDateChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QDateTime)>(&TransactionDetails1c4313::dateChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QDateTime)>(&TransactionDetails1c4313::Signal_DateChanged));
}

void TransactionDetails1c4313_DisconnectDateChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QDateTime)>(&TransactionDetails1c4313::dateChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QDateTime)>(&TransactionDetails1c4313::Signal_DateChanged));
}

void TransactionDetails1c4313_DateChanged(void* ptr, void* date)
{
	static_cast<TransactionDetails1c4313*>(ptr)->dateChanged(*static_cast<QDateTime*>(date));
}

int TransactionDetails1c4313_Status(void* ptr)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->status();
}

int TransactionDetails1c4313_StatusDefault(void* ptr)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->statusDefault();
}

void TransactionDetails1c4313_SetStatus(void* ptr, int status)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setStatus(status);
}

void TransactionDetails1c4313_SetStatusDefault(void* ptr, int status)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setStatusDefault(status);
}

void TransactionDetails1c4313_ConnectStatusChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::statusChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::Signal_StatusChanged));
}

void TransactionDetails1c4313_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::statusChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::Signal_StatusChanged));
}

void TransactionDetails1c4313_StatusChanged(void* ptr, int status)
{
	static_cast<TransactionDetails1c4313*>(ptr)->statusChanged(status);
}

int TransactionDetails1c4313_Type(void* ptr)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->type();
}

int TransactionDetails1c4313_TypeDefault(void* ptr)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->typeDefault();
}

void TransactionDetails1c4313_SetType(void* ptr, int ty)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setType(ty);
}

void TransactionDetails1c4313_SetTypeDefault(void* ptr, int ty)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setTypeDefault(ty);
}

void TransactionDetails1c4313_ConnectTypeChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::typeChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::Signal_TypeChanged));
}

void TransactionDetails1c4313_DisconnectTypeChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::typeChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::Signal_TypeChanged));
}

void TransactionDetails1c4313_TypeChanged(void* ptr, int ty)
{
	static_cast<TransactionDetails1c4313*>(ptr)->typeChanged(ty);
}

int TransactionDetails1c4313_Amount(void* ptr)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->amount();
}

int TransactionDetails1c4313_AmountDefault(void* ptr)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->amountDefault();
}

void TransactionDetails1c4313_SetAmount(void* ptr, int amount)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setAmount(amount);
}

void TransactionDetails1c4313_SetAmountDefault(void* ptr, int amount)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setAmountDefault(amount);
}

void TransactionDetails1c4313_ConnectAmountChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::amountChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::Signal_AmountChanged));
}

void TransactionDetails1c4313_DisconnectAmountChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::amountChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::Signal_AmountChanged));
}

void TransactionDetails1c4313_AmountChanged(void* ptr, int amount)
{
	static_cast<TransactionDetails1c4313*>(ptr)->amountChanged(amount);
}

int TransactionDetails1c4313_HoursReceived(void* ptr)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->hoursReceived();
}

int TransactionDetails1c4313_HoursReceivedDefault(void* ptr)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->hoursReceivedDefault();
}

void TransactionDetails1c4313_SetHoursReceived(void* ptr, int hoursReceived)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setHoursReceived(hoursReceived);
}

void TransactionDetails1c4313_SetHoursReceivedDefault(void* ptr, int hoursReceived)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setHoursReceivedDefault(hoursReceived);
}

void TransactionDetails1c4313_ConnectHoursReceivedChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::hoursReceivedChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::Signal_HoursReceivedChanged));
}

void TransactionDetails1c4313_DisconnectHoursReceivedChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::hoursReceivedChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::Signal_HoursReceivedChanged));
}

void TransactionDetails1c4313_HoursReceivedChanged(void* ptr, int hoursReceived)
{
	static_cast<TransactionDetails1c4313*>(ptr)->hoursReceivedChanged(hoursReceived);
}

int TransactionDetails1c4313_HoursBurned(void* ptr)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->hoursBurned();
}

int TransactionDetails1c4313_HoursBurnedDefault(void* ptr)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->hoursBurnedDefault();
}

void TransactionDetails1c4313_SetHoursBurned(void* ptr, int hoursBurned)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setHoursBurned(hoursBurned);
}

void TransactionDetails1c4313_SetHoursBurnedDefault(void* ptr, int hoursBurned)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setHoursBurnedDefault(hoursBurned);
}

void TransactionDetails1c4313_ConnectHoursBurnedChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::hoursBurnedChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::Signal_HoursBurnedChanged));
}

void TransactionDetails1c4313_DisconnectHoursBurnedChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::hoursBurnedChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(qint32)>(&TransactionDetails1c4313::Signal_HoursBurnedChanged));
}

void TransactionDetails1c4313_HoursBurnedChanged(void* ptr, int hoursBurned)
{
	static_cast<TransactionDetails1c4313*>(ptr)->hoursBurnedChanged(hoursBurned);
}

struct Moc_PackedString TransactionDetails1c4313_TransactionID(void* ptr)
{
	return ({ QByteArray tc8c5df = static_cast<TransactionDetails1c4313*>(ptr)->transactionID().toUtf8(); Moc_PackedString { const_cast<char*>(tc8c5df.prepend("WHITESPACE").constData()+10), tc8c5df.size()-10 }; });
}

struct Moc_PackedString TransactionDetails1c4313_TransactionIDDefault(void* ptr)
{
	return ({ QByteArray t299bde = static_cast<TransactionDetails1c4313*>(ptr)->transactionIDDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t299bde.prepend("WHITESPACE").constData()+10), t299bde.size()-10 }; });
}

void TransactionDetails1c4313_SetTransactionID(void* ptr, struct Moc_PackedString transactionID)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setTransactionID(QString::fromUtf8(transactionID.data, transactionID.len));
}

void TransactionDetails1c4313_SetTransactionIDDefault(void* ptr, struct Moc_PackedString transactionID)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setTransactionIDDefault(QString::fromUtf8(transactionID.data, transactionID.len));
}

void TransactionDetails1c4313_ConnectTransactionIDChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QString)>(&TransactionDetails1c4313::transactionIDChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QString)>(&TransactionDetails1c4313::Signal_TransactionIDChanged));
}

void TransactionDetails1c4313_DisconnectTransactionIDChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QString)>(&TransactionDetails1c4313::transactionIDChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QString)>(&TransactionDetails1c4313::Signal_TransactionIDChanged));
}

void TransactionDetails1c4313_TransactionIDChanged(void* ptr, struct Moc_PackedString transactionID)
{
	static_cast<TransactionDetails1c4313*>(ptr)->transactionIDChanged(QString::fromUtf8(transactionID.data, transactionID.len));
}

struct Moc_PackedString TransactionDetails1c4313_SentAddress(void* ptr)
{
	return ({ QByteArray te78448 = static_cast<TransactionDetails1c4313*>(ptr)->sentAddress().toUtf8(); Moc_PackedString { const_cast<char*>(te78448.prepend("WHITESPACE").constData()+10), te78448.size()-10 }; });
}

struct Moc_PackedString TransactionDetails1c4313_SentAddressDefault(void* ptr)
{
	return ({ QByteArray tf856c0 = static_cast<TransactionDetails1c4313*>(ptr)->sentAddressDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tf856c0.prepend("WHITESPACE").constData()+10), tf856c0.size()-10 }; });
}

void TransactionDetails1c4313_SetSentAddress(void* ptr, struct Moc_PackedString sentAddress)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setSentAddress(QString::fromUtf8(sentAddress.data, sentAddress.len));
}

void TransactionDetails1c4313_SetSentAddressDefault(void* ptr, struct Moc_PackedString sentAddress)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setSentAddressDefault(QString::fromUtf8(sentAddress.data, sentAddress.len));
}

void TransactionDetails1c4313_ConnectSentAddressChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QString)>(&TransactionDetails1c4313::sentAddressChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QString)>(&TransactionDetails1c4313::Signal_SentAddressChanged));
}

void TransactionDetails1c4313_DisconnectSentAddressChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QString)>(&TransactionDetails1c4313::sentAddressChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QString)>(&TransactionDetails1c4313::Signal_SentAddressChanged));
}

void TransactionDetails1c4313_SentAddressChanged(void* ptr, struct Moc_PackedString sentAddress)
{
	static_cast<TransactionDetails1c4313*>(ptr)->sentAddressChanged(QString::fromUtf8(sentAddress.data, sentAddress.len));
}

struct Moc_PackedString TransactionDetails1c4313_ReceivedAddress(void* ptr)
{
	return ({ QByteArray t3a4c8c = static_cast<TransactionDetails1c4313*>(ptr)->receivedAddress().toUtf8(); Moc_PackedString { const_cast<char*>(t3a4c8c.prepend("WHITESPACE").constData()+10), t3a4c8c.size()-10 }; });
}

struct Moc_PackedString TransactionDetails1c4313_ReceivedAddressDefault(void* ptr)
{
	return ({ QByteArray ta1cf30 = static_cast<TransactionDetails1c4313*>(ptr)->receivedAddressDefault().toUtf8(); Moc_PackedString { const_cast<char*>(ta1cf30.prepend("WHITESPACE").constData()+10), ta1cf30.size()-10 }; });
}

void TransactionDetails1c4313_SetReceivedAddress(void* ptr, struct Moc_PackedString receivedAddress)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setReceivedAddress(QString::fromUtf8(receivedAddress.data, receivedAddress.len));
}

void TransactionDetails1c4313_SetReceivedAddressDefault(void* ptr, struct Moc_PackedString receivedAddress)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setReceivedAddressDefault(QString::fromUtf8(receivedAddress.data, receivedAddress.len));
}

void TransactionDetails1c4313_ConnectReceivedAddressChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QString)>(&TransactionDetails1c4313::receivedAddressChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QString)>(&TransactionDetails1c4313::Signal_ReceivedAddressChanged));
}

void TransactionDetails1c4313_DisconnectReceivedAddressChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QString)>(&TransactionDetails1c4313::receivedAddressChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(QString)>(&TransactionDetails1c4313::Signal_ReceivedAddressChanged));
}

void TransactionDetails1c4313_ReceivedAddressChanged(void* ptr, struct Moc_PackedString receivedAddress)
{
	static_cast<TransactionDetails1c4313*>(ptr)->receivedAddressChanged(QString::fromUtf8(receivedAddress.data, receivedAddress.len));
}

void* TransactionDetails1c4313_Inputs(void* ptr)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->inputs();
}

void* TransactionDetails1c4313_InputsDefault(void* ptr)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->inputsDefault();
}

void TransactionDetails1c4313_SetInputs(void* ptr, void* inputs)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setInputs(static_cast<AddressList1c4313*>(inputs));
}

void TransactionDetails1c4313_SetInputsDefault(void* ptr, void* inputs)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setInputsDefault(static_cast<AddressList1c4313*>(inputs));
}

void TransactionDetails1c4313_ConnectInputsChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(AddressList1c4313*)>(&TransactionDetails1c4313::inputsChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(AddressList1c4313*)>(&TransactionDetails1c4313::Signal_InputsChanged));
}

void TransactionDetails1c4313_DisconnectInputsChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(AddressList1c4313*)>(&TransactionDetails1c4313::inputsChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(AddressList1c4313*)>(&TransactionDetails1c4313::Signal_InputsChanged));
}

void TransactionDetails1c4313_InputsChanged(void* ptr, void* inputs)
{
	static_cast<TransactionDetails1c4313*>(ptr)->inputsChanged(static_cast<AddressList1c4313*>(inputs));
}

void* TransactionDetails1c4313_Outputs(void* ptr)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->outputs();
}

void* TransactionDetails1c4313_OutputsDefault(void* ptr)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->outputsDefault();
}

void TransactionDetails1c4313_SetOutputs(void* ptr, void* outputs)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setOutputs(static_cast<AddressList1c4313*>(outputs));
}

void TransactionDetails1c4313_SetOutputsDefault(void* ptr, void* outputs)
{
	static_cast<TransactionDetails1c4313*>(ptr)->setOutputsDefault(static_cast<AddressList1c4313*>(outputs));
}

void TransactionDetails1c4313_ConnectOutputsChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(AddressList1c4313*)>(&TransactionDetails1c4313::outputsChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(AddressList1c4313*)>(&TransactionDetails1c4313::Signal_OutputsChanged));
}

void TransactionDetails1c4313_DisconnectOutputsChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(AddressList1c4313*)>(&TransactionDetails1c4313::outputsChanged), static_cast<TransactionDetails1c4313*>(ptr), static_cast<void (TransactionDetails1c4313::*)(AddressList1c4313*)>(&TransactionDetails1c4313::Signal_OutputsChanged));
}

void TransactionDetails1c4313_OutputsChanged(void* ptr, void* outputs)
{
	static_cast<TransactionDetails1c4313*>(ptr)->outputsChanged(static_cast<AddressList1c4313*>(outputs));
}

int TransactionDetails1c4313_TransactionDetails1c4313_QRegisterMetaType()
{
	return qRegisterMetaType<TransactionDetails1c4313*>();
}

int TransactionDetails1c4313_TransactionDetails1c4313_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<TransactionDetails1c4313*>(const_cast<const char*>(typeName));
}

int TransactionDetails1c4313_TransactionDetails1c4313_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TransactionDetails1c4313>();
#else
	return 0;
#endif
}

int TransactionDetails1c4313_TransactionDetails1c4313_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TransactionDetails1c4313>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* TransactionDetails1c4313___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetails1c4313___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetails1c4313___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* TransactionDetails1c4313___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TransactionDetails1c4313___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* TransactionDetails1c4313___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* TransactionDetails1c4313___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetails1c4313___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetails1c4313___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TransactionDetails1c4313___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetails1c4313___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetails1c4313___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TransactionDetails1c4313___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetails1c4313___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetails1c4313___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TransactionDetails1c4313_NewTransactionDetails(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new TransactionDetails1c4313(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new TransactionDetails1c4313(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new TransactionDetails1c4313(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new TransactionDetails1c4313(static_cast<QWindow*>(parent));
	} else {
		return new TransactionDetails1c4313(static_cast<QObject*>(parent));
	}
}

void TransactionDetails1c4313_DestroyTransactionDetails(void* ptr)
{
	static_cast<TransactionDetails1c4313*>(ptr)->~TransactionDetails1c4313();
}

void TransactionDetails1c4313_DestroyTransactionDetailsDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void TransactionDetails1c4313_ChildEventDefault(void* ptr, void* event)
{
	static_cast<TransactionDetails1c4313*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void TransactionDetails1c4313_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TransactionDetails1c4313*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void TransactionDetails1c4313_CustomEventDefault(void* ptr, void* event)
{
	static_cast<TransactionDetails1c4313*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void TransactionDetails1c4313_DeleteLaterDefault(void* ptr)
{
	static_cast<TransactionDetails1c4313*>(ptr)->QObject::deleteLater();
}

void TransactionDetails1c4313_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TransactionDetails1c4313*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char TransactionDetails1c4313_EventDefault(void* ptr, void* e)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char TransactionDetails1c4313_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<TransactionDetails1c4313*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void TransactionDetails1c4313_TimerEventDefault(void* ptr, void* event)
{
	static_cast<TransactionDetails1c4313*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString Wallet1c4313_Name(void* ptr)
{
	return ({ QByteArray t24fd5b = static_cast<Wallet1c4313*>(ptr)->name().toUtf8(); Moc_PackedString { const_cast<char*>(t24fd5b.prepend("WHITESPACE").constData()+10), t24fd5b.size()-10 }; });
}

struct Moc_PackedString Wallet1c4313_NameDefault(void* ptr)
{
	return ({ QByteArray t1d5631 = static_cast<Wallet1c4313*>(ptr)->nameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t1d5631.prepend("WHITESPACE").constData()+10), t1d5631.size()-10 }; });
}

void Wallet1c4313_SetName(void* ptr, struct Moc_PackedString name)
{
	static_cast<Wallet1c4313*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void Wallet1c4313_SetNameDefault(void* ptr, struct Moc_PackedString name)
{
	static_cast<Wallet1c4313*>(ptr)->setNameDefault(QString::fromUtf8(name.data, name.len));
}

void Wallet1c4313_ConnectNameChanged(void* ptr)
{
	QObject::connect(static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(QString)>(&Wallet1c4313::nameChanged), static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(QString)>(&Wallet1c4313::Signal_NameChanged));
}

void Wallet1c4313_DisconnectNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(QString)>(&Wallet1c4313::nameChanged), static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(QString)>(&Wallet1c4313::Signal_NameChanged));
}

void Wallet1c4313_NameChanged(void* ptr, struct Moc_PackedString name)
{
	static_cast<Wallet1c4313*>(ptr)->nameChanged(QString::fromUtf8(name.data, name.len));
}

int Wallet1c4313_EncryptionEnabled(void* ptr)
{
	return static_cast<Wallet1c4313*>(ptr)->encryptionEnabled();
}

int Wallet1c4313_EncryptionEnabledDefault(void* ptr)
{
	return static_cast<Wallet1c4313*>(ptr)->encryptionEnabledDefault();
}

void Wallet1c4313_SetEncryptionEnabled(void* ptr, int encryptionEnabled)
{
	static_cast<Wallet1c4313*>(ptr)->setEncryptionEnabled(encryptionEnabled);
}

void Wallet1c4313_SetEncryptionEnabledDefault(void* ptr, int encryptionEnabled)
{
	static_cast<Wallet1c4313*>(ptr)->setEncryptionEnabledDefault(encryptionEnabled);
}

void Wallet1c4313_ConnectEncryptionEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(qint32)>(&Wallet1c4313::encryptionEnabledChanged), static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(qint32)>(&Wallet1c4313::Signal_EncryptionEnabledChanged));
}

void Wallet1c4313_DisconnectEncryptionEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(qint32)>(&Wallet1c4313::encryptionEnabledChanged), static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(qint32)>(&Wallet1c4313::Signal_EncryptionEnabledChanged));
}

void Wallet1c4313_EncryptionEnabledChanged(void* ptr, int encryptionEnabled)
{
	static_cast<Wallet1c4313*>(ptr)->encryptionEnabledChanged(encryptionEnabled);
}

int Wallet1c4313_Sky(void* ptr)
{
	return static_cast<Wallet1c4313*>(ptr)->sky();
}

int Wallet1c4313_SkyDefault(void* ptr)
{
	return static_cast<Wallet1c4313*>(ptr)->skyDefault();
}

void Wallet1c4313_SetSky(void* ptr, int sky)
{
	static_cast<Wallet1c4313*>(ptr)->setSky(sky);
}

void Wallet1c4313_SetSkyDefault(void* ptr, int sky)
{
	static_cast<Wallet1c4313*>(ptr)->setSkyDefault(sky);
}

void Wallet1c4313_ConnectSkyChanged(void* ptr)
{
	QObject::connect(static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(qint32)>(&Wallet1c4313::skyChanged), static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(qint32)>(&Wallet1c4313::Signal_SkyChanged));
}

void Wallet1c4313_DisconnectSkyChanged(void* ptr)
{
	QObject::disconnect(static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(qint32)>(&Wallet1c4313::skyChanged), static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(qint32)>(&Wallet1c4313::Signal_SkyChanged));
}

void Wallet1c4313_SkyChanged(void* ptr, int sky)
{
	static_cast<Wallet1c4313*>(ptr)->skyChanged(sky);
}

int Wallet1c4313_CoinHours(void* ptr)
{
	return static_cast<Wallet1c4313*>(ptr)->coinHours();
}

int Wallet1c4313_CoinHoursDefault(void* ptr)
{
	return static_cast<Wallet1c4313*>(ptr)->coinHoursDefault();
}

void Wallet1c4313_SetCoinHours(void* ptr, int coinHours)
{
	static_cast<Wallet1c4313*>(ptr)->setCoinHours(coinHours);
}

void Wallet1c4313_SetCoinHoursDefault(void* ptr, int coinHours)
{
	static_cast<Wallet1c4313*>(ptr)->setCoinHoursDefault(coinHours);
}

void Wallet1c4313_ConnectCoinHoursChanged(void* ptr)
{
	QObject::connect(static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(qint32)>(&Wallet1c4313::coinHoursChanged), static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(qint32)>(&Wallet1c4313::Signal_CoinHoursChanged));
}

void Wallet1c4313_DisconnectCoinHoursChanged(void* ptr)
{
	QObject::disconnect(static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(qint32)>(&Wallet1c4313::coinHoursChanged), static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(qint32)>(&Wallet1c4313::Signal_CoinHoursChanged));
}

void Wallet1c4313_CoinHoursChanged(void* ptr, int coinHours)
{
	static_cast<Wallet1c4313*>(ptr)->coinHoursChanged(coinHours);
}

struct Moc_PackedString Wallet1c4313_FileName(void* ptr)
{
	return ({ QByteArray t9abfde = static_cast<Wallet1c4313*>(ptr)->fileName().toUtf8(); Moc_PackedString { const_cast<char*>(t9abfde.prepend("WHITESPACE").constData()+10), t9abfde.size()-10 }; });
}

struct Moc_PackedString Wallet1c4313_FileNameDefault(void* ptr)
{
	return ({ QByteArray t53e8d1 = static_cast<Wallet1c4313*>(ptr)->fileNameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t53e8d1.prepend("WHITESPACE").constData()+10), t53e8d1.size()-10 }; });
}

void Wallet1c4313_SetFileName(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<Wallet1c4313*>(ptr)->setFileName(QString::fromUtf8(fileName.data, fileName.len));
}

void Wallet1c4313_SetFileNameDefault(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<Wallet1c4313*>(ptr)->setFileNameDefault(QString::fromUtf8(fileName.data, fileName.len));
}

void Wallet1c4313_ConnectFileNameChanged(void* ptr)
{
	QObject::connect(static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(QString)>(&Wallet1c4313::fileNameChanged), static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(QString)>(&Wallet1c4313::Signal_FileNameChanged));
}

void Wallet1c4313_DisconnectFileNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(QString)>(&Wallet1c4313::fileNameChanged), static_cast<Wallet1c4313*>(ptr), static_cast<void (Wallet1c4313::*)(QString)>(&Wallet1c4313::Signal_FileNameChanged));
}

void Wallet1c4313_FileNameChanged(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<Wallet1c4313*>(ptr)->fileNameChanged(QString::fromUtf8(fileName.data, fileName.len));
}

int Wallet1c4313_Wallet1c4313_QRegisterMetaType()
{
	return qRegisterMetaType<Wallet1c4313*>();
}

int Wallet1c4313_Wallet1c4313_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Wallet1c4313*>(const_cast<const char*>(typeName));
}

int Wallet1c4313_Wallet1c4313_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Wallet1c4313>();
#else
	return 0;
#endif
}

int Wallet1c4313_Wallet1c4313_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Wallet1c4313>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* Wallet1c4313___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Wallet1c4313___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Wallet1c4313___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* Wallet1c4313___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void Wallet1c4313___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Wallet1c4313___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* Wallet1c4313___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Wallet1c4313___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Wallet1c4313___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Wallet1c4313___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Wallet1c4313___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Wallet1c4313___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Wallet1c4313___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Wallet1c4313___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Wallet1c4313___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Wallet1c4313_NewWallet(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Wallet1c4313(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Wallet1c4313(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Wallet1c4313(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Wallet1c4313(static_cast<QWindow*>(parent));
	} else {
		return new Wallet1c4313(static_cast<QObject*>(parent));
	}
}

void Wallet1c4313_DestroyWallet(void* ptr)
{
	static_cast<Wallet1c4313*>(ptr)->~Wallet1c4313();
}

void Wallet1c4313_DestroyWalletDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void Wallet1c4313_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Wallet1c4313*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void Wallet1c4313_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Wallet1c4313*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Wallet1c4313_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Wallet1c4313*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void Wallet1c4313_DeleteLaterDefault(void* ptr)
{
	static_cast<Wallet1c4313*>(ptr)->QObject::deleteLater();
}

void Wallet1c4313_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Wallet1c4313*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char Wallet1c4313_EventDefault(void* ptr, void* e)
{
	return static_cast<Wallet1c4313*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char Wallet1c4313_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Wallet1c4313*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void Wallet1c4313_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Wallet1c4313*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void WalletModel1c4313_AddWallet(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<WalletModel1c4313*>(ptr), "addWallet", Q_ARG(Wallet1c4313*, static_cast<Wallet1c4313*>(v0)));
}

void WalletModel1c4313_EditWallet(void* ptr, int row, struct Moc_PackedString name, char encryptionEnabled, int sky, int coinHours)
{
	QMetaObject::invokeMethod(static_cast<WalletModel1c4313*>(ptr), "editWallet", Q_ARG(qint32, row), Q_ARG(QString, QString::fromUtf8(name.data, name.len)), Q_ARG(bool, encryptionEnabled != 0), Q_ARG(qint32, sky), Q_ARG(qint32, coinHours));
}

void WalletModel1c4313_RemoveWallet(void* ptr, int row)
{
	QMetaObject::invokeMethod(static_cast<WalletModel1c4313*>(ptr), "removeWallet", Q_ARG(qint32, row));
}

void WalletModel1c4313_LoadModel(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<WalletModel1c4313*>(ptr), "loadModel");
}

struct Moc_PackedList WalletModel1c4313_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<WalletModel1c4313*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModel1c4313_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<WalletModel1c4313*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void WalletModel1c4313_SetRoles(void* ptr, void* roles)
{
	static_cast<WalletModel1c4313*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void WalletModel1c4313_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<WalletModel1c4313*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void WalletModel1c4313_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<WalletModel1c4313*>(ptr), static_cast<void (WalletModel1c4313::*)(QMap<qint32, QByteArray>)>(&WalletModel1c4313::rolesChanged), static_cast<WalletModel1c4313*>(ptr), static_cast<void (WalletModel1c4313::*)(QMap<qint32, QByteArray>)>(&WalletModel1c4313::Signal_RolesChanged));
}

void WalletModel1c4313_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<WalletModel1c4313*>(ptr), static_cast<void (WalletModel1c4313::*)(QMap<qint32, QByteArray>)>(&WalletModel1c4313::rolesChanged), static_cast<WalletModel1c4313*>(ptr), static_cast<void (WalletModel1c4313::*)(QMap<qint32, QByteArray>)>(&WalletModel1c4313::Signal_RolesChanged));
}

void WalletModel1c4313_RolesChanged(void* ptr, void* roles)
{
	static_cast<WalletModel1c4313*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList WalletModel1c4313_Wallets(void* ptr)
{
	return ({ QList<Wallet1c4313*>* tmpValue = new QList<Wallet1c4313*>(static_cast<WalletModel1c4313*>(ptr)->wallets()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModel1c4313_WalletsDefault(void* ptr)
{
	return ({ QList<Wallet1c4313*>* tmpValue = new QList<Wallet1c4313*>(static_cast<WalletModel1c4313*>(ptr)->walletsDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void WalletModel1c4313_SetWallets(void* ptr, void* wallets)
{
	static_cast<WalletModel1c4313*>(ptr)->setWallets(({ QList<Wallet1c4313*>* tmpP = static_cast<QList<Wallet1c4313*>*>(wallets); QList<Wallet1c4313*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void WalletModel1c4313_SetWalletsDefault(void* ptr, void* wallets)
{
	static_cast<WalletModel1c4313*>(ptr)->setWalletsDefault(({ QList<Wallet1c4313*>* tmpP = static_cast<QList<Wallet1c4313*>*>(wallets); QList<Wallet1c4313*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void WalletModel1c4313_ConnectWalletsChanged(void* ptr)
{
	QObject::connect(static_cast<WalletModel1c4313*>(ptr), static_cast<void (WalletModel1c4313::*)(QList<Wallet1c4313*>)>(&WalletModel1c4313::walletsChanged), static_cast<WalletModel1c4313*>(ptr), static_cast<void (WalletModel1c4313::*)(QList<Wallet1c4313*>)>(&WalletModel1c4313::Signal_WalletsChanged));
}

void WalletModel1c4313_DisconnectWalletsChanged(void* ptr)
{
	QObject::disconnect(static_cast<WalletModel1c4313*>(ptr), static_cast<void (WalletModel1c4313::*)(QList<Wallet1c4313*>)>(&WalletModel1c4313::walletsChanged), static_cast<WalletModel1c4313*>(ptr), static_cast<void (WalletModel1c4313::*)(QList<Wallet1c4313*>)>(&WalletModel1c4313::Signal_WalletsChanged));
}

void WalletModel1c4313_WalletsChanged(void* ptr, void* wallets)
{
	static_cast<WalletModel1c4313*>(ptr)->walletsChanged(({ QList<Wallet1c4313*>* tmpP = static_cast<QList<Wallet1c4313*>*>(wallets); QList<Wallet1c4313*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int WalletModel1c4313_WalletModel1c4313_QRegisterMetaType()
{
	return qRegisterMetaType<WalletModel1c4313*>();
}

int WalletModel1c4313_WalletModel1c4313_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<WalletModel1c4313*>(const_cast<const char*>(typeName));
}

int WalletModel1c4313_WalletModel1c4313_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletModel1c4313>();
#else
	return 0;
#endif
}

int WalletModel1c4313_WalletModel1c4313_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletModel1c4313>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int WalletModel1c4313_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel1c4313_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel1c4313_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModel1c4313_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel1c4313_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel1c4313_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModel1c4313_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel1c4313_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel1c4313_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* WalletModel1c4313___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel1c4313___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel1c4313___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel1c4313___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel1c4313___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel1c4313___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int WalletModel1c4313___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void WalletModel1c4313___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* WalletModel1c4313___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* WalletModel1c4313___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel1c4313___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* WalletModel1c4313___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList WalletModel1c4313___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel1c4313___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel1c4313___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* WalletModel1c4313___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* WalletModel1c4313___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel1c4313___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* WalletModel1c4313___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* WalletModel1c4313___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel1c4313___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel1c4313___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel1c4313___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel1c4313___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel1c4313___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel1c4313___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel1c4313___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel1c4313___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel1c4313___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void WalletModel1c4313___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel1c4313___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList WalletModel1c4313___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel1c4313___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel1c4313___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* WalletModel1c4313___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList WalletModel1c4313___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int WalletModel1c4313_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel1c4313_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel1c4313_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModel1c4313_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel1c4313_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel1c4313_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* WalletModel1c4313___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel1c4313___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel1c4313___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* WalletModel1c4313___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel1c4313___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* WalletModel1c4313___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* WalletModel1c4313___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel1c4313___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel1c4313___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModel1c4313___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel1c4313___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel1c4313___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModel1c4313___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel1c4313___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel1c4313___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModel1c4313___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel1c4313___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel1c4313___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModel1c4313___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel1c4313___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel1c4313___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel1c4313___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModel1c4313___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel1c4313___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel1c4313___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel1c4313___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModel1c4313___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel1c4313___wallets_atList(void* ptr, int i)
{
	return ({Wallet1c4313* tmp = static_cast<QList<Wallet1c4313*>*>(ptr)->at(i); if (i == static_cast<QList<Wallet1c4313*>*>(ptr)->size()-1) { static_cast<QList<Wallet1c4313*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel1c4313___wallets_setList(void* ptr, void* i)
{
	static_cast<QList<Wallet1c4313*>*>(ptr)->append(static_cast<Wallet1c4313*>(i));
}

void* WalletModel1c4313___wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Wallet1c4313*>();
}

void* WalletModel1c4313___setWallets_wallets_atList(void* ptr, int i)
{
	return ({Wallet1c4313* tmp = static_cast<QList<Wallet1c4313*>*>(ptr)->at(i); if (i == static_cast<QList<Wallet1c4313*>*>(ptr)->size()-1) { static_cast<QList<Wallet1c4313*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel1c4313___setWallets_wallets_setList(void* ptr, void* i)
{
	static_cast<QList<Wallet1c4313*>*>(ptr)->append(static_cast<Wallet1c4313*>(i));
}

void* WalletModel1c4313___setWallets_wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Wallet1c4313*>();
}

void* WalletModel1c4313___walletsChanged_wallets_atList(void* ptr, int i)
{
	return ({Wallet1c4313* tmp = static_cast<QList<Wallet1c4313*>*>(ptr)->at(i); if (i == static_cast<QList<Wallet1c4313*>*>(ptr)->size()-1) { static_cast<QList<Wallet1c4313*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel1c4313___walletsChanged_wallets_setList(void* ptr, void* i)
{
	static_cast<QList<Wallet1c4313*>*>(ptr)->append(static_cast<Wallet1c4313*>(i));
}

void* WalletModel1c4313___walletsChanged_wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Wallet1c4313*>();
}

int WalletModel1c4313_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel1c4313_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModel1c4313_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int WalletModel1c4313_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel1c4313_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModel1c4313_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int WalletModel1c4313_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel1c4313_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModel1c4313_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* WalletModel1c4313_NewWalletModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new WalletModel1c4313(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new WalletModel1c4313(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new WalletModel1c4313(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new WalletModel1c4313(static_cast<QWindow*>(parent));
	} else {
		return new WalletModel1c4313(static_cast<QObject*>(parent));
	}
}

void WalletModel1c4313_DestroyWalletModel(void* ptr)
{
	static_cast<WalletModel1c4313*>(ptr)->~WalletModel1c4313();
}

void WalletModel1c4313_DestroyWalletModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char WalletModel1c4313_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long WalletModel1c4313_FlagsDefault(void* ptr, void* index)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* WalletModel1c4313_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* WalletModel1c4313_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* WalletModel1c4313_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char WalletModel1c4313_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char WalletModel1c4313_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int WalletModel1c4313_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* WalletModel1c4313_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void WalletModel1c4313_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char WalletModel1c4313_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* WalletModel1c4313_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char WalletModel1c4313_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char WalletModel1c4313_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList WalletModel1c4313_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModel1c4313_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel1c4313_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString WalletModel1c4313_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t2c50d2 = static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t2c50d2.prepend("WHITESPACE").constData()+10), t2c50d2.size()-10 }; });
}

char WalletModel1c4313_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char WalletModel1c4313_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* WalletModel1c4313_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char WalletModel1c4313_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char WalletModel1c4313_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void WalletModel1c4313_ResetInternalDataDefault(void* ptr)
{
	static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::resetInternalData();
}

void WalletModel1c4313_RevertDefault(void* ptr)
{
	static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList WalletModel1c4313_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int WalletModel1c4313_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char WalletModel1c4313_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char WalletModel1c4313_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char WalletModel1c4313_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void WalletModel1c4313_SortDefault(void* ptr, int column, long long order)
{
	static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* WalletModel1c4313_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char WalletModel1c4313_SubmitDefault(void* ptr)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::submit();
}

long long WalletModel1c4313_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long WalletModel1c4313_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::supportedDropActions();
}

void WalletModel1c4313_ChildEventDefault(void* ptr, void* event)
{
	static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void WalletModel1c4313_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void WalletModel1c4313_CustomEventDefault(void* ptr, void* event)
{
	static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void WalletModel1c4313_DeleteLaterDefault(void* ptr)
{
	static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::deleteLater();
}

void WalletModel1c4313_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char WalletModel1c4313_EventDefault(void* ptr, void* e)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char WalletModel1c4313_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void WalletModel1c4313_TimerEventDefault(void* ptr, void* event)
{
	static_cast<WalletModel1c4313*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
