package main

import (
	"html/template"
	"os"
)

// User struct
type User struct {
	Name string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name: "John Smith",
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	data.Name = "Rob C"
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
