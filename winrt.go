package winrt

// common
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.IClosable
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.IAsyncOperation`1
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.AsyncOperationCompletedHandler`1
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.AsyncStatus
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.DateTime
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.TimeSpan
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.Deferral
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.DeferralCompletedHandler
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.IReference`1

// event
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.TypedEventHandler`2
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -debug -class Windows.Foundation.EventRegistrationToken

// storage file
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -class Windows.Storage.StorageFile -method-filter GetFileFromPathAsync -method-filter !*

// media
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -class Windows.Media.MediaPlaybackType
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -class Windows.Media.MediaPlaybackStatus
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -class Windows.Media.PlaybackPositionChangeRequestedEventArgs
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -class Windows.Media.SystemMediaTransportControlsButton
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -class Windows.Media.SystemMediaTransportControlsButtonPressedEventArgs
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -class Windows.Media.SystemMediaTransportControlsTimelineProperties
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -class Windows.Media.SystemMediaTransportControlsDisplayUpdater -method-filter CopyFromFileAsync -method-filter Update -method-filter !*
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -class Windows.Media.SystemMediaTransportControls -method-filter UpdateTimelineProperties -method-filter put_IsEnabled -method-filter put_IsPlayEnabled -method-filter put_IsPauseEnabled -method-filter put_IsPreviousEnabled -method-filter put_IsNextEnabled -method-filter put_IsRewindEnabled -method-filter put_PlaybackStatus -method-filter get_DisplayUpdater -method-filter add_ButtonPressed -method-filter remove_ButtonPressed -method-filter add_PlaybackPositionChangeRequested -method-filter remove_PlaybackPositionChangeRequested -method-filter !*

// media playback
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -class Windows.Media.Playback.MediaPlayer -method-filter Close -method-filter get_SystemMediaTransportControls -method-filter get_CommandManager -method-filter !*
//go:generate go run github.com/dece2183/media-winrt-go/cmd/winrt-go-gen -class Windows.Media.Playback.MediaPlaybackCommandManager -method-filter put_IsEnabled -method-filter !*
