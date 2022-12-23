package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}

	return t
}

func Parse(filePath string) (Template, error) {
	tpl, err := template.ParseFiles(filePath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing error: %w", err)
	}

	return Template{
		Html: tpl,
	}, nil
}

type Template struct {
	Html *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	err := t.Html.Execute(w, data)
	if err != nil {
		log.Printf("Parsing error: %v", err)
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		return
	}
}
