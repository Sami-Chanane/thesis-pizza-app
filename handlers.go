package main

import (
	"net/http"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.html")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.html")
}

// About is the handler for the menu page
func Menu(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "menu.html")
}

// About is the handler for the gallery page
func Gallery(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "gallery.html")
}

// About is the handler for the contact page
func Contact(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "contact.html")
}
