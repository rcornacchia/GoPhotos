package controllers

import "lenslocked.com/views"

// NewStatic returns a Static struct with all the static views
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
	}
}

// Static is a struct that holds all static views
type Static struct {
	Home    *views.View
	Contact *views.View
}
