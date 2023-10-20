package test

import (
	"context"
	"strings"
	"sync"
)

// An EventSink listens for various events we are able to capture during tests,
// currently just http requests/responses.
type EventSink interface {
	AddHTTPEvent(entry *LogEntry)
}

type httpEventSinkType int

var httpEventSinkKey httpEventSinkType

// EventSinksFromContext gets the EventSink listeners attached to the passed context.
func EventSinksFromContext(ctx context.Context) []EventSink {
	v := ctx.Value(httpEventSinkKey)
	if v == nil {
		return nil
	}
	return v.([]EventSink)
}

// AddSinkToContext attaches the sinks to the returned context.
func AddSinkToContext(ctx context.Context, sinks ...EventSink) context.Context {
	var eventSinks []EventSink
	v := ctx.Value(httpEventSinkKey)
	if v != nil {
		eventSinks = v.([]EventSink)
	}
	eventSinks = append(eventSinks, sinks...)
	return context.WithValue(ctx, httpEventSinkKey, eventSinks)
}

func NewMemoryEventSink() *MemoryEventSink {
	return &MemoryEventSink{}
}

// MemoryEventSink is an EventSink that stores events in memory
type MemoryEventSink struct {
	mutex      sync.Mutex
	HTTPEvents []*LogEntry `json:"httpEvents,omitempty"`
}

func (s *MemoryEventSink) AddHTTPEvent(entry *LogEntry) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.HTTPEvents = append(s.HTTPEvents, entry)
}

func (s *MemoryEventSink) FormatHTTP() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var eventStrings []string
	for _, entry := range s.HTTPEvents {
		s := entry.FormatHTTP()
		eventStrings = append(eventStrings, s)
	}
	return strings.Join(eventStrings, "\n---\n\n")
}

func (s *MemoryEventSink) PrettifyJSON(mutators ...JSONMutator) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, entry := range s.HTTPEvents {
		entry.PrettifyJSON(mutators...)
	}
}

func (s *MemoryEventSink) RemoveHTTPResponseHeader(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, entry := range s.HTTPEvents {
		entry.Response.RemoveHeader(key)
	}
}

func NewEventSinkMux() *EventSinkMux {
	return &EventSinkMux{}
}

// EventSinkMux is an EventSink that forwards events to a switchable destination.
// For single-threaded tests, this can reduce the amount of context-forwarding we need to do.
type EventSinkMux struct {
	mutex sync.Mutex
	sink  EventSink
}

func (s *EventSinkMux) AddHTTPEvent(entry *LogEntry) {
	s.mutex.Lock()
	sink := s.sink
	s.mutex.Unlock()

	if sink != nil {
		sink.AddHTTPEvent(entry)
	}
}

func (s *EventSinkMux) SetEventSink(sink EventSink) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.sink = sink
}
