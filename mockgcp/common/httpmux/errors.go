package httpmux

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

type wrappedStatus struct {
	Error *wrappedError `json:"error,omitempty"`
}

type wrappedError struct {
	Code    int            `json:"code,omitempty"`
	Message string         `json:"message,omitempty"`
	Status  string         `json:"status,omitempty"`
	Errors  []errorDetails `json:"errors,omitempty"`
}

type errorDetails struct {
	Domain  string `json:"domain,omitempty"`
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

// customErrorHandler wraps errors in an error blockk
func customErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	s := status.Convert(err)
	// pb := s.Proto()

	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	httpStatusCode := runtime.HTTPStatusFromCode(s.Code())
	wrapped := &wrappedStatus{
		Error: &wrappedError{
			Code:    httpStatusCode,
			Message: s.Message(),
		},
	}

	switch s.Code() {
	case codes.PermissionDenied:
		wrapped.Error.Status = "PERMISSION_DENIED"
	case codes.AlreadyExists:
		wrapped.Error.Status = "ALREADY_EXISTS"
	case codes.NotFound:
		wrapped.Error.Status = "NOT_FOUND"
		wrapped.Error.Errors = append(wrapped.Error.Errors, errorDetails{
			Domain:  "global",
			Message: wrapped.Error.Message,
			Reason:  "notFound",
		})
	}

	buf, merr := json.Marshal(wrapped)
	if merr != nil {
		klog.Warningf("Failed to marshal error message %q: %v", s, merr)
		runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
		return
	}

	if err := addGCPHeaders(ctx, w, nil); err != nil {
		klog.Warningf("unexpected error from header filter: %v", err)
	}

	w.WriteHeader(httpStatusCode)
	if _, err := w.Write(buf); err != nil {
		klog.Warningf("Failed to write response: %v", err)
	}

}
