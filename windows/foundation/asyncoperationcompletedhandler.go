// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package foundation

import (
	"sync"
	"time"
	"unsafe"

	"github.com/dece2183/media-winrt-go/internal/delegate"
	"github.com/dece2183/media-winrt-go/internal/kernel32"
	"github.com/go-ole/go-ole"
)

const GUIDAsyncOperationCompletedHandler string = "fcdcf02c-e5d8-4478-915a-4d90b74b83a5"
const SignatureAsyncOperationCompletedHandler string = "delegate({fcdcf02c-e5d8-4478-915a-4d90b74b83a5})"

type AsyncOperationCompletedHandler struct {
	ole.IUnknown
	sync.Mutex
	refs uintptr
	IID  ole.GUID
}

type AsyncOperationCompletedHandlerVtbl struct {
	ole.IUnknownVtbl
	Invoke uintptr
}

type AsyncOperationCompletedHandlerCallback func(instance *AsyncOperationCompletedHandler, asyncInfo *IAsyncOperation, asyncStatus AsyncStatus)

var callbacksAsyncOperationCompletedHandler = &asyncOperationCompletedHandlerCallbacks{
	mu:        &sync.Mutex{},
	callbacks: make(map[unsafe.Pointer]AsyncOperationCompletedHandlerCallback),
}

var releaseChannelsAsyncOperationCompletedHandler = &asyncOperationCompletedHandlerReleaseChannels{
	mu:    &sync.Mutex{},
	chans: make(map[unsafe.Pointer]chan struct{}),
}

func NewAsyncOperationCompletedHandler(iid *ole.GUID, callback AsyncOperationCompletedHandlerCallback) *AsyncOperationCompletedHandler {
	// create type instance
	size := unsafe.Sizeof(*(*AsyncOperationCompletedHandler)(nil))
	instPtr := kernel32.Malloc(size)
	inst := (*AsyncOperationCompletedHandler)(instPtr)

	// get the callbacks for the VTable
	callbacks := delegate.RegisterCallbacks(instPtr, inst)

	// the VTable should also be allocated in the heap
	sizeVTable := unsafe.Sizeof(*(*AsyncOperationCompletedHandlerVtbl)(nil))
	vTablePtr := kernel32.Malloc(sizeVTable)

	inst.RawVTable = (*interface{})(vTablePtr)

	vTable := (*AsyncOperationCompletedHandlerVtbl)(vTablePtr)
	vTable.IUnknownVtbl = ole.IUnknownVtbl{
		QueryInterface: callbacks.QueryInterface,
		AddRef:         callbacks.AddRef,
		Release:        callbacks.Release,
	}
	vTable.Invoke = callbacks.Invoke

	// Initialize all properties: the malloc may contain garbage
	inst.IID = *iid // copy contents
	inst.Mutex = sync.Mutex{}
	inst.refs = 0

	callbacksAsyncOperationCompletedHandler.add(unsafe.Pointer(inst), callback)

	// See the docs in the releaseChannelsAsyncOperationCompletedHandler struct
	releaseChannelsAsyncOperationCompletedHandler.acquire(unsafe.Pointer(inst))

	inst.addRef()
	return inst
}

func (r *AsyncOperationCompletedHandler) GetIID() *ole.GUID {
	return &r.IID
}

// addRef increments the reference counter by one
func (r *AsyncOperationCompletedHandler) addRef() uintptr {
	r.Lock()
	defer r.Unlock()
	r.refs++
	return r.refs
}

// removeRef decrements the reference counter by one. If it was already zero, it will just return zero.
func (r *AsyncOperationCompletedHandler) removeRef() uintptr {
	r.Lock()
	defer r.Unlock()

	if r.refs > 0 {
		r.refs--
	}

	return r.refs
}

func (instance *AsyncOperationCompletedHandler) Invoke(instancePtr, rawArgs0, rawArgs1, rawArgs2, rawArgs3, rawArgs4, rawArgs5, rawArgs6, rawArgs7, rawArgs8 unsafe.Pointer) uintptr {
	asyncInfoPtr := rawArgs0
	asyncStatusRaw := (int32)(uintptr(rawArgs1))

	// See the quote above.
	asyncInfo := (*IAsyncOperation)(asyncInfoPtr)
	asyncStatus := (AsyncStatus)(asyncStatusRaw)
	if callback, ok := callbacksAsyncOperationCompletedHandler.get(instancePtr); ok {
		callback(instance, asyncInfo, asyncStatus)
	}
	return ole.S_OK
}

func (instance *AsyncOperationCompletedHandler) AddRef() uintptr {
	return instance.addRef()
}

func (instance *AsyncOperationCompletedHandler) Release() uintptr {
	rem := instance.removeRef()
	if rem == 0 {
		// We're done.
		instancePtr := unsafe.Pointer(instance)
		callbacksAsyncOperationCompletedHandler.delete(instancePtr)

		// stop release channels used to avoid
		// https://github.com/golang/go/issues/55015
		releaseChannelsAsyncOperationCompletedHandler.release(instancePtr)

		kernel32.Free(unsafe.Pointer(instance.RawVTable))
		kernel32.Free(instancePtr)
	}
	return rem
}

type asyncOperationCompletedHandlerCallbacks struct {
	mu        *sync.Mutex
	callbacks map[unsafe.Pointer]AsyncOperationCompletedHandlerCallback
}

func (m *asyncOperationCompletedHandlerCallbacks) add(p unsafe.Pointer, v AsyncOperationCompletedHandlerCallback) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.callbacks[p] = v
}

func (m *asyncOperationCompletedHandlerCallbacks) get(p unsafe.Pointer) (AsyncOperationCompletedHandlerCallback, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, ok := m.callbacks[p]
	return v, ok
}

func (m *asyncOperationCompletedHandlerCallbacks) delete(p unsafe.Pointer) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.callbacks, p)
}

// typedEventHandlerReleaseChannels keeps a map with channels
// used to keep a goroutine alive during the lifecycle of this object.
// This is required to avoid causing a deadlock error.
// See this: https://github.com/golang/go/issues/55015
type asyncOperationCompletedHandlerReleaseChannels struct {
	mu    *sync.Mutex
	chans map[unsafe.Pointer]chan struct{}
}

func (m *asyncOperationCompletedHandlerReleaseChannels) acquire(p unsafe.Pointer) {
	m.mu.Lock()
	defer m.mu.Unlock()

	c := make(chan struct{})
	m.chans[p] = c

	go func() {
		// we need a timer to trick the go runtime into
		// thinking there's still something going on here
		// but we are only really interested in <-c
		t := time.NewTimer(time.Minute)
		for {
			select {
			case <-t.C:
				t.Reset(time.Minute)
			case <-c:
				t.Stop()
				return
			}
		}
	}()
}

func (m *asyncOperationCompletedHandlerReleaseChannels) release(p unsafe.Pointer) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if c, ok := m.chans[p]; ok {
		close(c)
		delete(m.chans, p)
	}
}
