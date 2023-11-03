package httpmux

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"google.golang.org/genproto/googleapis/longrunning"
)

// roundTripOperation serves the IAM policy verbs (e.g. :getIamPolicy)
// These are implemented on most resources, and rather than mock them
// per-resource, we implement them once here.
func (m *Mux) roundTripOperation(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	requestPath := req.URL.Path

	name := requestPath //lastComponent(requestPath)
	name = strings.TrimPrefix(name, "/")

	// Remove the version
	name = strings.TrimPrefix(name, "v1beta1/")
	name = strings.TrimPrefix(name, "v1/")

	getOperationRequest := &longrunning.GetOperationRequest{
		Name: name,
	}

	op, err := m.operations.GetOperation(ctx, getOperationRequest)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(op)
	if err != nil {
		return nil, err
	}
	body := io.NopCloser(bytes.NewReader(b))
	return &http.Response{StatusCode: http.StatusOK, Body: body}, nil
}

// func lastComponent(s string) string {
// 	i := strings.LastIndex(s, "/")
// 	return s[i+1:]
// }
