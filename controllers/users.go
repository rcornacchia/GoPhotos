package controllers

import "lenslocked.com/views"

// NewUsers is used to create a new Users controller
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "view/users/new.gohtml"),
	}
}

// Users exports the NewView
type Users struct {
	NewView *views.View
}

func (u *Users) {

}