package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Int   int
	Float float64
	Slice []string
	Map   map[string]string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name:  "Mahesh Ponnuru",
		Int:   123555,
		Float: 78.22,
		Slice: []string{"a", "b", "c"},
		Map: map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
			"key4": "value4",
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
