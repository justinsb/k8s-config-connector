package llm

import "os"

func AttachFile(p string) (Part, error) {
	b, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return &filePart{
		path:     p,
		contents: string(b),
	}, nil
}

type filePart struct {
	path     string
	contents string
}

var _ Part = &filePart{}

func (f *filePart) AsText() (string, bool) {
	return "", false
}

func (f *filePart) AsFunctionCalls() ([]FunctionCall, bool) {
	return nil, false
}
