// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package playback

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/dece2183/media-winrt-go/windows/media"
)

const SignatureMediaPlayer string = "rc(Windows.Media.Playback.MediaPlayer;{381a83cb-6fff-499b-8d64-2885dfc1249e})"

type MediaPlayer struct {
	ole.IUnknown
}

func NewMediaPlayer() (*MediaPlayer, error) {
	inspectable, err := ole.RoActivateInstance("Windows.Media.Playback.MediaPlayer")
	if err != nil {
		return nil, err
	}
	return (*MediaPlayer)(unsafe.Pointer(inspectable)), nil
}

func (impl *MediaPlayer) GetSystemMediaTransportControls() (*media.SystemMediaTransportControls, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlayer2))
	defer itf.Release()
	v := (*iMediaPlayer2)(unsafe.Pointer(itf))
	return v.GetSystemMediaTransportControls()
}

func (impl *MediaPlayer) GetCommandManager() (*MediaPlaybackCommandManager, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlayer3))
	defer itf.Release()
	v := (*iMediaPlayer3)(unsafe.Pointer(itf))
	return v.GetCommandManager()
}

const GUIDiMediaPlayer string = "381a83cb-6fff-499b-8d64-2885dfc1249e"
const SignatureiMediaPlayer string = "{381a83cb-6fff-499b-8d64-2885dfc1249e}"

type iMediaPlayer struct {
	ole.IInspectable
}

type iMediaPlayerVtbl struct {
	ole.IInspectableVtbl

	GetAutoPlay                      uintptr
	SetAutoPlay                      uintptr
	GetNaturalDuration               uintptr
	GetPosition                      uintptr
	SetPosition                      uintptr
	GetBufferingProgress             uintptr
	GetCurrentState                  uintptr
	GetCanSeek                       uintptr
	GetCanPause                      uintptr
	GetIsLoopingEnabled              uintptr
	SetIsLoopingEnabled              uintptr
	GetIsProtected                   uintptr
	GetIsMuted                       uintptr
	SetIsMuted                       uintptr
	GetPlaybackRate                  uintptr
	SetPlaybackRate                  uintptr
	GetVolume                        uintptr
	SetVolume                        uintptr
	GetPlaybackMediaMarkers          uintptr
	AddMediaOpened                   uintptr
	RemoveMediaOpened                uintptr
	AddMediaEnded                    uintptr
	RemoveMediaEnded                 uintptr
	AddMediaFailed                   uintptr
	RemoveMediaFailed                uintptr
	AddCurrentStateChanged           uintptr
	RemoveCurrentStateChanged        uintptr
	AddPlaybackMediaMarkerReached    uintptr
	RemovePlaybackMediaMarkerReached uintptr
	AddMediaPlayerRateChanged        uintptr
	RemoveMediaPlayerRateChanged     uintptr
	AddVolumeChanged                 uintptr
	RemoveVolumeChanged              uintptr
	AddSeekCompleted                 uintptr
	RemoveSeekCompleted              uintptr
	AddBufferingStarted              uintptr
	RemoveBufferingStarted           uintptr
	AddBufferingEnded                uintptr
	RemoveBufferingEnded             uintptr
	Play                             uintptr
	Pause                            uintptr
	SetUriSource                     uintptr
}

func (v *iMediaPlayer) VTable() *iMediaPlayerVtbl {
	return (*iMediaPlayerVtbl)(unsafe.Pointer(v.RawVTable))
}

const GUIDiMediaPlayerSource string = "bd4f8897-1423-4c3e-82c5-0fb1af94f715"
const SignatureiMediaPlayerSource string = "{bd4f8897-1423-4c3e-82c5-0fb1af94f715}"

type iMediaPlayerSource struct {
	ole.IInspectable
}

type iMediaPlayerSourceVtbl struct {
	ole.IInspectableVtbl

	GetProtectionManager uintptr
	SetProtectionManager uintptr
	SetFileSource        uintptr
	SetStreamSource      uintptr
	SetMediaSource       uintptr
}

