package transactions

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type TransactionDetails_ITF interface {
	std_core.QObject_ITF
	TransactionDetails_PTR() *TransactionDetails
}

func (ptr *TransactionDetails) TransactionDetails_PTR() *TransactionDetails {
	return ptr
}

func (ptr *TransactionDetails) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *TransactionDetails) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromTransactionDetails(ptr TransactionDetails_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.TransactionDetails_PTR().Pointer()
	}
	return nil
}

func NewTransactionDetailsFromPointer(ptr unsafe.Pointer) (n *TransactionDetails) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(TransactionDetails)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *TransactionDetails:
			n = deduced

		case *std_core.QObject:
			n = &TransactionDetails{QObject: *deduced}

		default:
			n = new(TransactionDetails)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackTransactionDetails1c4313_Constructor
func callbackTransactionDetails1c4313_Constructor(ptr unsafe.Pointer) {
	this := NewTransactionDetailsFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackTransactionDetails1c4313_Date
func callbackTransactionDetails1c4313_Date(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "date"); signal != nil {
		return std_core.PointerFromQDateTime((*(*func() *std_core.QDateTime)(signal))())
	}

	return std_core.PointerFromQDateTime(NewTransactionDetailsFromPointer(ptr).DateDefault())
}

func (ptr *TransactionDetails) ConnectDate(f func() *std_core.QDateTime) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "date"); signal != nil {
			f := func() *std_core.QDateTime {
				(*(*func() *std_core.QDateTime)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "date", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "date", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectDate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "date")
	}
}

func (ptr *TransactionDetails) Date() *std_core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQDateTimeFromPointer(C.TransactionDetails1c4313_Date(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) DateDefault() *std_core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQDateTimeFromPointer(C.TransactionDetails1c4313_DateDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

//export callbackTransactionDetails1c4313_SetDate
func callbackTransactionDetails1c4313_SetDate(ptr unsafe.Pointer, date unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setDate"); signal != nil {
		(*(*func(*std_core.QDateTime))(signal))(std_core.NewQDateTimeFromPointer(date))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetDateDefault(std_core.NewQDateTimeFromPointer(date))
	}
}

func (ptr *TransactionDetails) ConnectSetDate(f func(date *std_core.QDateTime)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDate"); signal != nil {
			f := func(date *std_core.QDateTime) {
				(*(*func(*std_core.QDateTime))(signal))(date)
				f(date)
			}
			qt.ConnectSignal(ptr.Pointer(), "setDate", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDate", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetDate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDate")
	}
}

func (ptr *TransactionDetails) SetDate(date std_core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetDate(ptr.Pointer(), std_core.PointerFromQDateTime(date))
	}
}

func (ptr *TransactionDetails) SetDateDefault(date std_core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetDateDefault(ptr.Pointer(), std_core.PointerFromQDateTime(date))
	}
}

//export callbackTransactionDetails1c4313_DateChanged
func callbackTransactionDetails1c4313_DateChanged(ptr unsafe.Pointer, date unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dateChanged"); signal != nil {
		(*(*func(*std_core.QDateTime))(signal))(std_core.NewQDateTimeFromPointer(date))
	}

}

func (ptr *TransactionDetails) ConnectDateChanged(f func(date *std_core.QDateTime)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dateChanged") {
			C.TransactionDetails1c4313_ConnectDateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dateChanged"); signal != nil {
			f := func(date *std_core.QDateTime) {
				(*(*func(*std_core.QDateTime))(signal))(date)
				f(date)
			}
			qt.ConnectSignal(ptr.Pointer(), "dateChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dateChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectDateChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DisconnectDateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "dateChanged")
	}
}

func (ptr *TransactionDetails) DateChanged(date std_core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DateChanged(ptr.Pointer(), std_core.PointerFromQDateTime(date))
	}
}

//export callbackTransactionDetails1c4313_Status
func callbackTransactionDetails1c4313_Status(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "status"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewTransactionDetailsFromPointer(ptr).StatusDefault()))
}

func (ptr *TransactionDetails) ConnectStatus(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "status"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "status", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "status", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectStatus() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "status")
	}
}

func (ptr *TransactionDetails) Status() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails1c4313_Status(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) StatusDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails1c4313_StatusDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetails1c4313_SetStatus
func callbackTransactionDetails1c4313_SetStatus(ptr unsafe.Pointer, status C.int) {
	if signal := qt.GetSignal(ptr, "setStatus"); signal != nil {
		(*(*func(int))(signal))(int(int32(status)))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetStatusDefault(int(int32(status)))
	}
}

func (ptr *TransactionDetails) ConnectSetStatus(f func(status int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setStatus"); signal != nil {
			f := func(status int) {
				(*(*func(int))(signal))(status)
				f(status)
			}
			qt.ConnectSignal(ptr.Pointer(), "setStatus", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setStatus", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetStatus() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setStatus")
	}
}

func (ptr *TransactionDetails) SetStatus(status int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetStatus(ptr.Pointer(), C.int(int32(status)))
	}
}

func (ptr *TransactionDetails) SetStatusDefault(status int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetStatusDefault(ptr.Pointer(), C.int(int32(status)))
	}
}

//export callbackTransactionDetails1c4313_StatusChanged
func callbackTransactionDetails1c4313_StatusChanged(ptr unsafe.Pointer, status C.int) {
	if signal := qt.GetSignal(ptr, "statusChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(status)))
	}

}

func (ptr *TransactionDetails) ConnectStatusChanged(f func(status int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "statusChanged") {
			C.TransactionDetails1c4313_ConnectStatusChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "statusChanged"); signal != nil {
			f := func(status int) {
				(*(*func(int))(signal))(status)
				f(status)
			}
			qt.ConnectSignal(ptr.Pointer(), "statusChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "statusChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "statusChanged")
	}
}

func (ptr *TransactionDetails) StatusChanged(status int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_StatusChanged(ptr.Pointer(), C.int(int32(status)))
	}
}

//export callbackTransactionDetails1c4313_Type
func callbackTransactionDetails1c4313_Type(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewTransactionDetailsFromPointer(ptr).TypeDefault()))
}

func (ptr *TransactionDetails) ConnectType(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *TransactionDetails) Type() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails1c4313_Type(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) TypeDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails1c4313_TypeDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetails1c4313_SetType
func callbackTransactionDetails1c4313_SetType(ptr unsafe.Pointer, ty C.int) {
	if signal := qt.GetSignal(ptr, "setType"); signal != nil {
		(*(*func(int))(signal))(int(int32(ty)))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetTypeDefault(int(int32(ty)))
	}
}

func (ptr *TransactionDetails) ConnectSetType(f func(ty int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setType"); signal != nil {
			f := func(ty int) {
				(*(*func(int))(signal))(ty)
				f(ty)
			}
			qt.ConnectSignal(ptr.Pointer(), "setType", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setType", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setType")
	}
}

func (ptr *TransactionDetails) SetType(ty int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetType(ptr.Pointer(), C.int(int32(ty)))
	}
}

func (ptr *TransactionDetails) SetTypeDefault(ty int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetTypeDefault(ptr.Pointer(), C.int(int32(ty)))
	}
}

//export callbackTransactionDetails1c4313_TypeChanged
func callbackTransactionDetails1c4313_TypeChanged(ptr unsafe.Pointer, ty C.int) {
	if signal := qt.GetSignal(ptr, "typeChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(ty)))
	}

}

func (ptr *TransactionDetails) ConnectTypeChanged(f func(ty int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "typeChanged") {
			C.TransactionDetails1c4313_ConnectTypeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "typeChanged"); signal != nil {
			f := func(ty int) {
				(*(*func(int))(signal))(ty)
				f(ty)
			}
			qt.ConnectSignal(ptr.Pointer(), "typeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "typeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectTypeChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "typeChanged")
	}
}

func (ptr *TransactionDetails) TypeChanged(ty int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_TypeChanged(ptr.Pointer(), C.int(int32(ty)))
	}
}

//export callbackTransactionDetails1c4313_Amount
func callbackTransactionDetails1c4313_Amount(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "amount"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewTransactionDetailsFromPointer(ptr).AmountDefault()))
}

func (ptr *TransactionDetails) ConnectAmount(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "amount"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "amount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "amount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectAmount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "amount")
	}
}

func (ptr *TransactionDetails) Amount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails1c4313_Amount(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) AmountDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails1c4313_AmountDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetails1c4313_SetAmount
func callbackTransactionDetails1c4313_SetAmount(ptr unsafe.Pointer, amount C.int) {
	if signal := qt.GetSignal(ptr, "setAmount"); signal != nil {
		(*(*func(int))(signal))(int(int32(amount)))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetAmountDefault(int(int32(amount)))
	}
}

func (ptr *TransactionDetails) ConnectSetAmount(f func(amount int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAmount"); signal != nil {
			f := func(amount int) {
				(*(*func(int))(signal))(amount)
				f(amount)
			}
			qt.ConnectSignal(ptr.Pointer(), "setAmount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAmount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetAmount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAmount")
	}
}

