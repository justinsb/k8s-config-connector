package runner

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

func BuildRunMockgcpTestCommand(ToolOptions *ToolOptions) *cobra.Command {
	var opts RunMockgcpTestOptions
	opts.ToolOptions = ToolOptions
	cmd := &cobra.Command{
		Use:   "run-mockgcp-test",
		Short: "Run a mockgcp test.",
		RunE: func(cmd *cobra.Command, args []string) error {
			out := cmd.OutOrStdout()

			results, err := RunMockgcpTest(cmd.Context(), opts)
			if err != nil {
				return err
			}
			fmt.Fprintf(out, "%d affectedPaths:\n", len(results.AffectedPaths))
			for _, affectedPath := range results.AffectedPaths {
				fmt.Fprintf(out, "  %s\n", affectedPath)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&opts.Resource, "resource", opts.Resource, "The resource name we are generating")
	cmd.Flags().StringVar(&opts.Group, "group", opts.Group, "The resource group we are generating")

	return cmd
}

// RunMockgcpTestOptions is the options for the RunMockgcpTest tool.
type RunMockgcpTestOptions struct {
	*ToolOptions

	// Group is the name of the proto service.
	Group string

	// Resource is the name of the proto resource.
	Resource string
}

func (o *RunMockgcpTestOptions) DefaultAndValidate() error {
	if err := o.ToolOptions.DefaultAndValidate(); err != nil {
		return err
	}
	if o.Resource == "" {
		return fmt.Errorf("resource is required")
	}
	if o.Group == "" {
		return fmt.Errorf("group is required")
	}
	return nil
}

// RunMockgcpTestResults holds the results from executing the RunMockgcpTest tool.
type RunMockgcpTestResults struct {
	AffectedPaths []string

	// TODO: This feels like a layering violation
	ExecResults *ExecResults
}

// RunMockgcpTest executes the mockgcp test.
func RunMockgcpTest(ctx context.Context, opts RunMockgcpTestOptions) (*RunMockgcpTestResults, error) {
	results := &RunMockgcpTestResults{}

	if err := opts.DefaultAndValidate(); err != nil {
		return nil, err
	}

	// Check to see if the script file exists
	scriptFullPath := filepath.Join(opts.RepoRoot, "mockgcp", fmt.Sprintf("mock%s", opts.Group), "testdata", opts.Resource, "crud", "script.yaml")
	if _, err := os.Stat(scriptFullPath); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("missing script %s", scriptFullPath)
	}

	// Check to see if the http log file already exists
	logFileRelativePath := filepath.Join("mockgcp", fmt.Sprintf("mock%s", opts.Group), "testdata", opts.Resource, "crud", "_http.log")
	logFilePath := filepath.Join(opts.RepoRoot, logFileRelativePath)
	if _, err := os.Stat(logFilePath); !errors.Is(err, os.ErrNotExist) && !opts.Force {
		return nil, fmt.Errorf("http log %s already exists", logFilePath)
	}

	// Real tests can take a while
	timeout := time.Hour

	// Current HTTP Log generation is determenistic not ML generated.

	// Run the test to generate the log.
	cfg := CommandConfig{
		Name: "Generate HTTP log",
		Cmd:  "go",
		Args: []string{
			"test", "./mockgcptests",
			"-v",
			"-run", fmt.Sprintf("TestScripts/mock%s/testdata/%s/crud", opts.Group, opts.Resource),
			"-timeout", timeout.String(),
		},
		WorkDir:     filepath.Join(opts.RepoRoot, "mockgcp"),
		Env:         map[string]string{"WRITE_GOLDEN_OUTPUT": "1", "E2E_GCP_TARGET": "real"},
		MaxAttempts: 1,
	}
	var execOpts RunnerOptions
	execResults, err := executeCommand(&execOpts, cfg)
	results.ExecResults = &execResults
	if err != nil {
		return results, err
	}
	results.AffectedPaths = append(results.AffectedPaths, logFileRelativePath)
	return results, err
}

func captureHttpLog(ctx context.Context, runnerOptions *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	var opts RunMockgcpTestOptions

	opts.RepoRoot = runnerOptions.branchRepoDir
	opts.ScratchDir = filepath.Join(runnerOptions.loggingDir, branch.Name)
	opts.Force = runnerOptions.force

	opts.Resource = branch.Resource
	opts.Group = branch.Group

	results, err := RunMockgcpTest(ctx, opts)
	if err != nil {
		return nil, nil, err
	}
	return results.AffectedPaths, results.ExecResults, nil

}