func (v *iMediaPlayerSource) VTable() *iMediaPlayerSourceVtbl {
	return (*iMediaPlayerSourceVtbl)(unsafe.Pointer(v.RawVTable))
}

const GUIDiMediaPlayerSource2 string = "82449b9f-7322-4c0b-b03b-3e69a48260c5"
const SignatureiMediaPlayerSource2 string = "{82449b9f-7322-4c0b-b03b-3e69a48260c5}"

type iMediaPlayerSource2 struct {
	ole.IInspectable
}

type iMediaPlayerSource2Vtbl struct {
	ole.IInspectableVtbl

	GetSource uintptr
	SetSource uintptr
}

func (v *iMediaPlayerSource2) VTable() *iMediaPlayerSource2Vtbl {
	return (*iMediaPlayerSource2Vtbl)(unsafe.Pointer(v.RawVTable))
}

const GUIDiMediaPlayer2 string = "3c841218-2123-4fc5-9082-2f883f77bdf5"
const SignatureiMediaPlayer2 string = "{3c841218-2123-4fc5-9082-2f883f77bdf5}"

type iMediaPlayer2 struct {
	ole.IInspectable
}

type iMediaPlayer2Vtbl struct {
	ole.IInspectableVtbl

	GetSystemMediaTransportControls uintptr
	GetAudioCategory                uintptr
	SetAudioCategory                uintptr
	GetAudioDeviceType              uintptr
	SetAudioDeviceType              uintptr
}

