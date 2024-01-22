package main

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templateName string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + templateName)
	if err != nil {
		panic("Template not found.")
	}
	parsedTemplate.Execute(w, nil)
}
