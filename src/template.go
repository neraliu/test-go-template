package main

import (
	"html/template"
	"log"
	"os"
)

/*
| Context                       | Contextual Filtering  |
---------------------------------------------------------
| HTML                          | Yes                   |
| Comment                       | No data binding       |
| Attr Name                     | Yes 			|

| Double Quoted Attr Value      | Yes                   |
| Single Quoted Attr Value      | Yes                   |
| UnQuoted Attr Value           | Yes                   |

| Quoted Attr Value / CSS       | Yes			|
| Quoted Attr Value / URI       | Yes                   |

| <script> tag                  | Yes			|
| <style> tag                   | Yes			|
*/

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
	// result: filtering as expected.
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
		check(err)
	*/

	// comment
	// result: no data binding.
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <!-- {{.}} -->!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "<>'&\"")
		check(err)
	*/

	// attribute name
	// result: 'ZgotmplZ' is injected
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <div {{.}}>!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "<>'&\"")
		check(err)
	*/

	// double-quoted attribute value
	// result: filtering as expected.
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <div class="{{.}}">!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "<>'&\"")
		check(err)
	*/

	// single-quoted attribute value
	// result: filtering as expected.
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <div class='{{.}}'>!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "<>'&\"")
		check(err)
	*/

	// un-quoted attribute value
	// result: filtering as expected.
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <div class={{.}}>!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "<> '&\"")
		check(err)
	*/

	// double-quoted attribute value / CSS
	// result: 'ZgotmplZ' is injected
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <div style="{{.}}">!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "x:expression(alert(0))")
		check(err)
	*/

	// double-quoted attribute value / CSS
	// result: 'ZgotmplZ' is injected
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <div style="color:{{.}}">!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "x:expression(alert(0))")
		check(err)
	*/

	// double-quoted attribute value / URI
	// result: '#ZgotmplZ' is injected
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <a href="{{.}}">!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "javascript:alert(0);")
		check(err)
	*/

	// double-quoted attribute value / URI
	// result: filtering as expected.
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <a href="javascript:alert('{{.}}')">!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "'();")
		check(err)
	*/

	// script tag
	// result: filtering as expected.
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <script>{{.}}</script>!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "'\"();\b")
		check(err)
	*/

	// style tag
	// result: 'ZgotmplZ' is injected
	/*
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, <style>{{.}}</style>!{{end}}`)
		check(err)
		err = t.ExecuteTemplate(os.Stdout, "T", "'\"();\b")
		check(err)
	*/
}
