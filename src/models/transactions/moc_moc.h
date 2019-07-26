/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.13.0)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <memory>
#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#include <QtCore/QList>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'moc.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.13.0. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_WalletModel1c4313_t {
    QByteArrayData data[19];
    char stringdata0[195];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_WalletModel1c4313_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_WalletModel1c4313_t qt_meta_stringdata_WalletModel1c4313 = {
    {
QT_MOC_LITERAL(0, 0, 17), // "WalletModel1c4313"
QT_MOC_LITERAL(1, 18, 12), // "rolesChanged"
QT_MOC_LITERAL(2, 31, 0), // ""
QT_MOC_LITERAL(3, 32, 10), // "type378cdd"
QT_MOC_LITERAL(4, 43, 5), // "roles"
QT_MOC_LITERAL(5, 49, 14), // "walletsChanged"
QT_MOC_LITERAL(6, 64, 20), // "QList<Wallet1c4313*>"
QT_MOC_LITERAL(7, 85, 7), // "wallets"
QT_MOC_LITERAL(8, 93, 9), // "addWallet"
QT_MOC_LITERAL(9, 103, 13), // "Wallet1c4313*"
QT_MOC_LITERAL(10, 117, 2), // "v0"
QT_MOC_LITERAL(11, 120, 10), // "editWallet"
QT_MOC_LITERAL(12, 131, 3), // "row"
QT_MOC_LITERAL(13, 135, 4), // "name"
QT_MOC_LITERAL(14, 140, 17), // "encryptionEnabled"
QT_MOC_LITERAL(15, 158, 3), // "sky"
QT_MOC_LITERAL(16, 162, 9), // "coinHours"
QT_MOC_LITERAL(17, 172, 12), // "removeWallet"
QT_MOC_LITERAL(18, 185, 9) // "loadModel"

    },
    "WalletModel1c4313\0rolesChanged\0\0"
    "type378cdd\0roles\0walletsChanged\0"
    "QList<Wallet1c4313*>\0wallets\0addWallet\0"
    "Wallet1c4313*\0v0\0editWallet\0row\0name\0"
    "encryptionEnabled\0sky\0coinHours\0"
    "removeWallet\0loadModel"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_WalletModel1c4313[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       6,   14, // methods
       2,   68, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       2,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   44,    2, 0x06 /* Public */,
       5,    1,   47,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
       8,    1,   50,    2, 0x0a /* Public */,
      11,    5,   53,    2, 0x0a /* Public */,
      17,    1,   64,    2, 0x0a /* Public */,
      18,    0,   67,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, 0x80000000 | 3,    4,
    QMetaType::Void, 0x80000000 | 6,    7,

 // slots: parameters
    QMetaType::Void, 0x80000000 | 9,   10,
    QMetaType::Void, QMetaType::Int, QMetaType::QString, QMetaType::Bool, QMetaType::Int, QMetaType::Int,   12,   13,   14,   15,   16,
    QMetaType::Void, QMetaType::Int,   12,
    QMetaType::Void,

 // properties: name, type, flags
       4, 0x80000000 | 3, 0x0049510b,
       7, 0x80000000 | 6, 0x0049510b,

 // properties: notify_signal_id
       0,
       1,

       0        // eod
};

void WalletModel1c4313::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<WalletModel1c4313 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->rolesChanged((*reinterpret_cast< type378cdd(*)>(_a[1]))); break;
        case 1: _t->walletsChanged((*reinterpret_cast< QList<Wallet1c4313*>(*)>(_a[1]))); break;
        case 2: _t->addWallet((*reinterpret_cast< Wallet1c4313*(*)>(_a[1]))); break;
        case 3: _t->editWallet((*reinterpret_cast< qint32(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< bool(*)>(_a[3])),(*reinterpret_cast< qint32(*)>(_a[4])),(*reinterpret_cast< qint32(*)>(_a[5]))); break;
        case 4: _t->removeWallet((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 5: _t->loadModel(); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<Wallet1c4313*> >(); break;
            }
            break;
        case 2:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< Wallet1c4313* >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (WalletModel1c4313::*)(type378cdd );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&WalletModel1c4313::rolesChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (WalletModel1c4313::*)(QList<Wallet1c4313*> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&WalletModel1c4313::walletsChanged)) {
                *result = 1;
                return;
            }
        }
    } else if (_c == QMetaObject::RegisterPropertyMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<Wallet1c4313*> >(); break;
        }
    }

