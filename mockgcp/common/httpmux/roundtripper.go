package httpmux

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type Mux struct {
	mux        *runtime.ServeMux
	operations *operations.Operations
}

var _ http.RoundTripper = &Mux{}

func (m *Mux) WithOperations(operations *operations.Operations) *Mux {
	m.operations = operations
	return m
}

func (m *Mux) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.operations != nil {
		log.Printf("request: %v %v", req.Method, req.URL)

		requestPath := req.URL.Path
		if strings.Contains(requestPath, "/operations/") {
			return m.roundTripOperation(req)
		}
	}

	var body bytes.Buffer
	w := &bufferedResponseWriter{body: &body, header: make(http.Header)}
	m.mux.ServeHTTP(w, req)
	response := &http.Response{}
	response.Body = ioutil.NopCloser(&body)
	response.Header = w.header
	if w.statusCode == 0 {
		w.statusCode = 200
	}
	response.Status = fmt.Sprintf("%d %s", w.statusCode, http.StatusText(w.statusCode))
	response.StatusCode = w.statusCode
	return response, nil
}

type FilterFunc func(w http.ResponseWriter, r *http.Request, pathParams map[string]string, inner func(w http.ResponseWriter, r *http.Request))

// HandlePath allows custom handling of paths
func (m *Mux) HandlePath(method string, pathPattern string, filter FilterFunc) error {
	inner := func(w http.ResponseWriter, r *http.Request) {
		m.mux.ServeHTTP(w, r)
	}
	h := func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		filter(w, r, pathParams, inner)
	}
	return m.mux.HandlePath(method, pathPattern, h)
}

// bufferedResponseWriter implements http.ResponseWriter and stores the response.
type bufferedResponseWriter struct {
	statusCode int
	body       io.Writer
	header     http.Header
}

var _ http.ResponseWriter = &bufferedResponseWriter{}

// Header implements http.ResponseWriter
func (w *bufferedResponseWriter) Header() http.Header {
	return w.header
}

// Write implements http.ResponseWriter
func (w *bufferedResponseWriter) Write(b []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = 200
	}
	return w.body.Write(b)
}

// WriteHeader implements http.ResponseWriter
func (w *bufferedResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}
