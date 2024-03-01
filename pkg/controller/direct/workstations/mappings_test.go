package workstations

import (
	"fmt"
	"os"
	"testing"
)

func TestMappings(t *testing.T) {
	os.Setenv("CHECK_COVERAGE", "1")
	errs := workstationClusterMapping.Validate()
	for _, err := range errs {
		t.Errorf("error mapping: %v", err)
		if err.Proposal != "" {
			fmt.Printf("%v\n", err.Proposal)
		}
	}
}
