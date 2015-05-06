package ledge

import (
	"bytes"
	"fmt"
	"sync/atomic"
	"time"
)

type fakeLogger struct {
	Logger
	BlockingEntryReader
	*fakeTimer
}

func newFakeLogger(
	specification *Specification,
) (*fakeLogger, error) {
	buffer := bytes.NewBuffer(nil)
	unmarshaller, err := NewJSONUnmarshaller(specification)
	if err != nil {
		return nil, err
	}
	entryReader, err := NewEntryReader(
		buffer,
		unmarshaller,
		RPCDecoder,
		EntryReaderOptions{},
	)
	if err != nil {
		return nil, err
	}
	fakeIDAllocator := newFakeIDAllocator()
	fakeTimer := newFakeTimer(0)
	logger, err := NewLogger(
		buffer,
		JSONMarshaller,
		specification,
		LoggerOptions{
			IDAllocator: fakeIDAllocator,
			Timer:       fakeTimer,
			Encoder:     RPCEncoder,
		},
	)
	if err != nil {
		return nil, err
	}
	return &fakeLogger{
		logger,
		NewBlockingEntryReader(entryReader),
		fakeTimer,
	}, nil
}

func (f *fakeLogger) CheckEntriesEqual(
	expected []*Entry,
	checkID bool,
	checkTime bool,
) error {
	entries, err := f.Entries()
	if err != nil {
		return err
	}
	if len(expected) != len(entries) {
		return fmt.Errorf("ledge: expected %d entries, got %d", len(expected), len(entries))
	}
	for i, elem := range expected {
		if !entriesEqual(elem, entries[i], checkID, checkTime) {
			return fmt.Errorf("ledge: expected %+v, got %+v", elem, entries[i])
		}
	}
	return nil
}

type fakeIDAllocator struct {
	value int64
}

func newFakeIDAllocator() *fakeIDAllocator {
	return &fakeIDAllocator{
		-1,
	}
}

func (ti *fakeIDAllocator) Allocate() string {
	return fmt.Sprintf("%d", atomic.AddInt64(&ti.value, 1))
}

type fakeTimer struct {
	now int64
}

func newFakeTimer(
	initialTimeUnixSec int64,
) *fakeTimer {
	return &fakeTimer{
		initialTimeUnixSec,
	}
}

func (tt *fakeTimer) AddTimeSec(delta int64) {
	tt.now += delta
}

func (tt *fakeTimer) Now() time.Time {
	return time.Unix(tt.now, 0)
}