func (ptr *TransactionDetails) SetAmount(amount int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetAmount(ptr.Pointer(), C.int(int32(amount)))
	}
}

func (ptr *TransactionDetails) SetAmountDefault(amount int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetAmountDefault(ptr.Pointer(), C.int(int32(amount)))
	}
}

//export callbackTransactionDetails1c4313_AmountChanged
func callbackTransactionDetails1c4313_AmountChanged(ptr unsafe.Pointer, amount C.int) {
	if signal := qt.GetSignal(ptr, "amountChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(amount)))
	}

}

func (ptr *TransactionDetails) ConnectAmountChanged(f func(amount int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "amountChanged") {
			C.TransactionDetails1c4313_ConnectAmountChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "amountChanged"); signal != nil {
			f := func(amount int) {
				(*(*func(int))(signal))(amount)
				f(amount)
			}
			qt.ConnectSignal(ptr.Pointer(), "amountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "amountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectAmountChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DisconnectAmountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "amountChanged")
	}
}

func (ptr *TransactionDetails) AmountChanged(amount int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_AmountChanged(ptr.Pointer(), C.int(int32(amount)))
	}
}

//export callbackTransactionDetails1c4313_HoursReceived
func callbackTransactionDetails1c4313_HoursReceived(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "hoursReceived"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewTransactionDetailsFromPointer(ptr).HoursReceivedDefault()))
}

func (ptr *TransactionDetails) ConnectHoursReceived(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hoursReceived"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hoursReceived", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hoursReceived", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectHoursReceived() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hoursReceived")
	}
}

func (ptr *TransactionDetails) HoursReceived() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails1c4313_HoursReceived(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) HoursReceivedDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails1c4313_HoursReceivedDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetails1c4313_SetHoursReceived
func callbackTransactionDetails1c4313_SetHoursReceived(ptr unsafe.Pointer, hoursReceived C.int) {
	if signal := qt.GetSignal(ptr, "setHoursReceived"); signal != nil {
		(*(*func(int))(signal))(int(int32(hoursReceived)))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetHoursReceivedDefault(int(int32(hoursReceived)))
	}
}

func (ptr *TransactionDetails) ConnectSetHoursReceived(f func(hoursReceived int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setHoursReceived"); signal != nil {
			f := func(hoursReceived int) {
				(*(*func(int))(signal))(hoursReceived)
				f(hoursReceived)
			}
			qt.ConnectSignal(ptr.Pointer(), "setHoursReceived", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setHoursReceived", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetHoursReceived() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setHoursReceived")
	}
}

func (ptr *TransactionDetails) SetHoursReceived(hoursReceived int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetHoursReceived(ptr.Pointer(), C.int(int32(hoursReceived)))
	}
}

func (ptr *TransactionDetails) SetHoursReceivedDefault(hoursReceived int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetHoursReceivedDefault(ptr.Pointer(), C.int(int32(hoursReceived)))
	}
}

//export callbackTransactionDetails1c4313_HoursReceivedChanged
func callbackTransactionDetails1c4313_HoursReceivedChanged(ptr unsafe.Pointer, hoursReceived C.int) {
	if signal := qt.GetSignal(ptr, "hoursReceivedChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(hoursReceived)))
	}

}

func (ptr *TransactionDetails) ConnectHoursReceivedChanged(f func(hoursReceived int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hoursReceivedChanged") {
			C.TransactionDetails1c4313_ConnectHoursReceivedChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hoursReceivedChanged"); signal != nil {
			f := func(hoursReceived int) {
				(*(*func(int))(signal))(hoursReceived)
				f(hoursReceived)
			}
			qt.ConnectSignal(ptr.Pointer(), "hoursReceivedChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hoursReceivedChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectHoursReceivedChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DisconnectHoursReceivedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hoursReceivedChanged")
	}
}

func (ptr *TransactionDetails) HoursReceivedChanged(hoursReceived int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_HoursReceivedChanged(ptr.Pointer(), C.int(int32(hoursReceived)))
	}
}

//export callbackTransactionDetails1c4313_HoursBurned
func callbackTransactionDetails1c4313_HoursBurned(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "hoursBurned"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewTransactionDetailsFromPointer(ptr).HoursBurnedDefault()))
}

func (ptr *TransactionDetails) ConnectHoursBurned(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hoursBurned"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hoursBurned", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hoursBurned", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectHoursBurned() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hoursBurned")
	}
}

func (ptr *TransactionDetails) HoursBurned() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails1c4313_HoursBurned(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) HoursBurnedDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails1c4313_HoursBurnedDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetails1c4313_SetHoursBurned
func callbackTransactionDetails1c4313_SetHoursBurned(ptr unsafe.Pointer, hoursBurned C.int) {
	if signal := qt.GetSignal(ptr, "setHoursBurned"); signal != nil {
		(*(*func(int))(signal))(int(int32(hoursBurned)))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetHoursBurnedDefault(int(int32(hoursBurned)))
	}
}

func (ptr *TransactionDetails) ConnectSetHoursBurned(f func(hoursBurned int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setHoursBurned"); signal != nil {
			f := func(hoursBurned int) {
				(*(*func(int))(signal))(hoursBurned)
				f(hoursBurned)
			}
			qt.ConnectSignal(ptr.Pointer(), "setHoursBurned", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setHoursBurned", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetHoursBurned() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setHoursBurned")
	}
}

func (ptr *TransactionDetails) SetHoursBurned(hoursBurned int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetHoursBurned(ptr.Pointer(), C.int(int32(hoursBurned)))
	}
}

func (ptr *TransactionDetails) SetHoursBurnedDefault(hoursBurned int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetHoursBurnedDefault(ptr.Pointer(), C.int(int32(hoursBurned)))
	}
}

//export callbackTransactionDetails1c4313_HoursBurnedChanged
func callbackTransactionDetails1c4313_HoursBurnedChanged(ptr unsafe.Pointer, hoursBurned C.int) {
	if signal := qt.GetSignal(ptr, "hoursBurnedChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(hoursBurned)))
	}

}

func (ptr *TransactionDetails) ConnectHoursBurnedChanged(f func(hoursBurned int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hoursBurnedChanged") {
			C.TransactionDetails1c4313_ConnectHoursBurnedChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hoursBurnedChanged"); signal != nil {
			f := func(hoursBurned int) {
				(*(*func(int))(signal))(hoursBurned)
				f(hoursBurned)
			}
			qt.ConnectSignal(ptr.Pointer(), "hoursBurnedChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hoursBurnedChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectHoursBurnedChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DisconnectHoursBurnedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hoursBurnedChanged")
	}
}

func (ptr *TransactionDetails) HoursBurnedChanged(hoursBurned int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_HoursBurnedChanged(ptr.Pointer(), C.int(int32(hoursBurned)))
	}
}

//export callbackTransactionDetails1c4313_TransactionID
func callbackTransactionDetails1c4313_TransactionID(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "transactionID"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewTransactionDetailsFromPointer(ptr).TransactionIDDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *TransactionDetails) ConnectTransactionID(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "transactionID"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "transactionID", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transactionID", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectTransactionID() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "transactionID")
	}
}

func (ptr *TransactionDetails) TransactionID() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails1c4313_TransactionID(ptr.Pointer()))
	}
	return ""
}

func (ptr *TransactionDetails) TransactionIDDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails1c4313_TransactionIDDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTransactionDetails1c4313_SetTransactionID
func callbackTransactionDetails1c4313_SetTransactionID(ptr unsafe.Pointer, transactionID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTransactionID"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(transactionID))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetTransactionIDDefault(cGoUnpackString(transactionID))
	}
}

func (ptr *TransactionDetails) ConnectSetTransactionID(f func(transactionID string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTransactionID"); signal != nil {
			f := func(transactionID string) {
				(*(*func(string))(signal))(transactionID)
				f(transactionID)
			}
			qt.ConnectSignal(ptr.Pointer(), "setTransactionID", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTransactionID", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetTransactionID() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTransactionID")
	}
}

