package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func TemplateMust() {
	course := Course{"Go", 40}

	t := template.Must(template.New("TemplateCourse").Parse("Course: {{.Name}} - Hours: {{.Hours}}"))

	err := t.Execute(os.Stdout, course)

	if err != nil {
		panic(err)
	}
}

func TemplateFile() {

	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Java", 20},
		{"Python", 10},
	})

	if err != nil {
		panic(err)
	}
}

func TemplateWebServer() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))

		err := t.Execute(w, Courses{
			{"Go", 40},
			{"Java", 20},
			{"Python", 10},
		})

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8282", nil)
}

func TemplateCompose() {
	templates := []string{"header.html", "content.html", "footer.html"}

	t := template.Must(template.New("content.html").ParseFiles(templates...))

	err := t.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Java", 20},
		{"Python", 10},
	})

	if err != nil {
		panic(err)
	}
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func TemplateMappingFunctions() {
	templates := []string{"header.html", "content.html", "footer.html"}

	t := template.New("content.html")

	t.Funcs(template.FuncMap{"ToUpper": ToUpper})

	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Java", 20},
		{"Python", 10},
	})

	if err != nil {
		panic(err)
	}
}

func main() {
	TemplateMappingFunctions()
}