func (v *iMediaPlayer2) VTable() *iMediaPlayer2Vtbl {
	return (*iMediaPlayer2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iMediaPlayer2) GetSystemMediaTransportControls() (*media.SystemMediaTransportControls, error) {
	var out *media.SystemMediaTransportControls
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetSystemMediaTransportControls,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out media.SystemMediaTransportControls
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

const GUIDiMediaPlayerEffects string = "85a1deda-cab6-4cc0-8be3-6035f4de2591"
const SignatureiMediaPlayerEffects string = "{85a1deda-cab6-4cc0-8be3-6035f4de2591}"

type iMediaPlayerEffects struct {
	ole.IInspectable
}

type iMediaPlayerEffectsVtbl struct {
	ole.IInspectableVtbl

	AddAudioEffectWithSettings uintptr
	RemoveAllEffects           uintptr
}

func (v *iMediaPlayerEffects) VTable() *iMediaPlayerEffectsVtbl {
	return (*iMediaPlayerEffectsVtbl)(unsafe.Pointer(v.RawVTable))
}

const GUIDiMediaPlayer3 string = "ee0660da-031b-4feb-bd9b-92e0a0a8d299"
const SignatureiMediaPlayer3 string = "{ee0660da-031b-4feb-bd9b-92e0a0a8d299}"

type iMediaPlayer3 struct {
	ole.IInspectable
}

type iMediaPlayer3Vtbl struct {
	ole.IInspectableVtbl

	AddIsMutedChanged                   uintptr
	RemoveIsMutedChanged                uintptr
	AddSourceChanged                    uintptr
	RemoveSourceChanged                 uintptr
	GetAudioBalance                     uintptr
	SetAudioBalance                     uintptr
	GetRealTimePlayback                 uintptr
	SetRealTimePlayback                 uintptr
	GetStereoscopicVideoRenderMode      uintptr
	SetStereoscopicVideoRenderMode      uintptr
	GetBreakManager                     uintptr
	GetCommandManager                   uintptr
	GetAudioDevice                      uintptr
	SetAudioDevice                      uintptr
	GetTimelineController               uintptr
	SetTimelineController               uintptr
	GetTimelineControllerPositionOffset uintptr
	SetTimelineControllerPositionOffset uintptr
	GetPlaybackSession                  uintptr
	StepForwardOneFrame                 uintptr
	StepBackwardOneFrame                uintptr
	GetAsCastingSource                  uintptr
}

func (v *iMediaPlayer3) VTable() *iMediaPlayer3Vtbl {
	return (*iMediaPlayer3Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iMediaPlayer3) GetCommandManager() (*MediaPlaybackCommandManager, error) {
	var out *MediaPlaybackCommandManager
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetCommandManager,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out MediaPlaybackCommandManager
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

const GUIDiMediaPlayer4 string = "80035db0-7448-4770-afcf-2a57450914c5"
const SignatureiMediaPlayer4 string = "{80035db0-7448-4770-afcf-2a57450914c5}"

type iMediaPlayer4 struct {
	ole.IInspectable
}

type iMediaPlayer4Vtbl struct {
	ole.IInspectableVtbl

	SetSurfaceSize uintptr
	GetSurface     uintptr
}

func (v *iMediaPlayer4) VTable() *iMediaPlayer4Vtbl {
	return (*iMediaPlayer4Vtbl)(unsafe.Pointer(v.RawVTable))
}

const GUIDiMediaPlayerEffects2 string = "fa419a79-1bbe-46c5-ae1f-8ee69fb3c2c7"
const SignatureiMediaPlayerEffects2 string = "{fa419a79-1bbe-46c5-ae1f-8ee69fb3c2c7}"

type iMediaPlayerEffects2 struct {
	ole.IInspectable
}

type iMediaPlayerEffects2Vtbl struct {
	ole.IInspectableVtbl

	AddVideoEffectWithSettings uintptr
}

func (v *iMediaPlayerEffects2) VTable() *iMediaPlayerEffects2Vtbl {
	return (*iMediaPlayerEffects2Vtbl)(unsafe.Pointer(v.RawVTable))
}

const GUIDiMediaPlayer5 string = "cfe537fd-f86a-4446-bf4d-c8e792b7b4b3"
const SignatureiMediaPlayer5 string = "{cfe537fd-f86a-4446-bf4d-c8e792b7b4b3}"

type iMediaPlayer5 struct {
	ole.IInspectable
}

type iMediaPlayer5Vtbl struct {
	ole.IInspectableVtbl

	AddVideoFrameAvailable                     uintptr
	RemoveVideoFrameAvailable                  uintptr
	GetIsVideoFrameServerEnabled               uintptr
	SetIsVideoFrameServerEnabled               uintptr
	CopyFrameToVideoSurface                    uintptr
	CopyFrameToVideoSurfaceWithTargetRectangle uintptr
	CopyFrameToStereoscopicVideoSurfaces       uintptr
}

func (v *iMediaPlayer5) VTable() *iMediaPlayer5Vtbl {
	return (*iMediaPlayer5Vtbl)(unsafe.Pointer(v.RawVTable))
}

const GUIDiMediaPlayer6 string = "e0caa086-ae65-414c-b010-8bc55f00e692"
const SignatureiMediaPlayer6 string = "{e0caa086-ae65-414c-b010-8bc55f00e692}"

type iMediaPlayer6 struct {
	ole.IInspectable
}

type iMediaPlayer6Vtbl struct {
	ole.IInspectableVtbl

	AddSubtitleFrameChanged                     uintptr
	RemoveSubtitleFrameChanged                  uintptr
	RenderSubtitlesToSurface                    uintptr
	RenderSubtitlesToSurfaceWithTargetRectangle uintptr
}

func (v *iMediaPlayer6) VTable() *iMediaPlayer6Vtbl {
	return (*iMediaPlayer6Vtbl)(unsafe.Pointer(v.RawVTable))
}

const GUIDiMediaPlayer7 string = "5d1dc478-4500-4531-b3f4-777a71491f7f"
const SignatureiMediaPlayer7 string = "{5d1dc478-4500-4531-b3f4-777a71491f7f}"

type iMediaPlayer7 struct {
	ole.IInspectable
}

type iMediaPlayer7Vtbl struct {
	ole.IInspectableVtbl

	GetAudioStateMonitor uintptr
}

func (v *iMediaPlayer7) VTable() *iMediaPlayer7Vtbl {
	return (*iMediaPlayer7Vtbl)(unsafe.Pointer(v.RawVTable))
}
