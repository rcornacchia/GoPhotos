package controllers

import (
	"net/http"

	"lenslocked.com/views"
)

// NewUsers is used to create a new Users controller
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// Users exports the NewView
type Users struct {
	NewView *views.View
}

// New method for rendering the signup view
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
