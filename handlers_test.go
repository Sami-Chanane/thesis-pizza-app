package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Home))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Log(err)
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("for %s expected %d but got %d", "Home", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()
}

func TestAbout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Home))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Log(err)
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("for %s expected %d but got %d", "About", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()
}

func TestMenu(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Home))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Log(err)
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("for %s expected %d but got %d", "Menu", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()
}

func TestGallery(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Home))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Log(err)
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("for %s expected %d but got %d", "Gallery", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()
}

func TestContact(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Home))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Log(err)
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("for %s expected %d but got %d", "Contact", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()
}