func (ptr *TransactionDetails) SetTransactionID(transactionID string) {
	if ptr.Pointer() != nil {
		var transactionIDC *C.char
		if transactionID != "" {
			transactionIDC = C.CString(transactionID)
			defer C.free(unsafe.Pointer(transactionIDC))
		}
		C.TransactionDetails1c4313_SetTransactionID(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

func (ptr *TransactionDetails) SetTransactionIDDefault(transactionID string) {
	if ptr.Pointer() != nil {
		var transactionIDC *C.char
		if transactionID != "" {
			transactionIDC = C.CString(transactionID)
			defer C.free(unsafe.Pointer(transactionIDC))
		}
		C.TransactionDetails1c4313_SetTransactionIDDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

//export callbackTransactionDetails1c4313_TransactionIDChanged
func callbackTransactionDetails1c4313_TransactionIDChanged(ptr unsafe.Pointer, transactionID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "transactionIDChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(transactionID))
	}

}

func (ptr *TransactionDetails) ConnectTransactionIDChanged(f func(transactionID string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transactionIDChanged") {
			C.TransactionDetails1c4313_ConnectTransactionIDChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "transactionIDChanged"); signal != nil {
			f := func(transactionID string) {
				(*(*func(string))(signal))(transactionID)
				f(transactionID)
			}
			qt.ConnectSignal(ptr.Pointer(), "transactionIDChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transactionIDChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectTransactionIDChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DisconnectTransactionIDChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "transactionIDChanged")
	}
}

func (ptr *TransactionDetails) TransactionIDChanged(transactionID string) {
	if ptr.Pointer() != nil {
		var transactionIDC *C.char
		if transactionID != "" {
			transactionIDC = C.CString(transactionID)
			defer C.free(unsafe.Pointer(transactionIDC))
		}
		C.TransactionDetails1c4313_TransactionIDChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

//export callbackTransactionDetails1c4313_SentAddress
func callbackTransactionDetails1c4313_SentAddress(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "sentAddress"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewTransactionDetailsFromPointer(ptr).SentAddressDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *TransactionDetails) ConnectSentAddress(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sentAddress"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "sentAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sentAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSentAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sentAddress")
	}
}

func (ptr *TransactionDetails) SentAddress() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails1c4313_SentAddress(ptr.Pointer()))
	}
	return ""
}

func (ptr *TransactionDetails) SentAddressDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails1c4313_SentAddressDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTransactionDetails1c4313_SetSentAddress
func callbackTransactionDetails1c4313_SetSentAddress(ptr unsafe.Pointer, sentAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setSentAddress"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(sentAddress))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetSentAddressDefault(cGoUnpackString(sentAddress))
	}
}

func (ptr *TransactionDetails) ConnectSetSentAddress(f func(sentAddress string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSentAddress"); signal != nil {
			f := func(sentAddress string) {
				(*(*func(string))(signal))(sentAddress)
				f(sentAddress)
			}
			qt.ConnectSignal(ptr.Pointer(), "setSentAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSentAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetSentAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSentAddress")
	}
}

func (ptr *TransactionDetails) SetSentAddress(sentAddress string) {
	if ptr.Pointer() != nil {
		var sentAddressC *C.char
		if sentAddress != "" {
			sentAddressC = C.CString(sentAddress)
			defer C.free(unsafe.Pointer(sentAddressC))
		}
		C.TransactionDetails1c4313_SetSentAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: sentAddressC, len: C.longlong(len(sentAddress))})
	}
}

func (ptr *TransactionDetails) SetSentAddressDefault(sentAddress string) {
	if ptr.Pointer() != nil {
		var sentAddressC *C.char
		if sentAddress != "" {
			sentAddressC = C.CString(sentAddress)
			defer C.free(unsafe.Pointer(sentAddressC))
		}
		C.TransactionDetails1c4313_SetSentAddressDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: sentAddressC, len: C.longlong(len(sentAddress))})
	}
}

//export callbackTransactionDetails1c4313_SentAddressChanged
func callbackTransactionDetails1c4313_SentAddressChanged(ptr unsafe.Pointer, sentAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "sentAddressChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(sentAddress))
	}

}

func (ptr *TransactionDetails) ConnectSentAddressChanged(f func(sentAddress string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sentAddressChanged") {
			C.TransactionDetails1c4313_ConnectSentAddressChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sentAddressChanged"); signal != nil {
			f := func(sentAddress string) {
				(*(*func(string))(signal))(sentAddress)
				f(sentAddress)
			}
			qt.ConnectSignal(ptr.Pointer(), "sentAddressChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sentAddressChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSentAddressChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DisconnectSentAddressChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sentAddressChanged")
	}
}

func (ptr *TransactionDetails) SentAddressChanged(sentAddress string) {
	if ptr.Pointer() != nil {
		var sentAddressC *C.char
		if sentAddress != "" {
			sentAddressC = C.CString(sentAddress)
			defer C.free(unsafe.Pointer(sentAddressC))
		}
		C.TransactionDetails1c4313_SentAddressChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: sentAddressC, len: C.longlong(len(sentAddress))})
	}
}

//export callbackTransactionDetails1c4313_ReceivedAddress
func callbackTransactionDetails1c4313_ReceivedAddress(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "receivedAddress"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewTransactionDetailsFromPointer(ptr).ReceivedAddressDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *TransactionDetails) ConnectReceivedAddress(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "receivedAddress"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "receivedAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "receivedAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectReceivedAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "receivedAddress")
	}
}

func (ptr *TransactionDetails) ReceivedAddress() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails1c4313_ReceivedAddress(ptr.Pointer()))
	}
	return ""
}

func (ptr *TransactionDetails) ReceivedAddressDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails1c4313_ReceivedAddressDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTransactionDetails1c4313_SetReceivedAddress
func callbackTransactionDetails1c4313_SetReceivedAddress(ptr unsafe.Pointer, receivedAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setReceivedAddress"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(receivedAddress))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetReceivedAddressDefault(cGoUnpackString(receivedAddress))
	}
}

func (ptr *TransactionDetails) ConnectSetReceivedAddress(f func(receivedAddress string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setReceivedAddress"); signal != nil {
			f := func(receivedAddress string) {
				(*(*func(string))(signal))(receivedAddress)
				f(receivedAddress)
			}
			qt.ConnectSignal(ptr.Pointer(), "setReceivedAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setReceivedAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetReceivedAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setReceivedAddress")
	}
}

func (ptr *TransactionDetails) SetReceivedAddress(receivedAddress string) {
	if ptr.Pointer() != nil {
		var receivedAddressC *C.char
		if receivedAddress != "" {
			receivedAddressC = C.CString(receivedAddress)
			defer C.free(unsafe.Pointer(receivedAddressC))
		}
		C.TransactionDetails1c4313_SetReceivedAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: receivedAddressC, len: C.longlong(len(receivedAddress))})
	}
}

func (ptr *TransactionDetails) SetReceivedAddressDefault(receivedAddress string) {
	if ptr.Pointer() != nil {
		var receivedAddressC *C.char
		if receivedAddress != "" {
			receivedAddressC = C.CString(receivedAddress)
			defer C.free(unsafe.Pointer(receivedAddressC))
		}
		C.TransactionDetails1c4313_SetReceivedAddressDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: receivedAddressC, len: C.longlong(len(receivedAddress))})
	}
}

//export callbackTransactionDetails1c4313_ReceivedAddressChanged
func callbackTransactionDetails1c4313_ReceivedAddressChanged(ptr unsafe.Pointer, receivedAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "receivedAddressChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(receivedAddress))
	}

}

func (ptr *TransactionDetails) ConnectReceivedAddressChanged(f func(receivedAddress string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "receivedAddressChanged") {
			C.TransactionDetails1c4313_ConnectReceivedAddressChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "receivedAddressChanged"); signal != nil {
			f := func(receivedAddress string) {
				(*(*func(string))(signal))(receivedAddress)
				f(receivedAddress)
			}
			qt.ConnectSignal(ptr.Pointer(), "receivedAddressChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "receivedAddressChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectReceivedAddressChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DisconnectReceivedAddressChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "receivedAddressChanged")
	}
}

func (ptr *TransactionDetails) ReceivedAddressChanged(receivedAddress string) {
	if ptr.Pointer() != nil {
		var receivedAddressC *C.char
		if receivedAddress != "" {
			receivedAddressC = C.CString(receivedAddress)
			defer C.free(unsafe.Pointer(receivedAddressC))
		}
		C.TransactionDetails1c4313_ReceivedAddressChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: receivedAddressC, len: C.longlong(len(receivedAddress))})
	}
}

//export callbackTransactionDetails1c4313_Inputs
func callbackTransactionDetails1c4313_Inputs(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputs"); signal != nil {
		return PointerFromAddressList((*(*func() *AddressList)(signal))())
	}

	return PointerFromAddressList(NewTransactionDetailsFromPointer(ptr).InputsDefault())
}

func (ptr *TransactionDetails) ConnectInputs(f func() *AddressList) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "inputs"); signal != nil {
			f := func() *AddressList {
				(*(*func() *AddressList)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "inputs", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "inputs", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectInputs() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "inputs")
	}
}

func (ptr *TransactionDetails) Inputs() *AddressList {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressListFromPointer(C.TransactionDetails1c4313_Inputs(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) InputsDefault() *AddressList {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressListFromPointer(C.TransactionDetails1c4313_InputsDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackTransactionDetails1c4313_SetInputs
func callbackTransactionDetails1c4313_SetInputs(ptr unsafe.Pointer, inputs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setInputs"); signal != nil {
		(*(*func(*AddressList))(signal))(NewAddressListFromPointer(inputs))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetInputsDefault(NewAddressListFromPointer(inputs))
	}
}

func (ptr *TransactionDetails) ConnectSetInputs(f func(inputs *AddressList)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setInputs"); signal != nil {
			f := func(inputs *AddressList) {
				(*(*func(*AddressList))(signal))(inputs)
				f(inputs)
			}
			qt.ConnectSignal(ptr.Pointer(), "setInputs", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setInputs", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetInputs() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setInputs")
	}
}

func (ptr *TransactionDetails) SetInputs(inputs AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetInputs(ptr.Pointer(), PointerFromAddressList(inputs))
	}
}

func (ptr *TransactionDetails) SetInputsDefault(inputs AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetInputsDefault(ptr.Pointer(), PointerFromAddressList(inputs))
	}
}

