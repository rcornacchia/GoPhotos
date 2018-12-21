package views

import "html/template"

// NewView parses all template files
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}

// View is struct that holds template
type View struct {
	Template *template.Template
}
