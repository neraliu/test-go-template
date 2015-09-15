package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	check(err)
	err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
	check(err)

	// basic
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
		check(err)
	*/

	// attribute name
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <div {{.}}>!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "<>'&\"")
		check(err)
	*/

	// double-quoted attribute value
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <div class="{{.}}">!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "<>'&\"")
		check(err)
	*/

	// single-quoted attribute value
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <div class='{{.}}'>!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "<>'&\"")
		check(err)
	*/

	// un-quoted attribute value
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <div class={{.}}>!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "<> '&\"")
		check(err)
	*/

	// double-quoted attribute value / CSS
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <div style="{{.}}">!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "<> '&\"")
		check(err)
	*/

	// double-quoted attribute value / CSS
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <div style="color:{{.}}">!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "<> '&\"")
		check(err)
	*/

	// double-quoted attribute value / URI
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <a href="{{.}}">!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "javascript:alert(0);")
		check(err)
	*/

	// double-quoted attribute value / URI
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <a href="javascript:alert('{{.}}')">!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "'();")
		check(err)
	*/

	// script tag
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <script>{{.}}</script>!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "'\"();\b")
		check(err)
	*/

	// style tag
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <style>{{.}}</style>!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "'\"();\b")
		check(err)
	*/
}
