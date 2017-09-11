package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Page represents a webpage. Body can contain a string or HTML.
type Page struct {
	Title string
	Body  []byte
}

type Result struct {
	Result float64
	Error  string
}

// homeHandler serves the HTML containing instructions for app usage.
func homeHandler(w http.ResponseWriter, req *http.Request) {
	displayPage("home", w)
}

// sineHandler displays the result of the sine operation for the given input.
func sineHandler(w http.ResponseWriter, req *http.Request) {
	r := Result{
		Result: 0.123456,
	}
	displayResult(r, w)
}

// loadPage retrieves the HTML file with the given title. Title should match
// filename.
func loadPage(title string) (*Page, error) {
	filename := "html/" + title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func displayResult(r Result, w http.ResponseWriter) {
	t, err := template.ParseFiles("html/result.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, r)
	if err != nil {
		log.Fatal(err)
	}
}

// displayPage loads the page with the given title and writes its content to
// w. An error message is displayed if an error occurs.
func displayPage(title string, w http.ResponseWriter) {
	page, err := loadPage(title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s", page.Body)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/sine", sineHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