//export callbackTransactionDetails1c4313_InputsChanged
func callbackTransactionDetails1c4313_InputsChanged(ptr unsafe.Pointer, inputs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputsChanged"); signal != nil {
		(*(*func(*AddressList))(signal))(NewAddressListFromPointer(inputs))
	}

}

func (ptr *TransactionDetails) ConnectInputsChanged(f func(inputs *AddressList)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "inputsChanged") {
			C.TransactionDetails1c4313_ConnectInputsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "inputsChanged"); signal != nil {
			f := func(inputs *AddressList) {
				(*(*func(*AddressList))(signal))(inputs)
				f(inputs)
			}
			qt.ConnectSignal(ptr.Pointer(), "inputsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "inputsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectInputsChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DisconnectInputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "inputsChanged")
	}
}

func (ptr *TransactionDetails) InputsChanged(inputs AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_InputsChanged(ptr.Pointer(), PointerFromAddressList(inputs))
	}
}

//export callbackTransactionDetails1c4313_Outputs
func callbackTransactionDetails1c4313_Outputs(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "outputs"); signal != nil {
		return PointerFromAddressList((*(*func() *AddressList)(signal))())
	}

	return PointerFromAddressList(NewTransactionDetailsFromPointer(ptr).OutputsDefault())
}

func (ptr *TransactionDetails) ConnectOutputs(f func() *AddressList) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "outputs"); signal != nil {
			f := func() *AddressList {
				(*(*func() *AddressList)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "outputs", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "outputs", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectOutputs() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "outputs")
	}
}

func (ptr *TransactionDetails) Outputs() *AddressList {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressListFromPointer(C.TransactionDetails1c4313_Outputs(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) OutputsDefault() *AddressList {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressListFromPointer(C.TransactionDetails1c4313_OutputsDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackTransactionDetails1c4313_SetOutputs
func callbackTransactionDetails1c4313_SetOutputs(ptr unsafe.Pointer, outputs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setOutputs"); signal != nil {
		(*(*func(*AddressList))(signal))(NewAddressListFromPointer(outputs))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetOutputsDefault(NewAddressListFromPointer(outputs))
	}
}

func (ptr *TransactionDetails) ConnectSetOutputs(f func(outputs *AddressList)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setOutputs"); signal != nil {
			f := func(outputs *AddressList) {
				(*(*func(*AddressList))(signal))(outputs)
				f(outputs)
			}
			qt.ConnectSignal(ptr.Pointer(), "setOutputs", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setOutputs", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetOutputs() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setOutputs")
	}
}

func (ptr *TransactionDetails) SetOutputs(outputs AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetOutputs(ptr.Pointer(), PointerFromAddressList(outputs))
	}
}

func (ptr *TransactionDetails) SetOutputsDefault(outputs AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_SetOutputsDefault(ptr.Pointer(), PointerFromAddressList(outputs))
	}
}

//export callbackTransactionDetails1c4313_OutputsChanged
func callbackTransactionDetails1c4313_OutputsChanged(ptr unsafe.Pointer, outputs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "outputsChanged"); signal != nil {
		(*(*func(*AddressList))(signal))(NewAddressListFromPointer(outputs))
	}

}

func (ptr *TransactionDetails) ConnectOutputsChanged(f func(outputs *AddressList)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "outputsChanged") {
			C.TransactionDetails1c4313_ConnectOutputsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "outputsChanged"); signal != nil {
			f := func(outputs *AddressList) {
				(*(*func(*AddressList))(signal))(outputs)
				f(outputs)
			}
			qt.ConnectSignal(ptr.Pointer(), "outputsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "outputsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectOutputsChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DisconnectOutputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "outputsChanged")
	}
}

func (ptr *TransactionDetails) OutputsChanged(outputs AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_OutputsChanged(ptr.Pointer(), PointerFromAddressList(outputs))
	}
}

func TransactionDetails_QRegisterMetaType() int {
	return int(int32(C.TransactionDetails1c4313_TransactionDetails1c4313_QRegisterMetaType()))
}

func (ptr *TransactionDetails) QRegisterMetaType() int {
	return int(int32(C.TransactionDetails1c4313_TransactionDetails1c4313_QRegisterMetaType()))
}

func TransactionDetails_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TransactionDetails1c4313_TransactionDetails1c4313_QRegisterMetaType2(typeNameC)))
}

func (ptr *TransactionDetails) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TransactionDetails1c4313_TransactionDetails1c4313_QRegisterMetaType2(typeNameC)))
}

func TransactionDetails_QmlRegisterType() int {
	return int(int32(C.TransactionDetails1c4313_TransactionDetails1c4313_QmlRegisterType()))
}

func (ptr *TransactionDetails) QmlRegisterType() int {
	return int(int32(C.TransactionDetails1c4313_TransactionDetails1c4313_QmlRegisterType()))
}

func TransactionDetails_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.TransactionDetails1c4313_TransactionDetails1c4313_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TransactionDetails) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.TransactionDetails1c4313_TransactionDetails1c4313_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TransactionDetails) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetails1c4313___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __children_newList() unsafe.Pointer {
	return C.TransactionDetails1c4313___children_newList(ptr.Pointer())
}

func (ptr *TransactionDetails) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TransactionDetails1c4313___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TransactionDetails) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.TransactionDetails1c4313___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *TransactionDetails) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetails1c4313___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __findChildren_newList() unsafe.Pointer {
	return C.TransactionDetails1c4313___findChildren_newList(ptr.Pointer())
}

func (ptr *TransactionDetails) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetails1c4313___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __findChildren_newList3() unsafe.Pointer {
	return C.TransactionDetails1c4313___findChildren_newList3(ptr.Pointer())
}

func (ptr *TransactionDetails) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetails1c4313___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __qFindChildren_newList2() unsafe.Pointer {
	return C.TransactionDetails1c4313___qFindChildren_newList2(ptr.Pointer())
}

func NewTransactionDetails(parent std_core.QObject_ITF) *TransactionDetails {
	tmpValue := NewTransactionDetailsFromPointer(C.TransactionDetails1c4313_NewTransactionDetails(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackTransactionDetails1c4313_DestroyTransactionDetails
func callbackTransactionDetails1c4313_DestroyTransactionDetails(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~TransactionDetails"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewTransactionDetailsFromPointer(ptr).DestroyTransactionDetailsDefault()
	}
}

func (ptr *TransactionDetails) ConnectDestroyTransactionDetails(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~TransactionDetails"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~TransactionDetails", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~TransactionDetails", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectDestroyTransactionDetails() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~TransactionDetails")
	}
}

func (ptr *TransactionDetails) DestroyTransactionDetails() {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DestroyTransactionDetails(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *TransactionDetails) DestroyTransactionDetailsDefault() {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DestroyTransactionDetailsDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTransactionDetails1c4313_ChildEvent
func callbackTransactionDetails1c4313_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewTransactionDetailsFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *TransactionDetails) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackTransactionDetails1c4313_ConnectNotify
func callbackTransactionDetails1c4313_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTransactionDetailsFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TransactionDetails) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTransactionDetails1c4313_CustomEvent
func callbackTransactionDetails1c4313_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewTransactionDetailsFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *TransactionDetails) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackTransactionDetails1c4313_DeleteLater
func callbackTransactionDetails1c4313_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewTransactionDetailsFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *TransactionDetails) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTransactionDetails1c4313_Destroyed
func callbackTransactionDetails1c4313_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackTransactionDetails1c4313_DisconnectNotify
func callbackTransactionDetails1c4313_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTransactionDetailsFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TransactionDetails) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTransactionDetails1c4313_Event
func callbackTransactionDetails1c4313_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionDetailsFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *TransactionDetails) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionDetails1c4313_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackTransactionDetails1c4313_EventFilter
func callbackTransactionDetails1c4313_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionDetailsFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *TransactionDetails) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionDetails1c4313_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackTransactionDetails1c4313_ObjectNameChanged
func callbackTransactionDetails1c4313_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackTransactionDetails1c4313_TimerEvent
func callbackTransactionDetails1c4313_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewTransactionDetailsFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *TransactionDetails) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails1c4313_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type AddressDetails_ITF interface {
	std_core.QObject_ITF
	AddressDetails_PTR() *AddressDetails
}

func (ptr *AddressDetails) AddressDetails_PTR() *AddressDetails {
	return ptr
}

func (ptr *AddressDetails) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *AddressDetails) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromAddressDetails(ptr AddressDetails_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AddressDetails_PTR().Pointer()
	}
	return nil
}

func NewAddressDetailsFromPointer(ptr unsafe.Pointer) (n *AddressDetails) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(AddressDetails)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *AddressDetails:
			n = deduced

		case *std_core.QObject:
			n = &AddressDetails{QObject: *deduced}

		default:
			n = new(AddressDetails)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackAddressDetails1c4313_Constructor
