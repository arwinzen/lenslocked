package main

import (
	"fmt"
	"net/http"
	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to your awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:arwingoo90@gmail.com\">arwingoo90@gmail.com</a>.")
}

// the Request struct has a URL field which contains a Path field
// this represents the path that the user entered in the browser search bar
// by using this we can build a simple router to understand what's going on
// under the hood
func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		// TODO: Handle the page not found error
		w.Header().Set("Content-Type", "text/html; charset=ut-8")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Page not found</h1>")
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	fmt.Fprintln(os.Stdout, "Hello World!")
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
