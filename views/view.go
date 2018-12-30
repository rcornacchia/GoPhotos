package views

import (
	"html/template"
	"path/filepath"
)

// LayoutDir and TemplateExt are consts for globbing goHTML files
var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

// NewView parses all template files
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// View is struct that holds template
type View struct {
	Template *template.Template
	Layout   string
}

// layoutFiles returns a slice of strings representing
// the layout files used in our application
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
