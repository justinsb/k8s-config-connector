package composerenvironment

import (
	"fmt"
	"os"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/mappings"
)

func TestMappings(t *testing.T) {
	ValidateMapping(t, mapping)
}

func ValidateMapping(t *testing.T, mapping *mappings.Mapping) {
	errs := mapping.Validate()
	if len(errs) != 0 {
		for _, err := range errs {
			t.Errorf("%v", err.Message)
			if err.Proposal != "" {
				fmt.Fprintf(os.Stderr, "    %v\n", err.Proposal)
			}
		}
	}
}