#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<WalletModel1c4313 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< type378cdd*>(_v) = _t->roles(); break;
        case 1: *reinterpret_cast< QList<Wallet1c4313*>*>(_v) = _t->wallets(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<WalletModel1c4313 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRoles(*reinterpret_cast< type378cdd*>(_v)); break;
        case 1: _t->setWallets(*reinterpret_cast< QList<Wallet1c4313*>*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject WalletModel1c4313::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_WalletModel1c4313.data,
    qt_meta_data_WalletModel1c4313,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *WalletModel1c4313::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *WalletModel1c4313::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_WalletModel1c4313.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int WalletModel1c4313::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QAbstractListModel::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 6)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 6;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 6)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 6;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 2;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void WalletModel1c4313::rolesChanged(type378cdd _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void WalletModel1c4313::walletsChanged(QList<Wallet1c4313*> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}
struct qt_meta_stringdata_AddressDetails1c4313_t {
    QByteArrayData data[8];
    char stringdata0[115];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_AddressDetails1c4313_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_AddressDetails1c4313_t qt_meta_stringdata_AddressDetails1c4313 = {
    {
QT_MOC_LITERAL(0, 0, 20), // "AddressDetails1c4313"
QT_MOC_LITERAL(1, 21, 14), // "addressChanged"
QT_MOC_LITERAL(2, 36, 0), // ""
QT_MOC_LITERAL(3, 37, 7), // "address"
QT_MOC_LITERAL(4, 45, 17), // "addressSkyChanged"
QT_MOC_LITERAL(5, 63, 10), // "addressSky"
QT_MOC_LITERAL(6, 74, 23), // "addressCoinHoursChanged"
QT_MOC_LITERAL(7, 98, 16) // "addressCoinHours"

    },
    "AddressDetails1c4313\0addressChanged\0"
    "\0address\0addressSkyChanged\0addressSky\0"
    "addressCoinHoursChanged\0addressCoinHours"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_AddressDetails1c4313[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       3,   14, // methods
       3,   38, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       3,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   29,    2, 0x06 /* Public */,
       4,    1,   32,    2, 0x06 /* Public */,
       6,    1,   35,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::Float,    5,
    QMetaType::Void, QMetaType::Int,    7,

 // properties: name, type, flags
       3, QMetaType::QString, 0x00495103,
       5, QMetaType::Float, 0x00495103,
       7, QMetaType::Int, 0x00495103,

 // properties: notify_signal_id
       0,
       1,
       2,

       0        // eod
};

void AddressDetails1c4313::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<AddressDetails1c4313 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->addressChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->addressSkyChanged((*reinterpret_cast< float(*)>(_a[1]))); break;
        case 2: _t->addressCoinHoursChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (AddressDetails1c4313::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressDetails1c4313::addressChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (AddressDetails1c4313::*)(float );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressDetails1c4313::addressSkyChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (AddressDetails1c4313::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressDetails1c4313::addressCoinHoursChanged)) {
                *result = 2;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<AddressDetails1c4313 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->address(); break;
        case 1: *reinterpret_cast< float*>(_v) = _t->addressSky(); break;
        case 2: *reinterpret_cast< qint32*>(_v) = _t->addressCoinHours(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<AddressDetails1c4313 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setAddress(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setAddressSky(*reinterpret_cast< float*>(_v)); break;
        case 2: _t->setAddressCoinHours(*reinterpret_cast< qint32*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject AddressDetails1c4313::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_AddressDetails1c4313.data,
    qt_meta_data_AddressDetails1c4313,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *AddressDetails1c4313::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *AddressDetails1c4313::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_AddressDetails1c4313.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int AddressDetails1c4313::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 3)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 3;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 3)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 3;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 3;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void AddressDetails1c4313::addressChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void AddressDetails1c4313::addressSkyChanged(float _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void AddressDetails1c4313::addressCoinHoursChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
struct qt_meta_stringdata_AddressList1c4313_t {
    QByteArrayData data[13];
    char stringdata0[170];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_AddressList1c4313_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_AddressList1c4313_t qt_meta_stringdata_AddressList1c4313 = {
    {
QT_MOC_LITERAL(0, 0, 17), // "AddressList1c4313"
QT_MOC_LITERAL(1, 18, 10), // "addAddress"
QT_MOC_LITERAL(2, 29, 0), // ""
QT_MOC_LITERAL(3, 30, 21), // "AddressDetails1c4313*"
QT_MOC_LITERAL(4, 52, 11), // "transaction"
QT_MOC_LITERAL(5, 64, 13), // "removeAddress"
QT_MOC_LITERAL(6, 78, 5), // "index"
QT_MOC_LITERAL(7, 84, 12), // "rolesChanged"
QT_MOC_LITERAL(8, 97, 10), // "type378cdd"
QT_MOC_LITERAL(9, 108, 5), // "roles"
QT_MOC_LITERAL(10, 114, 16), // "addressesChanged"
QT_MOC_LITERAL(11, 131, 28), // "QList<AddressDetails1c4313*>"
QT_MOC_LITERAL(12, 160, 9) // "addresses"

    },
    "AddressList1c4313\0addAddress\0\0"
    "AddressDetails1c4313*\0transaction\0"
    "removeAddress\0index\0rolesChanged\0"
    "type378cdd\0roles\0addressesChanged\0"
    "QList<AddressDetails1c4313*>\0addresses"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_AddressList1c4313[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       4,   14, // methods
       2,   46, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       4,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   34,    2, 0x06 /* Public */,
       5,    1,   37,    2, 0x06 /* Public */,
       7,    1,   40,    2, 0x06 /* Public */,
      10,    1,   43,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, 0x80000000 | 3,    4,
    QMetaType::Void, QMetaType::Int,    6,
    QMetaType::Void, 0x80000000 | 8,    9,
    QMetaType::Void, 0x80000000 | 11,   12,

 // properties: name, type, flags
       9, 0x80000000 | 8, 0x0049510b,
      12, 0x80000000 | 11, 0x0049510b,

 // properties: notify_signal_id
       2,
       3,

       0        // eod
};

void AddressList1c4313::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<AddressList1c4313 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->addAddress((*reinterpret_cast< AddressDetails1c4313*(*)>(_a[1]))); break;
        case 1: _t->removeAddress((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 2: _t->rolesChanged((*reinterpret_cast< type378cdd(*)>(_a[1]))); break;
        case 3: _t->addressesChanged((*reinterpret_cast< QList<AddressDetails1c4313*>(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 0:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< AddressDetails1c4313* >(); break;
            }
            break;
        case 3:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<AddressDetails1c4313*> >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (AddressList1c4313::*)(AddressDetails1c4313 * );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressList1c4313::addAddress)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (AddressList1c4313::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressList1c4313::removeAddress)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (AddressList1c4313::*)(type378cdd );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressList1c4313::rolesChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (AddressList1c4313::*)(QList<AddressDetails1c4313*> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressList1c4313::addressesChanged)) {
                *result = 3;
                return;
            }
        }
    } else if (_c == QMetaObject::RegisterPropertyMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<AddressDetails1c4313*> >(); break;
        }
    }

#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<AddressList1c4313 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< type378cdd*>(_v) = _t->roles(); break;
        case 1: *reinterpret_cast< QList<AddressDetails1c4313*>*>(_v) = _t->addresses(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<AddressList1c4313 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRoles(*reinterpret_cast< type378cdd*>(_v)); break;
        case 1: _t->setAddresses(*reinterpret_cast< QList<AddressDetails1c4313*>*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject AddressList1c4313::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_AddressList1c4313.data,
    qt_meta_data_AddressList1c4313,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *AddressList1c4313::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *AddressList1c4313::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_AddressList1c4313.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int AddressList1c4313::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QAbstractListModel::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 4)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 4;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 4)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 4;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 2;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void AddressList1c4313::addAddress(AddressDetails1c4313 * _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void AddressList1c4313::removeAddress(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void AddressList1c4313::rolesChanged(type378cdd _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void AddressList1c4313::addressesChanged(QList<AddressDetails1c4313*> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}
struct qt_meta_stringdata_TransactionDetails1c4313_t {
    QByteArrayData data[26];
    char stringdata0[339];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_TransactionDetails1c4313_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_TransactionDetails1c4313_t qt_meta_stringdata_TransactionDetails1c4313 = {
    {
QT_MOC_LITERAL(0, 0, 24), // "TransactionDetails1c4313"
QT_MOC_LITERAL(1, 25, 11), // "dateChanged"
QT_MOC_LITERAL(2, 37, 0), // ""
QT_MOC_LITERAL(3, 38, 4), // "date"
QT_MOC_LITERAL(4, 43, 13), // "statusChanged"
QT_MOC_LITERAL(5, 57, 6), // "status"
QT_MOC_LITERAL(6, 64, 11), // "typeChanged"
QT_MOC_LITERAL(7, 76, 2), // "ty"
QT_MOC_LITERAL(8, 79, 13), // "amountChanged"
QT_MOC_LITERAL(9, 93, 6), // "amount"
QT_MOC_LITERAL(10, 100, 20), // "hoursReceivedChanged"
QT_MOC_LITERAL(11, 121, 13), // "hoursReceived"
QT_MOC_LITERAL(12, 135, 18), // "hoursBurnedChanged"
QT_MOC_LITERAL(13, 154, 11), // "hoursBurned"
QT_MOC_LITERAL(14, 166, 20), // "transactionIDChanged"
QT_MOC_LITERAL(15, 187, 13), // "transactionID"
QT_MOC_LITERAL(16, 201, 18), // "sentAddressChanged"
QT_MOC_LITERAL(17, 220, 11), // "sentAddress"
QT_MOC_LITERAL(18, 232, 22), // "receivedAddressChanged"
QT_MOC_LITERAL(19, 255, 15), // "receivedAddress"
QT_MOC_LITERAL(20, 271, 13), // "inputsChanged"
QT_MOC_LITERAL(21, 285, 18), // "AddressList1c4313*"
QT_MOC_LITERAL(22, 304, 6), // "inputs"
QT_MOC_LITERAL(23, 311, 14), // "outputsChanged"
QT_MOC_LITERAL(24, 326, 7), // "outputs"
QT_MOC_LITERAL(25, 334, 4) // "type"

    },
    "TransactionDetails1c4313\0dateChanged\0"
    "\0date\0statusChanged\0status\0typeChanged\0"
    "ty\0amountChanged\0amount\0hoursReceivedChanged\0"
    "hoursReceived\0hoursBurnedChanged\0"
    "hoursBurned\0transactionIDChanged\0"
    "transactionID\0sentAddressChanged\0"
    "sentAddress\0receivedAddressChanged\0"
    "receivedAddress\0inputsChanged\0"
    "AddressList1c4313*\0inputs\0outputsChanged\0"
    "outputs\0type"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_TransactionDetails1c4313[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
      11,   14, // methods
      11,  102, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
      11,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   69,    2, 0x06 /* Public */,
       4,    1,   72,    2, 0x06 /* Public */,
       6,    1,   75,    2, 0x06 /* Public */,
       8,    1,   78,    2, 0x06 /* Public */,
      10,    1,   81,    2, 0x06 /* Public */,
      12,    1,   84,    2, 0x06 /* Public */,
      14,    1,   87,    2, 0x06 /* Public */,
      16,    1,   90,    2, 0x06 /* Public */,
      18,    1,   93,    2, 0x06 /* Public */,
      20,    1,   96,    2, 0x06 /* Public */,
      23,    1,   99,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QDateTime,    3,
    QMetaType::Void, QMetaType::Int,    5,
    QMetaType::Void, QMetaType::Int,    7,
    QMetaType::Void, QMetaType::Int,    9,
    QMetaType::Void, QMetaType::Int,   11,
    QMetaType::Void, QMetaType::Int,   13,
    QMetaType::Void, QMetaType::QString,   15,
    QMetaType::Void, QMetaType::QString,   17,
    QMetaType::Void, QMetaType::QString,   19,
    QMetaType::Void, 0x80000000 | 21,   22,
    QMetaType::Void, 0x80000000 | 21,   24,

 // properties: name, type, flags
       3, QMetaType::QDateTime, 0x00495103,
       5, QMetaType::Int, 0x00495103,
      25, QMetaType::Int, 0x00495103,
       9, QMetaType::Int, 0x00495103,
      11, QMetaType::Int, 0x00495103,
      13, QMetaType::Int, 0x00495103,
      15, QMetaType::QString, 0x00495103,
      17, QMetaType::QString, 0x00495103,
      19, QMetaType::QString, 0x00495103,
      22, 0x80000000 | 21, 0x0049510b,
      24, 0x80000000 | 21, 0x0049510b,

 // properties: notify_signal_id
       0,
       1,
       2,
       3,
       4,
       5,
       6,
       7,
       8,
       9,
      10,

       0        // eod
};

void TransactionDetails1c4313::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<TransactionDetails1c4313 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->dateChanged((*reinterpret_cast< QDateTime(*)>(_a[1]))); break;
        case 1: _t->statusChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 2: _t->typeChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 3: _t->amountChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 4: _t->hoursReceivedChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 5: _t->hoursBurnedChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 6: _t->transactionIDChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 7: _t->sentAddressChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 8: _t->receivedAddressChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 9: _t->inputsChanged((*reinterpret_cast< AddressList1c4313*(*)>(_a[1]))); break;
        case 10: _t->outputsChanged((*reinterpret_cast< AddressList1c4313*(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 9:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< AddressList1c4313* >(); break;
            }
            break;
        case 10:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< AddressList1c4313* >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (TransactionDetails1c4313::*)(QDateTime );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails1c4313::dateChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (TransactionDetails1c4313::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails1c4313::statusChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (TransactionDetails1c4313::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails1c4313::typeChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (TransactionDetails1c4313::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails1c4313::amountChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (TransactionDetails1c4313::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails1c4313::hoursReceivedChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (TransactionDetails1c4313::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails1c4313::hoursBurnedChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (TransactionDetails1c4313::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails1c4313::transactionIDChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (TransactionDetails1c4313::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails1c4313::sentAddressChanged)) {
                *result = 7;
                return;
            }
        }
        {
            using _t = void (TransactionDetails1c4313::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails1c4313::receivedAddressChanged)) {
                *result = 8;
                return;
            }
        }
        {
            using _t = void (TransactionDetails1c4313::*)(AddressList1c4313 * );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails1c4313::inputsChanged)) {
                *result = 9;
                return;
            }
        }
        {
            using _t = void (TransactionDetails1c4313::*)(AddressList1c4313 * );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails1c4313::outputsChanged)) {
                *result = 10;
                return;
            }
        }
    } else if (_c == QMetaObject::RegisterPropertyMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 10:
        case 9:
            *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< AddressList1c4313* >(); break;
        }
    }

#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<TransactionDetails1c4313 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QDateTime*>(_v) = _t->date(); break;
        case 1: *reinterpret_cast< qint32*>(_v) = _t->status(); break;
        case 2: *reinterpret_cast< qint32*>(_v) = _t->type(); break;
        case 3: *reinterpret_cast< qint32*>(_v) = _t->amount(); break;
        case 4: *reinterpret_cast< qint32*>(_v) = _t->hoursReceived(); break;
        case 5: *reinterpret_cast< qint32*>(_v) = _t->hoursBurned(); break;
        case 6: *reinterpret_cast< QString*>(_v) = _t->transactionID(); break;
        case 7: *reinterpret_cast< QString*>(_v) = _t->sentAddress(); break;
        case 8: *reinterpret_cast< QString*>(_v) = _t->receivedAddress(); break;
        case 9: *reinterpret_cast< AddressList1c4313**>(_v) = _t->inputs(); break;
        case 10: *reinterpret_cast< AddressList1c4313**>(_v) = _t->outputs(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<TransactionDetails1c4313 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setDate(*reinterpret_cast< QDateTime*>(_v)); break;
        case 1: _t->setStatus(*reinterpret_cast< qint32*>(_v)); break;
        case 2: _t->setType(*reinterpret_cast< qint32*>(_v)); break;
        case 3: _t->setAmount(*reinterpret_cast< qint32*>(_v)); break;
        case 4: _t->setHoursReceived(*reinterpret_cast< qint32*>(_v)); break;
        case 5: _t->setHoursBurned(*reinterpret_cast< qint32*>(_v)); break;
        case 6: _t->setTransactionID(*reinterpret_cast< QString*>(_v)); break;
        case 7: _t->setSentAddress(*reinterpret_cast< QString*>(_v)); break;
        case 8: _t->setReceivedAddress(*reinterpret_cast< QString*>(_v)); break;
        case 9: _t->setInputs(*reinterpret_cast< AddressList1c4313**>(_v)); break;
        case 10: _t->setOutputs(*reinterpret_cast< AddressList1c4313**>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject TransactionDetails1c4313::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_TransactionDetails1c4313.data,
    qt_meta_data_TransactionDetails1c4313,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *TransactionDetails1c4313::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *TransactionDetails1c4313::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_TransactionDetails1c4313.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int TransactionDetails1c4313::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 11)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 11;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 11)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 11;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 11;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 11;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 11;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 11;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 11;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 11;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void TransactionDetails1c4313::dateChanged(QDateTime _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void TransactionDetails1c4313::statusChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void TransactionDetails1c4313::typeChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void TransactionDetails1c4313::amountChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void TransactionDetails1c4313::hoursReceivedChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void TransactionDetails1c4313::hoursBurnedChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void TransactionDetails1c4313::transactionIDChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void TransactionDetails1c4313::sentAddressChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}

// SIGNAL 8
void TransactionDetails1c4313::receivedAddressChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 8, _a);
}

// SIGNAL 9
void TransactionDetails1c4313::inputsChanged(AddressList1c4313 * _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 9, _a);
}

// SIGNAL 10
void TransactionDetails1c4313::outputsChanged(AddressList1c4313 * _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 10, _a);
}
struct qt_meta_stringdata_Wallet1c4313_t {
    QByteArrayData data[12];
    char stringdata0[141];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_Wallet1c4313_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_Wallet1c4313_t qt_meta_stringdata_Wallet1c4313 = {
    {
QT_MOC_LITERAL(0, 0, 12), // "Wallet1c4313"
QT_MOC_LITERAL(1, 13, 11), // "nameChanged"
QT_MOC_LITERAL(2, 25, 0), // ""
QT_MOC_LITERAL(3, 26, 4), // "name"
QT_MOC_LITERAL(4, 31, 24), // "encryptionEnabledChanged"
QT_MOC_LITERAL(5, 56, 17), // "encryptionEnabled"
QT_MOC_LITERAL(6, 74, 10), // "skyChanged"
QT_MOC_LITERAL(7, 85, 3), // "sky"
QT_MOC_LITERAL(8, 89, 16), // "coinHoursChanged"
QT_MOC_LITERAL(9, 106, 9), // "coinHours"
QT_MOC_LITERAL(10, 116, 15), // "fileNameChanged"
QT_MOC_LITERAL(11, 132, 8) // "fileName"

    },
    "Wallet1c4313\0nameChanged\0\0name\0"
    "encryptionEnabledChanged\0encryptionEnabled\0"
    "skyChanged\0sky\0coinHoursChanged\0"
    "coinHours\0fileNameChanged\0fileName"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_Wallet1c4313[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       5,   14, // methods
       5,   54, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       5,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   39,    2, 0x06 /* Public */,
       4,    1,   42,    2, 0x06 /* Public */,
       6,    1,   45,    2, 0x06 /* Public */,
       8,    1,   48,    2, 0x06 /* Public */,
      10,    1,   51,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::Int,    5,
    QMetaType::Void, QMetaType::Int,    7,
    QMetaType::Void, QMetaType::Int,    9,
    QMetaType::Void, QMetaType::QString,   11,

 // properties: name, type, flags
       3, QMetaType::QString, 0x00495103,
       5, QMetaType::Int, 0x00495103,
       7, QMetaType::Int, 0x00495103,
       9, QMetaType::Int, 0x00495103,
      11, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       0,
       1,
       2,
       3,
       4,

       0        // eod
};

void Wallet1c4313::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<Wallet1c4313 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->nameChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->encryptionEnabledChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 2: _t->skyChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 3: _t->coinHoursChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 4: _t->fileNameChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (Wallet1c4313::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Wallet1c4313::nameChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (Wallet1c4313::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Wallet1c4313::encryptionEnabledChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (Wallet1c4313::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Wallet1c4313::skyChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (Wallet1c4313::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Wallet1c4313::coinHoursChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (Wallet1c4313::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Wallet1c4313::fileNameChanged)) {
                *result = 4;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<Wallet1c4313 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->name(); break;
        case 1: *reinterpret_cast< qint32*>(_v) = _t->encryptionEnabled(); break;
        case 2: *reinterpret_cast< qint32*>(_v) = _t->sky(); break;
        case 3: *reinterpret_cast< qint32*>(_v) = _t->coinHours(); break;
        case 4: *reinterpret_cast< QString*>(_v) = _t->fileName(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<Wallet1c4313 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setName(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setEncryptionEnabled(*reinterpret_cast< qint32*>(_v)); break;
        case 2: _t->setSky(*reinterpret_cast< qint32*>(_v)); break;
        case 3: _t->setCoinHours(*reinterpret_cast< qint32*>(_v)); break;
        case 4: _t->setFileName(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject Wallet1c4313::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_Wallet1c4313.data,
    qt_meta_data_Wallet1c4313,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *Wallet1c4313::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *Wallet1c4313::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_Wallet1c4313.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int Wallet1c4313::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 5)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 5;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 5)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 5;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 5;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void Wallet1c4313::nameChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void Wallet1c4313::encryptionEnabledChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void Wallet1c4313::skyChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void Wallet1c4313::coinHoursChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void Wallet1c4313::fileNameChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