func callbackAddressDetails1c4313_Constructor(ptr unsafe.Pointer) {
	this := NewAddressDetailsFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackAddressDetails1c4313_Address
func callbackAddressDetails1c4313_Address(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "address"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewAddressDetailsFromPointer(ptr).AddressDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *AddressDetails) ConnectAddress(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "address"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "address", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "address", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressDetails) DisconnectAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "address")
	}
}

func (ptr *AddressDetails) Address() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.AddressDetails1c4313_Address(ptr.Pointer()))
	}
	return ""
}

func (ptr *AddressDetails) AddressDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.AddressDetails1c4313_AddressDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackAddressDetails1c4313_SetAddress
func callbackAddressDetails1c4313_SetAddress(ptr unsafe.Pointer, address C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setAddress"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(address))
	} else {
		NewAddressDetailsFromPointer(ptr).SetAddressDefault(cGoUnpackString(address))
	}
}

func (ptr *AddressDetails) ConnectSetAddress(f func(address string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAddress"); signal != nil {
			f := func(address string) {
				(*(*func(string))(signal))(address)
				f(address)
			}
			qt.ConnectSignal(ptr.Pointer(), "setAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressDetails) DisconnectSetAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAddress")
	}
}

func (ptr *AddressDetails) SetAddress(address string) {
	if ptr.Pointer() != nil {
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		C.AddressDetails1c4313_SetAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))})
	}
}

func (ptr *AddressDetails) SetAddressDefault(address string) {
	if ptr.Pointer() != nil {
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		C.AddressDetails1c4313_SetAddressDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))})
	}
}

//export callbackAddressDetails1c4313_AddressChanged
func callbackAddressDetails1c4313_AddressChanged(ptr unsafe.Pointer, address C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "addressChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(address))
	}

}

func (ptr *AddressDetails) ConnectAddressChanged(f func(address string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressChanged") {
			C.AddressDetails1c4313_ConnectAddressChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addressChanged"); signal != nil {
			f := func(address string) {
				(*(*func(string))(signal))(address)
				f(address)
			}
			qt.ConnectSignal(ptr.Pointer(), "addressChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressDetails) DisconnectAddressChanged() {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_DisconnectAddressChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressChanged")
	}
}

func (ptr *AddressDetails) AddressChanged(address string) {
	if ptr.Pointer() != nil {
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		C.AddressDetails1c4313_AddressChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))})
	}
}

//export callbackAddressDetails1c4313_AddressSky
func callbackAddressDetails1c4313_AddressSky(ptr unsafe.Pointer) C.float {
	if signal := qt.GetSignal(ptr, "addressSky"); signal != nil {
		return C.float((*(*func() float32)(signal))())
	}

	return C.float(NewAddressDetailsFromPointer(ptr).AddressSkyDefault())
}

func (ptr *AddressDetails) ConnectAddressSky(f func() float32) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addressSky"); signal != nil {
			f := func() float32 {
				(*(*func() float32)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "addressSky", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressSky", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressDetails) DisconnectAddressSky() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addressSky")
	}
}

func (ptr *AddressDetails) AddressSky() float32 {
	if ptr.Pointer() != nil {
		return float32(C.AddressDetails1c4313_AddressSky(ptr.Pointer()))
	}
	return 0
}

func (ptr *AddressDetails) AddressSkyDefault() float32 {
	if ptr.Pointer() != nil {
		return float32(C.AddressDetails1c4313_AddressSkyDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackAddressDetails1c4313_SetAddressSky
func callbackAddressDetails1c4313_SetAddressSky(ptr unsafe.Pointer, addressSky C.float) {
	if signal := qt.GetSignal(ptr, "setAddressSky"); signal != nil {
		(*(*func(float32))(signal))(float32(addressSky))
	} else {
		NewAddressDetailsFromPointer(ptr).SetAddressSkyDefault(float32(addressSky))
	}
}

func (ptr *AddressDetails) ConnectSetAddressSky(f func(addressSky float32)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAddressSky"); signal != nil {
			f := func(addressSky float32) {
				(*(*func(float32))(signal))(addressSky)
				f(addressSky)
			}
			qt.ConnectSignal(ptr.Pointer(), "setAddressSky", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAddressSky", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressDetails) DisconnectSetAddressSky() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAddressSky")
	}
}

func (ptr *AddressDetails) SetAddressSky(addressSky float32) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_SetAddressSky(ptr.Pointer(), C.float(addressSky))
	}
}

func (ptr *AddressDetails) SetAddressSkyDefault(addressSky float32) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_SetAddressSkyDefault(ptr.Pointer(), C.float(addressSky))
	}
}

//export callbackAddressDetails1c4313_AddressSkyChanged
func callbackAddressDetails1c4313_AddressSkyChanged(ptr unsafe.Pointer, addressSky C.float) {
	if signal := qt.GetSignal(ptr, "addressSkyChanged"); signal != nil {
		(*(*func(float32))(signal))(float32(addressSky))
	}

}

func (ptr *AddressDetails) ConnectAddressSkyChanged(f func(addressSky float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressSkyChanged") {
			C.AddressDetails1c4313_ConnectAddressSkyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addressSkyChanged"); signal != nil {
			f := func(addressSky float32) {
				(*(*func(float32))(signal))(addressSky)
				f(addressSky)
			}
			qt.ConnectSignal(ptr.Pointer(), "addressSkyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressSkyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressDetails) DisconnectAddressSkyChanged() {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_DisconnectAddressSkyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressSkyChanged")
	}
}

func (ptr *AddressDetails) AddressSkyChanged(addressSky float32) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_AddressSkyChanged(ptr.Pointer(), C.float(addressSky))
	}
}

//export callbackAddressDetails1c4313_AddressCoinHours
func callbackAddressDetails1c4313_AddressCoinHours(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "addressCoinHours"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewAddressDetailsFromPointer(ptr).AddressCoinHoursDefault()))
}

func (ptr *AddressDetails) ConnectAddressCoinHours(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addressCoinHours"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "addressCoinHours", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressCoinHours", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressDetails) DisconnectAddressCoinHours() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addressCoinHours")
	}
}

func (ptr *AddressDetails) AddressCoinHours() int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressDetails1c4313_AddressCoinHours(ptr.Pointer())))
	}
	return 0
}

func (ptr *AddressDetails) AddressCoinHoursDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressDetails1c4313_AddressCoinHoursDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackAddressDetails1c4313_SetAddressCoinHours
func callbackAddressDetails1c4313_SetAddressCoinHours(ptr unsafe.Pointer, addressCoinHours C.int) {
	if signal := qt.GetSignal(ptr, "setAddressCoinHours"); signal != nil {
		(*(*func(int))(signal))(int(int32(addressCoinHours)))
	} else {
		NewAddressDetailsFromPointer(ptr).SetAddressCoinHoursDefault(int(int32(addressCoinHours)))
	}
}

func (ptr *AddressDetails) ConnectSetAddressCoinHours(f func(addressCoinHours int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAddressCoinHours"); signal != nil {
			f := func(addressCoinHours int) {
				(*(*func(int))(signal))(addressCoinHours)
				f(addressCoinHours)
			}
			qt.ConnectSignal(ptr.Pointer(), "setAddressCoinHours", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAddressCoinHours", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressDetails) DisconnectSetAddressCoinHours() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAddressCoinHours")
	}
}

func (ptr *AddressDetails) SetAddressCoinHours(addressCoinHours int) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_SetAddressCoinHours(ptr.Pointer(), C.int(int32(addressCoinHours)))
	}
}

func (ptr *AddressDetails) SetAddressCoinHoursDefault(addressCoinHours int) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_SetAddressCoinHoursDefault(ptr.Pointer(), C.int(int32(addressCoinHours)))
	}
}

//export callbackAddressDetails1c4313_AddressCoinHoursChanged
func callbackAddressDetails1c4313_AddressCoinHoursChanged(ptr unsafe.Pointer, addressCoinHours C.int) {
	if signal := qt.GetSignal(ptr, "addressCoinHoursChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(addressCoinHours)))
	}

}

func (ptr *AddressDetails) ConnectAddressCoinHoursChanged(f func(addressCoinHours int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressCoinHoursChanged") {
			C.AddressDetails1c4313_ConnectAddressCoinHoursChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addressCoinHoursChanged"); signal != nil {
			f := func(addressCoinHours int) {
				(*(*func(int))(signal))(addressCoinHours)
				f(addressCoinHours)
			}
			qt.ConnectSignal(ptr.Pointer(), "addressCoinHoursChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressCoinHoursChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressDetails) DisconnectAddressCoinHoursChanged() {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_DisconnectAddressCoinHoursChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressCoinHoursChanged")
	}
}

func (ptr *AddressDetails) AddressCoinHoursChanged(addressCoinHours int) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_AddressCoinHoursChanged(ptr.Pointer(), C.int(int32(addressCoinHours)))
	}
}

