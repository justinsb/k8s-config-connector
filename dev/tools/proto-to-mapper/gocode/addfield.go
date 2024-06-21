package gocode

import (
	"bytes"
	"fmt"
	"os"
)

func (p *Package) FindStruct(structName string) *GoStruct {
	for _, t := range p.Files {
		f := t.FindStruct(structName)
		if f != nil {
			return f
		}
	}
	return nil
}

func (p *GoFile) FindStruct(structName string) *GoStruct {
	for _, t := range p.Structs {
		if t.Name == structName {
			return t
		}
	}
	return nil
}

func (s *GoStruct) InsertField(addField string) error {
	fileSet := s.parent.parent.fileSet

	position := fileSet.PositionFor(s.ast.Fields.Closing, true)
	p := position.Filename

	b, err := os.ReadFile(p)
	if err != nil {
		return fmt.Errorf("reading file %q: %w", p, err)
	}

	var out bytes.Buffer
	out.Write(b[:position.Offset])
	out.WriteString("\n")
	out.WriteString(addField)
	out.WriteString("\n")
	out.Write(b[position.Offset:])

	// 	oldLines := strings.Split(string(b), "\n")
	// 	var newLines []string
	// 	for i := 0; i < position.Line; i++ {
	// 		newLines = append(newLines, oldLines[i])
	// 	}
	// 	for _, line := range strings.Split(addField, "\n") {
	// 		newLines = append(newLines, line)
	// 	}
	// 	for i := position.Line; i < len(oldLines); i++ {
	// 		newLines = append(newLines, oldLines[i])
	// 	}

	// 	newLines = append(newLines, addField)
	// 	for i := position.Line; i < len(oldLines); i++ {
	// 		newLines = append(newLines, oldLines[i])
	// 	}
	// 	lines = append(lines, addField)

	// 	newFileSet := token.NewFileSet()
	// 	newFile, err := parser.ParseFile(newFileSet, fileName, fileName, 0)
	// 	if err != nil {
	// 		return fmt.Errorf("error parsing file %q: %w", fileName, err)
	// 	}

	// 	newFile.
	// 		done := false

	// 	astutil.Apply(s.parent.ast, nil, func(c *astutil.Cursor) bool {
	// 		n := c.Node()
	// 		switch x := n.(type) {
	// 		case *ast.StructType:
	// 			if x == s.ast {
	// 				x.Fields.List = append(x.Fields.List, addField)
	// 				done = true
	// 			}
	// 		}

	// 		return true
	// 	})

	// 	if !done {
	// 		return fmt.Errorf("did not find expected struct ast.Node")
	// 	}
	// 	return nil
	// }

	// func (s *GoStruct) WriteFile() error {
	// 	goFile := s.parent
	// 	fileSet := goFile.parent.fileSet
	// 	// pos := s.ast.Pos()
	// 	// position := s.fileSet.PositionFor(pos, true)
	// 	// file := fileSet.File(pos)

	// 	var w bytes.Buffer
	// 	if err := format.Node(&w, fileSet, goFile.ast); err != nil {
	// 		return fmt.Errorf("building go code: %w", err)
	// 	}

	// p := goFile.fileName
	if err := os.WriteFile(p, out.Bytes(), 0644); err != nil {
		return fmt.Errorf("writing file %q: %w", p, err)
	}
	return nil
}

// func (s *GoStruct) InsertField(addField *ast.Field) error {
// 	done := false

// 	astutil.Apply(s.parent.ast, nil, func(c *astutil.Cursor) bool {
// 		n := c.Node()
// 		switch x := n.(type) {
// 		case *ast.StructType:
// 			if x == s.ast {
// 				x.Fields.List = append(x.Fields.List, addField)
// 				done = true
// 			}
// 		}

// 		return true
// 	})

// 	if !done {
// 		return fmt.Errorf("did not find expected struct ast.Node")
// 	}
// 	return nil
// }

// func (s *GoStruct) WriteFile() error {
// 	goFile := s.parent
// 	fileSet := goFile.parent.fileSet
// 	// pos := s.ast.Pos()
// 	// position := s.fileSet.PositionFor(pos, true)
// 	// file := fileSet.File(pos)

// 	var w bytes.Buffer
// 	if err := format.Node(&w, fileSet, goFile.ast); err != nil {
// 		return fmt.Errorf("building go code: %w", err)
// 	}

// 	p := goFile.fileName
// 	if err := os.WriteFile(p, w.Bytes(), 0644); err != nil {
// 		return fmt.Errorf("writing file %q: %w", p, err)
// 	}
// 	return nil
// }
