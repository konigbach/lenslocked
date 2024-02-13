package main

import (
	"html/template"
	"os"
)

type User struct {
	Name    string
	Address map[string]string
}

func main() {
	tpl, err := template.ParseFiles("test.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "John Doe",
		Address: map[string]string{
			"home": "1234 Elm St",
			"work": "5678 Oak St",
		},
	}
	err = tpl.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