func AddressDetails_QRegisterMetaType() int {
	return int(int32(C.AddressDetails1c4313_AddressDetails1c4313_QRegisterMetaType()))
}

func (ptr *AddressDetails) QRegisterMetaType() int {
	return int(int32(C.AddressDetails1c4313_AddressDetails1c4313_QRegisterMetaType()))
}

func AddressDetails_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AddressDetails1c4313_AddressDetails1c4313_QRegisterMetaType2(typeNameC)))
}

func (ptr *AddressDetails) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AddressDetails1c4313_AddressDetails1c4313_QRegisterMetaType2(typeNameC)))
}

func AddressDetails_QmlRegisterType() int {
	return int(int32(C.AddressDetails1c4313_AddressDetails1c4313_QmlRegisterType()))
}

func (ptr *AddressDetails) QmlRegisterType() int {
	return int(int32(C.AddressDetails1c4313_AddressDetails1c4313_QmlRegisterType()))
}

func AddressDetails_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.AddressDetails1c4313_AddressDetails1c4313_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *AddressDetails) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.AddressDetails1c4313_AddressDetails1c4313_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *AddressDetails) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressDetails1c4313___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressDetails) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressDetails) __children_newList() unsafe.Pointer {
	return C.AddressDetails1c4313___children_newList(ptr.Pointer())
}

func (ptr *AddressDetails) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressDetails1c4313___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressDetails) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressDetails) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.AddressDetails1c4313___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *AddressDetails) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressDetails1c4313___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressDetails) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressDetails) __findChildren_newList() unsafe.Pointer {
	return C.AddressDetails1c4313___findChildren_newList(ptr.Pointer())
}

func (ptr *AddressDetails) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressDetails1c4313___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressDetails) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressDetails) __findChildren_newList3() unsafe.Pointer {
	return C.AddressDetails1c4313___findChildren_newList3(ptr.Pointer())
}

func (ptr *AddressDetails) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressDetails1c4313___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressDetails) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressDetails) __qFindChildren_newList2() unsafe.Pointer {
	return C.AddressDetails1c4313___qFindChildren_newList2(ptr.Pointer())
}

func NewAddressDetails(parent std_core.QObject_ITF) *AddressDetails {
	tmpValue := NewAddressDetailsFromPointer(C.AddressDetails1c4313_NewAddressDetails(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackAddressDetails1c4313_DestroyAddressDetails
func callbackAddressDetails1c4313_DestroyAddressDetails(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~AddressDetails"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressDetailsFromPointer(ptr).DestroyAddressDetailsDefault()
	}
}

func (ptr *AddressDetails) ConnectDestroyAddressDetails(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~AddressDetails"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~AddressDetails", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~AddressDetails", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressDetails) DisconnectDestroyAddressDetails() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~AddressDetails")
	}
}

func (ptr *AddressDetails) DestroyAddressDetails() {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_DestroyAddressDetails(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *AddressDetails) DestroyAddressDetailsDefault() {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_DestroyAddressDetailsDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAddressDetails1c4313_ChildEvent
func callbackAddressDetails1c4313_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewAddressDetailsFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *AddressDetails) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackAddressDetails1c4313_ConnectNotify
func callbackAddressDetails1c4313_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAddressDetailsFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AddressDetails) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAddressDetails1c4313_CustomEvent
func callbackAddressDetails1c4313_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewAddressDetailsFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *AddressDetails) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackAddressDetails1c4313_DeleteLater
func callbackAddressDetails1c4313_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressDetailsFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *AddressDetails) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAddressDetails1c4313_Destroyed
func callbackAddressDetails1c4313_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackAddressDetails1c4313_DisconnectNotify
func callbackAddressDetails1c4313_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAddressDetailsFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AddressDetails) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAddressDetails1c4313_Event
func callbackAddressDetails1c4313_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressDetailsFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *AddressDetails) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressDetails1c4313_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackAddressDetails1c4313_EventFilter
func callbackAddressDetails1c4313_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressDetailsFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *AddressDetails) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressDetails1c4313_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackAddressDetails1c4313_ObjectNameChanged
func callbackAddressDetails1c4313_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackAddressDetails1c4313_TimerEvent
func callbackAddressDetails1c4313_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewAddressDetailsFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *AddressDetails) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails1c4313_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type AddressList_ITF interface {
	std_core.QAbstractListModel_ITF
	AddressList_PTR() *AddressList
}

func (ptr *AddressList) AddressList_PTR() *AddressList {
	return ptr
}

func (ptr *AddressList) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *AddressList) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromAddressList(ptr AddressList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AddressList_PTR().Pointer()
	}
	return nil
}

func NewAddressListFromPointer(ptr unsafe.Pointer) (n *AddressList) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(AddressList)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *AddressList:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &AddressList{QAbstractListModel: *deduced}

		default:
			n = new(AddressList)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *AddressList) Init() { this.init() }

//export callbackAddressList1c4313_Constructor
func callbackAddressList1c4313_Constructor(ptr unsafe.Pointer) {
	this := NewAddressListFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectAddAddress(this.addAddress)
	this.ConnectRemoveAddress(this.removeAddress)
	this.init()
}

//export callbackAddressList1c4313_AddAddress
func callbackAddressList1c4313_AddAddress(ptr unsafe.Pointer, transaction unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addAddress"); signal != nil {
		(*(*func(*AddressDetails))(signal))(NewAddressDetailsFromPointer(transaction))
	}

}

func (ptr *AddressList) ConnectAddAddress(f func(transaction *AddressDetails)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addAddress") {
			C.AddressList1c4313_ConnectAddAddress(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addAddress"); signal != nil {
			f := func(transaction *AddressDetails) {
				(*(*func(*AddressDetails))(signal))(transaction)
				f(transaction)
			}
			qt.ConnectSignal(ptr.Pointer(), "addAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressList) DisconnectAddAddress() {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_DisconnectAddAddress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addAddress")
	}
}

func (ptr *AddressList) AddAddress(transaction AddressDetails_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_AddAddress(ptr.Pointer(), PointerFromAddressDetails(transaction))
	}
}

//export callbackAddressList1c4313_RemoveAddress
func callbackAddressList1c4313_RemoveAddress(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "removeAddress"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *AddressList) ConnectRemoveAddress(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "removeAddress") {
			C.AddressList1c4313_ConnectRemoveAddress(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "removeAddress"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "removeAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressList) DisconnectRemoveAddress() {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_DisconnectRemoveAddress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "removeAddress")
	}
}

func (ptr *AddressList) RemoveAddress(index int) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_RemoveAddress(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackAddressList1c4313_Roles
func callbackAddressList1c4313_Roles(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roles"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__roles_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__roles_newList())
		for k, v := range NewAddressListFromPointer(ptr).RolesDefault() {
			tmpList.__roles_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *AddressList) ConnectRoles(f func() map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "roles"); signal != nil {
			f := func() map[int]*std_core.QByteArray {
				(*(*func() map[int]*std_core.QByteArray)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "roles", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "roles", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressList) DisconnectRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "roles")
	}
}

func (ptr *AddressList) Roles() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.AddressList1c4313_Roles(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

func (ptr *AddressList) RolesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.AddressList1c4313_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackAddressList1c4313_SetRoles
func callbackAddressList1c4313_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setRoles"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	} else {
		NewAddressListFromPointer(ptr).SetRolesDefault(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	}
}

func (ptr *AddressList) ConnectSetRoles(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRoles"); signal != nil {
			f := func(roles map[int]*std_core.QByteArray) {
				(*(*func(map[int]*std_core.QByteArray))(signal))(roles)
				f(roles)
			}
			qt.ConnectSignal(ptr.Pointer(), "setRoles", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressList) DisconnectSetRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRoles")
	}
}

func (ptr *AddressList) SetRoles(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *AddressList) SetRolesDefault(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackAddressList1c4313_RolesChanged
func callbackAddressList1c4313_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "rolesChanged"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__rolesChanged_roles_keyList() {
				out[v] = tmpList.__rolesChanged_roles_atList(v, i)
			}
			return out
		}(roles))
	}

}

func (ptr *AddressList) ConnectRolesChanged(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rolesChanged") {
			C.AddressList1c4313_ConnectRolesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rolesChanged"); signal != nil {
			f := func(roles map[int]*std_core.QByteArray) {
				(*(*func(map[int]*std_core.QByteArray))(signal))(roles)
				f(roles)
			}
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressList) DisconnectRolesChanged() {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *AddressList) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackAddressList1c4313_Addresses
func callbackAddressList1c4313_Addresses(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "addresses"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__addresses_newList())
			for _, v := range (*(*func() []*AddressDetails)(signal))() {
				tmpList.__addresses_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__addresses_newList())
		for _, v := range NewAddressListFromPointer(ptr).AddressesDefault() {
			tmpList.__addresses_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *AddressList) ConnectAddresses(f func() []*AddressDetails) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addresses"); signal != nil {
			f := func() []*AddressDetails {
				(*(*func() []*AddressDetails)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "addresses", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addresses", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressList) DisconnectAddresses() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addresses")
	}
}

