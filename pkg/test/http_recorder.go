// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode"

	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp,omitempty"`
	Request   Request   `json:"request,omitempty"`
	Response  Response  `json:"response,omitempty"`
	Error     string    `json:"error,omitempty"`
}

type Request struct {
	Method string      `json:"method,omitempty"`
	URL    string      `json:"url,omitempty"`
	Header http.Header `json:"header,omitempty"`
	Body   string      `json:"body,omitempty"`
}

type Response struct {
	Status     string      `json:"status,omitempty"`
	StatusCode int         `json:"statusCode,omitempty"`
	Header     http.Header `json:"header,omitempty"`
	Body       string      `json:"body,omitempty"`
}

type HTTPRecorder struct {
	outputDir string
	inner     http.RoundTripper

	// mutex to avoid concurrent writes to the same file
	mutex sync.Mutex

	eventSinks []EventSink
}

func NewHTTPRecorder(inner http.RoundTripper, outputDir string, eventSinks []EventSink) *HTTPRecorder {
	rt := &HTTPRecorder{outputDir: outputDir, inner: inner, eventSinks: eventSinks}
	return rt
}

func (r *HTTPRecorder) RoundTrip(req *http.Request) (*http.Response, error) {
	var entry LogEntry
	entry.Timestamp = time.Now()
	entry.Request.Method = req.Method
	entry.Request.URL = req.URL.String()

	entry.Request.Header = make(http.Header)
	for k, values := range req.Header {
		switch strings.ToLower(k) {
		case "authorization":
			entry.Request.Header[k] = []string{"(removed)"}
		default:
			entry.Request.Header[k] = values
		}
	}

	if req.Body != nil {
		requestBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic("failed to read request body")
		}
		entry.Request.Body = string(requestBody)
		req.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
	}

	response, err := r.inner.RoundTrip(req)

	if err != nil {
		entry.Error = fmt.Sprintf("%v", err)
	}

	if recordErr := r.record(&entry, req, response); recordErr != nil {
		klog.Warningf("failed to record HTTP request: %v", recordErr)
	}

	return response, err
}

func (r *HTTPRecorder) record(entry *LogEntry, req *http.Request, resp *http.Response) error {
	if resp != nil {
		entry.Response.Status = resp.Status
		entry.Response.StatusCode = resp.StatusCode

		entry.Response.Header = make(http.Header)
		for k, values := range resp.Header {
			switch strings.ToLower(k) {
			case "authorization":
				entry.Response.Header[k] = []string{"(removed)"}
			default:
				entry.Response.Header[k] = values
			}
		}

		if resp.Body != nil {
			requestBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic("failed to read response body")
			}
			entry.Response.Body = string(requestBody)
			resp.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
		}
	}

	ctx := req.Context()

	// If we have event sink(s), write to that sink also
	{
		for _, eventSink := range r.eventSinks {
			eventSink.AddHTTPEvent(entry)
		}
	}

	// Write to a log file
	// (We could consider moving this to a log sink also)
	{
		t := TestFromContext(ctx)
		testName := "unknown"
		if t != nil {
			testName = t.Name()
		}
		dirName := sanitizePath(testName)
		p := filepath.Join(r.outputDir, dirName, "requests.log")

		b, err := yaml.Marshal(entry)
		if err != nil {
			return fmt.Errorf("failed to marshal data: %w", err)
		}

		// Just in case we are writing to the same file concurrently
		r.mutex.Lock()
		defer r.mutex.Unlock()

		if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
			return fmt.Errorf("failed to create directory %q: %w", filepath.Dir(p), err)
		}
		f, err := os.OpenFile(p, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			return fmt.Errorf("failed to open file %q: %w", p, err)
		}
		defer f.Close()

		if _, err := f.Write(b); err != nil {
			return fmt.Errorf("failed to write to file %q: %w", p, err)
		}
		delimeter := "\n\n---\n\n"
		if _, err := f.Write([]byte(delimeter)); err != nil {
			return fmt.Errorf("failed to write to file %q: %w", p, err)
		}
	}

	return nil
}

func sanitizePath(s string) string {
	var out strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			out.WriteRune(r)
		} else {
			out.WriteRune('_')
		}
	}
	return out.String()
}

func (e *LogEntry) FormatHTTP() string {
	var b strings.Builder
	b.WriteString(e.Request.FormatHTTP())
	b.WriteString(e.Response.FormatHTTP())
	return b.String()
}

func (r *Request) FormatHTTP() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s %s\n", r.Method, r.URL))
	var keys []string
	for k := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		for _, v := range r.Header[k] {
			b.WriteString(fmt.Sprintf("%s: %s\n", k, v))
		}
	}
	b.WriteString("\n")
	if r.Body != "" {
		b.WriteString(r.Body)
		b.WriteString("\n\n")
	}
	return b.String()
}

func (r *Response) FormatHTTP() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s\n", r.Status))
	var keys []string
	for k := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		for _, v := range r.Header[k] {
			b.WriteString(fmt.Sprintf("%s: %s\n", k, v))
		}
	}
	b.WriteString("\n")
	if r.Body != "" {
		b.WriteString(r.Body)
		b.WriteString("\n")
	}
	return b.String()
}

type JSONMutator func(obj map[string]any)

func (r *LogEntry) PrettifyJSON(mutators ...JSONMutator) {
	r.Request.PrettifyJSON(mutators...)
	r.Response.PrettifyJSON(mutators...)
}

func (r *Response) PrettifyJSON(mutators ...JSONMutator) {
	r.Body = prettifyJSON(r.Body, mutators...)
}

func (r *Request) PrettifyJSON(mutators ...JSONMutator) {
	r.Body = prettifyJSON(r.Body, mutators...)
}

func prettifyJSON(s string, mutators ...JSONMutator) string {
	if s == "" {
		return s
	}

	obj := make(map[string]any)
	if err := json.Unmarshal([]byte(s), &obj); err != nil {
		klog.Fatalf("error from json.Unmarshal(%q): %v", s, err)
		return s
	}

	for _, mutator := range mutators {
		mutator(obj)
	}

	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		klog.Fatalf("error from json.MarshalIndent: %v", err)
		return s
	}
	return string(b)
}

func (r *Response) RemoveHeader(key string) {
	r.Header.Del(key)
}

func (r *Request) RemoveHeader(key string) {
	r.Header.Del(key)
}
