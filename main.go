package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Welcome to my go website!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1>Contact Page</h1><p>get in touch with my
	 email: <a href="mailto:behzad@engineer.com">behzad@engineer.com</a>`)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handleFunc(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found!", http.StatusNotFound)

	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Staring the server on port 3000...")
	http.ListenAndServe(":3000", nil)
}
