package gocode

import (
	"testing"
)

func TestAddField(t *testing.T) {
	goPackage := "github.com/GoogleCloudPlatform/apis"
	baseDir := "../../../../apis/"

	packages, err := LoadPackageTree(goPackage, baseDir)
	if err != nil {
		t.Fatalf("inspecting go code: %v", err)
	}

	for _, p := range packages {
		if p.GoPackage == "github.com/GoogleCloudPlatform/apis/monitoring/v1beta1" {
			s := p.FindStruct("Widget")
			if s == nil {
				t.Fatalf("could not find struct Widget")
			}

			// 			// A chart of alert policy data.
			// AlertChart *AlertChart `json:"alertChart,omitempty"`

			// addField := &ast.Field{}
			// addField.Type = &ast.Ident{Name: "*AlertChart"}
			// addField.Names = []*ast.Ident{{Name: "AlertChart"}}
			// addField.Tag = &ast.BasicLit{Kind: token.STRING, Value: `json:"alertChart,omitempty"`}
			// addField.Comment = &ast.CommentGroup{
			// 	List: []*ast.Comment{
			// 		{Text: "// A chart of alert policy data."},
			// 	},
			// }
			// s.parent.ast.Comments = append(s.parent.ast.Comments, &ast.CommentGroup{
			// 	List: []*ast.Comment{
			// 		{Text: "// A chart of alert policy data."},
			// 	},
			// }
			// s.parent.
			addField := "// A chart of alert policy data.\nAlertChart *AlertChart `json:\"alertChart,omitempty\"`"
			if err := s.InsertField(addField); err != nil {
				t.Fatalf("error inserting field: %v", err)
			}
			// if err := s.WriteFile(); err != nil {
			// 	t.Fatalf("error writing file: %v", err)
			// }
			return
		}
	}

	t.Fatalf("did not find Widget in monitoring package")
}