func (ptr *AddressList) Addresses() []*AddressDetails {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*AddressDetails {
			out := make([]*AddressDetails, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__addresses_atList(i)
			}
			return out
		}(C.AddressList1c4313_Addresses(ptr.Pointer()))
	}
	return make([]*AddressDetails, 0)
}

func (ptr *AddressList) AddressesDefault() []*AddressDetails {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*AddressDetails {
			out := make([]*AddressDetails, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__addresses_atList(i)
			}
			return out
		}(C.AddressList1c4313_AddressesDefault(ptr.Pointer()))
	}
	return make([]*AddressDetails, 0)
}

//export callbackAddressList1c4313_SetAddresses
func callbackAddressList1c4313_SetAddresses(ptr unsafe.Pointer, addresses C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setAddresses"); signal != nil {
		(*(*func([]*AddressDetails))(signal))(func(l C.struct_Moc_PackedList) []*AddressDetails {
			out := make([]*AddressDetails, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setAddresses_addresses_atList(i)
			}
			return out
		}(addresses))
	} else {
		NewAddressListFromPointer(ptr).SetAddressesDefault(func(l C.struct_Moc_PackedList) []*AddressDetails {
			out := make([]*AddressDetails, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setAddresses_addresses_atList(i)
			}
			return out
		}(addresses))
	}
}

func (ptr *AddressList) ConnectSetAddresses(f func(addresses []*AddressDetails)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAddresses"); signal != nil {
			f := func(addresses []*AddressDetails) {
				(*(*func([]*AddressDetails))(signal))(addresses)
				f(addresses)
			}
			qt.ConnectSignal(ptr.Pointer(), "setAddresses", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAddresses", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressList) DisconnectSetAddresses() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAddresses")
	}
}

func (ptr *AddressList) SetAddresses(addresses []*AddressDetails) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_SetAddresses(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__setAddresses_addresses_newList())
			for _, v := range addresses {
				tmpList.__setAddresses_addresses_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *AddressList) SetAddressesDefault(addresses []*AddressDetails) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_SetAddressesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__setAddresses_addresses_newList())
			for _, v := range addresses {
				tmpList.__setAddresses_addresses_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackAddressList1c4313_AddressesChanged
func callbackAddressList1c4313_AddressesChanged(ptr unsafe.Pointer, addresses C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "addressesChanged"); signal != nil {
		(*(*func([]*AddressDetails))(signal))(func(l C.struct_Moc_PackedList) []*AddressDetails {
			out := make([]*AddressDetails, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__addressesChanged_addresses_atList(i)
			}
			return out
		}(addresses))
	}

}

func (ptr *AddressList) ConnectAddressesChanged(f func(addresses []*AddressDetails)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressesChanged") {
			C.AddressList1c4313_ConnectAddressesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addressesChanged"); signal != nil {
			f := func(addresses []*AddressDetails) {
				(*(*func([]*AddressDetails))(signal))(addresses)
				f(addresses)
			}
			qt.ConnectSignal(ptr.Pointer(), "addressesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressList) DisconnectAddressesChanged() {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_DisconnectAddressesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressesChanged")
	}
}

func (ptr *AddressList) AddressesChanged(addresses []*AddressDetails) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_AddressesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__addressesChanged_addresses_newList())
			for _, v := range addresses {
				tmpList.__addressesChanged_addresses_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func AddressList_QRegisterMetaType() int {
	return int(int32(C.AddressList1c4313_AddressList1c4313_QRegisterMetaType()))
}

func (ptr *AddressList) QRegisterMetaType() int {
	return int(int32(C.AddressList1c4313_AddressList1c4313_QRegisterMetaType()))
}

func AddressList_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AddressList1c4313_AddressList1c4313_QRegisterMetaType2(typeNameC)))
}

func (ptr *AddressList) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AddressList1c4313_AddressList1c4313_QRegisterMetaType2(typeNameC)))
}

func AddressList_QmlRegisterType() int {
	return int(int32(C.AddressList1c4313_AddressList1c4313_QmlRegisterType()))
}

func (ptr *AddressList) QmlRegisterType() int {
	return int(int32(C.AddressList1c4313_AddressList1c4313_QmlRegisterType()))
}

func AddressList_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.AddressList1c4313_AddressList1c4313_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *AddressList) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.AddressList1c4313_AddressList1c4313_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *AddressList) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList1c4313_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____itemData_keyList_newList() unsafe.Pointer {
	return C.AddressList1c4313_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList1c4313_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.AddressList1c4313_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList1c4313_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.AddressList1c4313_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList1c4313___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressList) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.AddressList1c4313___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *AddressList) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList1c4313___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressList) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.AddressList1c4313___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *AddressList) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList1c4313___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) __dataChanged_roles_newList() unsafe.Pointer {
	return C.AddressList1c4313___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *AddressList) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressList1c4313___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *AddressList) __itemData_newList() unsafe.Pointer {
	return C.AddressList1c4313___itemData_newList(ptr.Pointer())
}

func (ptr *AddressList) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.AddressList1c4313___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.AddressList1c4313___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *AddressList) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.AddressList1c4313___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *AddressList) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.AddressList1c4313___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *AddressList) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.AddressList1c4313___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *AddressList) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList1c4313___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressList) __match_newList() unsafe.Pointer {
	return C.AddressList1c4313___match_newList(ptr.Pointer())
}

func (ptr *AddressList) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList1c4313___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressList) __mimeData_indexes_newList() unsafe.Pointer {
	return C.AddressList1c4313___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *AddressList) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList1c4313___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressList) __persistentIndexList_newList() unsafe.Pointer {
	return C.AddressList1c4313___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *AddressList) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressList1c4313___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressList) __roleNames_newList() unsafe.Pointer {
	return C.AddressList1c4313___roleNames_newList(ptr.Pointer())
}

func (ptr *AddressList) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.AddressList1c4313___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressList1c4313___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *AddressList) __setItemData_roles_newList() unsafe.Pointer {
	return C.AddressList1c4313___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *AddressList) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.AddressList1c4313___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList1c4313_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.AddressList1c4313_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList1c4313_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.AddressList1c4313_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressList1c4313___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressList) __children_newList() unsafe.Pointer {
	return C.AddressList1c4313___children_newList(ptr.Pointer())
}

func (ptr *AddressList) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressList1c4313___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressList) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.AddressList1c4313___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *AddressList) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressList1c4313___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressList) __findChildren_newList() unsafe.Pointer {
	return C.AddressList1c4313___findChildren_newList(ptr.Pointer())
}

func (ptr *AddressList) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressList1c4313___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressList) __findChildren_newList3() unsafe.Pointer {
	return C.AddressList1c4313___findChildren_newList3(ptr.Pointer())
}

func (ptr *AddressList) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressList1c4313___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressList) __qFindChildren_newList2() unsafe.Pointer {
	return C.AddressList1c4313___qFindChildren_newList2(ptr.Pointer())
}

func (ptr *AddressList) __roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressList1c4313___roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressList) __roles_newList() unsafe.Pointer {
	return C.AddressList1c4313___roles_newList(ptr.Pointer())
}

func (ptr *AddressList) __roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roles_keyList_atList(i)
			}
			return out
		}(C.AddressList1c4313___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) __setRoles_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressList1c4313___setRoles_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressList) __setRoles_roles_newList() unsafe.Pointer {
	return C.AddressList1c4313___setRoles_roles_newList(ptr.Pointer())
}

func (ptr *AddressList) __setRoles_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setRoles_roles_keyList_atList(i)
			}
			return out
		}(C.AddressList1c4313___setRoles_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) __rolesChanged_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressList1c4313___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressList) __rolesChanged_roles_newList() unsafe.Pointer {
	return C.AddressList1c4313___rolesChanged_roles_newList(ptr.Pointer())
}

