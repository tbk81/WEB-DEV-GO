package main

import (
	"fmt"
	"net/http"
)

// const (
// 	StatusNotFound = 404
// )

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my double test again awesome sight!<h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:dev@tbk81.com\">dev@tbk81.com</a>.")
}

// func errorHandler(w http.ResponseWriter, r *http.Request){
// 	w.(http.StatusNotFound)
// }

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found!", http.StatusNotFound) // Does the same as the 2 lines below
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprint(w, "Page not found!")
	}
	// if r.URL.Path == "/" {
	// 	homeHandler(w, r)
	// } else if r.URL.Path == "/contact" {
	// 	contactHandler(w, r)
	// }
	// fmt.Fprint(w, r.URL.Path)
}

func main() {
	http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
