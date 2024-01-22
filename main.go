package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const PORT string = "8080"

func main() {

	// routes
	http.HandleFunc("/home", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/menu", Menu)
	http.HandleFunc("/gallery", Gallery)
	http.HandleFunc("/contact", Contact)
	// handle static files
	http.Handle("/", http.FileServer(http.Dir("./static")))
	// starting server
	fmt.Println("Server listening on port", PORT)
	http.ListenAndServe(":"+PORT, nil)

}

func useless() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

// Simple implementation of an integer minimum
// Adapted from: https://gobyexample.com/testing-and-benchmarking
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