func (ptr *AddressList) __rolesChanged_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____rolesChanged_roles_keyList_atList(i)
			}
			return out
		}(C.AddressList1c4313___rolesChanged_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) __addresses_atList(i int) *AddressDetails {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressDetailsFromPointer(C.AddressList1c4313___addresses_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __addresses_setList(i AddressDetails_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___addresses_setList(ptr.Pointer(), PointerFromAddressDetails(i))
	}
}

func (ptr *AddressList) __addresses_newList() unsafe.Pointer {
	return C.AddressList1c4313___addresses_newList(ptr.Pointer())
}

func (ptr *AddressList) __setAddresses_addresses_atList(i int) *AddressDetails {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressDetailsFromPointer(C.AddressList1c4313___setAddresses_addresses_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __setAddresses_addresses_setList(i AddressDetails_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___setAddresses_addresses_setList(ptr.Pointer(), PointerFromAddressDetails(i))
	}
}

func (ptr *AddressList) __setAddresses_addresses_newList() unsafe.Pointer {
	return C.AddressList1c4313___setAddresses_addresses_newList(ptr.Pointer())
}

func (ptr *AddressList) __addressesChanged_addresses_atList(i int) *AddressDetails {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressDetailsFromPointer(C.AddressList1c4313___addressesChanged_addresses_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __addressesChanged_addresses_setList(i AddressDetails_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313___addressesChanged_addresses_setList(ptr.Pointer(), PointerFromAddressDetails(i))
	}
}

func (ptr *AddressList) __addressesChanged_addresses_newList() unsafe.Pointer {
	return C.AddressList1c4313___addressesChanged_addresses_newList(ptr.Pointer())
}

func (ptr *AddressList) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList1c4313_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____roles_keyList_newList() unsafe.Pointer {
	return C.AddressList1c4313_____roles_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) ____setRoles_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList1c4313_____setRoles_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____setRoles_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_____setRoles_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____setRoles_roles_keyList_newList() unsafe.Pointer {
	return C.AddressList1c4313_____setRoles_roles_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) ____rolesChanged_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList1c4313_____rolesChanged_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____rolesChanged_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_____rolesChanged_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____rolesChanged_roles_keyList_newList() unsafe.Pointer {
	return C.AddressList1c4313_____rolesChanged_roles_keyList_newList(ptr.Pointer())
}

func NewAddressList(parent std_core.QObject_ITF) *AddressList {
	tmpValue := NewAddressListFromPointer(C.AddressList1c4313_NewAddressList(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackAddressList1c4313_DestroyAddressList
func callbackAddressList1c4313_DestroyAddressList(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~AddressList"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressListFromPointer(ptr).DestroyAddressListDefault()
	}
}

func (ptr *AddressList) ConnectDestroyAddressList(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~AddressList"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~AddressList", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~AddressList", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressList) DisconnectDestroyAddressList() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~AddressList")
	}
}

func (ptr *AddressList) DestroyAddressList() {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_DestroyAddressList(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *AddressList) DestroyAddressListDefault() {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_DestroyAddressListDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAddressList1c4313_DropMimeData
func callbackAddressList1c4313_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList1c4313_Flags
func callbackAddressList1c4313_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewAddressListFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressList) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.AddressList1c4313_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackAddressList1c4313_Index
func callbackAddressList1c4313_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewAddressListFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *AddressList) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList1c4313_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressList1c4313_Sibling
func callbackAddressList1c4313_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewAddressListFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *AddressList) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList1c4313_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressList1c4313_Buddy
func callbackAddressList1c4313_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewAddressListFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressList) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList1c4313_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressList1c4313_CanDropMimeData
func callbackAddressList1c4313_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList1c4313_CanFetchMore
func callbackAddressList1c4313_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList1c4313_ColumnCount
func callbackAddressList1c4313_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewAddressListFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *AddressList) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList1c4313_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackAddressList1c4313_ColumnsAboutToBeInserted
func callbackAddressList1c4313_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList1c4313_ColumnsAboutToBeMoved
func callbackAddressList1c4313_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackAddressList1c4313_ColumnsAboutToBeRemoved
func callbackAddressList1c4313_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList1c4313_ColumnsInserted
func callbackAddressList1c4313_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList1c4313_ColumnsMoved
func callbackAddressList1c4313_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackAddressList1c4313_ColumnsRemoved
func callbackAddressList1c4313_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList1c4313_Data
func callbackAddressList1c4313_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(*std_core.QModelIndex, int) *std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewAddressListFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *AddressList) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressList1c4313_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackAddressList1c4313_DataChanged
func callbackAddressList1c4313_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		(*(*func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(signal))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackAddressList1c4313_FetchMore
func callbackAddressList1c4313_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*std_core.QModelIndex))(signal))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewAddressListFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *AddressList) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackAddressList1c4313_HasChildren
func callbackAddressList1c4313_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList1c4313_HeaderData
func callbackAddressList1c4313_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewAddressListFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *AddressList) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressList1c4313_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackAddressList1c4313_HeaderDataChanged
func callbackAddressList1c4313_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(std_core.Qt__Orientation, int, int))(signal))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList1c4313_InsertColumns
func callbackAddressList1c4313_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList1c4313_InsertRows
func callbackAddressList1c4313_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList1c4313_ItemData
func callbackAddressList1c4313_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__itemData_newList())
			for k, v := range (*(*func(*std_core.QModelIndex) map[int]*std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__itemData_newList())
		for k, v := range NewAddressListFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *AddressList) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.AddressList1c4313_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackAddressList1c4313_LayoutAboutToBeChanged
func callbackAddressList1c4313_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackAddressList1c4313_LayoutChanged
func callbackAddressList1c4313_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackAddressList1c4313_Match
func callbackAddressList1c4313_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__match_newList())
		for _, v := range NewAddressListFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *AddressList) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.AddressList1c4313_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackAddressList1c4313_MimeData
func callbackAddressList1c4313_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData((*(*func([]*std_core.QModelIndex) *std_core.QMimeData)(signal))(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewAddressListFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewAddressListFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *AddressList) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.AddressList1c4313_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackAddressList1c4313_MimeTypes
func callbackAddressList1c4313_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewAddressListFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *AddressList) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.AddressList1c4313_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackAddressList1c4313_ModelAboutToBeReset
func callbackAddressList1c4313_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackAddressList1c4313_ModelReset
func callbackAddressList1c4313_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackAddressList1c4313_MoveColumns
func callbackAddressList1c4313_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *AddressList) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackAddressList1c4313_MoveRows
func callbackAddressList1c4313_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *AddressList) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackAddressList1c4313_Parent
func callbackAddressList1c4313_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewAddressListFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressList) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList1c4313_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressList1c4313_RemoveColumns
func callbackAddressList1c4313_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList1c4313_RemoveRows
func callbackAddressList1c4313_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList1c4313_ResetInternalData
func callbackAddressList1c4313_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressListFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *AddressList) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackAddressList1c4313_Revert
func callbackAddressList1c4313_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressListFromPointer(ptr).RevertDefault()
	}
}

func (ptr *AddressList) RevertDefault() {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_RevertDefault(ptr.Pointer())
	}
}

//export callbackAddressList1c4313_RoleNames
func callbackAddressList1c4313_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__roleNames_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__roleNames_newList())
		for k, v := range NewAddressListFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *AddressList) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.AddressList1c4313_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackAddressList1c4313_RowCount
func callbackAddressList1c4313_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewAddressListFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *AddressList) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList1c4313_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackAddressList1c4313_RowsAboutToBeInserted
func callbackAddressList1c4313_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackAddressList1c4313_RowsAboutToBeMoved
func callbackAddressList1c4313_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackAddressList1c4313_RowsAboutToBeRemoved
func callbackAddressList1c4313_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList1c4313_RowsInserted
func callbackAddressList1c4313_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList1c4313_RowsMoved
func callbackAddressList1c4313_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackAddressList1c4313_RowsRemoved
func callbackAddressList1c4313_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList1c4313_SetData
func callbackAddressList1c4313_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(signal))(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *AddressList) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackAddressList1c4313_SetHeaderData
func callbackAddressList1c4313_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *AddressList) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackAddressList1c4313_SetItemData
func callbackAddressList1c4313_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(signal))(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewAddressListFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *AddressList) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackAddressList1c4313_Sort
func callbackAddressList1c4313_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, std_core.Qt__SortOrder))(signal))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewAddressListFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *AddressList) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackAddressList1c4313_Span
func callbackAddressList1c4313_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize((*(*func(*std_core.QModelIndex) *std_core.QSize)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewAddressListFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressList) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.AddressList1c4313_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackAddressList1c4313_Submit
func callbackAddressList1c4313_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).SubmitDefault())))
}

func (ptr *AddressList) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackAddressList1c4313_SupportedDragActions
func callbackAddressList1c4313_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewAddressListFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *AddressList) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.AddressList1c4313_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackAddressList1c4313_SupportedDropActions
func callbackAddressList1c4313_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewAddressListFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *AddressList) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.AddressList1c4313_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackAddressList1c4313_ChildEvent
func callbackAddressList1c4313_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewAddressListFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *AddressList) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackAddressList1c4313_ConnectNotify
func callbackAddressList1c4313_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAddressListFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AddressList) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAddressList1c4313_CustomEvent
func callbackAddressList1c4313_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewAddressListFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *AddressList) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackAddressList1c4313_DeleteLater
func callbackAddressList1c4313_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressListFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *AddressList) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAddressList1c4313_Destroyed
func callbackAddressList1c4313_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackAddressList1c4313_DisconnectNotify
func callbackAddressList1c4313_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAddressListFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AddressList) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAddressList1c4313_Event
func callbackAddressList1c4313_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *AddressList) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackAddressList1c4313_EventFilter
func callbackAddressList1c4313_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *AddressList) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList1c4313_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackAddressList1c4313_ObjectNameChanged
func callbackAddressList1c4313_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackAddressList1c4313_TimerEvent
func callbackAddressList1c4313_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewAddressListFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *AddressList) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList1c4313_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
