// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package media

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const SignatureSystemMediaTransportControls string = "rc(Windows.Media.SystemMediaTransportControls;{99fa3ff4-1742-42a6-902e-087d41f965ec})"

type SystemMediaTransportControls struct {
	ole.IUnknown
}

func (impl *SystemMediaTransportControls) SetIsEnabled(value bool) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControls))
	defer itf.Release()
	v := (*iSystemMediaTransportControls)(unsafe.Pointer(itf))
	return v.SetIsEnabled(value)
}

func (impl *SystemMediaTransportControls) SetIsPlayEnabled(value bool) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControls))
	defer itf.Release()
	v := (*iSystemMediaTransportControls)(unsafe.Pointer(itf))
	return v.SetIsPlayEnabled(value)
}

func (impl *SystemMediaTransportControls) SetIsPauseEnabled(value bool) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControls))
	defer itf.Release()
	v := (*iSystemMediaTransportControls)(unsafe.Pointer(itf))
	return v.SetIsPauseEnabled(value)
}

func (impl *SystemMediaTransportControls) SetIsRewindEnabled(value bool) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControls))
	defer itf.Release()
	v := (*iSystemMediaTransportControls)(unsafe.Pointer(itf))
	return v.SetIsRewindEnabled(value)
}

func (impl *SystemMediaTransportControls) SetIsPreviousEnabled(value bool) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControls))
	defer itf.Release()
	v := (*iSystemMediaTransportControls)(unsafe.Pointer(itf))
	return v.SetIsPreviousEnabled(value)
}

func (impl *SystemMediaTransportControls) SetIsNextEnabled(value bool) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControls))
	defer itf.Release()
	v := (*iSystemMediaTransportControls)(unsafe.Pointer(itf))
	return v.SetIsNextEnabled(value)
}

const GUIDiSystemMediaTransportControls string = "99fa3ff4-1742-42a6-902e-087d41f965ec"
const SignatureiSystemMediaTransportControls string = "{99fa3ff4-1742-42a6-902e-087d41f965ec}"

type iSystemMediaTransportControls struct {
	ole.IInspectable
}

type iSystemMediaTransportControlsVtbl struct {
	ole.IInspectableVtbl

	GetPlaybackStatus       uintptr
	SetPlaybackStatus       uintptr
	GetDisplayUpdater       uintptr
	GetSoundLevel           uintptr
	GetIsEnabled            uintptr
	SetIsEnabled            uintptr
	GetIsPlayEnabled        uintptr
	SetIsPlayEnabled        uintptr
	GetIsStopEnabled        uintptr
	SetIsStopEnabled        uintptr
	GetIsPauseEnabled       uintptr
	SetIsPauseEnabled       uintptr
	GetIsRecordEnabled      uintptr
	SetIsRecordEnabled      uintptr
	GetIsFastForwardEnabled uintptr
	SetIsFastForwardEnabled uintptr
	GetIsRewindEnabled      uintptr
	SetIsRewindEnabled      uintptr
	GetIsPreviousEnabled    uintptr
	SetIsPreviousEnabled    uintptr
	GetIsNextEnabled        uintptr
	SetIsNextEnabled        uintptr
	GetIsChannelUpEnabled   uintptr
	SetIsChannelUpEnabled   uintptr
	GetIsChannelDownEnabled uintptr
	SetIsChannelDownEnabled uintptr
	AddButtonPressed        uintptr
	RemoveButtonPressed     uintptr
	AddPropertyChanged      uintptr
	RemovePropertyChanged   uintptr
}

func (v *iSystemMediaTransportControls) VTable() *iSystemMediaTransportControlsVtbl {
	return (*iSystemMediaTransportControlsVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iSystemMediaTransportControls) SetIsEnabled(value bool) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetIsEnabled,
		uintptr(unsafe.Pointer(v)),                // this
		uintptr(*(*byte)(unsafe.Pointer(&value))), // in bool
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iSystemMediaTransportControls) SetIsPlayEnabled(value bool) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetIsPlayEnabled,
		uintptr(unsafe.Pointer(v)),                // this
		uintptr(*(*byte)(unsafe.Pointer(&value))), // in bool
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iSystemMediaTransportControls) SetIsPauseEnabled(value bool) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetIsPauseEnabled,
		uintptr(unsafe.Pointer(v)),                // this
		uintptr(*(*byte)(unsafe.Pointer(&value))), // in bool
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iSystemMediaTransportControls) SetIsRewindEnabled(value bool) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetIsRewindEnabled,
		uintptr(unsafe.Pointer(v)),                // this
		uintptr(*(*byte)(unsafe.Pointer(&value))), // in bool
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iSystemMediaTransportControls) SetIsPreviousEnabled(value bool) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetIsPreviousEnabled,
		uintptr(unsafe.Pointer(v)),                // this
		uintptr(*(*byte)(unsafe.Pointer(&value))), // in bool
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iSystemMediaTransportControls) SetIsNextEnabled(value bool) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetIsNextEnabled,
		uintptr(unsafe.Pointer(v)),                // this
		uintptr(*(*byte)(unsafe.Pointer(&value))), // in bool
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

const GUIDiSystemMediaTransportControls2 string = "ea98d2f6-7f3c-4af2-a586-72889808efb1"
const SignatureiSystemMediaTransportControls2 string = "{ea98d2f6-7f3c-4af2-a586-72889808efb1}"

type iSystemMediaTransportControls2 struct {
	ole.IInspectable
}

type iSystemMediaTransportControls2Vtbl struct {
	ole.IInspectableVtbl

	GetAutoRepeatMode                     uintptr
	SetAutoRepeatMode                     uintptr
	GetShuffleEnabled                     uintptr
	SetShuffleEnabled                     uintptr
	GetPlaybackRate                       uintptr
	SetPlaybackRate                       uintptr
	UpdateTimelineProperties              uintptr
	AddPlaybackPositionChangeRequested    uintptr
	RemovePlaybackPositionChangeRequested uintptr
	AddPlaybackRateChangeRequested        uintptr
	RemovePlaybackRateChangeRequested     uintptr
	AddShuffleEnabledChangeRequested      uintptr
	RemoveShuffleEnabledChangeRequested   uintptr
	AddAutoRepeatModeChangeRequested      uintptr
	RemoveAutoRepeatModeChangeRequested   uintptr
}

func (v *iSystemMediaTransportControls2) VTable() *iSystemMediaTransportControls2Vtbl {
	return (*iSystemMediaTransportControls2Vtbl)(unsafe.Pointer(v.RawVTable))
}
