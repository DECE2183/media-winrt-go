// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package genericattributeprofile

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
)

const SignatureGattClientNotificationResult string = "rc(Windows.Devices.Bluetooth.GenericAttributeProfile.GattClientNotificationResult;{506d5599-0112-419a-8e3b-ae21afabd2c2})"

type GattClientNotificationResult struct {
	ole.IUnknown
}

func (impl *GattClientNotificationResult) GetSubscribedClient() (*GattSubscribedClient, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattClientNotificationResult))
	defer itf.Release()
	v := (*iGattClientNotificationResult)(unsafe.Pointer(itf))
	return v.GetSubscribedClient()
}

func (impl *GattClientNotificationResult) GetStatus() (GattCommunicationStatus, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattClientNotificationResult))
	defer itf.Release()
	v := (*iGattClientNotificationResult)(unsafe.Pointer(itf))
	return v.GetStatus()
}

func (impl *GattClientNotificationResult) GetProtocolError() (*foundation.IReference, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattClientNotificationResult))
	defer itf.Release()
	v := (*iGattClientNotificationResult)(unsafe.Pointer(itf))
	return v.GetProtocolError()
}

func (impl *GattClientNotificationResult) GetBytesSent() (uint16, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattClientNotificationResult2))
	defer itf.Release()
	v := (*iGattClientNotificationResult2)(unsafe.Pointer(itf))
	return v.GetBytesSent()
}

const GUIDiGattClientNotificationResult string = "506d5599-0112-419a-8e3b-ae21afabd2c2"
const SignatureiGattClientNotificationResult string = "{506d5599-0112-419a-8e3b-ae21afabd2c2}"

type iGattClientNotificationResult struct {
	ole.IInspectable
}

type iGattClientNotificationResultVtbl struct {
	ole.IInspectableVtbl

	GetSubscribedClient uintptr
	GetStatus           uintptr
	GetProtocolError    uintptr
}

func (v *iGattClientNotificationResult) VTable() *iGattClientNotificationResultVtbl {
	return (*iGattClientNotificationResultVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iGattClientNotificationResult) GetSubscribedClient() (*GattSubscribedClient, error) {
	var out *GattSubscribedClient
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetSubscribedClient,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out GattSubscribedClient
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattClientNotificationResult) GetStatus() (GattCommunicationStatus, error) {
	var out GattCommunicationStatus
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetStatus,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out GattCommunicationStatus
	)

	if hr != 0 {
		return GattCommunicationStatusSuccess, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattClientNotificationResult) GetProtocolError() (*foundation.IReference, error) {
	var out *foundation.IReference
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetProtocolError,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.IReference
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

const GUIDiGattClientNotificationResult2 string = "8faec497-45e0-497e-9582-29a1fe281ad5"
const SignatureiGattClientNotificationResult2 string = "{8faec497-45e0-497e-9582-29a1fe281ad5}"

type iGattClientNotificationResult2 struct {
	ole.IInspectable
}

type iGattClientNotificationResult2Vtbl struct {
	ole.IInspectableVtbl

	GetBytesSent uintptr
}

func (v *iGattClientNotificationResult2) VTable() *iGattClientNotificationResult2Vtbl {
	return (*iGattClientNotificationResult2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iGattClientNotificationResult2) GetBytesSent() (uint16, error) {
	var out uint16
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetBytesSent,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint16
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}
