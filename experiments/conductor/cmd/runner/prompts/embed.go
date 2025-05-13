package prompts

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"
)

//go:embed *.txt
var fs embed.FS

func GetTemplate(key string) (*template.Template, error) {
	b, err := fs.ReadFile(key)
	if err != nil {
		return nil, fmt.Errorf("reading embedded file %q: %w", key, err)
	}

	t, err := template.New(key).Parse(string(b))
	if err != nil {
		return nil, fmt.Errorf("parsing template %q: %w", key, err)
	}
	return t, nil
}

type MockgcpGenerateScriptPromptOptions struct {
	GcloudCommand string
	Group         string
	Resource      string
}

func GenerateMockgcpGenerateScriptPrompt(opts MockgcpGenerateScriptPromptOptions) ([]byte, error) {
	key := "mockgcp-generate-script.txt"
	t, err := GetTemplate(key)
	if err != nil {
		return nil, err
	}

	var bb bytes.Buffer
	if err := t.Execute(&bb, opts); err != nil {
		return nil, fmt.Errorf("running %q template: %w", key, err)
	}

	return bb.Bytes(), nil
}
