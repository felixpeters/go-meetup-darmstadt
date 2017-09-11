package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	page, err := loadPage("home")
	if err != nil {
		io.WriteString(w, "unable to load page")
	}
	fmt.Fprintf(w, "%s", page.Body)
}

func sineHandler(w http.ResponseWriter, req *http.Request) {
	page, err := loadPage("sine")
	if err != nil {
		io.WriteString(w, "unable to load page")
	}
	fmt.Fprintf(w, "%s", page.Body)
}

func loadPage(title string) (*Page, error) {
	filename := "html/" + title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/sine", sineHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
