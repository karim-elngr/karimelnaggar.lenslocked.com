package main

import (
	"html/template"
	"os"
)

type User struct {
	Name    string
	Age     int
	IsMale  bool
	Friends []string
	Games   map[string]string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name:    "John Smith",
		Age:     24,
		IsMale:  true,
		Friends: []string{"Mike", "Brad"},
		Games: map[string]string{
			"league of legends": "professional",
		},
	}
	printUser(&data, t)

	data = User{
		Name:    "Jane Smith",
		Age:     23,
		IsMale:  false,
		Friends: []string{"Lena", "Sarah", "Monica"},
		Games:   map[string]string{},
	}
	printUser(&data, t)

	data = User{
		Name:    "Jack Smith",
		Age:     8,
		IsMale:  true,
		Friends: []string{},
	}
	printUser(&data, t)
}

func printUser(u *User, t *template.Template) {
	err := t.Execute(os.Stdout, u)
	if err != nil {
		panic(err)
	}
}
